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
	//fs holds our static web server content
	//go:embed web/templates/*.html
	//go:embed web/static/favicon/*.ico
	//go:embed web/static/css/*.css
	fs embed.FS

	// Parse Web UI template
	tmpl = template.Must(template.ParseFS(fs, "web/templates/index.html"))

	// Used for flags.
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "run as a webserver",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := cmd.Flags().GetInt("port")
			if err != nil {
				fmt.Errorf("Error: %v", err)
			}
			serve(p, tmpl, fs) // Start server
		},
	}
)

func serve(port int, tmpl *template.Template, fs embed.FS) {

	// http.FS can be used to create a http Filesystem
	var staticFS = http.FS(fs)
	fsStatics := http.FileServer(staticFS)

	http.Handle("web/static/css/", fsStatics)

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		ico, err := fs.ReadFile("web/static/favicon/favicon.ico")
		if err != nil {
			fmt.Errorf("Error: %v", err)
		}
		w.Write(ico)
	})

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
