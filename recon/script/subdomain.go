package script

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Subdomain struct {
	Subdomains []string

	Result string
}

func (s *Subdomain) Execute(host string) {
	fmt.Println()
	fmt.Println("Starting subdomain scan...")

	searchUrl := fmt.Sprintf("https://www.google.com/search?q=site:%s", host)
	fetchUrl(searchUrl)

	s.createResult(host)
}

func fetchUrl(url string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "aut0rec0n by h1d3k1 15h1gur0 repository -> https://github.com/hideckies/aut0rec0n")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	html, _ := ioutil.ReadAll(resp.Body)
	htmlStr := string(html)
	fmt.Println(htmlStr)
}

func (s *Subdomain) createResult(host string) {
	subdomains := []string{}
	for _, subdomain := range s.Subdomains {
		subdomains = append(subdomains, subdomain)
	}

	s.Result = fmt.Sprintf(`
=================================================================
Subdomain Scanning for %s
=================================================================

%s

=================================================================
`,
		host,
		strings.Join(subdomains, "\n"))
}
