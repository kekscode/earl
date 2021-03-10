package cmd

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
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

		//tmpl, err := template.New("index.html").ParseFS(fs, "templates/index.html")
		//if err != nil {
		//	fmt.Errorf("Error: %v", err)
		//}

		w.Header().Add("Content-Type", "text/html")

		// Debug
		fmt.Printf("%v", tmpl)
		tmpl.Execute(os.Stdout, nil)
		tmpl.ExecuteTemplate(os.Stdout, "index.html", nil)
		tmpl.ExecuteTemplate(os.Stdout, "index.html", struct{ test string }{test: "stefan"})
		tmpl.ExecuteTemplate(os.Stdout, "index.html", b.ListMarks())
		// End Debug

		tmpl.ExecuteTemplate(w, "index.html", nil)
		//	Marks: b.ListMarks(),
		//})
		//fmt.Printf("%v", b.ListMarks())
		//err = tmpl.Execute(w, struct{}{})
		//if err != nil {
		//fmt.Errorf("Error: %v", err)
		//}

	})

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
