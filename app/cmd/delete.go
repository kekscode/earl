package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete an URL",
	}
)
