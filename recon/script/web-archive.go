package script

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type WebArchive struct {
	Archives []string
	Result   string
}

func (w *WebArchive) Execute(host string, subdomains []string) {
	fmt.Println()
	fmt.Println("Starting web archives reconnaissance...")

	domains := append(subdomains, host)
	for _, domain := range domains {
		apiUrl := fmt.Sprintf("http://archive.org/wayback/available?url=%s", domain)

		res, err := http.Get(apiUrl)
		if err != nil {
			fmt.Printf("%s", err)
		}

		raw, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Printf("%s", err)
		}

		fmt.Println(raw)

		max := 6
		min := 2
		randNum := rand.Intn(max-min) + min
		time.Sleep(time.Duration(randNum * int(time.Second)))
	}

	w.createResult(host)
}

// Create a result
func (w *WebArchive) createResult(host string) {
	w.Result = fmt.Sprintf(`
=================================================================
Web archives for %s
=================================================================
■ 
=================================================================
`,
		host)
}
