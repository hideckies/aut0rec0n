package script

import (
	"fmt"

	"github.com/domainr/whois"
)

type WHOIS struct {
	Records string
	Result  string
}

func (w *WHOIS) Execute(host string) {
	fmt.Println()
	fmt.Println("Start WHOIS reconnaissance...")

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

	w.createResult(host)
}

func (w *WHOIS) createResult(host string) {
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
