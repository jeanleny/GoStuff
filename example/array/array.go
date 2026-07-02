package main

import (
	"fmt"
	"array/lib/arrayFunc"
)

func main (){
	var a[5] int
	b:= make ([]int, 8)
	fmt.Printf("le gran b la sous make %x\n", b)
	fmt.Print("-----------------Int iterator---------------------\n")
	for i := 0; i < len(a); i++ {
		fmt.Printf("elem nb %d : %d\n",i, a[i] );
	}
		
	fmt.Print("-----------------In range---------------------\n")
	for i, y := range a {
		fmt.Printf("elem nb %d : %d\n",i , y);
	}
	fmt.Print("-----------------random fills---------------------\n")
	array.RandomFill(b)
	fmt.Println(b)
	c := make([]int, 18)
	array.RangeFill(c, 10)
	fmt.Println(c)
	c = append(c, b[4])
	fmt.Println(c)
	d := make([]int, len(b))
	copy(d, b)
	fmt.Println("print the copied array", d);
	fmt.Print("----------------- maps---------------------\n")
	m := make (map[string]int);
	m["1st"] = 4;
	m["2nd"] = 8;
	fmt.Println(m["1st"])
	fmt.Println(m["2nd"])
	delete(m, "1st")
	fmt.Println(m["1st"])
}
