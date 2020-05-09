package filesid

import (
	"fmt"
	"log"

	"github.com/hsmtkk/virustotal_go/pkg/apikey"
	"github.com/hsmtkk/virustotal_go/pkg/filesid"
	"github.com/spf13/cobra"
)

var FilesIDCommand = &cobra.Command{
	Use:  "filesid",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hash := args[0]
		apiKey, err := apikey.New().LoadAPIKey()
		if err != nil {
			log.Fatal(err)
		}
		rtr := filesid.New(apiKey)
		stats, err := rtr.RetrieveInformation(hash)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("harmless: %d\n", stats.HarmLess)
		fmt.Printf("undetected: %d\n", stats.Undetected)
		fmt.Printf("suspicious: %d\n", stats.Suspicious)
		fmt.Printf("malicious: %d\n", stats.Malicious)
	},
}
