package util

import (
	"log"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	stream, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(stream), "\n")
	return lines
}
