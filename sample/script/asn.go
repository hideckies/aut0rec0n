package script

import (
	"fmt"
)

type ASN struct {
	Result         string
	ResultColor    string
	ResultContents string
}

func (a *ASN) Execute(host string, ip string) {
	if ip == "" {
		fmt.Println("ASN: ip address is not specified.")
	}

	a.createResultContents()
}

func (a *ASN) createResultContents() {
	a.ResultContents = fmt.Sprintf(``)
}
