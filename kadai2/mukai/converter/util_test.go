package converter

import "testing"

func TestExt(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "filename", args:args{path: "file.jpg"}, want: "jpg"},
		{name: "filename", args:args{path: "jpg"}, want: ""},
		{name: "filename", args:args{path: ""}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extension(tt.args.path); got != tt.want {
				t.Errorf("ext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSameExt(t *testing.T) {
	type args struct {
		path string
		ext  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{path: "dir/file.png", ext: "png"}, want: true},
		{name: "", args: args{path: "dir/file.png", ext: "gif"}, want: false},
		{name: "", args: args{path:"", ext:"png"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameExt(tt.args.path, tt.args.ext); got != tt.want {
				t.Errorf("isSameExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArbitraryExtAbsPath(t *testing.T) {
	type args struct {
		filePath string
		ext      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{filePath: "dir/file.jpg", ext: "png"}, want: "dir/file.png"},
		{name: "", args: args{filePath: "dir/", ext: "png"}, want: "dir/"},
		{name: "", args: args{filePath: "dir/file", ext: "png"}, want: "dir/file"},
		{name: "", args: args{filePath: "", ext: "png"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arbitraryExtAbsPath(tt.args.filePath, tt.args.ext); got != tt.want {
				t.Errorf("arbitraryExtAbsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
