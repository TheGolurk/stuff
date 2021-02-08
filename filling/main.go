package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("file.html", os.O_APPEND|os.O_RDWR|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	line := 0
	for scanner.Scan() {
		line++
		text := scanner.Text()
		if strings.Contains(text, "<svg") {
			fmt.Println(text, line)
			n, err := file.WriteAt([]byte("HOLA"), int64(line))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("N: ", n)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
