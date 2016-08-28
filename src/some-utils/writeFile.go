package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println(" hello")
	content := "Hello World"
	file, error := os.Create("./StringTOFile.txt")
	defer file.Close()
	hadleError(error)
	_, error2 := io.WriteString(file, content)
	hadleError(error2)

}

func hadleError(theError error) {
	if theError != nil {
		panic(theError)
	}
}
