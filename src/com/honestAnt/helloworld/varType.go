package main

import "fmt"

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

}
