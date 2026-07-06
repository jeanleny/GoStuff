package main

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
)

func display() tea.View{

	v := tea.NewView("Hello, World!")
	return (v)
}

type Object interface{
	getPrice()	int
	getEarn()   int
	getAmount() int
	Buy()
}

type ObjStats struct {
	price int
	earn int
	amount int
}

func (b *ObjStats) Buy() {} //le truc cheloouuuuuu type func buy() = 0 en cPP

func (obj *business) Buy (name string){
	choice, check := obj.stock[name]
	if !check {
		fmt.Println("Item doesn't exist")
	}
	//choice.amount++
	fmt.Println("item : ", name)
	fmt.Println("amount : ", choice.getAmount())
	fmt.Println("earn : ", choice.getEarn())
	fmt.Println("price : ", choice.getPrice())
	fmt.Println("BUYINGGGG")
}

func (obj *ObjStats) getEarn() int {return obj.earn}
func (obj *ObjStats) getPrice() int {return obj.price}
func (obj *ObjStats) getAmount() int {return obj.amount}

type Worker struct {ObjStats}
type Factory struct {ObjStats}
type Company struct {ObjStats}

type business struct {
	stock map[string]Object
	money int
}

type model struct {
	cursor 	int
	items	business
	actions []string
}

func (m model) View() tea.View {
	s := "Kessafou ?\n"

	for i, choice := range m.actions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s : %d\n", cursor, choice, m.items.stock[choice].getAmount())
	}
	return (tea.NewView(s))
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg:= msg.(type){
		case tea.KeyPressMsg :
			switch msg.String() {
				case "ctrl+c" :
					return m, tea.Quit
				case "up" :
					if m.cursor > 0 {
						m.cursor--
					}
				case "down" :
					if m.cursor < len(m.actions) - 1 {
						m.cursor++
					}
				case "enter" :
					m.items.Buy(m.actions[m.cursor])
			}
	}

	return m, nil
}

func makeModel() model {
	return model {
		actions: []string {"worker", "factory", "company"},
		cursor : 0,
		items : business{
			money : 100,
			stock : map[string]Object {
				"worker": &Worker{ObjStats{price : 0, earn : 0, amount : 0}},
				"factory": &Factory{ObjStats{price : 0, earn : 0, amount : 0}},
				"company": &Company{ObjStats{price : 0, earn : 0, amount : 0}},
			},
		},
	}
}

func main ()  {
	m:= makeModel()
	p := tea.NewProgram(m)
	p.Run()
	if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
