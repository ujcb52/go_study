//반복문2

package main

import "fmt"

func main() {
	//예제1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println("ex1 : ", sum1)

	//예제2
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		// j := i++ //Go에서 후치연산은 반환 값 X
	}
	fmt.Println("ex2 : ", sum2)

	//예제3
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3 : ", sum3)
}
