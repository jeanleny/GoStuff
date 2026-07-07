package items

import (
	"fmt"
)

type Object interface{
	GetPrice()	int
	GetEarn()   int
	GetAmount() int
	AddAmount(int) 
	Buy()
}

type ObjStats struct {
	Price int
	Earn int
	Amount int
}

type Worker struct {ObjStats}
type Factory struct {ObjStats}
type Company struct {ObjStats}

func (obj *ObjStats) GetEarn() int {return obj.Earn}
func (obj *ObjStats) GetPrice() int {return obj.Price}
func (obj *ObjStats) GetAmount() int {return obj.Amount}

func (b *ObjStats) Buy() {} //le truc cheloouuuuuu type func buy() = 0 en cPP

func (obj *Business) Buy (name string){
	choice, check := obj.Stock[name]
	if !check {
		fmt.Println("Item doesn't exist")
	}
	choice.AddAmount(1);
}

type Business struct {
	Stock map[string]Object
	Money int
}

func (obj *ObjStats) AddAmount(n int){
	obj.Amount += n
}
