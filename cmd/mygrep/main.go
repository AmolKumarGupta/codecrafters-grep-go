package main

import (
	// Uncomment this to pass the first stage
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}

	os.Exit(0)
}

func matchLine(line []byte, pattern string) (bool, error) {
	if utf8.RuneCountInString(pattern) == 0 {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	var ok bool

	if pattern == "\\d" {
		pattern = "1234567890"

	} else if pattern == "\\w" {
		pattern = "1234567890qwertyiuopasdfghjklzxcvbnmQWERTYIUOPASDFGHJKLZXCVBNM_"

	} else if strings.HasPrefix(pattern, "[") && strings.HasSuffix(pattern, "]") {
		if string(pattern[1]) == "^" {
			pattern = pattern[2 : len(pattern)-1]
			return negativeCharactorGroup(line, pattern)
		}

		pattern = pattern[1 : len(pattern)-1]
	}

	ok = bytes.ContainsAny(line, pattern)

	return ok, nil
}

func negativeCharactorGroup(line []byte, pattern string) (bool, error) {
	ok := bytes.ContainsAny(line, pattern)
	return !ok, nil
}
