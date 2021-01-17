package main

import (
	"log"
	"net/url"
	"time"

	"github.com/kekscode/earl/app/cmd"
)

func main() {
	cmd.Execute()

	b := Book{}
	u, err := url.Parse("https://google.de")
	if err != nil {
		log.Fatalf("%q", err)
	}
	m := Mark{
		URL:     u,
		Title:   "Google Web Search",
		Tags:    []string{"a", "ab", "abc", "blubber lutsch"},
		Comment: "Blubber ist bla",
		Added:   time.Now()}

	b.AddMark(m)
	b.ListMarks()
}
