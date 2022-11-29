package script

import (
	"fmt"
)

type ASN struct {
	Result string
}

func (a *ASN) Execute(host string) {
	fmt.Println()
	fmt.Println("Starting ASN reconnaissance...")

	a.createResult(host)
}

func (a *ASN) createResult(host string) {
	a.Result = fmt.Sprintf(`
=================================================================
ASN reconnaissance for %s
=================================================================
■ 
=================================================================
`,
		host)
}
