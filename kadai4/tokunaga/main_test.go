package main

import (
	"testing"
	"time"
)


func TestIsOsyougatu(t *testing.T) {
	cases := []struct{
		name string
		input time.Time
		expected bool
	}{
		{name: "12/31", input: time.Date(2017, 12, 31,23,59,11,11, time.UTC), expected: false},
		{name: "1/1", input: time.Date(2018,1,1,9,11,11,11, time.UTC), expected: true},
		{name: "1/2", input: time.Date(2018,1,2,10,12,11,11, time.UTC), expected: true},
		{name: "1/3", input: time.Date(2018,1,3,11,12,11,11, time.UTC), expected: true},
		{name: "1/4", input: time.Date(2018,1,4,0,0,0,0, time.UTC), expected: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := isOsyougatu(c.input); actual != c.expected {
				t.Errorf("want isOsyougatu(%s) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}

func TestGetDaikiti(t *testing.T) {
	expected := "大吉"
	actual := getDaikiti()
	if actual != expected {
		t.Errorf("want: getDaikiti() = %s, got: %s ", expected, actual)
	}
}
