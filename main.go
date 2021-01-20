package main

import (
	"fmt"

	"github.com/kekscode/earl/app/cmd"
	"github.com/kekscode/earl/book"
)

func main() {

	book := book.New()
	book.ReadFromJSON()
	book.AddMark("https://user:pass@google.de/test?bla=blubb%20#chapter!one", "", []string{}, "")
	fmt.Print(book.ListMarks())
	book.SaveToJSON()

	cmd.Execute()
}
