package main

import (
	"github.com/kekscode/earl/app/cmd"
	"github.com/kekscode/earl/book"
)

func main() {
	cmd.Execute()

	book := book.New()
	book.AddMark("https://user:pass@google.de/test?bla=blubb%20#chapter!one", "", []string{}, "")
	book.ListMarks()
	book.SaveToJSON()
}
