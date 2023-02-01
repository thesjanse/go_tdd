package hello

import (
	"bytes"
	"fmt"
)

func Hello(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}
