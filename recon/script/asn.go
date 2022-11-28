package script

import "fmt"

type ASN struct {
	Result string
}

func (a *ASN) Execute(host string) {
	fmt.Println()
	fmt.Println("Starting ASN reconnaissance...")
}

func (a *ASN) createResult(host string) {

}
