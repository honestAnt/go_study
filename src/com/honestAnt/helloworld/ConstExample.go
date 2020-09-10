package main

import (
	"fmt"
	"math"
)

/*
https://studygolang.com/articles/11872
在 Go 语言中，术语"常量"用于表示固定的值。比如 5 、-89、 I love Go、67.89 等等。
*/
func main() {
	const a = 55
	//a = 89 //不允许重新赋值
	testImport()
	testNoTypeConst()
	testBoolanConst()
	testNumConst()
}

func testImport() {
	fmt.Println("hello world")
	var a = math.Sqrt(4) //
	//const b = math.Sqrt(4) //不允许
	fmt.Println("a", a)
}

/*
什么类型的字符串属于常量？答案是他们是无类型的。
答案是无类型的常量有一个与它们相关联的默认类型，并且当且仅当一行代码需要时才提供它。在声明中 var name = "Sam" ， name 需要一个类型，它从字符串常量 Sam 的默认类型中获取

*/
func testNoTypeConst() {
	//无类型常量
	var name = "Sam"
	fmt.Println("type %T value %V", name, name)
	const hello = "hello World"
	//有类型常量
	const typeHello string = "hello world"
	fmt.Println("typeHello ", typeHello)
	//混合类型
	var defaultName = "Same"
	type myString string
	var customName myString = "Sam"
	//customName = defaultName // 不被允许
	fmt.Println("defaultname ", defaultName)
	fmt.Println("customName %T, value %V", customName, customName)
}

func testBoolanConst() {
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst
	var customBool myBool = trueConst //允许
	//defaultBool = customBool //不允许
	fmt.Println("defaultBool ", defaultBool, "custom Bool ", customBool)
}

//数字常量可以在表达式中自由混合和匹配，只有当它们被分配给变量或者在需要类型的代码中的任何地方使用时，才需要类型。
func testNumConst() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Println("i is type %T, fis TYpe %T, c is type %T", i, f, c)
}
