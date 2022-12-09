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

	// WHOIS
	if recon.sWHOIS != nil {
		filenameWHOIS := fmt.Sprintf("%s/whois.txt", dirname)
		createFile(filenameWHOIS, recon.sWHOIS.Result)
	}

	// DNS
	if recon.sDNS != nil {
		filenameDNS := fmt.Sprintf("%s/dns.txt", dirname)
		createFile(filenameDNS, recon.sDNS.Result)
	}

	// Subdomain
	if recon.sSubdomain != nil {
		filenameSubdomain := fmt.Sprintf("%s/subdomain.txt", dirname)
		createFile(filenameSubdomain, recon.sSubdomain.Result)
	}

	// SSL certificates
	if recon.sSSL != nil {
		filenameSSL := fmt.Sprintf("%s/ssl.txt", dirname)
		createFile(filenameSSL, recon.sSSL.Result)
	}

	// Web archive
	if recon.sWebArchive != nil {
		filenameWebArchive := fmt.Sprintf("%s/web-archive.txt", dirname)
		createFile(filenameWebArchive, recon.sWebArchive.Result)
	}

	// Port
}

func createFile(filename string, content string) {
	d := []byte(content)
	os.WriteFile(filename, d, 0644)
}
