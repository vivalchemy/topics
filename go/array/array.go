package array

import "fmt"

func PrintArray() {
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	var arr2d [2][2]int = [2][2]int{{1, 2}, {3, 4}}
	var anotherArr = [5]int{1, 2, 3, 4, 5} // another way of initializing an array
	println("hi")
	fmt.Println(arr)
	fmt.Println(arr2d)
	fmt.Println(anotherArr)

	fmt.Println("Printing by value")
	printArrayByValue(arr)
	print2dArrayByValue(arr2d)

	fmt.Println("Printing by reference")
	printArrayByReference(&arr)
	print2dArrayByReference(&arr2d)
}

// works
func printArrayByValue(arr [5]int) {
	fmt.Println(arr)
}

// works
func print2dArrayByValue(arr [2][2]int) {
	fmt.Println(arr)
}

// works
// made to test how arrays are passed around in golang
func printArrayByReference(a *[5]int) {
	fmt.Println(*a)
	fmt.Println(a[3])
}

// works
// made to test how multi dimensional arrays work in golang
func print2dArrayByReference(a *[2][2]int) {
	fmt.Println(*a)
	fmt.Println(a[1])
	fmt.Println(a[1][0])
}
