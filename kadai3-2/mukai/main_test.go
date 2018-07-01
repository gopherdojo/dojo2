package main

import (
	"reflect"
	"testing"
)

func Test_splitRange(t *testing.T) {
	type args struct {
		fileSize int
		split    int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name:"", args:args{10, 3}, want:[]string{"0-3", "4-6", "7-9", "10-10"}},
		{name:"", args:args{9, 3}, want:[]string{"0-3", "4-6", "7-9"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitRange(tt.args.fileSize, tt.args.split); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
