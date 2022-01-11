package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	bytesRead, err := r.r.Read(b)

	if err != nil{
		return bytesRead, io.EOF
	}

	input := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz "
	
	output := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm "

	for i := 0; i < bytesRead; i++ {
		position := strings.IndexByte(input, b[i])

		if position == -1 {
			continue
		}

		b[i] = output[position]
	}

	return bytesRead, nil


}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}