package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var in string
	_, _ = fmt.Scanln(&in)
	fmt.Println("选择你要实现的功能:按1查找字符位置，按2替换字符")
	var sle int
	_, _ = fmt.Scanln(&sle)
	switch sle {
	case 1: //查找字符串位置
		fmt.Println("请输入你要查找的字符")
		var temp string
		_, _ = fmt.Scanln(&temp)
		search(in, temp)
	case 2: //替换字符串
		fmt.Println("请输入查找内容")
		var a string
		_, _ = fmt.Scanln(&a)
		fmt.Println("替换为:")
		var b string
		_, _ = fmt.Scanln(&b)
		fmt.Println("请输入要替换的个数n？（n为负数时替换全部）")
		var temp int
		_, _ = fmt.Scanln(&temp)
		replace(in, a, b, temp)
	}
}
func search(in string, temp string) { //查找模块
	a := strings.Index(in, temp)
	fmt.Println(temp + "内容首次出现在位置" + strconv.Itoa(a))
}
func replace(in string, a string, b string, temp int) { //替换模块
	result := strings.Replace(in, a, b, temp)
	fmt.Println(result)
}
