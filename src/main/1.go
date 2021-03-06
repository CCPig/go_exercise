package main

import "fmt"

//this is my first go program

type student struct {
	name string
	age  int
}

func main() {
	//打印语句
	fmt.Println("hello world!")
	fmt.Println("hello" + " world!")

	//变量的使用
	var stockcode = "000001.CS"
	var enddate = "2021-03-06"
	var url = "Code=%s|Enddate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	// ：=变量初始化,而不是用于变量赋值，会自动识别类型
	stu := student{age: 24, name: "chaochen"}
	fmt.Println(stu)

	var stu1 student
	stu1 = student{age: 15, name: "kaichen"}
	fmt.Println(stu1)

	stu2 := 10
	fmt.Println(stu2)

	const a, b, c = 1, 2, "haha"
	fmt.Println(a, b, c)
}
