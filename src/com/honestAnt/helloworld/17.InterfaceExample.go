package main

import (
	"fmt"
)

/*
https://studygolang.com/articles/12266
什么是接口？
在面向对象的领域里，接口一般这样定义：接口定义一个对象的行为。接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定
*/

func main() {
	interTest1()
	interTest2()
	interTest3()
	interTest4()
	interTest5()
	interTest6()
}

//定义一个接口，包含一个方法
type InterVoiwelsFinder interface {
	InterFinderVowels() []rune
}

type InterMyString string

func (ms InterMyString) InterFinderVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

//接口就是方法签名（Method Signature）的集合。当一个类型定义了接口中的所有方法，我们称它实现了该接口。这与面向对象编程（OOP）的说法很类似。
// 接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法
//如果一个类型包含了接口中声明的所有方法，那么它就隐式地实现了 Go 接口
func interTest1() {
	name := InterMyString("Sam Anderson")
	var v InterVoiwelsFinder
	v = name // possible since InterMyString implements InterVoiwelsFinder
	fmt.Printf("Vowels are %c", v.InterFinderVowels())
}

type InterSalaryCaculator interface {
	InterCalculateSalary() int
}
type InterPermanent struct {
	empId    int
	basicpay int
	pf       int
}
type InterContract struct {
	empId    int
	basicpay int
}

func (p InterPermanent) InterCalculateSalary() int {
	return p.basicpay + p.pf
}

func (c InterContract) InterCalculateSalary() int {
	return c.basicpay
}

func InterTotalExpense(s []InterSalaryCaculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.InterCalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func interTest2() {
	pemp1 := InterPermanent{1, 5000, 20}
	pemp2 := InterPermanent{2, 6000, 30}
	cemp1 := InterContract{3, 3000}
	employees := []InterSalaryCaculator{pemp1, pemp2, cemp1}
	InterTotalExpense(employees)
}

type InterTest interface {
	InterTester()
}

type InterMyFloat float64

func (m InterMyFloat) InterTester() {
	fmt.Println(m)
}

func InterDescirbe(t InterTest) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

//空接口
func InterDescirbe2(t interface{}) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

//我们可以把接口看作内部的一个元组 (type, value)。 type 是接口底层的具体类型（Concrete Type），而 value 是具体类型的值。
//空接口
//没有包含方法的接口称为空接口。空接口表示为 interface{}。由于空接口没有方法，因此所有类型都实现了空接口。
func interTest3() {
	var t InterTest
	f := InterMyFloat(89.7)
	t = f
	InterDescirbe(t)
	InterDescirbe2(t)
	t.InterTester()
}

//类型断言
//类型断言用于提取接口的底层值（Underlying Value）。如果类型不是int会报错
//v, ok := i.(T)
func InterAssert(i interface{}) {
	//s := i.(int) //get the underlying int value from i； 如果类型不为int的会报错
	v, ok := i.(int) //当类型不为 int时不会报错，ok为false
	fmt.Println(v, ok)
}

func interTest4() {
	var s interface{} = 56
	InterAssert(s)
	var s2 interface{} = "Steven Paul"
	InterAssert(s2)
}

func InterFindType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

//类型选择（Type Switch）
//类型选择用于将接口的具体类型与很多 case 语句所指定的类型进行比较。它与一般的 switch 语句类似。唯一的区别在于类型选择指定的是类型，而一般的 switch 指定的是值。
//类型选择的语法类似于类型断言。类型断言的语法是 i.(T)，而对于类型选择，类型 T 由关键字 type 代替。下面看看程序是如何工作的。
func interTest5() {
	InterFindType("Naveen")
	InterFindType(77)
	InterFindType(89.98)
}

//还可以将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。
type InterDescriber interface {
	InterDescribe()
}
type InterPerson struct {
	name string
	age  int
}

func (p InterPerson) InterDescribe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}
func InterFindType2(i interface{}) {
	switch v := i.(type) {
	case InterDescriber:
		v.InterDescribe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func interTest6() {
	InterFindType2("Naveen")
	p := InterPerson{
		name: "Naveen R",
		age:  25,
	}
	InterFindType2(p)
}
