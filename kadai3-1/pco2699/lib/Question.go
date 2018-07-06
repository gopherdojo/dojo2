package Question

import (
	"io"
	"bufio"
	"os"
	"math/rand"
	"time"
)

type Reader interface
{
	Read(r io.Reader) error
}

type Randomizer interface {
	Randomize() string
}

type Opener interface {
	Open(file_path string) (io.Reader, error)
}

type Question struct {
	words []string
	Reader
	Randomizer
	Opener
}

func (q *Question) Open(filePath string) (io.Reader, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (q *Question) Read(r io.Reader) error {
	reader := bufio.NewReader(r)
	w := make([]string, 3, 10)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		w = append(w, string(line))
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	q.words = w

	return nil
}

func (q *Question) Randomize() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(q.words))
	return q.words[random]
}
