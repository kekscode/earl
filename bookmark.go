package main

import (
	"fmt"
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
