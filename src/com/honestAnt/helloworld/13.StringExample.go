package main

import (
	"fmt"
	"unicode/utf8"
)

/*
https://studygolang.com/articles/12261
Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，我们可以创建一个字符串
*/

func main() {
	stringTest1()
	stringTest2()
	stringTest3()
	stringTest4()
	stringTest5()
	stringTest6()
	stringTest7()
}

func stringTest1() {
	name := "hello world"
	fmt.Println(name)
}

//单独获取字符串的每一个字节
func stringTest2() {
	name := "Hello World"
	stringPrintByte(name)
	stringPrintChar(name)
}

//打印字符
func stringPrintChar(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

//打印字节
func stringPrintByte(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

//打印字节;
//而我们打印字符时，却假定每个字符的编码只会占用一个字节，这是错误的。在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间
//rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示
func stringPrintCharByRun(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%x ", runes[i])
	}
}

//特殊字符
func stringTest3() {
	name := "Hello World"
	stringPrintByte(name)
	fmt.Printf("\n")
	stringPrintChar(name)
	fmt.Printf("\n")
	name = "Señor"
	stringPrintByte(name)
	fmt.Printf("\n")
	stringPrintChar(name)
	fmt.Printf("\n")
	stringPrintCharByRun(name)
}

//字符串的 for range 循环
func stringTest4() {
	name := "Señor"
	for index, rune := range name {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

//用字节切片构造字符串
func stringTest5() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
	byteSlice = []byte{67, 97, 102, 195, 169} //10进制的
	str = string(byteSlice)
	fmt.Println(str)
}

//用 rune 切片构造字符串
func stringTest6() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)
	//字符串长度 func RuneCountInString(s string) (n int) 方法用来获取字符串的长度
	fmt.Printf("length of %s is %d\n", str, utf8.RuneCountInString(str))
}

//字符串是不可变的
//Go 中的字符串是不可变的。一旦一个字符串被创建，那么它将无法被修改
//为了修改字符串，可以把字符串转化为一个 rune 切片。然后这个切片可以进行任何想要的改变，然后再转化为一个字符串
func stringTest7() {
	h := "hello"
	fmt.Println(stringMutate(h))
	fmt.Println(stringMutate2([]rune(h)))
}

func stringMutate(s string) string {
	//s[0] = 'a' //cannot assign to s[0]
	return string(s)
}

func stringMutate2(s []rune) string {
	s[0] = 'a'
	return string(s)
}
