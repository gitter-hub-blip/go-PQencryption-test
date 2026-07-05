// 每个可执行的 Go 程序都必须属于 main 包（类比 C 语言的 main 函数所在文件）。
package main

// import 引入标准库，类似 Python 的 import 或 C 的 #include。
// Go 要求：引入了就必须使用，否则编译报错。
import (
	"bufio" // 带缓冲的 I/O，用来按行读取输入
	"fmt"   // 格式化输出，类似 C 的 printf 家族
	"os"    // 操作系统接口，这里用到标准输入 os.Stdin
)

// 程序入口。注意 Go 的函数没有分号，左花括号必须和 func 同一行。
func main() {
	fmt.Print("请输入一行内容：")

	// bufio.NewScanner 创建一个按行扫描器，包住标准输入。
	// := 是"声明并赋值"的简写，类型由右边自动推断（类似 Python 的动态赋值，
	// 但 Go 是静态类型，scanner 的类型在编译期就确定了）。
	scanner := bufio.NewScanner(os.Stdin)

	// Scan() 读取一行，成功返回 true。相当于 Python 的 input()，
	// 但读到的内容要再调用 Text() 才能拿到（不含换行符）。
	if scanner.Scan() {
		line := scanner.Text()
		// Println 自动在末尾加换行，类似 Python 的 print()。
		fmt.Println("你输入的是：", line)
	}
}
