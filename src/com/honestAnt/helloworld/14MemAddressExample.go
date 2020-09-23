package main

import "fmt"

/*
https://studygolang.com/articles/12262
指针是一种存储变量内存地址（Memory Address）的变量。
指针变量的类型为 *T，该指针指向一个 T 类型的变量
*/

func main() {
	addressTest()
	addressTest2()
	addressTest3()
	addressTest4()
	addressTest5()
	addressTest6()
}

//由于 b 可能处于内存的任何位置，你应该会得到一个不同的地址。
func addressTest() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is $T\n", a)
	fmt.Printf("address of b is ", a)
}

//指针的零值（Zero Value）
func addressTest2() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is ", b)
		b = &a
		fmt.Println("b after initialization is ", b)
	}
}

//指针的解引用
//指针的解引用可以获取指针所指向的变量的值。将 a 解引用的语法是 *a。
func addressTest3() {
	b := 255
	a := &b
	fmt.Println("address of b is ", a)
	fmt.Println("value of b is ", *a)
	//用指针修改值
	*a++
	fmt.Println("value of b is ", b)
}

func changeAddress(val *int) {
	*val = 55
}

//向函数传递指针参数
func addressTest4() {
	a := 58
	fmt.Println("value of a before function call is ", a)
	b := &a
	changeAddress(b)
	fmt.Println("value of a after function call is ", a)
}

//a[x] 是 (*a)[x] 的简写形式，因此上面代码中的 (*arr)[0] 可以替换为 arr[0]
func modifyAddressByArray(arr *[3]int) { //数组；不推荐
	(*arr)[0] = 90
}
func modifyAddressBySlice(arr []int) { //切片
	arr[0] = 90
}

//不要向函数传递数组的指针，而应该使用切片
//假如我们想要在函数内修改一个数组，并希望调用函数的地方也能得到修改后的数组，一种解决方案是把一个指向数组的指针传递给这个函数。
func addressTest5() {
	a := [3]int{89, 90, 91}
	modifyAddressByArray(&a)
	fmt.Println(a)
	//切片方式
	b := [3]int{89, 90, 91}
	modifyAddressBySlice(b[:])
	fmt.Println(b)
}

//Go 不支持指针运算
func addressTest6() {
	b := [...]int{109, 110, 111}
	p := &b
	//p++  上面的程序会抛出编译错误：main.go:6: invalid operation: p++ (non-numeric type *[3]int)
	fmt.Println("b address is ", p)
}
