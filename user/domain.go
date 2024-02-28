package user

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func DomainCheck(domain string) error {
	log.Println("Domain:", domain)

	splitDomain := strings.Split(domain, ".")
	domainToCheck := "_mirror-user-verification." + strings.Join(splitDomain[len(splitDomain)-2:], ".")
	log.Println("2nd Level Domain:", domainToCheck)

	txt, err := net.LookupTXT(domainToCheck)

	if err.(*net.DNSError).IsNotFound {
		return fmt.Errorf("We couldn't find the Mirror verification record for " + domainToCheck)
	}
	if err != nil {
		return fmt.Errorf("Error getting your DNS Text Entries: " + err.Error())
	}

	log.Println("Text Records for", domainToCheck, "are:", strings.Join(txt, ", "))
	return nil
}
