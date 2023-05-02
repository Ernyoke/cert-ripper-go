package validate

import (
	"cert-ripper-go/pkg/cert"
	"cert-ripper-go/pkg/host"
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
)

var (
	Cmd = &cobra.Command{
		Use:   "validate",
		Short: "Validate the certificate",
		Long: `Validate the certificate using the following checks:
1. Check the expiration date
2. Check if the certificate is trusted using the trust store from the host machine
3. Check if the certificate is not part of a revocation list
`,
		Run: runValidate,
	}

	rawUrl string
)

func runValidate(cmd *cobra.Command, args []string) {
	var u *url.URL
	if host.IsValidHostname(rawUrl) {
		u = &url.URL{
			Host: rawUrl,
		}
	} else {
		var parseErr error
		u, parseErr = url.ParseRequestURI(rawUrl)
		if parseErr != nil {
			fmt.Printf("Failed to parse URL %s\nError: %s", rawUrl, parseErr)
			return
		}
	}

	certs, fetchErr := cert.GetCertificateChain(u)
	if fetchErr != nil {
		fmt.Println("Failed to fetch certificate chain", fetchErr)
		return
	}

	if len(certs) <= 0 {
		fmt.Println("No certificates in the chain")
		return
	}

	isValid, validationErr := cert.ValidateCertificate(u.Host, certs[0])
	if validationErr != nil {
		fmt.Printf("Server certificate validation failed. Reason: %s", validationErr)
		return
	}
	if isValid {
		fmt.Printf("Certificate for host %s is valid", u.Host)
	} else {
		fmt.Println("Server certificate validation failed. Reason: none")
	}
}

func init() {
	includeValidateFlags(Cmd)
}

func includeValidateFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&rawUrl, "url", "u", "",
		"URL or hostname for which we would want to grab the certificate chain.")

	if err := cmd.MarkFlagRequired("url"); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
