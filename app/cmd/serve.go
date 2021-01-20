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
			serve()
		},
	}
)

func serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := book.New()
		b.ReadFromJSON()
		fmt.Fprintf(w, b.ListMarksHTML())
	})

	fmt.Print("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
