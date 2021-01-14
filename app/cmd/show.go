package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "show an URL",
	}
)
