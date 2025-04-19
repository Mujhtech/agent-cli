package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/mujhtech/agent-cli/pkg/agent"
	"github.com/mujhtech/agent-cli/pkg/views"
)

var gradientSigil string

type model struct {
	form   *huh.Form
	width  int
	choice string
}

func initialModel() model {
	m := model{}

	var options []huh.Option[string]

	for index, model := range agent.AvailableModelCatalogs {
		options = append(options, huh.Option[string]{Key: fmt.Sprintf("%s %s", fmt.Sprintf("agent-cli models -d=%v", index), lipgloss.NewStyle().Foreground(views.LightGray).Render(model.Name)), Value: fmt.Sprintf("%s %v", "models -d=", index)})
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("command").
				Options(options...).
				Value(&m.choice).Height(11),
		),
	).WithShowHelp(false)

	m.form = form

	return m
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = 100
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			m.choice = ""
			return m, tea.Quit
		}
	}
	var cmds []tea.Cmd

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		command := m.form.GetString("command")
		m.choice = command
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.form.State {
	case huh.StateCompleted:
		return ""
	default:
		v := strings.TrimSuffix(m.form.View(), "\n\n")
		formHeader := "\n"
		formHeader += lipgloss.NewStyle().Foreground(views.Green).Render("Agent Models")
		formHeader += "\n"
		formHeader += "List of available models\n\n"
		formHeader += "-------------------------------------------------------------\n\n"
		form := v

		body := lipgloss.JoinHorizontal(lipgloss.Top, gradientSigil, formHeader+form)

		return lipgloss.NewStyle().Padding(2, 0, 2, 2).Render(body)
	}
}

func Run() (string, error) {
	m := initialModel()
	p, err := tea.NewProgram(m).Run()
	if err != nil {
		return "", err
	}

	if m, ok := p.(model); ok && m.choice != "" {
		return m.choice, nil
	}

	return "", nil
}
