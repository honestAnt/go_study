package packageExample

import (
	"fmt"
	"math"
)

/*
我们将 rectangle 包中的函数 Area 和 Diagonal 首字母大写。在 Go 中这具有特殊意义。
在 Go 中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量。
在这里，我们需要在 main 包中访问 Area 和 Diagonal 函数，因此会将它们的首字母大写。
*/
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt(len*len) + (wid * wid)
	return diagonal
	//return math.Sqrt(len * len) +  (wid * wid)
}

/*
仅内部可访问，外部无法访问的函数（名字开头是小写的）
*/
func test(len, wid float64) float64 {
	diagonal := math.Sqrt(len*len) + (wid * wid)
	return diagonal
}

func init() {
	fmt.Println("rectangle package initialized")
}
