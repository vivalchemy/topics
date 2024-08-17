package maps

import "fmt"

func PrintMap() {
	var m map[string]int = map[string]int{
		"a": 1,
		"b": 2,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
