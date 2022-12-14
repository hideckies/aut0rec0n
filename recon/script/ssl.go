package script

import (
	"crypto/tls"
	"fmt"
	"time"
)

type SSL struct {
	SupportSSL           string
	HostnameVerification string
	Issuer               string
	ExpirationDate       string

	Result         string
	ResultColor    string
	ResultContents string
}

func (s *SSL) Execute(host string) {
	port := 443

	targetAddr := fmt.Sprintf("%s:%d", host, port)
	conn, err := tls.Dial("tcp", targetAddr, nil)
	if err == nil {
		s.SupportSSL = "Server support SSL certificate."
	} else {
		s.SupportSSL = fmt.Sprintf("Server doesn't support SSL certificates: %s\n", err)
		s.createResultContents()
		return
	}

	// Hostname
	err = conn.VerifyHostname(host)
	if err == nil {
		s.HostnameVerification = "Hostname is correctly listed in the certicate."
	} else {
		s.HostnameVerification = fmt.Sprintf("Hostname doesn't match with the certificate: %s\n", err)
	}

	peerCert := conn.ConnectionState().PeerCertificates[0]

	// Issuer
	s.Issuer = peerCert.Issuer.String()

	// Expiration date
	expiry := peerCert.NotAfter
	s.ExpirationDate = expiry.Format(time.RFC850)

	s.createResultContents()
}

// Create result
func (s *SSL) createResultContents() {
	s.ResultContents = fmt.Sprintf("%s\n■ Hostname\n%s\n■ Issuer\n%s\n■ Expiration Date\n%s",
		s.SupportSSL,
		s.HostnameVerification,
		s.Issuer,
		s.ExpirationDate)
}
