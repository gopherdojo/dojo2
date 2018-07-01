package conversion

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo2/kadai2/Khigashiguchi/2-2/format"
)

// Conversion represents an image conversion
type Conversion struct {
	Decoder format.Decoder
	Encoder format.Encoder
	ToExt   string
}

// ReplaceExt returns filename replaced extension from filepath to ToExt
func (c *Conversion) ReplaceExt(filename string) string {
	tail := filepath.Ext(filename)
	head := strings.TrimSuffix(filename, tail)
	return fmt.Sprintf("%s.%s", head, c.ToExt)
}

// Do convert conversion
func (c *Conversion) Do(filepath, destination string) error {
	sf, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("Error happen while opening file %s: %s", filepath, err)
	}
	defer sf.Close()
	d, err := c.Decoder.Decode(sf)
	if err != nil {
		return fmt.Errorf("Error happen while decoding file %s: %s", filepath, err)
	}
	wf, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("Error happen while creating file %s: %s", destination, err)
	}
	defer wf.Close()
	if err := c.Encoder.Encode(wf, d); err != nil {
		return fmt.Errorf("Error happen while encoding to file %s: %s", destination, err)
	}
	return nil
}
