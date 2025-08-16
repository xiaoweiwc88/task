package main

import "fmt"

//1. 定义Person结构体
type Person struct {
	Name string
	Age int
}


//2.定义Employee结构体，组合Person
type Employee struct {
	Person //匿名字段，实现组合
	EmployeeID  string
}

//3.为 Employee 实现PrintInfo() 方法
func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s\n", e.Name)   //直接访问组合的数组
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Emplyee ID: %s\n", e.EmployeeID)
}

func main() {
	//创建Employee实例
	e := Employee {
		Person:  Person{Name: "Alice", Age: 30},
		EmployeeID: "E1001",
	}

	//调用方法
	e.PrintInfo()
}
