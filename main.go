package main

import (
	"fmt"
	"log"
	"net/url"
	"time"
)

// Book represents a container for marks
type Book struct {
	Marks []Mark
}

// Mark represents a persistent item
type Mark struct {
	URL     *url.URL
	Title   string
	Tags    []string
	Comment string
	Added   time.Time
}

func (b *Book) ListMarks() {
	fmt.Printf("%q", b.Marks)
}
func (b *Book) GetMark() {}
func (b *Book) AddMark(m Mark) {
	b.Marks = append(b.Marks, m)
}
func (b *Book) ModifyMark() {}
func (b *Book) DeleteMark() {}

func main() {
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
