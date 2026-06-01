package windows

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func RunData(name string) {

	pid := os.Getpid()

	_ = WritePID(name, pid)

	defer RemovePID(name)

	fmt.Println("window:", name)
	fmt.Println("pid:", pid)

	p := tea.NewProgram(NewDataModel(name))

	_, _ = p.Run()
}

func RunConfigChild(name string) {

	pid := os.Getpid()

	_ = WritePID(name, pid)

	defer RemovePID(name)

	p := tea.NewProgram(NewConfigModel("config"))

	_, _ = p.Run()
}
