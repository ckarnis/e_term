package config

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func AddSource() {
	file, err := os.OpenFile(
		configFile,
		os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		fmt.Println("config.toml not found, run -> eterm init")
		file.Close()
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel())
	form, runerr := p.Run()
	if runerr != nil {
		fmt.Println("runtime error")
		file.Close()
		os.Exit(1)
	}

	m := form.(model)

	name := strings.TrimSpace(m.inputs[0].Value())
	base_url := strings.TrimSpace(m.inputs[1].Value())
	api_key := strings.TrimSpace(m.inputs[2].Value())
	s := strings.TrimSpace(m.inputs[3].Value())
	time_sec, _ := strconv.Atoi(s)

	table := formatTableName(name)

	testNew := fmt.Sprintf(`
[%s]
	base_url = "%s"
	api_key = "%s"
	timeout_seconds = %d
	`, table, base_url, api_key, time_sec)

	if _, err := file.WriteString(testNew); err != nil {
		fmt.Println("config.toml could not be edited")
		file.Close()
		os.Exit(1)
	}

	file.Close()

}

// move to general
func formatTableName(name string) string {
	name = strings.TrimSpace(name)

	bareKey := regexp.MustCompile(`^[A-Za-z0-9_-]+$`)
	if bareKey.MatchString(name) {
		return name
	}

	escaped := strings.ReplaceAll(name, `"`, `\"`)
	return `"` + escaped + `"`
}
