package main

import (
	"fmt"
	"os"
	items "clicker/internal/items"
	tea "charm.land/bubbletea/v2"
)

func display() tea.View{

	v := tea.NewView("Hello, World!")
	return (v)
}

type model struct {
	cursor 	int
	items	items.Business
	actions []string
}

func (m model) View() tea.View {
	s := "Kessafou ?\n"

	for i, choice := range m.actions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s : %d\n", cursor, choice, m.items.Stock[choice].GetAmount())
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
		items : items.Business{
			Money : 100,
			Stock : map[string]items.Object {
				"worker":   &items.Worker{items.ObjStats{Price : 0, Earn : 0, Amount : 0}},
				"factory": &items.Factory{items.ObjStats{Price : 0, Earn : 0, Amount : 0}},
				"company": &items.Company{items.ObjStats{Price : 0, Earn : 0, Amount : 0}},
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
