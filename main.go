package main

import "fmt"

func main() {
	fmt.Printf("%s\n", "Hello World!!")
	// 포맷 지정자를 string으로 정하고 개행문자를 포함했습니다.
	// %s에 들어갈 값으로 "Hello World!!"를 넣습니다.
	
	fmt.Printf("Hello World!!\n")
	// 출력할 포맷을 제외하고 문자열 출력만 하도록 간소화했습니다.
	// Printf는 기본 string 타입을 받기 때문에 string 타입은 포맷 지정 없이도 사용이 가능합니다.
	
	var s string = "hi\n"
	fmt.Printf(s)
	// hi를 출력하고 개행을 합니다.
	// 단, 변수를 포맷 지정 없이 Printf로 출력하는 것은 권장되지 않습니다. (go-staticcheck)
	// 에러는 발생하지 않으며 정상 출력됩니다. 다만 가능할 뿐, 이런 사용 방식은 지양합니다.
	
	xByte, err := fmt.Printf("Hello World!!\n")
	fmt.Printf("The value of xByte : %d\nError : %v\n", xByte, err)
	// The value of xByte : 14
	// Error : <nil>
	// string은 글자 하나가 8bit, 1byte로 저장되기 때문에 "Hello World!!"에 대한 Byte 수는 13이 됩니다.
	// \n 개행 문자도 1byte입니다.
}