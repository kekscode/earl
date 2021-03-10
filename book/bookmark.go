package book

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"sync"
	"time"
)

// Book represents a container for marks
type Book struct {
	sync.Mutex
	Marks []Mark
}

// Mark represents a persistent item
type Mark struct {
	ID      string
	URL     url.URL
	Tags    []string
	Comment string
	Added   time.Time
}

// New returns a new Book containing a first bookmark
func New() Book {
	b := Book{}
	return b
}

func (b *Book) SaveToJSON() error {
	b.Lock()
	defer b.Unlock()
	file, _ := json.MarshalIndent(b, "", " ")
	err := ioutil.WriteFile("earl.json", file, 0644)
	if err != nil {
		err := errors.New(fmt.Sprintf("Could not write to disk %v", file))
		return err
	}
	return nil
}

func (b *Book) ReadFromJSON() error {
	file, err := ioutil.ReadFile("earl.json")
	if err != nil {
		err := errors.New(fmt.Sprintf("Could not read from disk %v", file))
		return err
	}
	err = json.Unmarshal(file, b)
	if err != nil {
		err := errors.New(fmt.Sprintf("Could not unmarshal json from file %v", file))
		return err
	}

	return nil
}

func (b *Book) ListMarks() []Mark {
	return b.Marks
}

func (b *Book) GetMark() {}
func (b *Book) AddMark(addr string, tags []string, comment string) {
	u, err := url.Parse(addr)
	// TODO: Return this error
	if err != nil {
		log.Fatalf("%q", err)
	}
	m := Mark{
		ID:      u.String(),
		URL:     *u,
		Tags:    tags,
		Comment: comment,
		Added:   time.Now(),
	}
	b.Marks = append(b.Marks, m)
}
func (b *Book) UpdateMark() {}
func (b *Book) DeleteMark() {}
