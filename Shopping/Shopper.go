package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type thingy struct {
	items    []string         //the text that renders
	cursor   int              //the moving thing
	selected map[int]struct{} //write out on paper why this is used

}

func initialModel() thingy { //where does this get run?
	return thingy{
		items:    []string{"A piece of text", "another piece of text", "a final piece of text"},
		selected: make(map[int]struct{}),
	}
}

func (a thingy) Init() tea.Cmd { //what is this function for?
	return nil
}

func (a thingy) Update(msg tea.Msg) (tea.Model, tea.Cmd) { //what are the many sets of parentheticals doing?
	switch msg := msg.(type) { //there is the keyword msg in the parameters, in the first part of the switch, and in the second part. What is happening?
	case tea.KeyMsg: //self-explanatory: if the tea.Msg is a key press, we'll switch through that
		switch msg.String() { //that msg variable is evaluated as a String. this makes sense because we know it is a key press

		case "ctrl+c", "q":
			return a, tea.Quit //why does a need to be returned?
		case "up", "k":
			if a.cursor > 0 { //so a thingy is a type. where is it initialized?
				a.cursor--
			}
		case "down", "j":
			if a.cursor < len(a.items)-1 {
				a.cursor++
			}
		case "enter", " ":
			_, ok := a.selected[a.cursor] //what is the _? so I think a variable is created representing the row the cursor is on
			if ok {                       //so a.selected @a.cursor is 1
				delete(a.selected, a.cursor) //so delete is a function that takes in a map & key, then the key type. still don't understand
			} else {
				a.selected[a.cursor] = struct{}{} //if its empty, then the selected item at a.cursor is set equal to a struct? no idea
			}
		}
	}
	return a, nil //return the model, and nil? why two return values?
}

func (a thingy) View() string { //ok. takes in model, than calls View, then returns a string
	//header
	s := "What should we buy at the market?\n\n" //where does this get printed? how is it printed then never printed again?

	//Iterate over choices (view)
	for i, item := range a.items { //so new variable item, which holds what? what is i for?

		cursor := " " //a string showing what no cursor looks like
		if a.cursor == i {
			cursor = ">"
		}

		checked := " "                  //another string showing what no checked looks like
		if _, ok := a.selected[i]; ok { //need to figure out what this if means
			checked = "x" //what it looks like in selected form.
		}

		//render

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, item) //so this variable s from the top has the choices tacked on, every loop

	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("alas, there's been an error: %")
		os.Exit(1)
	}
}
