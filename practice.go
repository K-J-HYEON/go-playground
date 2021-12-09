package main // Hello World!!를 입력 및 출력합니다.

import "fmt"

func main() {
	var x string
	var y string
	fmt.Scan(&x, &y) // string 두 개의 값을 입력 받아, 각각 x와 y에 저장합니다.
	// Hello World!! 
	// 또는
	// Hello
	// World!!
	// z, _ := fmt.Scan(&x, &y) // 인자의 개수를 z에 초기화합니다.
	fmt.Println(x, "and", y)
	// Hello and World!!
	// fmt.Println(z)
}