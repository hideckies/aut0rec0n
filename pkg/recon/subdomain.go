package recon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/hideckies/aut0rec0n/pkg/config"
	"github.com/hideckies/aut0rec0n/pkg/output"
	shodan "github.com/hideckies/aut0rec0n/pkg/sources"
	virusTotal "github.com/hideckies/aut0rec0n/pkg/sources"
	"github.com/hideckies/aut0rec0n/pkg/util"
	googlesearch "github.com/rocketlaunchr/google-search"
)

type SubdomainConfig struct {
	Host      string
	UserAgent string

	ApiKeys config.ApiKeys
}

type SubdomainResult struct {
	Subdomains []string
}

type Subdomain struct {
	Config SubdomainConfig
	Result SubdomainResult
}

// Initialize a new Subdomain
func NewSubdomain(host string, conf config.Config) Subdomain {
	var s Subdomain
	s.Config = SubdomainConfig{
		Host:      host,
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		ApiKeys:   conf.ApiKeys,
	}
	s.Result = SubdomainResult{}
	return s
}

// Execute enumerating subdomains
func (s *Subdomain) Execute() error {
	err := s.getFromGoogle()
	if err != nil {
		return err
	}

	err = s.getFromShodan()
	if err != nil {
		return err
	}

	err = s.getFromVirusTotal()
	if err != nil {
		return err
	}

	s.Print()
	return nil
}

// Search Google for enumerating subdomains
func (s *Subdomain) getFromGoogle() error {
	searchTxt := fmt.Sprintf("site:%s -www", s.Config.Host)
	result, err := googlesearch.Search(
		nil,
		searchTxt,
		googlesearch.SearchOptions{
			Limit:     100,
			UserAgent: s.Config.UserAgent,
		})
	if err != nil {
		return err
	}

	subdomains := make([]string, 0)
	for _, result := range result {
		resultUrl := result.URL
		separatedUrls := strings.Split(resultUrl, "/")
		subdomain := strings.Join(separatedUrls[2:3], "/")
		// Remove port strings
		rePort := regexp.MustCompile(`\:\d+`)
		subdomain = rePort.ReplaceAllString(subdomain, "")

		if subdomain != s.Config.Host && !util.StrArrContains(subdomains, subdomain) {
			subdomains = append(subdomains, subdomain)
		}
	}

	s.Result.Subdomains = append(s.Result.Subdomains, subdomains...)
	return nil
}

// Fetch from Shodan API
func (s *Subdomain) getFromShodan() error {
	fetchUrl := fmt.Sprintf("https://api.shodan.io/dns/domain/%s?key=%s", s.Config.Host, s.Config.ApiKeys.Shodan)
	resp, err := http.Get(fetchUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode == 401 {
		color.Red("Shodan: 401 Unauthorized\nDid you set the Shodan API Key in ~/.config/aut0rec0n/config.yaml ?")
		return nil
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse the JSON
	var respData shodan.Shodan
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return err
	}

	subdomains := make([]string, 0)
	for _, newSubdomain := range respData.Subdomains {
		if newSubdomain != s.Config.Host && !util.StrArrContains(s.Result.Subdomains, newSubdomain) {
			subdomains = append(subdomains, fmt.Sprintf("%s.%s", newSubdomain, s.Config.Host))
		}
	}

	s.Result.Subdomains = append(s.Result.Subdomains, subdomains...)
	return nil
}

// Fetch from VirusTotal API
func (s *Subdomain) getFromVirusTotal() error {
	fetchUrl := fmt.Sprintf("https://www.virustotal.com/api/v3/domains/%s/subdomains", s.Config.Host)

	client := &http.Client{}
	req, err := http.NewRequest("GET", fetchUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("x-apikey", s.Config.ApiKeys.VirusTotal)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode == 401 {
		color.Red("VirusTotal: 401 Unauthorized\nDid you set the VirusTotal API Key in ~/.config/aut0rec0n/config.yaml ?")
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse the JSON
	var respData virusTotal.VirusTotal
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return err
	}

	subdomains := make([]string, 0)
	for _, data := range respData.Data {
		if data.Id != s.Config.Host && !util.StrArrContains(s.Result.Subdomains, data.Id) {
			subdomains = append(subdomains, data.Id)
		}
	}

	s.Result.Subdomains = append(s.Result.Subdomains, subdomains...)

	return nil
}

// Print the result
func (s *Subdomain) Print() {
	output.Headline("SUBDOMAIN")
	if len(s.Result.Subdomains) > 0 {
		for _, subdomain := range s.Result.Subdomains {
			color.Green(subdomain)
		}
	} else {
		color.Yellow("could not find subdomains")
	}
}
