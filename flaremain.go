package main

import (
	"fmt"
	"strings"
	"os"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

const (
	user   = "llearnell@gmail.com"
	domain = "safenetapp.com"
	apiKey = "e53145712923cb7cdd3284d14a7b981f926ce"
)

func main() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch the zone ID for zone example.org
	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch all DNS records for example.org
	records, err := api.DNSRecords(zoneID, cloudflare.DNSRecord{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range records {
		time.Sleep(1)
		// fmt.Printf("%s: %s\n", r.Name, r.Content)
		if strings.Contains(r.Name, "auto-") {
			fmt.Printf("Deleting %s\n", r.Name)
			time.Sleep(1)
			err = api.DeleteDNSRecord(zoneID, r.ID)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error deleting DNS record: ", err)
			return
			}

		}
	}
}
