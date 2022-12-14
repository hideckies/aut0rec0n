package script

import (
	"fmt"

	"github.com/domainr/whois"
)

type WHOIS struct {
	Records string

	Result         string
	ResultColor    string
	ResultContents string
}

func (w *WHOIS) Execute(host string) {
	req, err := whois.NewRequest(host)
	if err != nil {
		fmt.Printf("! %v\n", err)
		return
	}
	res, err := whois.DefaultClient.Fetch(req)
	if err != nil {
		fmt.Printf("! %v\n", err)
	}
	w.Records = string(res.Body)

	w.ResultContents = w.Records
}
