//IF문(2)

package main

import "fmt"

func main() {
	var a int = 50
	b := 70

	//예제1
	if a >= 65 {
		fmt.Println("65이상")
	} else {
		fmt.Println("65미만")
	}

	//예제2
	if b >= 70 {
		fmt.Println("70이상")
	} else {
		fmt.Println("70미만")
	}

	//에러발생
	/*
			if b >= 70 {
				fmt.Println("70이상")
			} else
		   {
			   fmt.Println("70미만")
			 }
	*/

}
