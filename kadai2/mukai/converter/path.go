package converter

import (
	"io/ioutil"
	"path/filepath"
)

type pather interface {
	files(dir string) ([]converterFileInterface, error)
}

type Path struct {
}

func (path Path) files(dir string) ([]converterFileInterface, error) {
	infos, e := ioutil.ReadDir(dir)
	if e != nil {
		return nil, e
	}
	var paths []converterFileInterface
	for _, v := range infos {
		file := converterFile{absPath: filepath.Join(dir, v.Name()), isDir: v.IsDir()}
		paths = append(paths, file)
	}
	return paths, nil
}
