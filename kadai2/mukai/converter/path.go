package converter

import (
	"io/ioutil"
	"path/filepath"
)

type Pather interface {
	files(dir string) ([]convertFile, error)
}

type Path struct {
}

func (path Path) files(dir string) ([]convertFile, error) {
	infos, e := ioutil.ReadDir(dir)
	if e != nil {
		return nil, e
	}
	var paths []convertFile
	for _, v := range infos {
		file := convertFile{absPath: filepath.Join(dir, v.Name()), isDir: v.IsDir()}
		paths = append(paths, file)
	}
	return paths, nil
}
