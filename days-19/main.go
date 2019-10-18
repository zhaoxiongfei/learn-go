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

/*
练习：rot13Reader
一个常见模式是 io.Reader 包裹另一个 `io.Reader`，然后通过某种形式修改数据流。

例如，gzip.NewReader 函数接受 `io.Reader`（压缩的数据流）并且返回同样实现了 io.Reader 的 `*gzip.Reader`（解压缩后的数据流）。

编写一个实现了 io.Reader 的 `rot13Reader`， 并从一个 io.Reader 读取， 利用 rot13 代换密码对数据流进行修改。

已经帮你构造了 rot13Reader 类型。 通过实现 Read 方法使其匹配 `io.Reader`。
*/
