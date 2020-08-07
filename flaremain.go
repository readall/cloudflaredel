package main

import (
	"fmt"
	"strings"
	"os"
	// "time"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

const (
	user   = "llearnell@gmail.com"
	domain = "safenetapp.com"
	apiKey = "e53145712923cb7cdd3284d14a7b981f926ce"
)

func main() {
	// argsToCmd := os.Args
	if (len(os.Args) < 2){
		fmt.Printf("Provide the substring that you want to bulk delete")
		fmt.Printf("Usage :\n %s <%s> \n", os.Args[0], "auto-")
		os.Exit(0)
	}
	suStr := os.Args[1]

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

	found := false
	for _, r := range records {
		// time.Sleep(1)
		// fmt.Printf("%s: %s\n", r.Name, r.Content)
		if strings.Contains(r.Name, suStr) {
			fmt.Printf("Deleting %s\n", r.Name)
			// time.Sleep(1)
			found = true
			err = api.DeleteDNSRecord(zoneID, r.ID)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error deleting DNS record: ", err)
			return
			}

		}
	}
	if found == false {
		fmt.Printf("  Did not find any dns record with pattern %s \n", suStr)
	}
}
