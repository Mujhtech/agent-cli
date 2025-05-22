package views

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/gamut"
)

var (
	LightGray = lipgloss.AdaptiveColor{Light: "#828282", Dark: "#828282"}
	Primary   = lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"}
	Green     = lipgloss.AdaptiveColor{Light: "#23cc71", Dark: "#23cc71"}
	White     = lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"}
	Subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	Blends    = gamut.Blends(lipgloss.Color("#F25D94"), lipgloss.Color("#EDFF82"), 50)
)

func GetCustomTheme() *huh.Theme {
	t := huh.ThemeCharm()

	t.Blurred.FocusedButton = t.Blurred.FocusedButton.Background(Primary)
	t.Blurred.FocusedButton = t.Blurred.FocusedButton.Bold(true)
	// t.Blurred.TextInput.Prompt = t.Blurred.TextInput.Prompt.Foreground(Light)
	// t.Blurred.TextInput.Cursor = t.Blurred.TextInput.Cursor.Foreground(Light)
	t.Blurred.SelectSelector = t.Blurred.SelectSelector.Foreground(Primary)
	t.Blurred.Title = t.Blurred.Title.Foreground(LightGray).Bold(true)
	t.Blurred.Description = t.Blurred.Description.Foreground(LightGray)

	t.Focused.Title = t.Focused.Title.Foreground(Primary).Bold(true)
	t.Focused.Description = t.Focused.Description.Foreground(LightGray).Bold(true)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Bold(true)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Background(Primary)
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(Primary)
	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(LightGray)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(Primary)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(Primary)

	// t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(Red)
	// t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(Red)

	t.Focused.Base = t.Focused.Base.BorderForeground(Primary)
	t.Focused.Base = t.Focused.Base.BorderBottomForeground(Primary)

	// t.Focused.Base = t.Focused.Base.MarginTop(DefaultLayoutMarginTop)
	// t.Blurred.Base = t.Blurred.Base.MarginTop(DefaultLayoutMarginTop)

	return t
}
