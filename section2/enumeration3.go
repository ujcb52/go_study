//열거형3
package main

import "fmt"

func main() {
	const (
		DEFAULT = iota
		SILVER
		_ //GOLD
		PLATINUM
	)

	fmt.Println("D : ", DEFAULT)
	fmt.Println("S : ", SILVER)
	//	fmt.Println("G : ", GOLD)
	fmt.Println("P : ", PLATINUM)
}
