package sget

import (
	"fmt"
	"net/http"
	"os"
)

// Sget structs
type Sget struct {
	URL string
}

// New for sget package
func New() *Sget {
	return &Sget{}
}

// Run execute methods in sget package
func (sget *Sget) Run() error {
	if err := sget.Prepare(os.Args[1:]); err != nil {
		return err
	}

	if err := sget.RequestHeader(); err != nil {
		return err
	}

	return nil
}

// Prepare set up necessary parameters
// set sget.URL
func (sget *Sget) Prepare(argv []string) error {
	if len(argv) != 1 {
		return fmt.Errorf("failed because of too or few arguments, given %s", argv)
	}
	sget.URL = argv[0]
	return nil
}

// RequestHeader request head
func (sget *Sget) RequestHeader() error {
	head, err := http.Head(sget.URL)
	if err != nil {
		return fmt.Errorf("failed to request head %s", err)
	}
	fmt.Println(head)
	return nil
}
