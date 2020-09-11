package main

import "fmt"

/*
https://studygolang.com/articles/11902
if 是条件语句。if 语句的语法是

if condition {
}
如果 condition 为真，则执行 { 和 } 之间的代码。

不同于其他语言，例如 C 语言，Go 语言里的 { } 是必要的，即使在 { } 之间只有一条语句。
*/

func main() {
	ifTest()
	stateIfTest()
	ifElseTest()
}

func ifTest() {
	num := 10
	if num%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}

func stateIfTest() {
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

/*
else 语句应该在 if 语句的大括号 } 之后的同一行中。如果不是，编译器会不通过。
error example:
if condition {
	//xxx
}
else {
	//xxx2
}
*/
func ifElseTest() {
	num := 99
	if num <= 50 {
		fmt.Println("number is less then or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}
}
