package main

import "fmt"

/*
https://studygolang.com/articles/12264
方法其实就是一个函数，在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。
func (t Type) methodName(parameter list) {
}
*/

func main() {
	methodTest1()
	methodTest2()
	methodTest3()
	methodTest4()
	methodTest5()
	methodTest6()
}

type MethodExample struct {
	name     string
	age      int
	salary   int
	currency string
}

//methodDisplaySalary() 方法将 Employee 做为接收器类型
func (e MethodExample) methodDisplaySalary() {
	fmt.Printf("salary of %s is %s%d", e.name, e.currency, e.salary)
}

/*
methodDisplaySalary()方法被转化为一个函数，把 Employee 当做参数传入。
*/
func methodDisplaySalary2(e MethodExample) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func methodTest1() {
	emp1 := MethodExample{
		name:     "Sam adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.methodDisplaySalary()
	methodDisplaySalary2(emp1)
	//Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
	//相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的
}

//指针接收器与值接收器
//到目前为止，我们只看到了使用值接收器的方法。还可以创建使用指针接收器的方法。值接收器和指针接收器之间的区别在于，在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的。
/*
那么什么时候使用指针接收器，什么时候使用值接收器？
一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。

指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。

在其他的所有情况，值接收器都可以被使用。
*/
func (e MethodExample) methodChangeName(newName string) {
	e.name = newName
}

func (e *MethodExample) methodChangeAge(newAge int) {
	e.age = newAge
}

func methodTest2() {
	e := MethodExample{
		name: "Mark Andrew",
		age:  50,
	}

	fmt.Printf("Employee name before change: %s", e.name)
	e.methodChangeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).methodChangeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
}

//匿名字段的方法
//属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样。
type methodAddress struct {
	city  string
	state string
}

func (a methodAddress) methdoFullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type methodPerson struct {
	firstName string
	lastName  string
	methodAddress
}

func methodTest3() {
	p := methodPerson{
		firstName: "Elon",
		lastName:  "Musk",
		methodAddress: methodAddress{
			city:  "los angles",
			state: "california",
		},
	}
	p.methdoFullAddress()
}

type methodRectangle struct {
	length int
	width  int
}

func methodArea(r methodRectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r methodRectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

//在方法中使用值接收器 与 在函数中使用值参数
//这个话题很多Go语言新手都弄不明白。我会尽量讲清楚。
//当一个函数有一个值参数，它只能接受一个值参数。
//当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
func methodTest4() {
	r := methodRectangle{
		length: 10,
		width:  5,
	}
	methodArea(r)
	r.area()
	p := &r
	//methodArea(p)  //compilation error
	p.area()
}

//在方法中使用指针接收器 与 在函数中使用指针参数
//和值参数相类似，函数使用指针参数只接受指针，而使用指针接收器的方法可以使用值接收器和指针接收器。

func methodPerimeter(r *methodRectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (r *methodRectangle) methodPerimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func methodTest5() {
	r := methodRectangle{
		length: 10,
		width:  5,
	}
	p := &r
	methodPerimeter(p)
	p.methodPerimeter()

	//methodPerimeter(r)// cannot use r；只接受 指针类型的
	r.methodPerimeter() // 使用值调用指针类型的接收器
}

//在非结构体上的方法
//为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。到目前为止，我们定义的所有结构体和结构体上的方法都是在同一个 main 包中，因此它们是可以运行的。

type methodMyInt int

func (a methodMyInt) methodAdd(b methodMyInt) methodMyInt {
	return a + b
}

func methodTest6() {
	num1 := methodMyInt(5)
	num2 := methodMyInt(10)
	sum := num1.methodAdd(num2)
	fmt.Println("sum is ", sum)
}
