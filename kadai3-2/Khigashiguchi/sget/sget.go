package sget

import (
	"fmt"
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

	fmt.Println(sget.URL)

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
