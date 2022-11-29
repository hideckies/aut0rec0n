package script

import (
	"fmt"

	"github.com/domainr/whois"
)

type Whois struct {
	Records string
	Result  string
}

func (w *Whois) Execute(host string) {
	fmt.Println()
	fmt.Println("Starting WHOIS reconnaissance...")

	req, err := whois.NewRequest(host)
	if err != nil {
		fmt.Println(err)
	}
	res, err := whois.DefaultClient.Fetch(req)
	if err != nil {
		fmt.Println(err)
	}
	w.Records = string(res.Body)

	w.createResult(host)
}

func (w *Whois) createResult(host string) {
	w.Result = fmt.Sprintf(`
=================================================================
WHOIS for %s
=================================================================
%s
=================================================================
`,
		host,
		w.Records)
}
