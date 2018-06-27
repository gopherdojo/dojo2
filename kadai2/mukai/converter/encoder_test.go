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
		want encoder
	}{
		{name: "", args: args{format: "jpg"}, want: jpegEncoder{}},
		{name: "", args: args{format: "png"}, want: pngEncoder{}},
		{name: "", args: args{format: "gif"}, want: gifEncoder{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEncoder(tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
