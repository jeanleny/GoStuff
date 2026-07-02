package main

import (
	"fmt"
	"test/lib/menu"
)

func main (){
	midi:= menu.Menu{ Starter : "coleslaw",	MainCourse :"poulet basquaise", Dessert : "tiramisu"}
	for  i := 0; i < 4; i++ {
		fmt.Printf("result : %d\n", i)
	}
	fmt.Printf("element :%s\n", midi.Starter)
	fmt.Printf("element :%s\n", midi.MainCourse)
	fmt.Printf("element :%s\n", midi.Dessert)
	soir := menu.Create("bite", "boule", "docteur")
	fmt.Printf("element :%s\n", soir.Starter)
	fmt.Printf("element :%s\n", soir.MainCourse)
	fmt.Printf("element :%s\n", soir.Dessert)
}
