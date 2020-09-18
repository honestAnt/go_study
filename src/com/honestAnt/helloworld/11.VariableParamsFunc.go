package main

import "fmt"

/*
https://studygolang.com/articles/12173
可变参数函数是一种参数个数可变的函数。
eg: func append(slice []Type, elems ...Type) []Type
可变参数函数的工作原理是把可变参数转换为一个新的切片。
*/
func main() {
	varFuncTest1()
	varFuncTest2()
	varFuncTest3()
}

func varFuncTest1() {
	varFindTest1(89, 89, 90, 95)
	varFindTest1(45, 56, 67, 45, 90, 109)
	varFindTest1(78, 38, 56, 98)
	varFindTest1(87)
}

func varFindTest1(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index ", i, "in", nums)
			found = true
		}
		if !found {
			fmt.Println(num, "not found in ", nums)
		}
		fmt.Printf("\n")
	}
}

//切片传入可变参数
func varFuncTest2() {
	nums := []int{89, 90, 95}
	//varFindTest1(89, nums) // 编译报错
	varFindTest1(89, nums...)
}

func varFuncTest3() {
	welcomeSlice := []string{"hello", "slice world"}
	varChangeTest(welcomeSlice...)
	fmt.Println(welcomeSlice)
}

//如果是切片会改变原始的值；如果是数组不会
func varChangeTest(s ...string) {
	s[0] = "Go"
}
