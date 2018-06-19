package images

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Conversion represents an image conversion between given formats.
type Conversion struct {
	Decoder        Decoder
	Encoder        Encoder
	DestinationExt string
}

// ReplaceExt returns filename replaced the extension with DestinationExt.
// For example, if `DestinationExt` is `png`, `ReplaceExt("hello.jpg")` will return `"hello.png"`.
func (c *Conversion) ReplaceExt(filename string) string {
	tail := filepath.Ext(filename)
	head := strings.TrimSuffix(filename, tail)
	return fmt.Sprintf("%s.%s", head, c.DestinationExt)
}

// Do converts the source file to destination.
// `source` and `destination` must be file path.
func (c *Conversion) Do(source string, destination string) error {
	r, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("Error while opening source file %s: %s", source, err)
	}
	defer r.Close()
	m, err := c.Decoder.Decode(r)
	if err != nil {
		return fmt.Errorf("Error while decoding file %s: %s", source, err)
	}
	w, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("Error while opening destination file %s: %s", destination, err)
	}
	defer w.Close()
	if err := c.Encoder.Encode(w, m); err != nil {
		return fmt.Errorf("Error while encoding to file %s: %s", destination, err)
	}
	return nil
}
