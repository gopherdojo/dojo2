package converter

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Convert(beforeExtension string, afterExtension string, targetDir string) {
	targetPathList := []string{}

	err := filepath.WalkDir(targetDir, func(path string, info fs.DirEntry, err error)error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		targetPathList = append(targetPathList, path)

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	for _, targetPath := range targetPathList {
		presentExtension := filepath.Ext(targetPath)

		if presentExtension != "."+beforeExtension {
			return
		}
	
		dir, fileName := filepath.Split(targetPath)
	
		newFileName := strings.Replace(fileName, "."+beforeExtension, "."+afterExtension, 1)
	
		newPath := filepath.Join(dir, newFileName)
		if err := os.Rename(targetPath, newPath); err != nil {
			log.Fatal(err)
		}
	}
}
