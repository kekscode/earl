package cmd

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/kekscode/earl/book"
	"github.com/spf13/cobra"
)

var (

	// templatesFs holds our dynamic templates
	//go:embed templates/*.html
	templatesFs embed.FS

	// staticFs holds our static web server content
	//go:embed css/*.css
	//go:embed favicon/*.ico
	staticFs embed.FS

	// Parse Web UI template
	tmpl = template.Must(template.ParseFS(templatesFs, "templates/index.html"))

	// Used for flags.
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "run as a webserver",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := cmd.Flags().GetInt("port")
			if err != nil {
				fmt.Errorf("Error: %v", err)
			}
			serve(p, tmpl, staticFs) // Start server
		},
	}
)

// serve is the main function for earls integrated webserver
func serve(port int, tmpl *template.Template, staticFs embed.FS) {

	// Our Muxer
	r := mux.NewRouter()

	// http.FS can be used to create a http Filesystem
	statics := http.FS(staticFs)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(statics)))

	// Handle index UI template
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		b := book.New()
		b.ReadFromJSON()

		tmpl.ExecuteTemplate(w, "index.html", b.ListMarks())
	})

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server at %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
