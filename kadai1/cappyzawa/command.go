package conv

import (
	"fmt"
	"os"
	"strings"
)

// Command - interface defining a method for command
type Command interface {
	Run(dir string, from string, to string) error
}

type command struct {
	iFilePath IFilePath
	decoder   Converter
	encoder   Converter
}

// NewCommand - initialize Command
func NewCommand(iFilePath IFilePath, decoder, encoder Converter) Command {
	return &command{
		iFilePath: iFilePath,
		decoder:   decoder,
		encoder:   encoder,
	}
}

// Run - execute command
func (c *command) Run(dir, from, to string) error {
	if _, err := os.Stat(dir); err != nil {
		return err
	}
	fExt := fmt.Sprintf(".%s", from)
	err := c.iFilePath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if c.iFilePath.Ext(path) == fExt {
			return c.convert(path, to)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
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
	baseDir := c.iFilePath.Dir(path)
	baseFile := c.iFilePath.Base(path)
	baseExt := c.iFilePath.Ext(path)
	newFileName := strings.Replace(baseFile, baseExt, tExt, 1)
	newFilePath := c.iFilePath.Join(baseDir, newFileName)
	return os.Create(newFilePath)
}
