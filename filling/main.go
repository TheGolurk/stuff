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
			w, h := "", ""
			expr := regexp.MustCompile(`(\d+.\d+)`)
			values := expr.FindAll([]byte(v), -1)

			for i, k := range values {
				if i == 2 {
					w = string(k)
				} else if i == 3 {
					h = string(k)
				}
			}

			v = strings.Replace(v, ">", "", 1)
			v = fmt.Sprintf(`%s width="%s" height="%s">`, v, w, h)
		}
	}
	fmt.Println("lines", lines)
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
