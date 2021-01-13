package main

import (
	"net/url"
	"time"
)

// Book represents a container for marks
type Book struct {
	Marks []Mark
}

// Mark represents a persistent item
type Mark struct {
	URL     url.URL
	Title   string
	Tags    []string
	Comment string
	Added   time.Time
}

func main() {}
