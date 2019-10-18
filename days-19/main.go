package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func convert(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b += 13
	case 'M' < b && b <= 'Z':
		b -= 13
	case 'a' <= b && b <= 'm':
		b += 13
	case 'm' < b && b <= 'z':
		b -= 13
	}
	return b
}

func (rot13 *rot13Reader) Read(buf []byte) (s int, e error) {
	s, e = rot13.r.Read(buf)

	for i := 0; i < s; i++ {
		buf[i] = convert(buf[i])
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
