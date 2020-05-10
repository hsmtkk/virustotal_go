package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/hsmtkk/virustotal_go/pkg/apikey"
	"github.com/hsmtkk/virustotal_go/pkg/files"
	"github.com/spf13/cobra"
)

var FilesCommand = &cobra.Command{
	Use:  "files",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		fileName := filepath.Base(path)
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		apiKey, err := apikey.New().LoadAPIKey()
		if err != nil {
			log.Fatal(err)
		}
		uploader := files.New(apiKey)
		id, err := uploader.UploadAnalyze(fileName, contents)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("analyze ID: %s\n", id)
	},
}
