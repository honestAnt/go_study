package main

import "fmt"

/*
https://studygolang.com/articles/11924
循环语句是用来重复执行某一段代码。

for 是 Go 语言唯一的循环语句。Go 语言中并没有其他语言比如 C 语言中的 while 和 do while 循环。
*/

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	breakTest()
	continueTest()
	otherTest()
	otherTest2()
	multiVarTest()
}

func breakTest() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminater if i > 5
		}
		fmt.Printf(" %d", i)
	}
	fmt.Printf("\nline after for loop")
}

func continueTest() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf(" %d", i)
	}
}

func otherTest() {
	i := 0
	for i <= 10 {
		fmt.Printf(" %d", i)
		i += 2
	}
}
func otherTest2() {
	i := 0
	for i <= 10 {
		fmt.Printf(" %d", i)
		i += 2
	}
}

//循环中可以声明和操作多个变量
func multiVarTest() {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}

/*
无限循环
for {
	//
}
*/
