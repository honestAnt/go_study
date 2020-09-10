package main

import "fmt"

/*
https://studygolang.com/articles/11892
函数是一块执行特定任务的代码。一个函数是在输入源基础上，通过执行一系列的算法，生成预期的输出。

func functionName(parametername type) returntype {
	//函数体（具体实现的功能）
}
函数的声明以关键词 func 开始，后面紧跟自定义的函数名 functionname (函数名)。函数的参数列表定义在 ( 和 ) 之间，返回值的类型则定义在之后的 returntype (返回值类型)处。
声明一个参数的语法采用 参数名 参数类型 的方式，任意多个参数采用类似 (parameter1 type, parameter2 type) 即(参数1 参数1的类型,参数2 参数2的类型)的形式指定。之后包含在 { 和 } 之间的代码，就是函数体。
函数中的参数列表和返回值并非是必须的
func functionname() {
	//该函数不需要参数，也没有返回值
}
*/

func main() {
	bill := calculateBill(10, 5)
	bill2 := calculateBill2(10, 5)
	fmt.Println("bill", bill, "bill2 ", bill2)
}

func calculateBill(price int, no int) int {
	var totalPrice = price * no // 商品总价=单价* 数量
	return totalPrice
}

//连续参数如果类型一致可以只需在最后添加该类型
func calculateBill2(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

//多返回值

func rectProps(legth, width float64) (float64, float64) {

}
