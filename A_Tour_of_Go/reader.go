package main

import (
	"fmt"
	"io"
	"strings"
)

func re() {
	r := strings.NewReader("hello world")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}

	}
}

type MyReader struct {
}

func (r MyReader) Read([]byte) (int, error) {
	return 0, nil
}
func main() {
	re()
	//reader.Validate(MyReader{})

}
