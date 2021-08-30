package main

import (
	"fmt"
	"log"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	// whoisparser "github.com/icamys/whois-parser"
)

func main() {
	whoisInfo, err := whois.Whois("cmb888.app")
	if err != nil {
		log.Panicf("%v\n", err)
	}

	result, err := whoisparser.Parse(whoisInfo)
	if err == nil {
		// Print the domain status
		fmt.Println(result.Domain.Status)

		// Print the domain created date
		fmt.Println(result.Domain.CreatedDate)

		// Print the domain expiration date
		fmt.Println(result.Domain.ExpirationDate)

		// Print the registrar name
		fmt.Println(result.Registrar.Name)

		// Print the registrant name
		fmt.Println(result.Registrant.Name)

		// Print the registrant email address
		fmt.Println(result.Registrant.Email)
	}
}
