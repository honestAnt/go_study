package main

import (
	"fmt"
	"github.com/xuri/excelize"
)

const filePath string = "src/data/wExcel.xlsx"
const sheetName string = "sheetName"

//参考 https://github.com/qax-os/excelize
func readExcel() {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := f.GetRows(sheetName)
	for idx_r, row := range rows {
		fmt.Println("-----------第 ", idx_r, " 行数据------------")
		for idx_c, cell := range row {
			fmt.Print("		第 ", idx_r, "行,第", idx_c, " 列数据：", cell)
			fmt.Print("|")
		}
		fmt.Println("")
	}
}

func writeExcel() {
	f := excelize.NewFile()
	sheetName := sheetName
	index := f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "name")
	f.SetCellValue(sheetName, "B1", "age")
	f.SetCellValue(sheetName, "C1", "address")
	f.SetActiveSheet(index)
	for i := 1; i < 100; i++ {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+1), fmt.Sprintf("name %d", i+1))
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+1), fmt.Sprintf("age %d", i+1))
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", i+1), fmt.Sprintf("address %d", i+1))
	}
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println(err)
	}
}

func main() {
	writeExcel()
	readExcel()
}
