package main

import (
	"log"

	"github.com/hsmtkk/virustotal_go/cmd/virustotal/files"
	"github.com/hsmtkk/virustotal_go/cmd/virustotal/filesid"
	"github.com/spf13/cobra"
)

func main() {
	rootCommand := &cobra.Command{
		Use: "virustotal",
	}
	rootCommand.AddCommand(files.FilesCommand)
	rootCommand.AddCommand(filesid.FilesIDCommand)
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
