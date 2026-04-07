package dev

import (
	"fmt"
	"io"
	"os"
)

func StreamOutput(prefix string, reader io.Reader) {
	buf := make([]byte, 4096)

	for {
		n, err := reader.Read(buf)
		if n > 0 {
			fmt.Printf("[%s] ", prefix)
			os.Stdout.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
}
