package script

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type WebArchive struct {
	Archives []archive

	Result         string
	ResultColor    string
	ResultContents string
}

type archive struct {
	Url               string   `json:"url"`
	ArchivedSnapshots snapshot `json:"archived_snapshots"`
}

type snapshot struct {
	Closest closest `json:"closest"`
}

type closest struct {
	Status    string `json:"status"`
	Available bool   `json:"available"`
	Url       string `json:"url"`
	Timestamp string `json:"timestamp"`
}

// Execute
func (w *WebArchive) Execute(host string, subdomains []string) {
	domains := append(subdomains, host)
	for _, domain := range domains {
		// Wayback Machine API: https://archive.org/help/wayback_api.php
		apiUrl := fmt.Sprintf("http://archive.org/wayback/available?url=%s", domain)

		res, err := http.Get(apiUrl)
		if err != nil {
			fmt.Printf("Web Archives Error: %s", err)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Web Archives Error: %s", err)
		}

		var ar archive

		json.Unmarshal(body, &ar)

		if ar.ArchivedSnapshots.Closest.Available {
			w.Archives = append(w.Archives, ar)
		}

		interval(2, 6)
	}

	w.createResultContents()
}

// Create a result
func (w *WebArchive) createResultContents() {
	subResults := []string{}

	for _, ar := range w.Archives {
		clos := ar.ArchivedSnapshots.Closest

		// Parse timestamp to date
		var tm time.Time
		i, err := strconv.ParseInt(clos.Timestamp, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		tm = time.Unix(i, 0)

		subResult := fmt.Sprintf("■ %s\nStatus: %s\nURL: %s\nDate: %s", ar.Url, clos.Status, clos.Url, tm.String())
		subResults = append(subResults, subResult)
	}

	w.ResultContents = fmt.Sprintf("%v", strings.Join(subResults, "\n\n"))
}

// Interval
func interval(min int, max int) {
	randNum := rand.Intn(max-min) + min
	time.Sleep(time.Duration(randNum * int(time.Second)))
}
