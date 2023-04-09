package cmd

import (
	"cert-ripper-go/pkg"
	"github.com/spf13/cobra"
	"log"
	"net/url"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the certificates from the chain to the standard output",
	Long:  ``,
	Run:   runPrint,
}

var printRawUrl string

func runPrint(cmd *cobra.Command, args []string) {
	var u *url.URL
	if pkg.IsValidHostname(printRawUrl) {
		u = &url.URL{
			Host: printRawUrl,
		}
	} else {
		var parseErr error
		u, parseErr = url.ParseRequestURI(printRawUrl)
		if parseErr != nil {
			log.Printf("Failed to parse URL %s\nError: %s", exportRawUrl, parseErr)
		}
	}

	certs, fetchErr := pkg.GetCertificateChain(u)
	if fetchErr != nil {
		log.Println("Failed to fetch certificate chain", fetchErr)
	}

	if ioErr := pkg.PrintCertificates(u.Host, certs); ioErr != nil {
		log.Println("Failed to print certificate to the standard output", ioErr)
	}
}

func init() {
	includePrintFlags(printCmd)
}

func includePrintFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&exportRawUrl, "url", "u", "www.example.com",
		"URL or hostname for which we would want to grab the certificate chain.")
}
