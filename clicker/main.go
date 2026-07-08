package main

import (
	"fmt"
	"os"
	items "clicker/internal/items"
	tea "charm.land/bubbletea/v2"
	render "clicker/internal/render"
	lipgloss "charm.land/lipgloss/v2"
)


func display() tea.View{

	v := tea.NewView("Hello, World!")
	return (v)
}

type model struct {
	cursor 	int
	items	items.Business
	width	int
	height	int
	actions []string
	ErrText string
}

func (m model) View() tea.View {
	s := []string {"Kessafou ?\n"}
	err := ""
	for i, choice := range m.actions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		line := fmt.Sprintf("%s %s : %d %s\n", cursor, choice, m.items.Stock[choice].GetAmount(), err)
		lineBlock := render.Menu.Render(line)
		if m.cursor == i && m.ErrText != "" {
			errBlock := render.ErrUi.Render(m.ErrText)
			lineBlock = lipgloss.JoinHorizontal(lipgloss.Top, lineBlock, errBlock)
		}
		s = append(s, lineBlock)
	}
		menuBlock := lipgloss.JoinVertical(lipgloss.Left, s...)
		moneyBlock := render.Money.Render(fmt.Sprintf("$$$ : %d\n", m.items.Money))
		moneyPlaced := lipgloss.Place(m.width - lipgloss.Width(menuBlock) ,lipgloss.Height(menuBlock), lipgloss.Right, lipgloss.Top, moneyBlock)
		full := lipgloss.JoinHorizontal(lipgloss.Top, menuBlock, moneyPlaced)
		return tea.NewView(full)
	}	

func (m model) Init() tea.Cmd {
	return nil
}

func setWindowSize(h, w int, m *model) {
	m.height = h
	m.width = w
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	m.ErrText = ""
	switch msg:= msg.(type){
	case tea.WindowSizeMsg :
		setWindowSize(msg.Height, msg.Width, &m)
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
					if (!m.items.Buy(m.actions[m.cursor])){
						m.ErrText = "Can't afford"
					} 						
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
				"worker":   &items.Worker{items.ObjStats{Price : 10, Earn : 1, Amount : 0}},
				"factory": &items.Factory{items.ObjStats{Price : 50, Earn : 5, Amount : 0}},
				"company": &items.Company{items.ObjStats{Price : 100, Earn : 10, Amount : 0}},
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
