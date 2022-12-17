package script

import (
	"fmt"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
)

type Subdomain struct {
	Subdomains []string

	Result         string
	ResultColor    string
	ResultContents string
}

func (s *Subdomain) Execute(host string) {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.3"

	var subdomains []string
	subdomains = append(subdomains, enumFromGoogle(host, userAgent)...)

	s.Subdomains = subdomains

	s.createResultContents()
}

func enumFromGoogle(host string, userAgent string) []string {
	searchTxt := fmt.Sprintf("site:%s -www", host)
	results, err := googlesearch.Search(
		nil,
		searchTxt,
		googlesearch.SearchOptions{
			Limit:     100,
			UserAgent: userAgent,
		})
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	var subdomains []string

	for _, result := range results {
		resultUrl := result.URL
		separatedUrls := strings.Split(resultUrl, "/")
		subdomain := strings.Join(separatedUrls[2:3], "/")

		if subdomain != host && !domainContains(subdomains, subdomain) {
			subdomains = append(subdomains, subdomain)
		}
	}

	return subdomains
}

func domainContains(domains []string, targetDomain string) bool {
	for _, domain := range domains {
		if domain == targetDomain {
			return true
		}
	}
	return false
}

func (s *Subdomain) createResultContents() {
	subdomains := []string{}
	subdomains = append(subdomains, s.Subdomains...)

	s.ResultContents = strings.Join(subdomains, "\n")
}
