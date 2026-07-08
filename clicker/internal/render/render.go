package render

import (
	lipgloss "charm.land/lipgloss/v2"
)

var Money = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FFFFFF")).
    Background(lipgloss.Color("#105a37")).
    PaddingTop(0).
   	PaddingLeft(0)
	

var Menu = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FFFFFF")).
    Background(lipgloss.Color("#00836d")).
    PaddingTop(0).
   	PaddingLeft(0)
    //Width(22)

var ErrUi = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#FFFFFF")).
    Background(lipgloss.Color("#960019"))
