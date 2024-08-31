package ssl

import (
	"fmt"
	"os/exec"
)

func CreateSSLCertificate(domain string) error {
	cmd := exec.Command("certbot", "certonly", "--nginx", "-d", domain, "-d", "www."+domain)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error creating SSL certificate: %v\n%s\n", err, output)
		return err
	}
	fmt.Println("SSL certificate created successfully")
	return nil
}
