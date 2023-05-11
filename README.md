# cert-ripper-go

Fetch certificate chain for a hostname or URL.

## Example of Usage

```bash
$ cert-ripper -h
Retrieve the certificate chain for a URL or a hostname.

Usage:
  cert-ripper [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  export      Export the certificates from the chain and save them into a folder
  help        Help about any command
  print       Print the certificates from the chain to the standard output
  validate    Validate the certificate

Flags:
  -h, --help      help for cert-ripper
  -v, --version   version for cert-ripper

Use "cert-ripper [command] --help" for more information about a command.
```

### `print` command:

Displays a certificate on the standard output in OpenSSL format.

Example of usage:

```bash
cert-ripper print --url=stackoverflow.com
```

Expected output:

```
Found 4 certificates in the certificate chain for stackoverflow.com 
===========================
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 304654416220459475292011240363286494111390 (0x37f4c8249f4024f228c33cba3198752e69e)
    Signature Algorithm: SHA256-RSA
        Issuer: C=US,O=Let's Encrypt,CN=R3
        Validity
...
```

### `export` command:

Saves the whole certificate chain in a folder. The certificates from the chain can be saved in different formats.

Example of usage:

- With shorthands:
```bash
cert-ripper export -u ervinszilagyi.dev -p certs -f pem
```

- With long commands
```bash
cert-ripper export --url=ervinszilagyi.dev --path=certs --format=txt
```

### `validate` command:

Validates the server certificate using the following steps:

1. Check the expiration date
2. Check if the certificate is trusted using the trust store from the host machine
3. Check if the certificate is not part of a revocation list

Example of usage:

```bash
cert-ripper validate -u ervinszilagyi.dev
```

### `request` command:

`request` command can be used to work with Certificate Signing Requests (CSR). It has the following subcommands:

```bash
Create and decode CSRs (Certificate Signing Request)

Usage:
  cert-ripper request [command]

Available Commands:
  create      Create a CSR (certificate signing request)
  decode      Decode and print CSR file to the STDOUT in OpenSSL text format
```

- `create` command - used to create a CSR request with a private key

```bash
Create a CSR (certificate signing request)

Usage:
  cert-ripper request create [flags]

Flags:
      --city string                 Locality/City (example: New-York)
      --commonName string           Common name (example: domain.com).
      --country string              Country code (example: US).
      --email string                Email address
  -h, --help                        help for create
      --organization string         Organization (example: Acme)
      --organizationUnit string     Organization unit (example: IT)
      --signatureAlg signatureAlg   Signature Algorithm (allowed values: SHA256WithRSA, SHA384WithRSA, SHA512WithRSA,SHA256WithECDSA, SHA384WithECDSA, SHA512WithECDSA) (default SHA256WithRSA)
      --state string                Province/State (example: California)
      --targetPath string           Target path for the CSR to be saved. (default ".")

```

Example:

```bash
cert-ripper request create --commonName esz.dev --country RO --state Mures --city Targu-Mures --organization ACME --organizationUnit IT --targetPath certs/req.csr --email esz@dev.com --signatureAlg ED25519
```

- `decode` command - used to decode a CSR and display it in OpenSSL format 

```bash
Decode and print CSR file to the STDOUT in OpenSSL text format

Usage:
  cert-ripper request decode [flags]

Flags:
  -h, --help          help for decode
      --path string   Path for of the CSR file.
```

Example:

```bash
cert-ripper request decode --path="certs/request.csr"
```

### `generate` command

```bash
cert-ripper generate fromstdio --host=example.com --validFrom="2023-05-09 15:04:05" --validFor=3600 --targetPath=certs --isCa
```

## Download and Install

### MacOS

Install with homebrew:

```bash
brew tap recon-tools/homebrew-recon-tools
brew install cert-ripper-go
```

### Debian/Ubuntu

ppa coming, for now download the executable from the [release](https://github.com/recon-tools/cert-ripper-go/releases) page

### Windows/Other

Download the executable from the releases page: https://github.com/recon-tools/cert-ripper-go/releases

## Building

Go 1.19 is required.

```bash
cd cert-ripper-go
go build -o target/cert-ripper
```

### Build with ldflags

```bash
go build -ldflags "-X 'cert-ripper-go/cmd.appVersion=0.0.1'" -o target/cert-ripper
```