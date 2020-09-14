package main

import "fmt"

/*
https://studygolang.com/articles/11957
switch 是一个条件语句，用于将表达式的值与可能匹配的选项列表进行比较，并根据匹配情况执行相应的代码块。它可以被认为是替代多个 if else 子句的常用方式。
*/
func main() {
	swTest1(4)
	swTest1(5)
	swTest2(5)
	swTest2(25)
	swTest3("e")
	swTest3("h")
	swTest4(50)
	swTest4(99)
	swTest4(101)
	swFallthrough(23)
	swFallthrough(59)
	swFallthrough(102)
}

func swTest1(finger int8) {
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}
}

func swTest2(finger int) {
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: //默认情况
		fmt.Println("incorrect finger number")

	}
}

//一个case 多个条件
func swTest3(finger string) {
	switch finger {
	case "a", "e", "i", "o", "u": //一个选项多个表达式
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}

//无表达式
func swTest4(num int) {
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num > 100:
		fmt.Println("num is greater than 100")
	}
}

//Fallthrough;在 Go 中，每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行。
// 使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中。
func swFallthrough(num int) {
	switch {
	case num < 50:
		fmt.Printf(" %d num is lesser than 50;", num)
		fallthrough
	case num < 100:
		fmt.Printf(" %d num is lesser than 100;", num)
		fallthrough
	case num < 200:
		fmt.Printf(" %d num is lesser than 200;", num)
	}
}
