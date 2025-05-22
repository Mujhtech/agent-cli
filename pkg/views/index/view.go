package index

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
	"github.com/mujhtech/agent-cli/pkg/views"
	"github.com/mujhtech/agent-cli/pkg/views/utils"
)

func Run() error {

	pwd, err := os.Getwd()

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(views.White).
		Padding(0, 1).
		Width(96).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	workdir := lipgloss.NewStyle().
		Foreground(views.LightGray).
		Render("* workdir: " + pwd)

	model := lipgloss.NewStyle().
		Foreground(views.LightGray).
		Render("* model: gpt-3.5-turbo")

	borderBorder := lipgloss.JoinVertical(lipgloss.Top, workdir, model)

	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingRight(1)
	itemStyle := lipgloss.NewStyle().Foreground(views.Primary).PaddingRight(1)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	t := tree.Root(pwd).
		EnumeratorStyle(enumeratorStyle).
		RootStyle(itemStyle).
		ItemStyle(itemStyle)

	if err := addBranches(t, ".", true); err != nil {
		fmt.Fprintf(os.Stderr, "Error building tree: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(lipgloss.Place(0, 0,
		lipgloss.Center, lipgloss.Center,
		borderStyle.Render(borderBorder),
	))
	fmt.Println(t)

	err = utils.WithInlineSpinner("Indexing...", func() error {
		// Simulate some work
		time.Sleep(5 * time.Second)

		return nil
	})

	return err
}

func addBranches(root *tree.Tree, path string, skipRoot bool) error {
	items, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, item := range items {
		if item.IsDir() {
			// Skip directories that start with a dot.
			if strings.HasPrefix(item.Name(), ".") {
				continue
			}

			var treeBranch *tree.Tree
			if skipRoot {
				// For root level, add contents directly to the main tree
				treeBranch = root
				skipRoot = false // Reset for subdirectories
			} else {
				// For other directories, create new branches
				treeBranch = tree.Root(item.Name())
				root.Child(treeBranch)
			}

			// Recurse
			branchPath := filepath.Join(path, item.Name())
			if err := addBranches(treeBranch, branchPath, false); err != nil {
				return err
			}
		} else {
			// Skip files that start with a dot
			if strings.HasPrefix(item.Name(), ".") {
				continue
			}

			root.Child(item.Name())
		}
	}

	return nil
}
