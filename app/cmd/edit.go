package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	editCmd = &cobra.Command{
		Use:   "edit",
		Short: "edit an URL",
	}
)
