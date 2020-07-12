//함수 기초(2)

package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) {
	fmt.Println("ex1 : ", a+b)
}

func main() {
	//함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)
	//예제1
	sum(100, add) //함수 전달
}
