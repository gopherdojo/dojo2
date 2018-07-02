package game

import (
	"fmt"
)

func ExampleVocabulary_NextWord() {
	v := Vocabulary([]string{"foo"})
	fmt.Print(v.NextWord())
	// Out: foo
}
