package conv

import (
	"path/filepath"
)

// IFilePath - interface defining methods for path/filepath
type IFilePath interface {
	Walk(root string, walkFn filepath.WalkFunc) error
	Ext(path string) string
	Base(path string) string
	Dir(path string) string
	Join(elem ...string) string
}

type iFilePath struct {
}

// NewIFilePath - initialize IFilePath
func NewIFilePath() IFilePath {
	return &iFilePath{}
}

func (fp *iFilePath) Walk(root string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(root, walkFn)
}

func (fp *iFilePath) Ext(path string) string {
	return filepath.Ext(path)
}

func (fp *iFilePath) Base(path string) string {
	return filepath.Base(path)
}

func (fp *iFilePath) Dir(path string) string {
	return filepath.Dir(path)
}

func (fp *iFilePath) Join(elem ...string) string {
	return filepath.Join(elem...)
}
