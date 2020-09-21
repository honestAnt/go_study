package main

import "fmt"

/*
https://studygolang.com/articles/12251
map 是在 Go 中将值（value）与键（key）关联的内置类型。通过相应的键可以获取到值。
通过向 make 函数传入键和值的类型，可以创建 map。make(map[type of key]type of value) 是创建 map 的语法。
*/
func main() {

}

func mapTest1() {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. going to make one.")
		personSalary = make(map[string]int)
	}
}

//给 map 添加元素
func mapTest2() {
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)
}

//声明并初始化
func mapTest3() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)
}

//获取map中的元素
//如果获取一个不存在的元素，会发生什么呢？map 会返回该元素类型的零值。在 personSalary 这个 map 里，如果我们获取一个不存在的元素，会返回 int 类型的零值 0。
// 判断map 中到底是不是存在这个 key: value, ok := map[key]；ok ==true 存在，false则不存在
func mapTest4(mapKey string) {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	employee := "jamie"
	value, ok := personSalary[mapKey]
	if ok == true {
		fmt.Println("salary of ", mapKey, " is ", value)
	} else {
		fmt.Println(mapKey, " is not found")
	}
	fmt.Println("salary of", employee, "is", personSalary[employee], "steve is ", personSalary["steve"])
}

//遍历 map 中所有的元素需要用 for range 循环。
func mapTest5() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
}

//删除 map 中 key 的语法是 delete(map, key)。这个函数没有返回值。
//获取 map 的长度使用 len 函数。
func mapTest6() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("map before deletion", personSalary, " length is ", len(personSalary))
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)
}

//Map 是引用类型
//和 slices 类似，map 也是引用类型。当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。因此，改变其中一个变量，就会影响到另一变量。
//Map 的相等性
//map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil。
func mapTest7() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000

	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)
	//if personSalary == newPersonSalary { //invalid operation: map1 == map2 (map can only be compared to nil)
	//	//
	//}
}
