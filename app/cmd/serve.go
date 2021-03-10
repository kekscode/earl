package cmd

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/kekscode/earl/book"
	"github.com/spf13/cobra"
)

var (
	//go:embed templates/*
	fs   embed.FS
	tmpl = template.Must(template.ParseFS(fs, "templates/index.html"))

	// Used for flags.
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "run as a webserver",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := cmd.Flags().GetInt("port")
			if err != nil {
				fmt.Errorf("Error: %v", err)
			}
			serve(p, tmpl) // Start server
		},
	}
)

func serve(port int, tmpl *template.Template) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		b := book.New()
		b.ReadFromJSON()

		tmpl.ExecuteTemplate(w, "index.html", b.ListMarks())
	})

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
