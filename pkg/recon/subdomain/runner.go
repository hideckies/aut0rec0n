package subdomain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/hideckies/aut0rec0n/pkg/config"
	"github.com/hideckies/aut0rec0n/pkg/output"
	"github.com/hideckies/aut0rec0n/pkg/progress"
	github "github.com/hideckies/aut0rec0n/pkg/sources"
	shodan "github.com/hideckies/aut0rec0n/pkg/sources"
	virusTotal "github.com/hideckies/aut0rec0n/pkg/sources"
	"github.com/hideckies/aut0rec0n/pkg/util"
	googlesearch "github.com/rocketlaunchr/google-search"
)

type Config struct {
	Host      string
	UserAgent string

	ApiKeys config.ApiKeys
}

type Result struct {
	Subdomains []string
}

type Runner struct {
	Config Config
	Result Result
}

// Initialize a new Subdomain
func NewRunner(host string, conf config.Config) Runner {
	var r Runner
	r.Config = Config{
		Host:      host,
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		ApiKeys:   conf.ApiKeys,
	}
	r.Result = Result{}
	return r
}

// Execute enumerating subdomains
func (r *Runner) Run() error {
	threadsNum := 4

	var wg sync.WaitGroup
	wg.Add(threadsNum)

	ch := make(chan int)

	go r.getFromGoogle(&wg, ch)
	go r.getFromGitHub(&wg, ch)
	go r.getFromShodan(&wg, ch)
	go r.getFromVirusTotal(&wg, ch)

	wg.Wait()

	// Unique subdomains slice
	uniqSubdomains := make([]string, 0)
	for i := 0; i < len(r.Result.Subdomains); i++ {
		if r.Result.Subdomains[i] != r.Config.Host && !util.StrArrContains(uniqSubdomains, r.Result.Subdomains[i]) {
			uniqSubdomains = append(uniqSubdomains, r.Result.Subdomains[i])
		}
	}

	// Finally add the unique subdomains to the result
	r.Result.Subdomains = uniqSubdomains

	r.Print()
	return nil
}

// Search Google for enumerating subdomains
func (r *Runner) getFromGoogle(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	searchTxt := fmt.Sprintf("site:%s", r.Config.Host)
	result, err := googlesearch.Search(
		nil,
		searchTxt,
		googlesearch.SearchOptions{
			Limit:     100,
			UserAgent: r.Config.UserAgent,
		})
	if err != nil {
		color.Red("%s", err)
		return
	}

	for _, result := range result {
		resultUrl := result.URL
		separatedUrls := strings.Split(resultUrl, "/")
		newSubdomain := strings.Join(separatedUrls[2:3], "/")
		// Remove port strings
		rePort := regexp.MustCompile(`\:\d+`)
		newSubdomain = rePort.ReplaceAllString(newSubdomain, "")

		r.Result.Subdomains = append(r.Result.Subdomains, newSubdomain)
	}
}

// Fetch from GitHub API
func (r *Runner) getFromGitHub(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	fetchUrl := fmt.Sprintf("https://api.github.com/search/code?q=%s&per_page=100&sort=created&order=asc", r.Config.Host)

	client := &http.Client{}
	req, err := http.NewRequest("GET", fetchUrl, nil)
	if err != nil {
		color.Red("%s", err)
		return
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.Config.ApiKeys.GitHub))

	resp, err := client.Do(req)
	if err != nil {
		color.Red("%s", err)
		return
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode == 401 {
		color.Red("GitHub: 401 authorized\nDid you set the GitHub access token in ~/.config/aut0rec0n/config.yaml ?")
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Red("%s", err)
		return
	}

	// Parse the JSON
	var respData github.GitHub
	err = json.Unmarshal(body, &respData)
	if err != nil {
		color.Red("%s", err)
		return
	}

	// Enumerate subdomains in source codes in each repository
	reSubdomain := regexp.MustCompile(fmt.Sprintf("[a-zA-Z0-9]+\\.%s", r.Config.Host))

	// Set progress bar
	bar := *progress.NewProgressBar(len(respData.Items), "Fetching GitHub API...")

	for _, item := range respData.Items {
		bar.Add(1)
		defer time.Sleep(200 * time.Millisecond)

		targetUrl := item.HtmlUrl
		req, err := http.NewRequest("GET", targetUrl, nil)
		if err != nil {
			color.Red("%s", err)
			return
		}

		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.Config.ApiKeys.GitHub))

		resp, err := client.Do(req)
		if err != nil {
			color.Red("%s", err)
			return
		}
		defer resp.Body.Close()

		// Circumvent rate limiting
		rateLimitRemaining, _ := strconv.ParseInt(resp.Header.Get("X-Ratelimit-Remaining"), 10, 64)
		if rateLimitRemaining == 0 {
			// If the rate limit remaining is 0, sleep until the resty after seconds
			retryAfterSeconds, _ := strconv.ParseInt(resp.Header.Get("Retry-After"), 10, 64)
			if retryAfterSeconds > 0 {
				time.Sleep(time.Duration(retryAfterSeconds+1) * time.Second)
			}
		}

		// Check the status code
		if resp.StatusCode == 401 {
			color.Red("GitHub: 401 authorized\nDid you set the GitHub access token in ~/.config/aut0rec0n/config.yaml ?")
			return
		}

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			color.Red("%s", err)
			return
		}

		// Find subdomains
		results := reSubdomain.FindAllString(string(body), -1)
		for i := range results {
			// Remove "%2F" (URL encoded "/") from a subdomain
			if results[i][:2] == "2F" {
				results[i] = results[i][2:]
			}
			r.Result.Subdomains = append(r.Result.Subdomains, results[i])
		}
	}
}

// Fetch from Shodan API
func (r *Runner) getFromShodan(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	fetchUrl := fmt.Sprintf("https://api.shodan.io/dns/domain/%s?key=%s", r.Config.Host, r.Config.ApiKeys.Shodan)
	resp, err := http.Get(fetchUrl)
	if err != nil {
		color.Red("%s", err)
		return
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode == 401 {
		color.Red("Shodan: 401 Unauthorized\nDid you set the Shodan API Key in ~/.config/aut0rec0n/config.yaml ?")
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Red("%s", err)
		return
	}

	// Parse the JSON
	var respData shodan.Shodan
	err = json.Unmarshal(body, &respData)
	if err != nil {
		color.Red("%s", err)
		return
	}

	subdomains := make([]string, 0)
	for _, newSubdomain := range respData.Subdomains {
		newSubdomain = fmt.Sprintf("%s.%s", newSubdomain, r.Config.Host)
		r.Result.Subdomains = append(r.Result.Subdomains, newSubdomain)
	}

	r.Result.Subdomains = append(r.Result.Subdomains, subdomains...)
}

// Fetch from VirusTotal API
func (r *Runner) getFromVirusTotal(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	fetchUrl := fmt.Sprintf("https://www.virustotal.com/api/v3/domains/%s/subdomains", r.Config.Host)

	client := &http.Client{}
	req, err := http.NewRequest("GET", fetchUrl, nil)
	if err != nil {
		return
	}

	req.Header.Set("x-apikey", r.Config.ApiKeys.VirusTotal)

	resp, err := client.Do(req)
	if err != nil {
		color.Red("%s", err)
		return
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode == 401 {
		color.Red("VirusTotal: 401 Unauthorized\nDid you set the VirusTotal API Key in ~/.config/aut0rec0n/config.yaml ?")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Red("%s", err)
		return
	}

	// Parse the JSON
	var respData virusTotal.VirusTotal
	err = json.Unmarshal(body, &respData)
	if err != nil {
		color.Red("%s", err)
		return
	}

	subdomains := make([]string, 0)
	for _, data := range respData.Data {
		r.Result.Subdomains = append(r.Result.Subdomains, data.Id)
	}

	r.Result.Subdomains = append(r.Result.Subdomains, subdomains...)
}

// Print the result
func (r *Runner) Print() {
	output.Headline("SUBDOMAIN")
	if len(r.Result.Subdomains) > 0 {
		for _, subdomain := range r.Result.Subdomains {
			color.Green(subdomain)
		}
	} else {
		color.Yellow("could not find subdomains")
	}
}
