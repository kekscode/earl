package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kekscode/earl/book"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "run as a webserver",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := cmd.Flags().GetInt("port")
			if err != nil {
				fmt.Errorf("Error: %v", err)
			}
			serve(p)
		},
	}
)

func serve(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := book.New()
		b.ReadFromJSON()
		fmt.Fprintf(w, b.ListMarksHTML())
	})

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
