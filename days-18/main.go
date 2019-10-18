package main

import (
	"fmt"
	"io"
)

type MyReader struct {
}

func (r *MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	b := make([]byte, 8)
	mr := &MyReader{}
	for {
		s, err := mr.Read(b)
		fmt.Printf("read length: %d, conent: %q\n", s, b[:s])
		if err == io.EOF {
			break
		}
	}
}

/*
练习：Reader
实现一个 Reader 类型，它不断生成 ASCII 字符 'A' 的流。
*/
