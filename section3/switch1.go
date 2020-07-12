//Switch문(1)
package main

import "fmt"

func main() {
	//제어문(조건문) - switch
	//Switch 뒤 표현식(Expression) 생략 가능
	//case 뒤 표현식(Expression) 사용 가능
	//자동 break 때문에 failthrouth whswo
	//Type 분기 -> 값이 아닌 변수 Type으로 분기 가

	//예제1
	a := -7
	switch {
	case a < 0:
		fmt.Println("a는 음수")
	case a == 0:
		fmt.Println("a는 0")
	case a > 0:
		fmt.Println("a는 양수")
	}

	//예제2
	switch b := 0; {
	case b < 0:
		fmt.Println("b는 음수")
	case b == 0:
		fmt.Println("b는 0")
	case a > 0:
		fmt.Println("a는 양수")
	}

	//예제3
	switch c := "go"; c {
	case "go":
		fmt.Println("GO!")
	case "java":
		fmt.Println("JAVA!")
	default:
		fmt.Println("일치하는 값 없음!")
	}

	//예제4
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("golang!")
	case "java":
		fmt.Println("java!")
	default:
		fmt.Println("None")
	}

	//예제5
	switch i, j := 20, 30; {
	case i > j:
		fmt.Println("i는 j보다 크다")
	case i == j:
		fmt.Println("i와 j는 같다")
	case i < j:
		fmt.Println("i는 j보다 작다")
	}
}
