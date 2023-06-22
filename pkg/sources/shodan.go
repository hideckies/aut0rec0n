package sources

type Data struct {
	Subdomain string `json:"subdomain"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	LastSeen  string `json:"last_seen"`
}

type Shodan struct {
	Domain     string   `json:"domain"`
	Tags       []string `json:"tags"`
	Data       []Data   `json:"data"`
	Subdomains []string `json:"subdomains"`
}
