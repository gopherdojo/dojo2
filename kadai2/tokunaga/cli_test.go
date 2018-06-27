package tokunaga

import "testing"

// 引数の拡張子が許可されているものならばtrue, それ以外なら false を返す
func TestCheckExtPermmited(t *testing.T) {
	cases := []struct {
		name          string
		inputExt      string
		permittedExts []string
		expected      bool
	}{
		{name: "ext is permitted", inputExt: "png", permittedExts: []string{"png", "jpeg"}, expected: true},
		{name: "ext is not permitted", inputExt: "gif", permittedExts: []string{"png", "jpeg"}, expected: false},
		{name: "ext is not permitted", inputExt: "", permittedExts: []string{"png", "jpeg"}, expected: false},
		{name: "ext is not permitted", inputExt: "png", permittedExts: []string{}, expected: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := checkExtPermmited(c.inputExt, c.permittedExts); actual != c.expected {
				t.Errorf("want checkExtPermmited(%s, %v) = %v, got %v", c.inputExt, c.permittedExts, c.expected, actual)
			}
		})
	}
}
