package main

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
)

type Items struct {
	nb int
	earn int
}

func display() tea.View{

	v := tea.NewView("Hello, World!")
	return (v)
}

type model struct {
	cursor 	int
	actions []string
}

func (m model) View() tea.View {
	s := "Kessafou ?\n"

	for i, choice := range m.actions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
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
