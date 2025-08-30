package main

import "fmt"

// 题目一
func addNumber(x *int) {
	*x = *x + 10
}

// 题目二
func sliceMulti(b *[]int) {
	for i := 0; i < len(*b); i++ {
		(*b)[i] = (*b)[i] * 2
	}
}

func main() {
	// a := 0
	// addNumber(&a)
	// fmt.Println(a)

	b := []int{1, 2, 3, 4, 5}
	sliceMulti(&b)
	fmt.Println(b)

}
