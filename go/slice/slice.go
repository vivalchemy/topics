package slice

import "fmt"

func PrintSlice() {
	var slice []int = []int{1, 2, 3, 4, 5}
	var slice2d [][]int = [][]int{{1, 2}, {3, 4}}
	var anotherSlice = []int{1, 2, 3, 4, 5} // another way of initializing an array
	idealSlice := make([]int, 4, 8)         // slice having length 4 and capacity 8
	println("hi")
	fmt.Println(slice)
	fmt.Println(slice2d)
	fmt.Println(anotherSlice)
	fmt.Println(cap(idealSlice))
	fmt.Println(len(idealSlice))

	fmt.Println("Printing slice by value")
	printSliceByValue(slice)
	fmt.Println(slice)
	print2dSliceByValue(slice2d)

	fmt.Println("Printing by reference")
	// this is not possible because slice are always passed by reference and not value
	// though the commented code won't give any error
	printSliceByReference(slice)
	// print2dSliceByReference(&slice2d)
}

// works
func printSliceByValue(arr []int) {
	fmt.Println(arr)
	arr = append(arr, []int{2, 3, 4}...)
	fmt.Println(arr)
}

// works
func print2dSliceByValue(arr [][]int) {
	fmt.Println(arr)
}

// // works
// // made to test how arrays are passed around in golang
func printSliceByReference(a []int) {
	fmt.Println(a)
	// fmt.Println(a[3])
}

// // works
// // made to test how multi dimensional arrays work in golang
// func print2dSliceByReference(a *[][]int) {
// 	fmt.Println(*a)
// 	fmt.Println()
// 	// fmt.Println(a[1][0])
// }
