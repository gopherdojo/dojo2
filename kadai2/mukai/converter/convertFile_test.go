package converter

import "testing"

func Test_convertFile_ext(t *testing.T) {
	type fields struct {
		absPath string
		isDir   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "filename", fields: fields{absPath: "file.jpg", isDir: true}, want: ""},
		{name: "filename", fields: fields{absPath: "file.jpg", isDir: false}, want: "jpg"},
		{name: "filename", fields: fields{absPath: "jpg", isDir: false}, want: ""},
		{name: "filename", fields: fields{absPath: "", isDir: false}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := convertFile{
				absPath: tt.fields.absPath,
				isDir:   tt.fields.isDir,
			}
			if got := f.ext(); got != tt.want {
				t.Errorf("convertFile.ext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertFile_arbitraryExtAbsPath(t *testing.T) {
	type fields struct {
		absPath string
		isDir   bool
	}
	type args struct {
		ext string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{name: "", fields: fields{absPath: "dir/file.jpg", isDir: false}, args: args{ext: "png"}, want: "dir/file.png"},
		{name: "", fields: fields{absPath: "dir/file.jpg", isDir: true}, args: args{ext: "png"}, want: "dir/"},
		{name: "", fields: fields{absPath: "dir/file", isDir: false}, args: args{ext: "png"}, want: "dir/file"},
		{name: "", fields: fields{absPath: "", isDir: false}, args: args{ext: "png"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := convertFile{
				absPath: tt.fields.absPath,
				isDir:   tt.fields.isDir,
			}
			if got := f.arbitraryExtAbsPath(tt.args.ext); got != tt.want {
				t.Errorf("convertFile.arbitraryExtAbsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertFile_isSameExt(t *testing.T) {
	type fields struct {
		absPath string
		isDir   bool
	}
	type args struct {
		ext string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name:"", fields:fields{absPath: "dir/file.png"}, args:args{ext: "png"}, want: true},
		{name:"", fields:fields{absPath: "dir/file.png"}, args:args{ext: "gif"}, want: false},
		{name:"", fields:fields{absPath: ""}, args:args{ext: "gif"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := convertFile{
				absPath: tt.fields.absPath,
				isDir:   tt.fields.isDir,
			}
			if got := f.isSameExt(tt.args.ext); got != tt.want {
				t.Errorf("convertFile.isSameExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
