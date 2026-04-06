package dev

import (
	"bufio"
	"fmt"
	"io"
)

func StreamOutput(prefix string, reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Printf("[%s] %s\n", prefix, scanner.Text())
	}
}
