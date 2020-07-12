//Go 특징1

package main

import "fmt"

func main() {
	// Go:  모호하거나 애메한 문법 금지
	//후치(후위) 연산자 허용(i++), 전치(전위) 연산자 비허용 (++i) -> x
	//증감연산 변환 값 없음.
	//포인터(Pointer 허용, 연산 비허용)
	//주석 (//, /**/) 사용법 숙지

	//예제1
	sum, i := 0, 0

	for i <= 100 {
		// sum += i++ // 예외 발생
		sum += i
		i++
	}
	fmt.Println("ex1 : ", sum)
}
