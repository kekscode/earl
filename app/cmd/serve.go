package cmd

import (
	"embed"
	"fmt"
	"html/template"
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

	//go:embed templates/*
	f embed.FS
)

func serve(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := book.New()
		b.ReadFromJSON()

		tmpl, err := template.New("index").ParseFS(f, "templates/index.html")
		if err != nil {
			fmt.Errorf(err.Error())
		}

		tmpl.ExecuteTemplate(w, "index.html", struct{ Content string }{
			Content: "help",
		})
	})

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
