package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	rootCmd = &cobra.Command{
		Use:   "earl",
		Short: "Earl manages URLs",
		Long: `Earl is used to add, manage and retrieve URLs.
It may be used on the command line and as a server.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// earl add <url> [<title>] [<tags>] [<comment>]
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("url", "u", "", "https://myurl... ending with a whitespace")
	addCmd.Flags().StringSliceP("tags", "t", []string{}, "Comma-separated list")
	addCmd.Flags().StringP("comment", "c", "", "\"A comment in quotes\"")

	// earl show "<url> [<title>] [<tags>] [<comment>]"
	rootCmd.AddCommand(showCmd)

	// earl list <url> [<title>] [<tags>] [<comment>]
	rootCmd.AddCommand(listCmd)

	// earl edit <url> [<title>] [<tags>] [<comment>]
	rootCmd.AddCommand(editCmd)

	// earl del,delete,remove,rm <url>
	rootCmd.AddCommand(deleteCmd)

	// earl serve
	rootCmd.AddCommand(serveCmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
