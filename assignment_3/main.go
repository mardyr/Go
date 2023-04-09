package main

/*
Create a program that reads the contents of a text file then
prints its contents to the terminal

The file to open should be provided as an argument to the program
when it is executed at the terminal. For example "go run main.go myfile.txt"
should open up the myfile.txt file

To read in the arguments provided to a program, you can reference
the variable 'os.Args' which is a slice of type string

To open a file check out the documentaion for the `Open`
function in the os package

What intefaces does the the 'File' implement?

If the 'File' type implements the 'Reader' interface you might be able
to use io.Copy function
*/

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(readFileFromTerminal(os.Args))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func readFileFromTerminal(args []string) string {
	return args[len(args)-1]
}
