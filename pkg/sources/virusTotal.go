package sources

type Attributes struct {
	Categories               struct{}   `json:"categories"`
	CreationDate             string     `json:"creationDate"`
	LastAnalysisResults      struct{}   `json:"lastAnalysisResults"`
	LastAnalysisStats        struct{}   `json:"lastAnalysisStats"`
	LastDnsRecords           []struct{} `json:"lastDnsRecords"`
	LastDnsRecordsDate       int        `json:"lastDnsRecordsDate"`
	LastHttpsCertificate     struct{}   `json:"lastHttpsCertificate"`
	LastHttpsCertificateDate int        `json:"lastHttpsCertificateDate"`
	LastModificationDate     int        `json:"lastModificationDate"`
	LastUpdateDate           int        `json:"lastUpdateDate"`
	PopularityRanks          struct{}   `json:"popularityRanks"`
	Registrar                string     `json:"registrar"`
	Reputation               int        `json:"reputation"`
	Tags                     []string   `json:"tags"`
	TotalVotes               struct{}   `json:"totalVotes"`
	Whois                    string     `json:"whois"`
}

type Domain struct {
	Attributes Attributes `json:"attributes"`
	Id         string     `json:"id"`
	Links      struct{}   `json:"links"`
	Type       string     `json:"type"`
}

type Links struct {
	Next string `json:"next"`
	Self string `json:"string"`
}

type Meta struct {
	Count  int    `json:"count"`
	Cursor string `json:"cursor"`
}

type VirusTotal struct {
	Data  []Domain `json:"data"`
	Links Links    `json:"links"`
	Meta  Meta     `json:"meta"`
}
