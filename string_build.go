package main

import (
	"bytes"
)

func main() {
	buf := new(bytes.Buffer)

	p1 := [...]byte{'A', 'B'}
	buf.Write(p1[:])

	p2 := []byte{'C', 'D'}
	buf.Write(p2)

	str := "EF"
	buf.WriteString(str)

	println(buf.String())
}
