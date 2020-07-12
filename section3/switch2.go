//
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i < 100 && i >= 50:
		fmt.Println("i -> ", i, "i는 50이상 100미만")
	case i < 50 && i >= 25:
		fmt.Println("i -> ", i, "i는 25이상 50 미만")
	default:
		fmt.Println("i -> ", i, "기본 값")
	}
}
