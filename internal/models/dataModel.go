package models

import tea "github.com/charmbracelet/bubbletea"

type DataModel struct {
	Name string
}

func NewDataModel(name string) DataModel {
	return DataModel{
		Name: name,
	}
}

func (m DataModel) Init() tea.Cmd {
	return nil
}

func (m DataModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m DataModel) View() string {
	return "Data window: " + m.Name + "\nctrl+c to quit"
}
