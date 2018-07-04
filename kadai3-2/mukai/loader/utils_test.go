package loader

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
		{name: "be divisible", args: args{10, 3}, want: []string{"0-3", "4-6", "7-9", "10-10"}},
		{name: "not be divisible", args: args{9, 3}, want: []string{"0-3", "4-6", "7-9"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitRange(tt.args.fileSize, tt.args.split); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseContentRange(t *testing.T) {
	type args struct {
		contentRange string
	}
	tests := []struct {
		name     string
		args     args
		wantMax  int
		wantHigh int
		wantErr  bool
	}{
		{name:"success", args:args{contentRange:"1-10/100"}, wantMax: 100, wantHigh:10, wantErr:false},
		{name:"parse error", args:args{contentRange:"1-10100"}, wantMax: -1, wantHigh:-1, wantErr:true},
		{name:"parse error blank", args:args{contentRange:""}, wantMax: -1, wantHigh:-1, wantErr:true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotHigh, err := parseContentRange(tt.args.contentRange)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseContentRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMax != tt.wantMax {
				t.Errorf("parseContentRange() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
			if gotHigh != tt.wantHigh {
				t.Errorf("parseContentRange() gotHigh = %v, want %v", gotHigh, tt.wantHigh)
			}
		})
	}
}
