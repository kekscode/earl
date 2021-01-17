package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "run as a webserver",
	}
)
