package simpleBubble

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type textInputModel struct {
	Title       string
	Placeholder string
	Result      string
	textInput   textinput.Model
	err         error
}

func NewTextInput(Title, Placeholder string) (string, error) {
	defer clearTerminal()
	ti := textinput.New()
	ti.Placeholder = Placeholder
	ti.Focus()

	m := textInputModel{
		textInput:   ti,
		Title:       Title,
		Placeholder: Placeholder,
		err:         nil,
	}

	p := tea.NewProgram(&m)
	if _, err := p.Run(); err != nil {
		return "", err
	}
	if m.err != nil {
		return "", m.err
	}
	return m.Result, nil
}

func (m *textInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *textInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.Result = m.textInput.Value()
			return m, tea.Quit
		}
	case error:
		m.err = msg
		return m, nil
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *textInputModel) View() string {
	msg := fmt.Sprintf(
		"%v\n\n%v\n\n%v",
		m.Title,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
	return msg
}
