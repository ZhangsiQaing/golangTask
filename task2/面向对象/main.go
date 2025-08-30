package main

import "fmt"

//题目一

// type Shape interface {
// 	Area()
// 	Perimeter()
// }

// type Rectangle struct {
// 	x float64
// 	y float64
// }

// func (t Rectangle) Area() {
// 	fmt.Println(t.x * t.y)
// }
// func (t Rectangle) Perimeter() {
// 	fmt.Println(2 * (t.x + t.y))
// }

// type Circle struct {
// 	r float64
// }

// func (t Circle) Area() {
// 	fmt.Println(math.Pi * float64(t.r) * float64(t.r))
// }
// func (t Circle) Perimeter() {
// 	fmt.Println(2 * math.Pi * math.Pi)
// }

// 题目二
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	P          Person
	EmployeeId int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("个人信息：%v，雇员id：%d", e.P, e.EmployeeId)
}

func main() {
	// rectangle := Rectangle{
	// 	x: 2,
	// 	y: 3,
	// }

	// circle := Circle{
	// 	r: 2.0,
	// }

	// shape := []Shape{rectangle, circle}
	// for _, v := range shape {
	// 	v.Area()
	// 	v.Perimeter()
	// }

	em := Employee{
		P:          Person{Age: 12, Name: "zsq"},
		EmployeeId: 12,
	}
	em.PrintInfo()

}
