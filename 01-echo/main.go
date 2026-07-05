// 每个可执行的 Go 程序都必须属于 main 包（类比 C 语言的 main 函数所在文件）。
package main

// import 引入标准库，类似 Python 的 import 或 C 的 #include。
// Go 要求：引入了就必须使用，否则编译报错。
import (
	"bufio"   // 带缓冲的 I/O，用来按行读取输入
	"fmt"     // 格式化输出，类似 C 的 printf 家族
	"os"      // 操作系统接口，这里用到标准输入 os.Stdin
	"strconv" // 字符串转换
)

// 程序入口。注意 Go 的函数没有分号，左花括号必须和 func 同一行。
func main() {

	// bufio.NewScanner 创建一个按行扫描器，包住标准输入。
	// := 是"声明并赋值"的简写，类型由右边自动推断（类似 Python 的动态赋值，
	// 但 Go 是静态类型，scanner 的类型在编译期就确定了）。
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("请输入一行内容：")
	scanner.Scan()
	line := scanner.Text()

	fmt.Print("请输入偏移量：")
	scanner.Scan()
	offset := scanner.Text()
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		fmt.Printf("非法字符\n")
		return
	}

	// Scan() 读取一行，成功返回 true。相当于 Python 的 input()，
	// 但读到的内容要再调用 Text() 才能拿到（不含换行符）。
	if 1 == 1 {

		final := []byte(line)
		for i := 0; i < len(line); i++ {
			// Println 自动在末尾加换行，类似 Python 的 print()。
			currentByte := int(line[i])

			if currentByte >= 65 && currentByte <= 90 {

				if currentByte+offsetInt > 90 {
					final[i] = byte(int(line[i]) + offsetInt - 26)
					continue
				}
				final[i] = byte(int(line[i]) + offsetInt)

			} else if currentByte >= 97 && currentByte <= 122 {

				if currentByte+offsetInt > 122 {
					final[i] = byte(int(line[i]) + offsetInt - 26)
					continue
				}
				final[i] = byte(int(line[i]) + offsetInt)

			} else {
				final[i] = byte(int(line[i]))
			}
		}

		fmt.Println("你输入的是：", string(final))
	}
}
