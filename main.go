package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,hasMX,sprRecord,hasDMARC,dmarcRecord")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input : %v \n ", err)
	}
}
func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	maxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("error : %v \n", err)
	}
	if len(maxRecord) > 0 {
		hasMX = true
	}
	textRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("error : %v \n", err)
	}
	for _, record := range textRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasMX = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("error : %v \n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v %v %v %v %v %v ", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
