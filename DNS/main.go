package main

import (
	"fmt"
	"net"
)

func main() {
	hostname := "google.com"

	v, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Println("Error", err)
		return

	}

	fmt.Printf("IP address of %s:\n,", hostname)
	for _, sa := range v {
		fmt.Println(sa)
	}
}

// Some commonly used functions include:

// net.LookupHost: Looks up the IP addresses associated with a given hostname.

// net.LookupCNAME: Looks up the canonical name (CNAME) for a given hostname.

// net.LookupMX: Looks up the mail exchange (MX) records for a domain.

// net.LookupTXT: Looks up the text (TXT) records for a domain.

// net.LookupSRV: Looks up service (SRV) records for a domain.
