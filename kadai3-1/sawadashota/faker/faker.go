package faker

import "github.com/icrowley/fake"

// ShortWordLength is maximum chars to type easy
const ShortWordLength = 5

type Easy struct{}

// Word returns shorter than 5 chars
func (q *Easy) Word() string {
	for {
		if w := fake.Word(); len([]rune(w)) <= ShortWordLength {
			return w
		}
	}
}

type Hard struct{}

// Word returns shorter than 5 chars
func (q *Hard) Word() string {
	for {
		if w := fake.Word(); len([]rune(w)) > ShortWordLength {
			return w
		}
	}
}
