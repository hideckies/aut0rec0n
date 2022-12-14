package script

import (
	"fmt"
)

type ASN struct {
	Result         string
	ResultColor    string
	ResultContents string
}

func (a *ASN) Execute(host string) {
	a.createResultContents()
}

func (a *ASN) createResultContents() {
	a.ResultContents = fmt.Sprintf(``)
}
