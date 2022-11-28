package script

import (
	"fmt"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
)

type Subdomain struct {
	Subdomains []string

	Result string
}

func (s *Subdomain) Execute(host string) {
	fmt.Println()
	fmt.Println("Starting subdomain scan...")

	// userAgent := "aut0rec0n by h1d3k1 15h1gur0 repository -> https://github.com/hideckies/aut0rec0n"
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.3"

	var subdomains []string
	subdomains = append(subdomains, enumFromGoogle(host, userAgent)...)

	s.Subdomains = subdomains

	s.createResult(host)
}

func enumFromGoogle(host string, userAgent string) []string {
	searchTxt := fmt.Sprintf("site:%s", host)
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

func (s *Subdomain) createResult(host string) {
	subdomains := []string{}
	for _, subdomain := range s.Subdomains {
		subdomains = append(subdomains, subdomain)
	}

	s.Result = fmt.Sprintf(`
=================================================================
Subdomains for %s
=================================================================
%s
=================================================================
`,
		host,
		strings.Join(subdomains, "\n"))
}
