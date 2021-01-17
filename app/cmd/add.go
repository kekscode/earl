package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add an URL",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}
)
