package main

import (
	"fmt"
	"os"

	greip "github.com/greipio/go/v2"
)

func main() {
	token := os.Getenv("GREIP_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "GREIP_TOKEN required")
		os.Exit(1)
	}
	client := greip.NewGreip(token)

	lookup, err := client.Lookup("1.1.1.1", nil)
	if err != nil {
		fail("Lookup", err)
	}
	if lookup == nil || lookup.IP.Address != "1.1.1.1" {
		fail("Lookup", fmt.Errorf("unexpected lookup response"))
	}

	if _, err := client.Threats("1.1.1.1"); err != nil {
		fail("Threats", err)
	}

	ips := []string{"1.1.1.1", "8.8.8.8"}
	if _, err := client.BulkLookup(ips, nil); err != nil {
		fail("BulkLookup", err)
	}

	if _, err := client.Country("PS", nil); err != nil {
		fail("Country", err)
	}

	if _, err := client.AsnLookup("AS13335"); err != nil {
		fail("ASN", err)
	}

	if _, err := client.Email("name@domain.com"); err != nil {
		fail("Email", err)
	}

	if _, err := client.Phone("1234567890", "TR"); err != nil {
		fail("Phone", err)
	}

	if _, err := client.IBAN("DE89370400440532013000"); err != nil {
		fail("IBAN", err)
	}

	if _, err := client.Profanity("This is a test message"); err != nil {
		fail("Profanity", err)
	}

	if _, err := client.Payment(map[string]interface{}{
		"action":               "purchase",
		"website_domain":       "greip.io",
		"transaction_amount":   1000,
		"transaction_currency": "GBP",
		"customer_email":       "name@domain.com",
		"customer_ip":          "1.1.1.1",
	}); err != nil {
		fail("Payment", err)
	}

	fmt.Println("All go live checks passed.")
}

func fail(name string, err error) {
	fmt.Fprintf(os.Stderr, "FAIL %s: %v\n", name, err)
	os.Exit(1)
}
