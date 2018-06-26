package converter

import (
	"io/ioutil"
	"path/filepath"
)

type Pather interface {
	files(dir string) ([]Converter, error)
}

type Path struct {
}

func (path Path) files(dir string) ([]Converter, error) {
	infos, e := ioutil.ReadDir(dir)
	if e != nil {
		return nil, e
	}
	var paths []Converter
	for _, v := range infos {
		file := convertFile{absPath: filepath.Join(dir, v.Name()), isDir: v.IsDir()}
		paths = append(paths, file)
	}
	return paths, nil
}
