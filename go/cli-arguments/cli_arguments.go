package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("### CLI Arguments Example\n\nThis example shows how to use the os.Args global variable to get the command line arguments.\n\nusage: go run cli_arguments.go -mini_flag --big_flag -1\n\n")
	for index, arg := range os.Args {
		if arg[0] == '-' {
			if arg[1] == '-' {
				fmt.Printf("Mini flag: %s Type: %T Index: %d\n", arg, arg, index)
				goto outside
			}
			fmt.Printf("Big Flag: %s Type: %T Index: %d\n", arg, arg, index)
			goto outside
		}
		fmt.Printf("!Flag: %s Type: %T Index: %d\n", arg, arg, index)
	outside:
	}
	os.Exit(0)
}
