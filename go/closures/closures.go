package closures

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func GetClosure() {
	var getNextInt func() int = intSeq()

	fmt.Println(getNextInt())
	fmt.Println(getNextInt())
	fmt.Println(getNextInt())
	fmt.Println(getNextInt())
	fmt.Println(getNextInt())

	var makeAnotherInt func() int = intSeq()
	fmt.Print(makeAnotherInt())
}
