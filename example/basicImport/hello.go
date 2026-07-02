package main

import (
	"fmt"
	"test/message"
)

func main(){
	var result string
	result = message.Text("bite")
	fmt.Println(result)
	test:= message.Sex()
	fmt.Println(test)
	fmt.Println(message.MeltingPot())
}
