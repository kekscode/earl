package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add an URL",
	}
)
