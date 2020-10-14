package main

import (
	"fmt"
)

/*
https://studygolang.com/articles/12325
上的所有示例中，我们都是使用值接受者（Value Receiver）来实现接口的。我们同样可以使用指针接受者（Pointer Receiver）来实现接口。
*/

func main() {
	inter2Test1()
}

type Inter2Describer interface {
	Inter2Describe()
}

type Inter2Person struct {
	name string
	age  int
}

func (p Inter2Person) Inter2Describe() { // 使用值接受者实现；也可以接收指针
	fmt.Printf("%s is %d years old \n", p.name, p.age)
}

type Inter2Address struct {
	state   string
	country string
}

func (a *Inter2Address) Inter2Describe() { // 使用指针接受者实现；仅接收指针
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func inter2Test1() {
	var d1 Inter2Describer
	p1 := Inter2Person{"sam", 25}
	d1 = p1
	d1.Inter2Describe()
	p2 := Inter2Person{"james", 32}
	d1 = &p2
	d1.Inter2Describe()

	var d2 Inter2Describer
	a := Inter2Address{"washington", "USA"}

	//d2 = a  非法，因为是指针实现
	d2 = &a
	d2.Inter2Describe()
}

//实现多个接口

type Inter2SalaryCalculator interface {
	Inter2DisplaySalary()
}

type Inter2LeaveCalculator interface {
	Inter2CalculateLeavesLeft() int
}

type Inter2Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Inter2Employee) Inter2DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Inter2Employee) Inter2CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func inter2Test2() {
	e := Inter2Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s Inter2SalaryCalculator = e
	s.Inter2DisplaySalary()
	var l Inter2LeaveCalculator = e
	fmt.Println("\nLeaves left =", l.Inter2CalculateLeavesLeft())
}

//接口的嵌套
type Inter2EmployeeOperations interface {
	Inter2SalaryCalculator
	Inter2LeaveCalculator
}

func inter2Test3() {
	e := Inter2Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp Inter2EmployeeOperations = e
	empOp.Inter2DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.Inter2CalculateLeavesLeft())
}
