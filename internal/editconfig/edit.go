package editconfig

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const configFile = "config.toml"

func EditSource() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
