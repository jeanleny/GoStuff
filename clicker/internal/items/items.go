package items

type Object interface{
	GetPrice()	float64	
	GetEarn()   float64
	GetAmount() int
	AddAmount(int) 
	Buy()
}

type ObjStats struct {
	Price	float64
	Earn 	float64
	Amount	int
}

type Worker struct {ObjStats}
type Factory struct {ObjStats}
type Company struct {ObjStats}

func (obj *ObjStats) GetEarn() float64 {return obj.Earn}
func (obj *ObjStats) GetPrice() float64 {return obj.Price}
func (obj *ObjStats) GetAmount() int {return obj.Amount}

func (b *ObjStats) Buy() {} //le truc cheloouuuuuu type func buy() = 0 en cPP
func (obj *Business) Work() {
	obj.Money++
}

func (obj *Business) Buy (name string) bool {
	choice := obj.Stock[name]
	price := obj.Stock[name].GetPrice()
	if (obj.Money < price) {
		return (false)
	}
	obj.Money -= price
	choice.AddAmount(1);
	obj.Earning += choice.GetEarn()
	return (true)
}

type Business struct {
	Stock map[string]Object
	Money float64
	Earning float64
}

func (obj *ObjStats) AddAmount(n int){
	obj.Amount += n
}
