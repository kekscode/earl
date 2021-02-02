package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kekscode/earl/app/cmd"
	"github.com/kekscode/earl/book"
)

func main() {

	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		s := <-sig
		log.Fatalf("Signal (%v) received, stopping\n", s)
	}()

	book := book.New()
	book.ReadFromJSON()
	//book.AddMark("https://user:pass@google.de/test?bla=blubb%20#chapter!one", []string{}, "")
	//fmt.Print(book.ListMarks())
	book.SaveToJSON()

	cmd.Execute()

}
