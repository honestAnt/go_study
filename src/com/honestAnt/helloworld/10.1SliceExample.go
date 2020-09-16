package main

import "fmt"

/*
https://studygolang.com/articles/12121
切片是由数组建立的一种方便、灵活且功能强大的包装（Wrapper）。切片本身不拥有任何数据。它们只是对现有数组的引用。
*/

func main() {

}

func sliceTest1() {
	c := []int{6, 7, 8} // creates and array and returns a slice reference
	fmt.Println(c)
}
