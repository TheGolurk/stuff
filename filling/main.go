package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("file.html")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "<svg") {
			fmt.Println(text)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
