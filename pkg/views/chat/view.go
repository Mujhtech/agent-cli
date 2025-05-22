package chat

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/mujhtech/agent-cli/pkg/views"
)

const (
	width = 96
)

type model struct {
	form   *huh.Form
	width  int
	prompt string
}

func initialModel() model {
	m := model{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Value(&m.prompt).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("prompt cannot be empty")
					}
					return nil
				}),
		),
	).WithWidth(200).WithTheme(views.GetCustomTheme()).WithShowHelp(false)

	m.form = form

	return m
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = 90
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
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
		// 	command := m.form.GetString("command")
		// 	m.prompt = command
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.form.State {
	case huh.StateCompleted:
		return ""
	default:

		dialogBoxStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(views.White).
			Padding(1, 0).
			Width(96).
			Height(1).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

		formFooter := lipgloss.NewStyle().
			Foreground(views.LightGray).
			Render("press Ctrl+C or esc to exit | press enter to send")

		return lipgloss.Place(width, 7,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(m.form.View()),
		) + "\n" + formFooter
	}
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func Run() (string, error) {
	m := initialModel()
	p, err := tea.NewProgram(m).Run()
	if err != nil {
		return "", err
	}

	if m, ok := p.(model); ok && m.prompt != "" {
		return m.prompt, nil
	}

	return m.prompt, nil
}
