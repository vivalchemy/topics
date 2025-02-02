// package filemanipulation
package main

import (
	"errors"
	"fmt"
	"os"
)

// create a new file in the specified location
func Create(path, fileName, ext string) {
	fileToEdit := fmt.Sprint(path, fileName, ".", ext)
	_, err := os.Create(fileToEdit)
	if err != nil {
		errors.New(fmt.Sprint("Enable to create a file", err))
	}
	file, err := os.OpenFile(fileToEdit, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		errors.New(fmt.Sprint("Enable to open the file", err))
	}
	defer file.Close()
	fmt.Println(file)
	file.WriteString("This is the text being appended")
}

func main() {
	Create("/home/shadow/Public/", "fileFromGo", "txt")
}
