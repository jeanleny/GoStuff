package main

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
)

var Stock = business{}

type business struct {
	stock 	[]Object
	w		int
	f		int
	c		int
	m		int
}

func (b *business) Add(o Object) {
	//o.Buy()
	b.stock = append(b.stock, o);
	switch o.(type) {
		case  Worker :
			b.w++;
		case  Factory :
			b.f++;
		case  Company :
			b.c++;
	}
}

type Object interface {
	Buy()
}

type Worker struct {
	price int
	earn int
}

type Factory struct {
	price int
	earn int
}

type Company struct {
	price int
	earn int
}

func (w Worker) Buy() {
	fmt.Printf("buy a worker for %d\n", w.price)
}

func (f Factory) Buy() {
	fmt.Println("buy a factory")
}

func (f Company) Buy() {
	fmt.Println("buy a company")
}

func display() tea.View{

	v := tea.NewView("Hello, World!")
	return (v)
}

func newWorker () Worker{
	return Worker{price : 10, earn : 1}
}

func newFactory () Factory{
	return Factory{price : 100, earn : 10}
}

func newCompany () Company{
	return Company{price : 500, earn : 50}
}

type model struct {
	cursor 	int
	actions []string
}

func getStock(choice string) int {
	switch choice {
	case "Worker" :
		return Stock.w
	case "Factory" :
		return Stock.f
	case "Company" :
		return Stock.c
	}
	return (0);
}

func (m model) View() tea.View {
	s := "Kessafou ?\n"

	for i, choice := range m.actions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s : %d\n", cursor, choice, getStock(choice))
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
					switch m.cursor {
						case 1 :
							Stock.Add(newWorker())
						case 2 :
							Stock.Add(newFactory())
						case 3 :
							Stock.Add(newCompany())
						}

			}
	}

	return m, nil
}

func makeModel() model {
	return model {
		actions : []string {"Click", "Worker", "Factory", "Company"},
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
