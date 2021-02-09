package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	lines, err := File2lines("file.html")
	if err != nil {
		fmt.Println("11", err)
	}
	for _, v := range lines {
		if strings.Contains(v, "<svg") {
			expr := regexp.MustCompile(`(\d+.\d+)`)
			values := expr.FindAll([]byte(v), -1)

			for i, k := range values {
				fmt.Printf("%q %d /n", k, i)
			}

			return
		}
	}
}

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
