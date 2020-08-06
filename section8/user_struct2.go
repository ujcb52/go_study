//사용자 정의 타입(2)

package main

import "fmt"

type cnt int

func main() {
	//기본 자료형 사용자 정의 타입
	//예제1

	a := cnt(15)
	fmt.Println("ex1 : ", a)

	//예제2

	var b cnt = 15
	fmt.Println("ex1 : ", b)

	//예제3
	//testConvertT(b) //예외 발생(중요)! 사용자 정의 타입 <-> 기본 타입 : 매개 변수 전달 시에 변환해야 사용 가능(int(변수))
	testConvertT(int(b)) //형 변환

	testConvertD(b)
}

func testConvertT(i int) {
	fmt.Println("ex3 : (Default Type) : ", i)
}

func testConvertD(i cnt) {
	fmt.Println("ex4 : (Custom Type) : ", i)
}
