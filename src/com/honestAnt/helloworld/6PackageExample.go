package main

import (
	rectangle "com/honestAnt/helloworld/packageExample" //导入自定义包
	//_ "com/honestAnt/helloworld/packageExample" //屏蔽不需要在代码中使用的导入
	"fmt"
)

/*
https://studygolang.com/articles/11893
到目前为止，我们看到的 Go 程序都只有一个文件，文件里包含一个 main 函数和几个其他的函数。在实际中，这种把所有源代码编写在一个文件的方法并不好用。
以这种方式编写，代码的重用和维护都会很困难。而包（Package）解决了这样的问题。

package packagename 这行代码指定了某一源文件属于一个包。它应该放在每一个源文件的第一行。

下面开始为我们的程序创建一个 main 函数和 main 包。在 Go 工作区内的 src 文件夹中创建一个文件夹，命名为 helloworld。
在 helloworld 文件夹中创建一个 6PackageExampl.go 文件。


导入了包，却不在代码中使用它，这在 Go 中是非法的。当这么做时，编译器是会报错的。其原因是为了避免导入过多未使用的包，从而导致编译时间显著增加
*/

func main() {
	fmt.Println("hello world")
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("area of rectangle %.2f", rectangle.Area(rectLen, rectWidth))
	fmt.Println("diagonal of the rectangle %.2f", rectangle.Diagonal(rectLen, rectWidth))
	//fmt.Println("diagonal of the rectangle %.2f", rectangle.test(rectLen, rectWidth)) error
}

/*
所有包都可以包含一个 init 函数。init 函数不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它。init 函数的形式如下：
*/
func init() {
	fmt.Println("package main initialized")
}
