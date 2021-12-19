package main

import "goprojects/ch20/koreaPost"

type Sender interface {
	Send(parcel string)
}





func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var sender Sender = &koreaPost.PostSender{}
	SendBook("나의 라임 오렌지 나무", sender)
	SendBook("마법천자문", sender)
}