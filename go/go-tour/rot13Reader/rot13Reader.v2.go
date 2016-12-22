package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		v := p[i]
		switch {
			case v > 'A' && v < 'Z':
				v -= 'A'
				v = (v + 13) % 26
				v += 'A'
			case v > 'a' && v < 'z':
				v -= 'a'
				v = (v + 13) % 26
				v += 'a'
		}
		p[i] = v
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
