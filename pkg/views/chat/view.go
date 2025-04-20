package chat

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mujhtech/agent-cli/pkg/views"
)

const (
	width = 96

	columnWidth = 30
)

func Run() (string, error) {

	doc := strings.Builder{}

	docStyle := lipgloss.NewStyle().Padding(1, 2, 1, 2)

	dialogBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	buttonStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#888B7E")).
		Padding(0, 3).
		MarginTop(1)

	activeButtonStyle := buttonStyle.
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#F25D94")).
		MarginRight(2).
		Underline(true)

	okButton := activeButtonStyle.Render("Yes")
	cancelButton := buttonStyle.Render("Maybe")

	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(rainbow(lipgloss.NewStyle(), "Are you sure you want to eat marmalade?", views.Blends))
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	dialog := lipgloss.Place(width, 9,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(views.Subtle),
	)

	doc.WriteString(dialog + "\n\n")

	fmt.Println(docStyle.Render(doc.String()))

	return "", nil
}

func rainbow(base lipgloss.Style, s string, colors []color.Color) string {
	var str string
	for i, ss := range s {
		color, _ := colorful.MakeColor(colors[i%len(colors)])
		str = str + base.Foreground(lipgloss.Color(color.Hex())).Render(string(ss))
	}
	return str
}
