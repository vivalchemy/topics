package compositedts

import "fmt"

func Slice() {
	var a []int = []int{}
	b := []int{}
	c := make([]int, 4, 6) //len 4 capacity 6. Elements till length can be accessed and are zeroed
	d := make([]int, 5)    // len 5 capacity 5
	a = append(a, 5, 4, 6)
	b = append(b, a...) // this spreads the elements in a so that instead of []int only ints are passed
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func Array() {
	var a [5]int = [5]int{0, 1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5, 0}
	fmt.Println(a)
	fmt.Println(b)
}

func Map() {
	var a map[bool]int = map[bool]int{true: 2, false: 4}
	b := map[bool]int{}
	c := make(map[string]int)    // empty map. Assign a space in the second argument to preallocate certain space
	d := make(map[string]int, 4) // still an empty map
	b[true] = 1
	b[false] = 0
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func main() {
	// Slice()
	// Array()
	Map()
}
