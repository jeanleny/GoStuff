package menu

import "fmt"

type Menu struct{
	Starter		string
	MainCourse	string
	Dessert		string
}

type Fruit interface {
	Print()
}

type baseFruit struct {
	seed int
}

type apple struct {

}

func (baseFruit apple) Print() {
	fmt.Printf("fruit with %d seeds\n", a.seed)

}

func (a apple) Print() {
	fmt.Printf("apple with %d seeds\n", a.seed)
}


func Create (s,m,d string) Menu {
	return Menu{Starter : s, MainCourse : m, Dessert : d}
}
