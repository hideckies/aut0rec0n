package recon

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/hideckies/aut0rec0n/pkg/output"
	"github.com/hideckies/aut0rec0n/pkg/util"
	googlesearch "github.com/rocketlaunchr/google-search"
)

type SubdomainConfig struct {
	Host      string
	UserAgent string
}

type SubdomainResult struct {
	Subdomains []string
}

type Subdomain struct {
	Config SubdomainConfig
	Result SubdomainResult
}

// Initialize a new Subdomain
func NewSubdomain(host string) Subdomain {
	var s Subdomain
	s.Config = SubdomainConfig{
		Host:      host,
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.3",
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
