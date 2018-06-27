package converter

import (
	"fmt"
	"reflect"
	"testing"
)

type testPath struct {
}

func (t testPath) files(dir string) ([]converterFileInterface, error) {
	if dir == "images" {
		var files []converterFileInterface
		files = append(files, testConvertFile{absPath: "images/file1.jpg", isDir: false})
		files = append(files, testConvertFile{absPath: "images/file2.jpg", isDir: false})
		files = append(files, testConvertFile{absPath: "images/file1.png", isDir: false})
		files = append(files, testConvertFile{absPath: "images/dir1", isDir: true})
		return files, nil
	} else if dir == "images/dir1" {
		var files []converterFileInterface
		files = append(files, testConvertFile{absPath: "images/dir1/file3.jpg", isDir: false})
		files = append(files, testConvertFile{absPath: "images/dir1/file4.jpg", isDir: false})
		files = append(files, testConvertFile{absPath: "images/dir1/dir2", isDir: true})
		return files, nil
	} else if dir == "images/dir1/dir2" {
		var files []converterFileInterface
		files = append(files, testConvertFile{absPath: "images/dir1/dir2/file5.jpg", isDir: false})
		files = append(files, testConvertFile{absPath: "images/dir1/dir2/file1.gif", isDir: false})
		return files, nil
	}
	return nil, fmt.Errorf("no such directory")
}

type testConvertFile struct {
	absPath string
	isDir   bool
}

func (f testConvertFile) convert(outputFormat string) (string, error) {
	path := arbitraryExtAbsPath(f.absPath, outputFormat)
	return path, nil
}

func (f testConvertFile) isDirectory() bool {
	return f.isDir
}

func (f testConvertFile) absolutePath() string {
	return f.absPath
}

func TestRecursiveConvert(t *testing.T) {
	type args struct {
		dir          string
		inputFormat  string
		outputFormat string
		pather       pather
	}
	expected := []string {
		"images/file1.png",
		"images/file2.png",
		"images/dir1/file3.png",
		"images/dir1/file4.png",
		"images/dir1/dir2/file5.png",
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name:"", args:args{inputFormat:"jpg", outputFormat:"png", dir:"images", pather: testPath{}}, want: expected, wantErr:false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RecursiveConvert(tt.args.dir, tt.args.inputFormat, tt.args.outputFormat, tt.args.pather)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecursiveConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecursiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}
