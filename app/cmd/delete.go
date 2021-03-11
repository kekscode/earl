package cmd

import (
	"fmt"

	"github.com/kekscode/earl/book"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete an URL",
		Run: func(cmd *cobra.Command, args []string) {
			b := book.New()
			b.ReadFromJSON()
			fmt.Printf("%v", "Not implemented yet")
		},
	}
)
