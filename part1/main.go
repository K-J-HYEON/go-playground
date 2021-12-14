package main

import (
	"fmt"
)

func addNum(slice []int) []int  {
	slice = append(slice, 4)
	return slice
}

func main() {
	slice := []int{1, 2, 3}
	slice = addNum(slice)

	fmt.Println(slice)
}



// func addNum(a, b, c int, /*output*/ slice *[]int) (int, bool, int)  {
// 	*slice = append(*slice, 4)
// }

// func main() {
// 	slice := []int{1, 2, 3}
// 	addNum(&slice)

// 	fmt.Println(slice)
// }