package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize"
	"io/ioutil"
	"os"
)

//https://draveness.me/golang/docs/part4-advanced/ch09-stdlib/golang-json/

type Persons struct {
	Persons []Emp `json:"person"`
}

type Emp struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Email string `json:"cpsnemail"`
}

func convert2Json() {
	emp := Emp{
		Name: "test11",
		Code: "011",
	}
	jsonEmp, err := json.Marshal(emp)
	if err != nil {
		fmt.Println("生成json数据错误")
	}
	fmt.Println(string(jsonEmp))

}

func parseJson() {
	jsonFile, err := os.Open("src/data/test.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var persons Persons
	json.Unmarshal(byteValue, &persons)

	//for i := 0; i< len(persons.Persons); i++ {
	//	fmt.Println("name: " + persons.Persons[i].Name)
	//	fmt.Println("code: " + persons.Persons[i].Code)
	//}
	write2Excel(persons)
}

func write2Excel(persons Persons) {
	f := excelize.NewFile()
	sheetName := "Sheet2"
	index := f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "name")
	f.SetCellValue(sheetName, "B1", "code")
	f.SetCellValue(sheetName, "C1", "email")

	for i, v := range persons.Persons {
		axis := fmt.Sprintf("A%d", i+2)
		axis2 := fmt.Sprintf("B%d", i+2)
		axis3 := fmt.Sprintf("C%d", i+2)
		fmt.Println("第%s 行 name  %s，code %s, email %s", axis, v.Name, v.Code, v.Email)
		f.SetCellValue(sheetName, axis, v.Name)
		f.SetCellValue(sheetName, axis2, v.Code)
		f.SetCellValue(sheetName, axis3, v.Email)
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs("src/data/person.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("开始转换json")
	convert2Json()
	fmt.Println("结束转换json")
	parseJson()
}
