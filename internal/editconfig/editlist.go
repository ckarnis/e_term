package editconfig

import (
	"fmt"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// rewrite using simple list instead of default tea list

const configPath = "config.toml"

var docStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)
var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#fb0000"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#0064f9"))
	cursorStyle  = focusedStyle
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", helpStyle.Render("Submit"))
)

type mode int

const (
	modeList mode = iota
	modeEdit
)

type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

type model struct {
	mode     mode
	selected string
	list     []string
	tree     map[string]any
	form     formModel
	cursor   int
}

type formModel struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	submitted  bool
}

func initialModel() model {
	var fileData map[string]any
	if _, err := toml.DecodeFile(configFile, &fileData); err != nil {
		fmt.Println("can't read config.toml")
		os.Exit(1)
	}

	items := make([]string, 0, len(fileData))
	for key := range fileData {
		items = append(items, key)

	}

	return model{
		mode: modeList,
		tree: fileData,
		list: items,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.mode {
	case modeList:
		switch msg := msg.(type) {

		case tea.KeyMsg:
			switch msg.String() {

			case "ctrl+c", "esc":
				return m, tea.Quit

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor < len(m.list)-1 {
					m.cursor++
				}

			case "enter":
				m.selected = m.list[m.cursor]

				sub, ok := m.tree[m.selected].(map[string]any)
				if !ok {
					return m, nil
				}

				baseURL, _ := sub["base_url"].(string)
				apiKey, _ := sub["api_key"].(string)

				timeout := 0
				switch v := sub["timeout_seconds"].(type) {
				case int64:
					timeout = int(v)
				case int:
					timeout = v
				}

				m.form = newForm(baseURL, apiKey, timeout)
				m.mode = modeEdit
			}
		}

		return m, nil
	case modeEdit:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				m.mode = modeList
				return m, nil

			case "enter":
				if m.form.focusIndex == len(m.form.inputs) {
					m.save()
					m.mode = modeList
					return m, nil
				}

				m.form.focusIndex++
			}
		}

		var cmd tea.Cmd
		m.form, cmd = m.form.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	switch m.mode {
	case modeList:
		s := "Select a source to edit\n\n"

		for i, item := range m.list {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("%s %s\n", cursor, item)
		}

		s += "\nArrows to move | Enter to edit | esc or ctrl+c to quit\n"

		return docStyle.Render(s)
	case modeEdit:
		return docStyle.Render(m.form.View())
	default:
		return ""
	}
}

// make this like the add form, it is too not good
func newForm(baseURL, apiKey string, timeout int) formModel {
	inputs := make([]textinput.Model, 3)

	for i := range inputs {
		t := textinput.New()
		t.CharLimit = 128

		switch i {
		case 0:
			t.Placeholder = "Base URL"
			t.SetValue(baseURL)
			t.Focus()
		case 1:
			t.Placeholder = "API Key"
			t.SetValue(apiKey)
		case 2:
			t.Placeholder = "Timeout Seconds"
			t.SetValue(strconv.Itoa(timeout))
		}

		inputs[i] = t
	}

	return formModel{inputs: inputs}
}

func (f formModel) Update(msg tea.Msg) (formModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "esc":
			return f, tea.Quit

		case "tab", "shift+tab", "up", "down":
			if msg.String() == "up" || msg.String() == "shift+tab" {
				f.focusIndex--
			} else {
				f.focusIndex++
			}

			if f.focusIndex > len(f.inputs) {
				f.focusIndex = 0
			} else if f.focusIndex < 0 {
				f.focusIndex = len(f.inputs)
			}

		case "enter":
			if f.focusIndex == len(f.inputs) {
				f.submitted = true
				return f, nil
			}
			f.focusIndex++
		}
	}

	// Update inputs and focus styles
	for i := range f.inputs {

		if i == f.focusIndex {
			f.inputs[i].Focus()
			f.inputs[i].PromptStyle = focusedStyle
			f.inputs[i].TextStyle = focusedStyle
		} else {
			f.inputs[i].Blur()
			f.inputs[i].PromptStyle = blurredStyle
			f.inputs[i].TextStyle = blurredStyle
		}

		var cmd tea.Cmd
		f.inputs[i], cmd = f.inputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return f, tea.Batch(cmds...)
}

func (f formModel) View() string {
	s := "\nEdit values (Enter to submit, Esc to cancel)\n\n"

	for _, in := range f.inputs {
		s += in.View() + "\n"
	}

	if f.focusIndex == len(f.inputs) {
		s += "\n[ Save ]\n"
	} else {
		s += "\nSave\n"
	}

	return s
}

func (m *model) save() error {
	sub, ok := m.tree[m.selected].(map[string]any)
	if !ok {
		return fmt.Errorf("invalid table: %s", m.selected)
	}

	baseURL := m.form.inputs[0].Value()
	apiKey := m.form.inputs[1].Value()
	timeout, err := strconv.Atoi(m.form.inputs[2].Value())
	if err != nil {
		return err
	}

	sub["base_url"] = baseURL
	sub["api_key"] = apiKey
	sub["timeout_seconds"] = timeout

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := toml.NewEncoder(file)
	return enc.Encode(m.tree)
}
