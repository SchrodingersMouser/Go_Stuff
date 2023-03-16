package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type board struct { //why is this a struct?
	squares  [3][3]string             //the squares of the board
	cursor   [4]int                   //the square the cursor is pointing at. index is row, value is column
	selected map[int]map[int]struct{} //which items are selected. 1 = unselected, 2 = O, 3 = X
	//	varName map keyTypes val Type
}

func initialModel() board {
	return board{
		squares: [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}},
		//squares:  "[ ] [ ] [ ]\n[ ] [ ] [ ]\n[ ] [ ] [ ]",
		selected: make(map[int]map[int]struct{}),
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
			if m.cursor[1] > 0 {
				m.cursor[1]--
			}
		case "down", "s":
			if m.cursor[1] < len(m.cursor)-1 {
				m.cursor[1]++
			}
		case "right", "d":
			if m.cursor[1] < len(m.cursor)-1 {
				m.cursor[1]++
			}
		case "left", "a":
			if m.cursor[1] > 0 {
				m.cursor[1]--
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor[1]]
			if ok {
				delete(m.selected, m.cursor[1])
			} else {
				m.selected[m.cursor[1]][1] = struct{}{} //started work updating here
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
			cursor := " " //no cursor
			if m.cursor[l] == i {
				cursor = ">" //cursor!
			}
			choice = choice

			// Is this choice selected?
			checked := " " //not selected
			if _, ok := m.selected[i]; ok {
				checked = "x" //selected
			}

			// render the row
			s += fmt.Sprintf("%s [%s]", cursor, checked)
		}
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
