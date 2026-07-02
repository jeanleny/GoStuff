package main

import (
	"fmt"
)

type Fruit interface {
	Print()
}

type apple struct {
	seed int
}

func (a apple) Print() {
	fmt.Printf("apple with %d seeds\n", a.seed)
}


type banana struct {
	peelType string
}

func (b *banana) Print() {
	fmt.Printf("banana with a %s peel type\n", b.peelType)
}


type FruitType = int

const (
	AppleType FruitType = iota
	BananaType 
)

var FruitList = map[FruitType]func() Fruit {
	AppleType : func() Fruit { return apple{ seed: 3} },
	BananaType: func() Fruit { return &banana{ peelType : "roten perfect for a banana bread" } },
}


func NewFruit(FT FruitType) (Fruit, error) {
	name, ok := FruitList[FT]
	if !ok {
		return nil, fmt.Errorf("wrong FruitType %d\n", FT)
	}

	return name(), nil
}

func guessType(michel Fruit){
	switch michel.(type) {
	case apple : 
		fmt.Println("ta mer la pom")
		return 
	case *banana :
		fmt.Println("ton per la banan")
	}
}

func main() {

	f, err := NewFruit(FruitType(12))
	if err != nil {
		fmt.Println(err)
	} else {
		f.Print()
	}

	f, err = NewFruit(BananaType)
	if err != nil {
		fmt.Println(err)
	} else {
		f.Print()
		guessType(f)
	}

}
