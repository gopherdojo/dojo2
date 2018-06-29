package main

import (
	"bufio"
	"compress/gzip"
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func ec(err error) {
	if err != nil {
		panic(err)
	}
}

func loadDic(path string) {
	file, err := os.Open(path)
	ec(err)
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	ec(err)
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	ec(err)

	//配列に入れる
	t, err := tarReader.Next()
	ec(err)
	size := make([]byte, t.Size)
	dictionary := make([]string, 3000, size)
	for{
		row, err := tarReader.Read(size)
                if err == io.EOF {
                    break
				}
		
	}
	

}

func input(r io.Reader) <-chan string {
	// TODO: チャネルを作る
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
			ch <- s.Text()
		}
		// TODO: チャネルを閉じる
		close(ch)
	}()
	// TODO: チャネルを返す
	return ch
}

func main() {
	loadDic("../dictionary.txt.gz")
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}

}
