package converter

import (
	"reflect"
	"testing"
)

func TestGetEncoder(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		args args
		want Encoder
	}{
		{name: "", args: args{format: "jpg"}, want: jpegEncoder{}},
		{name: "", args: args{format: "png"}, want: pngEncoder{}},
		{name: "", args: args{format: "gif"}, want: gifEncoder{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEncoder(tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
