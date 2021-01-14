package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list URLs",
	}
)
