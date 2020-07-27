package main

import "fmt"

func main() {
	//예제1
	cnt := increaseCnt()

	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
}

func increaseCnt() func() int {
	n := 0 //지역변수(켑쳐됌)
	return func() int {
		n += 1
		return n
	}
}
