package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
	"os"
)

type board struct { //why is this a struct?
	squares  [3][3]string //the squares of the board
	col      int          //the square the cursor is pointing at. index is row, value is column
	row      int
	selected [3][3]int //which items are selected. 1 = unselected, 2 = O, 3 = X
	//	varName map keyTypes val Type
}

var xStyle = gloss.NewStyle().
	Bold(true).
	Foreground(gloss.Color("5"))

func initialModel() board {
	return board{
		squares: [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}},
		//squares:  "[ ] [ ] [ ]\n[ ] [ ] [ ]\n[ ] [ ] [ ]",
		selected: [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
	}
}

func (m board) Init() tea.Cmd {
	return nil
}

func (m board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "w":
			if m.row > 0 {
				m.row--
			}
		case "down", "s":
			if m.row < 2 {
				m.row++
			}
		case "right", "d":
			if m.col < 2 {
				m.col++
			}
		case "left", "a":
			if m.col > 0 {
				m.col--
			}

		case "enter", " ":
			ok := m.selected[m.row][m.col]
			if ok == 0 {
				m.selected[m.row][m.col] = 1
			}
		}
	}
	return m, nil
}

func (m board) View() string {
	//header
	s := "want to play tic tac toe?\n\n"

	// Iterate over the board
	for l, h := range m.squares { //row

		for i, choice := range h { //column

			// Is the cursor pointing at this choice?
			open := "[" //no cursor
			close := "]"
			if m.col == i && m.row == l {
				open = xStyle.Render("[") //cursor!
				close = xStyle.Render("]")
			}
			choice = choice
			open = open
			close = close

			// Is this choice selected?
			checked := " " //not selected
			if m.selected[l][i] == 1 {
				checked = "x" //selected
			}

			// render the row

			if m.col == i && m.row == l {
				s += xStyle.Render(fmt.Sprintf("[%s]", checked))
			} else {
				s += fmt.Sprintf("[%s]", checked)
			}
		}
		s += "\n"
	}
	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %")
		os.Exit(1)
	}
}
