package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileHandle, _ := os.OpenFile("output.txt", os.O_APPEND, 0666)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, "\r")
	fmt.Fprintln(writer, "String I want to append")
	writer.Flush()
}
