package converter

import (
	"testing"
	"fmt"
)

type testPath struct {
}

func (t testPath) files(dir string) ([]Converter, error) {
	if dir == "images" {
		var files []Converter
		files = append(files, testConvertFile{absPath: "images/file1.jpg", isDir: false})
		files = append(files, testConvertFile{absPath: "images/file2.jpg", isDir: false})
		files = append(files, testConvertFile{absPath: "images/dir1", isDir: true})
		return files, nil
	} else if dir == "images/dir1" {
		var files []Converter
		files = append(files, testConvertFile{absPath: "images/dir1/file3.jpg", isDir: false})
		files = append(files, testConvertFile{absPath: "images/dir1/file4.jpg", isDir: false})
		files = append(files, testConvertFile{absPath: "images/dir1/dir2", isDir: true})
		return files, nil
	} else if dir == "images/dir1/dir2" {
		var files []Converter
		files = append(files, testConvertFile{absPath: "images/dir1/file5.jpg", isDir: false})
		return files, nil
	}
	return nil, fmt.Errorf("no such directory")
}

type testConvertFile struct {
	absPath string
	isDir bool
}

func (f testConvertFile) Convert(outputFormat string) error {
	fmt.Println("convert " + f.absPath + " to " + outputFormat)
	return nil
}

func (f testConvertFile) IsDir() bool {
	return f.isDir
}

func (f testConvertFile) AbsPath() string {
	return f.absPath
}

func TestRecursiveConvert(t *testing.T) {
	type args struct {
		dir          string
		inputFormat  string
		outputFormat string
		pather       Pather
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name:"", args:args{inputFormat:"jpg", outputFormat:"png", dir:"images", pather: testPath{}}, wantErr:false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RecursiveConvert(tt.args.dir, tt.args.inputFormat, tt.args.outputFormat, tt.args.pather); (err != nil) != tt.wantErr {
				t.Errorf("RecursiveConvert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
