package recon

import (
	"fmt"
	"os"
	"time"
)

func Output(recon *Recon) {
	dirname := recon.Conf.OutputDir
	if recon.Conf.OutputDir == "./aut0rec0n-result" {
		dNow := time.Now()
		dYear, dMonth, dDay := dNow.Date()
		dHour, dMinute, dSecond := dNow.Clock()
		dDate := fmt.Sprintf("%d%d%d%d%d%d", dYear, dMonth, dDay, dHour, dMinute, dSecond)

		dirname = fmt.Sprintf("%s_%s_%s", recon.Conf.OutputDir, recon.Conf.Host, dDate)
	}

	err := os.Mkdir(dirname, 0755)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	// DNS
	if recon.sDns != nil {
		filenameDns := fmt.Sprintf("%s/dns.txt", dirname)
		createFile(filenameDns, recon.sDns.Result)
	}

	// Port

	// Subdomain
}

func createFile(filename string, content string) {
	d := []byte(content)
	os.WriteFile(filename, d, 0644)
}
