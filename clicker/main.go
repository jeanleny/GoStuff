package main

import (
	"fmt"
	"os"
	"time"
	"net/http"
	items "clicker/internal/items"
	tea "charm.land/bubbletea/v2"
	render "clicker/internal/render"
	lipgloss "charm.land/lipgloss/v2"
)

const url = "http://localhost:8080/"

func display() tea.View{
	v := tea.NewView("Hello, World!")
	return (v)
}

type model struct {
	name 			string
	cursor			int
	items			items.Business
	width			int
	height			int
	actions			[]string
	ErrText			string
	Working			bool
	WorkingTimer	int
	UpdateTimer		int
	err				error
	status			int
}

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e errMsg) Error() string { return e.err.Error() }

type errMsg struct{ err error }

type statusMsg int

func checkServer() tea.Msg {
    // Create an HTTP client and make a GET request.
    c := &http.Client{Timeout: 10 * time.Second}
    res, err := c.Get(url)

    if err != nil {
        // There was an error making our request. Wrap the error we received
        // in a message and return it.
        return errMsg{err}
    }
    // We received a response from the server. Return the HTTP status code
    // as a message.
    return statusMsg(res.StatusCode)
}

func (m model) View() tea.View {
	s := []string {"Kessafou ?"}
	err := ""
	for i, choice := range m.actions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		line := fmt.Sprintf("%s %s : %d %s", cursor, choice, m.items.Stock[choice].GetAmount(), err)
		lineBlock := render.Menu.Render(line)
		if m.cursor == i && m.ErrText != "" {
			errBlock := render.ErrUi.Render(m.ErrText)
			lineBlock = lipgloss.JoinHorizontal(lipgloss.Top, lineBlock, errBlock)
		}
		s = append(s, lineBlock)
	}
		menuBlock := lipgloss.JoinVertical(lipgloss.Left, s...)
		moneyBlock := render.Money.Render(fmt.Sprintf("$$$ : %0.2f\n", m.items.Money))
		moneyPlaced := lipgloss.Place(m.width - lipgloss.Width(menuBlock) ,lipgloss.Height(menuBlock), lipgloss.Right, lipgloss.Top, moneyBlock)
		full := lipgloss.JoinHorizontal(lipgloss.Top, menuBlock, moneyPlaced)
		return tea.NewView(full)
	}	

func (m model) Init() tea.Cmd {
	return tea.Batch(checkServer, tickCmd())
}

func setWindowSize(h, w int, m *model) {
	m.height = h
	m.width = w
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {return TickMsg(t)})
}

type TickMsg time.Time

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	m.ErrText = ""
	switch msg:= msg.(type){
	case statusMsg :
		// The server returned a status message. Save it to our model. Also
        // tell the Bubble Tea runtime we want to exit because we have nothing
        // else to do. We'll still be able to render a final view with our
        // status message.
        m.status = int(msg)
		fmt.Println("pas du tou")
        return m, nil
	case errMsg :
		// There was an error. Note it in the model. And tell the runtime
        // we're done and want to quit.
        m.err = msg
        return m, nil
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
					checkServer()
					if (!m.items.Buy(m.actions[m.cursor])){
						m.ErrText = "Can't afford"
					} 
				case "space" :
					if (!m.Working) {
						m.items.Work()
						m.Working = true 
						return m ,tickCmd()
					}
			}
		case TickMsg :
			m.WorkingTimer --
			m.UpdateTimer --
			if m.UpdateTimer <= 0 {
				m.UpdateTimer = 1
				m.items.Money += m.items.Earning
				return m, tickCmd()
			}
			if m.WorkingTimer <= 0 {
				m.Working = false
				m.WorkingTimer = 1
			} else {
				return m, tickCmd()
			}
			break
	}
	return m, nil
}


func makeModel() model {
	return model {
		name : "le caca prousti",
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
	if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
