package main

import (
	"github.com/kekscode/earl/app/cmd"
	"github.com/kekscode/earl/book"
)

func main() {
	cmd.Execute()

	book := book.New()
	book.AddMark("https://google.de", "", []string{}, "")
	book.ListMarks()
}
