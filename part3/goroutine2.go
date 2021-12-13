package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
	fmt.Println(<- sendch)
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)
}