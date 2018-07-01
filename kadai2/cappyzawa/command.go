package conv

import (
	"fmt"
	"os"
	"path/filepath"
)

type Command struct {
	decoder Decoder
	encoder Encoder
}

func NewCommand(decoder Decoder, encoder Encoder) *Command {
	return &Command{
		decoder: decoder,
		encoder: encoder,
	}
}

// Run - execute Command
func (c *Command) Run(dir, from, to string) ([]string, error) {
	fExt := fmt.Sprintf(".%s", from)
	var createdFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == fExt {
			newFile, err := c.convert(path, to)
			if err != nil {
				return err
			}
			createdFiles = append(createdFiles, newFile.Name())
			return nil
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return createdFiles, nil
}

func (c *Command) convert(path, to string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := c.decoder.Decode(file)
	if err != nil {
		return nil, err
	}

	newFile, err := c.createOutputFile(path, to)
	if err != nil {
		return nil, err
	}
	defer newFile.Close()
	if err := c.encoder.Encode(newFile, img); err != nil {
		return nil, err
	}

	return newFile, nil
}

func (c *Command) createOutputFile(path, to string) (*os.File, error) {
	tExt := fmt.Sprintf(".%s", to)
	baseExt := filepath.Ext(path)
	newFile := path[:len(path)-len(baseExt)] + tExt
	return os.Create(newFile)
}
