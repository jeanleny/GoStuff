package items

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

func (obj *Business) Buy (name string) bool {
	choice := obj.Stock[name]
	price := obj.Stock[name].GetPrice()
	if (obj.Money < price) {
		return (false)
	}
	obj.Money -= price
	choice.AddAmount(1);
	return (true)
}

type Business struct {
	Stock map[string]Object
	Money int
}

func (obj *ObjStats) AddAmount(n int){
	obj.Amount += n
}
