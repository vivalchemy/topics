package someDummyTest

import "fmt"

func Testing() {
	fmt.Println("hello dummy world")
	etst := new(string)
	some, err := fmt.Scanln(etst)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(some, *etst)
}
