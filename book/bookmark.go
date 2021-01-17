package book

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

// New returns a new Book containing a first bookmark
func New() Book {
	b := Book{}
	return b
}

func (b *Book) ListMarks() {
	fmt.Printf("%q", b.Marks)
}
func (b *Book) GetMark() {}
func (b *Book) AddMark(addr string, title string, tags []string, comment string) {
	u, err := url.Parse(addr)
	if err != nil {
		log.Fatalf("%q", err)
	}
	m := Mark{
		URL:     u,
		Title:   title,
		Tags:    tags,
		Comment: comment,
		Added:   time.Now(),
	}
	b.Marks = append(b.Marks, m)
}
func (b *Book) ModifyMark() {}
func (b *Book) DeleteMark() {}
