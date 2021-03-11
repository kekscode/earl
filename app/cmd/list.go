package cmd

import (
	"fmt"

	"github.com/kekscode/earl/book"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list URLs",
		Run: func(cmd *cobra.Command, args []string) {
			b := book.New()
			b.ReadFromJSON()
			fmt.Printf("%v", b.ListMarks())
		},
	}
)
