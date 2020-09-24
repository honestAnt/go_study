package main

import "fmt"

/*
https://studygolang.com/articles/12263
构体是用户定义的类型，表示若干个字段（Field）的集合。有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。
type Employee struct {
    firstName string
    lastName  string
    age       int
}

声明结构体时也可以不用声明一个新类型，这样的结构体类型称为 匿名结构体（Anonymous Structure）。
var employee struct {
    firstName, lastName string
    age int
}
*/

func main() {
	structTest1()
	structTest2()
	structTest3()
	structTest4()
	structTest5()
	structTest6()
	structTest7()
	structTest8()
	structTest9()
	structTest10()
}

type StructExample struct {
	firstName, lastName string
	age, salary         int
}

func structTest1() {
	emp1 := StructExample{
		firstName: "sam",
		age:       25,
		lastName:  "Anderson",
		salary:    500,
	}

	//creating structure without using field names;必须要保证顺序
	emp2 := StructExample{"Thomas", "Paul", 29, 800}
	fmt.Println("Emample1 is ", emp1)
	fmt.Println("Emample2 is ", emp2)
}

//创建匿名结构体
func structTest2() {
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("employee 3 ", emp3)
}

//结构体的零值（Zero Value）
//当定义好的结构体并没有被显式地初始化时，该结构体的字段将默认赋为零值。
//当然还可以为某些字段指定初始值，而忽略其他字段。这样，忽略的字段名会赋值为零值。
func structTest3() {
	var emp4 StructExample
	emp5 := StructExample{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("emp4 ", emp4)
	fmt.Println("emp5 ", emp5)
}

//访问结构体的字段
//点号操作符 . 用于访问结构体的字段。
func structTest4() {
	emp6 := StructExample{"sam", "anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)
}

//还可以创建零值的 struct，以后再给各个字段赋值。
func structTest5() {
	var emp7 StructExample
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7:", emp7)
}

//结构体的指针
//还可以创建指向结构体的指针。
//Go 语言允许我们在访问 firstName 字段时，可以使用 emp8.firstName 来代替显式的解引用 (*emp8).firstName。
func structTest6() {
	emp8 := &StructExample{"sam", "anderson", 55, 6000}
	fmt.Println("address First Name:", emp8.firstName)
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("address Last Name:", emp8.lastName)
	fmt.Println("Last Name:", (*emp8).lastName)
}

//虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型。比如在上面的 Person 结构体里，虽说字段是匿名的，
// 但 Go 默认这些字段名是它们各自的类型。所以 Person 结构体有两个名为 string 和 int 的字段。
type StructPerson struct {
	string
	int
}

//匿名字段
//当我们创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段（Anonymous Field）。
func structTest7() {
	p := StructPerson{"naveen", 50}
	fmt.Println("p", p)
	var p1 StructPerson
	p1.string = "naveen1"
	p1.int = 51
	fmt.Println("p1", p1)
}

//嵌套结构体（Nested Structs）
//结构体的字段有可能也是一个结构体。这样的结构体称为嵌套结构体。
type StructAddress struct {
	city, state string
}
type StructPerson2 struct {
	name    string
	age     int
	address StructAddress
}

func structTest8() {
	var p StructPerson2
	p.name = "Naveen"
	p.age = 50
	p.address = StructAddress{
		city:  "shanghai/china",
		state: "nanjingdonglu",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
}

//提升字段（Promoted Fields）
//如果是结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段。这是因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问
type StructPerson3 struct {
	name          string
	age           int
	StructAddress //注意此处为匿名
}

func structTest9() {
	var p StructPerson3
	p.name = "Naveen"
	p.age = 50
	p.StructAddress = StructAddress{
		city:  "shanghai/china",
		state: "nanjingdonglu",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)
	fmt.Println("promoted City:", p.StructAddress.city)
	fmt.Println("State:", p.state)
	fmt.Println("promoted State:", p.StructAddress.state)
}

//导出结构体和字段
//如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型（Exported Type）。同样，如果结构体里的字段首字母大写，它也能被其他包访问到。
/*
package computer

type Spec struct { //exported struct
    Maker string //exported field
    model string //unexported field,其他文件无法访问到
    Price int //exported field
}

*/

//结构体相等性（Structs Equality）
//结构体是值类型。如果它的每一个字段都是可比较的，则该结构体也是可比较的。如果两个结构体变量的对应字段相等，则这两个变量也是相等的。
type StructName struct {
	firstName string
	lastName  string
}

func structTest10() {
	name1 := StructName{"steve", "jobs"}
	name2 := StructName{"steve", "jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name2 and name2 are not equal")
	}

	name3 := StructName{"steve", "jobs"}
	name4 := StructName{"steve", "jobs2"}
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}

//如果结构体包含不可比较的字段，则结构体变量也不可比较。
type structImage struct {
	data map[int]int //map类型不能直接用==比较
}

func structTest11() {
	/*image1 := structImage{data: map[int]int{
		0:155,
	}}
	image2 := structImage{data: map[int]int{
		0:155,
	}}
	if image1 == image2 { // invalid operation: image1 == image2 (struct containing map[int]int
		fmt.Println("iamge1 and image2 are equal")
	}*/
}
