package simpleBubble

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type checkListModel struct {
	Title         string
	err           error
	items         *[]checkListItem
	selectedStyle *lipgloss.Style
	cursorIndex   int
}

func NewCheckList(Title string, items *[]checkListItem) ([]string, error) {
	defer clearTerminal()
	s := lipgloss.NewStyle()
	initialModel := &checkListModel{
		Title:         Title,
		items:         items,
		selectedStyle: &s,
	}

	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	checkedItems := []string{}
	for _, v := range *initialModel.items {
		if v.Checked {
			checkedItems = append(checkedItems, v.Text)
		}
	}
	return checkedItems, nil
}

func (m checkListModel) Init() tea.Cmd {
	return nil
}

func (m checkListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		key := msg.(tea.KeyMsg).String()
		switch key {
		case "q", "ctrl+c":
			m.err = fmt.Errorf("user quit")
			return m, tea.Quit
		case "j", "down":
			if m.cursorIndex < len(*m.items)-1 {
				m.cursorIndex++
			}
		case "k", "up":
			if m.cursorIndex > 0 {
				m.cursorIndex--
			}
		case " ":
			if m.cursorIndex >= 0 && m.cursorIndex < len(*m.items) {
				(*m.items)[m.cursorIndex].Checked = !(*m.items)[m.cursorIndex].Checked
			}
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m checkListModel) View() string {
	var view string
	view += fmt.Sprintf("%v\n", m.Title)
	for i, item := range *m.items {
		t := fmt.Sprintf("[%v] %v", checked(item), item.Text)
		if i == m.cursorIndex {
			view += m.selectedStyle.Bold(true).Foreground(lipgloss.Color("#ca4e6c")).Render(t)
			view += "\n"
		} else {
			view += t + "\n"
		}
	}
	return view
}

func checked(i checkListItem) string {
	if i.Checked {
		return "x"
	}
	return " "
}
