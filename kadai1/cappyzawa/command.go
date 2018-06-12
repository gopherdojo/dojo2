package conv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Command - interface defining a method for command
type Command interface {
	Run(dir, from, to string) ([]string, error)
}

type command struct {
	decoder Converter
	encoder Converter
}

// NewCommand - initialize Command
func NewCommand(decoder, encoder Converter) Command {
	return &command{
		decoder: decoder,
		encoder: encoder,
	}
}

// Run - execute command
func (c *command) Run(dir, from, to string) ([]string, error) {
	if _, err := os.Stat(dir); err != nil {
		return nil, err
	}
	fExt := fmt.Sprintf(".%s", from)
	var createdFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == fExt {
			if err := c.convert(path, to); err !=nil {
				return err
			}
			createdFiles = append(createdFiles, path)
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return createdFiles, nil
}

func (c *command) convert(path, to string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	img, err := c.decoder.Decode(file)
	if err != nil {
		return err
	}

	newFile, err := c.createOutputFile(path, to)
	if err != nil {
		return err
	}
	defer newFile.Close()
	if err := c.encoder.Encode(newFile, img); err != nil {
		return err
	}

	return nil
}

func (c *command) createOutputFile(path, to string) (*os.File, error) {
	tExt := fmt.Sprintf(".%s", to)
	baseDir := filepath.Dir(path)
	baseFile := filepath.Base(path)
	baseExt := filepath.Ext(path)
	newFileName := strings.Replace(baseFile, baseExt, tExt, 1)
	newFilePath := filepath.Join(baseDir, newFileName)
	return os.Create(newFilePath)
}
