package main

import "fmt"

/*
https://studygolang.com/articles/11755
*/
func main() {
	fmt.Println("helloworld")
	test1("test")
	fmt.Println(result("result"))
}

func test1(a string) {
	fmt.Println(a)
}

func result(a string) (b string) {
	fmt.Println(a)
	return "bbb"
}
