package main

import (
	"fmt"
	"math"
)

/*
https://studygolang.com/articles/11756
*/
func main() {
	var age int //声明变量
	fmt.Println("age is ", age)
	age = 29 //赋值
	fmt.Println("age is: ", age)
	age = 54
	fmt.Println("age new is: ", age)

	var age2 int = 30 //初始化赋值
	fmt.Println("age2 is : ", age2)

	var age3 = 31 //类型推断
	fmt.Println("age3: ", age3)

	var wight, height int = 200, 500 //声明多个类型
	fmt.Println("wight : ", wight, "height: ", height)

	var wight2, height2 int //声明，不给默认值
	fmt.Println("wight2: ", wight2, "height2: ", height2)

	//一次声明多个不同类型的变量
	var (
		name    = "naveen"
		age4    = 20
		height3 int
	)
	fmt.Println("my name is :", name, ", age is: ", age4, ",height3 is: ", height3)
	//简短声明 使用 `:=`
	name2, age5 := "naveen2", 29 //简短声明
	fmt.Println("name2 is : ", name2, ",age5 is ", age5)

	a, b := 20, 30 //声明变量a和b
	fmt.Println("a is ", a, " b is", b)
	b, c := 40, 50 // b已经声明，但C尚未声明;这种情况是正常的
	fmt.Println("b is ", b, "c is ", c)
	b, c = 80, 90 //给已经声明的变量b和c赋新值
	fmt.Println("changed b is ", b, "c is ", c)

	a2, b2 := 30, 40 //声明变量a2,b2
	fmt.Println("a2 is ", a2, "b2 is ", b2)
	//a2, b2 := 50,60 //错误，没有尚未声明的变量 （）

	a3, b3 := 145.8, 543.8
	c3 := math.Min(a3, b3)
	fmt.Println("mininum value is ", c3)

	age6 := 29 //age是int类型
	//age4 = "test" //错误，尝试复制一个字符串给int类型变量
	fmt.Println("age6 is ", age6)
}
