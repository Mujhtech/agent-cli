package initals

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/mujhtech/agent-cli/pkg/cmd/version"
	"github.com/mujhtech/agent-cli/pkg/views"
)

var sigilParts = []string{
	`      -       ____ 	_____     _   _ 	 _____		`,
	`     / \     / ___|	| ____|   | \ | |  	|_   _|		`,
	`    / - \    | |  _   	|  _|	  |  \| |	  | |		`,
	`   / --- \   | |_| |  	| |___	  | .   |	  | |		`,
	`  /_/   \_\  \_____|  	|_____|	  |_| \_|     |_|		`,
}

type commandView struct {
	Name        string
	Command     string
	Description string
}

var gradientSigil string

type model struct {
	form   *huh.Form
	width  int
	choice string
}

func initialModel() model {
	m := model{}

	commands := []commandView{
		{Name: "agent-cli models", Description: "List of available models", Command: "models"},
		{Name: "agent-cli create [NAME]", Description: "Start new project with agent", Command: "create"},
		{Name: "agent-cli index [CODESPACE]", Description: "Index the current workspace", Command: "index"},
		{Name: "agent-cli code", Description: "Generate/Edit code in the current workspace", Command: "code"},
		{Name: "agent-cli review", Description: "Review code in the current workspace", Command: "review"},
		{Name: "agent-cli chat", Description: "Chat in the current workspace", Command: "chat"},
		{Name: "agent-cli help", Description: "Get help", Command: "help"},
		{Name: "q", Description: "Quit", Command: "q"},
	}

	var options []huh.Option[string]

	for _, command := range commands {
		options = append(options, huh.Option[string]{Key: fmt.Sprintf("%s %s", command.Name, lipgloss.NewStyle().Foreground(views.LightGray).Render(command.Description)), Value: command.Command})
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("command").
				Options(options...).
				Value(&m.choice).Height(9).
				WithTheme(views.GetCustomTheme()),
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
		m.width = 200
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
		formHeader := "\n\n"
		formHeader += lipgloss.NewStyle().Foreground(views.Primary).Render("Agent CLI")
		formHeader += "\n"
		formHeader += "An open-source alternative to OpenAI Codex & Claude Code\n\n"
		formHeader += "Version: " + version.Version + "\n\n"
		formHeader += "-------------------------------------------------------------\n\n"
		formHeader += "Get started\n"
		form := v

		body := lipgloss.JoinHorizontal(lipgloss.Top, gradientSigil, formHeader+form)

		return lipgloss.NewStyle().Padding(2, 0, 2, 2).Render(body)
	}
}

func Run() (string, error) {

	for _, line := range sigilParts {
		gradientSigil += lipgloss.NewStyle().Foreground(views.Primary).Render(line) + "\n"
	}

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
