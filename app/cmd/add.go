package cmd

import (
	"fmt"

	"github.com/kekscode/earl/book"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add an URL",
		Run: func(cmd *cobra.Command, args []string) {
			book := book.New()
			book.ReadFromJSON()

			url, err := cmd.Flags().GetString("url")
			if err != nil {
				fmt.Errorf("Error: %v", err)
			}

			tags, err := cmd.Flags().GetStringSlice("tags")
			if err != nil {
				fmt.Errorf("Error: %v", err)
			}

			comment, err := cmd.Flags().GetString("comment")
			if err != nil {
				fmt.Errorf("Error: %v", err)
			}

			book.AddMark(url, tags, comment)

			book.SaveToJSON()
		},
	}
)
