package main

import (
	"bufio"
	"fmt"
	"helper/helper"
	"os"
	"strings"
)

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := args[1]
	outputFile := args[2]

	read, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("error creating file")
		return
	}
	defer read.Close()

	outPutread, _ := os.Create(outputFile)

	scanner := bufio.NewScanner(read)
	write := bufio.NewWriter(outPutread)
	defer outPutread.Close()

	linecount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == strings.ToUpper(line) {
			line = helper.Lower(line)
			linecount++
		} else if line == strings.ToLower(line) {
			line = helper.Upper(line)
			linecount++

		} else if strings.Contains(line, "TODO") {
			line = helper.TODO(line)
			fmt.Println(line)
			linecount++

		} else if strings.Contains(line, "REVERSE") {
			line = helper.Reverse(line)
			linecount++

		}

		_, err := write.WriteString(line + "\n")
		if err != nil {
			fmt.Println("error writing file")
			return
		}

	}
	write.Flush()
	fmt.Println("Transformation completed")

	// for scanner.Scan() {
	// 	// line := Scan.Text()
	// 	linecount++
	// }

	fmt.Println("Number of lines:", linecount)
}
