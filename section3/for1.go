package main

import "fmt"

func main() {
	//반복문 - for
	//Go에서 유일하게 제공되는 반복문
	//다양한 사용법 숙지

	//예제1
	for i := 0; i < 5; i++ {
		fmt.Println("ex : ", i)
	}

	//에러 발생1
	/*
	  for i := 0; i < 5; i++
	  {
	    fmt.Println("ex : ", i)
	  }
	*/

	//에러 발생2
	/*
	  for i := 0; i < 5; i++
	    fmt.Println("ex : ", i)
	*/

	//예제2 (무한 루프)
	/*
	  for {
	    fmt.Println("ex2 : Hello Golang!")
	    fmt.Println("ex2 : Infinite Loop!")
	  }
	*/

	//예제3 (Range용법)
	loc := []string{"seoul", "busan", "daejeon"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
}
