package simpleBubble

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FancyItemSelectorModel struct {
	err    error
	choice fancyItem
	style  lipgloss.Style
	list   list.Model
}

func (m *FancyItemSelectorModel) Init() tea.Cmd {
	return nil
}

func (m *FancyItemSelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.err = fmt.Errorf("user quit")
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(fancyItem)
			if ok {
				m.choice = i
			}

			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := m.style.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *FancyItemSelectorModel) View() string {
	return m.style.Render(m.list.View())
}

func NewFancyItemSelector(windowTitle string, items *[]fancyItem) (*fancyItem, error) {
	listDocStyle := lipgloss.NewStyle().Margin(1, 2)
	var listItems []list.Item = make([]list.Item, len(*items))
	for i, v := range *items {
		listItems[i] = v
	}
	m := FancyItemSelectorModel{
		list:  list.New(listItems, list.NewDefaultDelegate(), 0, 0),
		style: listDocStyle,
	}
	m.list.Title = windowTitle

	p := tea.NewProgram(&m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return nil, fmt.Errorf("could not select item")
	}
	if m.err != nil {
		return nil, m.err
	}
	return &m.choice, nil
}
