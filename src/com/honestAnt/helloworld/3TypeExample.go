package main

import (
	"fmt"
	"unsafe"
)

/**
https://studygolang.com/articles/11869
*Go 支持的基本类型：
bool
数字类型
	int8, int16, int32, int64, int
	uint8, uint16, uint32, uint64, uint
	float32, float64
	complex64, complex128
	byte
	rune
string
*/
func main() {
	booleanTest()
	intTest()
	floatTest()
	complexTest()
	stringTest()
	typeConvertTest()
}

func booleanTest() {
	a := true
	b := false
	fmt.Println("a: ", a, " b: ", b)
	c := a && b //且
	fmt.Println("c:", c)
	d := a || b //或
	fmt.Println("d:", d)
}

/*
有符号整型
int8：表示 8 位有符号整型
大小：8 位
范围：-128～127

int16：表示 16 位有符号整型
大小：16 位
范围：-32768～32767

int32：表示 32 位有符号整型
大小：32 位
范围：-2147483648～2147483647

int64：表示 64 位有符号整型
大小：64 位
范围：-9223372036854775808～9223372036854775807

int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型。
大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
范围：在 32 位系统下是 -2147483648～2147483647，而在 64 位系统是 -9223372036854775808～9223372036854775807。
*/
/*无符号整型
uint8：表示 8 位无符号整型
大小：8 位
范围：0～255

uint16：表示 16 位无符号整型
大小：16 位
范围：0～65535

uint32：表示 32 位无符号整型
大小：32 位
范围：0～4294967295

uint64：表示 64 位无符号整型
大小：64 位
范围：0～18446744073709551615

uint：根据不同的底层平台，表示 32 或 64 位无符号整型。
大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
范围：在 32 位系统下是 0～4294967295，而在 64 位系统是 0～18446744073709551615。
*/
func intTest() {
	var a int = 89
	b := 95
	fmt.Println("value of a is: ", a, " and b is :", b)
	fmt.Println("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) //a的类型和大小
	fmt.Println("type of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //b的类型和大小
}

/*浮点型*/
/*float32：32 位浮点数*/
/*float64：64 位浮点数*/

func floatTest() {
	a, b := 5.67, 8.97
	fmt.Println("type of a %T, b %T ", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum ", sum, "diff ", diff)
	no1, no2 := 56, 89
	fmt.Println("sum ", no1+no2, "diff ", no1-no2)
}

/*复数类型
complex64：实部和虚部都是 float32 类型的的复数。
complex128：实部和虚部都是 float64 类型的的复数。*/

/*
其他数字类型
byte 是 uint8 的别名。
rune 是 int32 的别名。
*/
func complexTest() {
	c := 6 + 7i
	fmt.Println("c ", c)
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum", cadd)
	cmul := c1 * c2
	fmt.Println("cmul ", cmul)

}

func stringTest() {
	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("my name is ", name)
}

func typeConvertTest() {
	//i := 55
	j := 67.8
	h := 67.8
	//sum := i + j  //非法
	sum2 := h + j
	//fmt.Println("sum", sum)
	fmt.Println("sum2 ", sum2)
	i := 10
	var j2 float64 = float64(i)
	fmt.Println("j2 ", j2)
}
