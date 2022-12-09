package script

import (
	"crypto/tls"
	"fmt"
	"time"
)

type SSL struct {
	HostnameVerification string
	Issuer               string
	ExpirationDate       string

	Result string
}

func (s *SSL) Execute(host string) {
	fmt.Println()
	fmt.Println("Start validating SSL certificate...")

	port := 443

	targetAddr := fmt.Sprintf("%s:%d", host, port)
	conn, err := tls.Dial("tcp", targetAddr, nil)
	if err != nil {
		fmt.Printf("Server doesn't support SSL certificates: %s", err)
	}

	// Hostname
	err = conn.VerifyHostname(host)
	if err == nil {
		s.HostnameVerification = "Hostname is correctly listed in the certicate."
	} else {
		s.HostnameVerification = fmt.Sprintf("Hostname doesn't match with the certificate: %s", err)
	}

	peerCert := conn.ConnectionState().PeerCertificates[0]

	// Issuer
	s.Issuer = peerCert.Issuer.String()

	// Expiration date
	expiry := peerCert.NotAfter
	s.ExpirationDate = expiry.Format(time.RFC850)

	s.createResult(host)
}

// Create result
func (s *SSL) createResult(host string) {
	s.Result = fmt.Sprintf(`
=================================================================
SSL certificates for %s
=================================================================
■ Hostname
%s
■ Issuer
%s
■ Expiration Date
%s
=================================================================
`,
		host,
		s.HostnameVerification,
		s.Issuer,
		s.ExpirationDate)
}
