package main

import (
	"fmt"
	"math"
)

//1.定义Shap接口
type Shape interface {
	Area() float64
	Perimeter() float64
}


//2.定义Rectangle的结构体
type Rectangle struct {
	Width, Height float64
}

//实现Shap接口的Area方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//4.实现Shap接口的Perimeter方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

//3.定义Circle结构体
type Circle struct {
	Radius float64
}

//实现 Shape 实现Shap接口的Area方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//实现 实现Shap接口的Perimeter方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

//4. 主函数测试
func main() {
	var s Shape

	//矩形
	r := Rectangle{Width: 5, Height: 3}
	s = r
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())


	//圆形
	c := Circle{Radius: 4}
	s = c
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())
}
