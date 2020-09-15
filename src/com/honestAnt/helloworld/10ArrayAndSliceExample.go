package main

import "fmt"

/*
https://studygolang.com/articles/12121
数组是同一类型元素的集合。例如，整数集合 5,8,9,79,76 形成一个数组。Go 语言中不允许混合不同类型的元素，例如包含字符串和整数的数组。（译者注：当然，如果是 interface{} 类型数组，可以包含任意类型）
*/

func main() {
	arrayTest1()
	arrayTest2()
	arrayTest3()
	arrayTest4()
	arrayTest5()
}

func arrayTest1() {
	var a [3]int // int array with length 3. 数组中的所有元素都被自动赋值为数组类型的零值
	fmt.Println(a)
}

//声明 后赋值
func arrayTest2() {
	var a [3]int
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println(a)
}

//简略声明
func arrayTest3() {
	a := [3]int{12, 78, 50}
	b := [3]int{12}               //只赋值第一个
	c := [...]int{12, 50, 78, 62} // makes the compiler determine the length
	fmt.Println("a", a, "b: ", b, "c: ", c)
}

/*
Go 中的数组是值类型而不是引用类型。这意味着当数组赋值给一个新的变量时，该变量会得到一个原始数组的一个副本。如果对新变量进行更改，则不会影响原始数组。
*/
func arrayTest4() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a              //把a的值赋给b
	b[0] = "Singapore"  //修改b的第一个元素值
	fmt.Println("a", a) //a的第一个元素并没有改变
	fmt.Println("b", b)
}

//当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变。
//数组的长度
//通过将数组作为参数传递给 len 函数，可以得到数组的长度。
func arrayTest5() {
	num := [...]int{5, 6, 7, 8, 9}

	fmt.Println("length of a is", len(num))
	fmt.Println("before passing to function ", num)
	arrayChangeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
}

func arrayChangeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}
