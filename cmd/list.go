package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"igen/api"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available languages",
	Long:  "List all available languages that can be used with igen",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			languages []string
			err       error
		)
		languages, err = api.ListAvailableLanguages()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}

		for _, language := range languages {
			fmt.Println(language)
		}
	},
}
