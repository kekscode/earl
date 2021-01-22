package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add an URL",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Flags().GetString("url"))
			fmt.Println(cmd.Flags().GetString("tags"))
			fmt.Println(cmd.Flags().GetString("comment"))
		},
	}
)
