// package function
package main

import "fmt"

func typing(a, b int, c, d string) {
	fmt.Println(a, b, c, d)
	fmt.Printf("%T %T %T %T\n", a, b, c, d)
}

func variadic(str ...string) {
	fmt.Println(str)
	fmt.Println(len(str))
}

func main() {
	typing(1, 2, "3", "4")
	variadic("1", "2", "3", "4")
}
