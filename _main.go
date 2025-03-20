package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type procesFinishedMsg struct{ err error }

func runProcess() tea.Cmd {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		editor = "vim"
	}
	cmd := exec.Command(editor) //nolint:gosec

	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return procesFinishedMsg{err}
	})
}

type model struct {
	altscreenActive bool
	err             error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			m.altscreenActive = !m.altscreenActive
			cmd := tea.EnterAltScreen
			if !m.altscreenActive {
				cmd = tea.ExitAltScreen
			}
			return m, cmd
		case "e":
			return m, runProcess()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case procesFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error() + "\n"
	}
	return "Press 'e' to open your EDITOR.\nPress 'a' to toggle the altscreen\nPress 'q' to quit.\n"
}

// func main() {
// 	m := model{}
// 	if _, err := tea.NewProgram(m).Run(); err != nil {
// 		fmt.Println("Error running program:", err)
// 		os.Exit(1)
// 	}
// }

func NewCmdEventBrowser() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Create a new event browser",
		Example: "inngest init",
		Run:     runBrowser,
	}
	return cmd
}

func runBrowser(cmd *cobra.Command, args []string) {
	m := model{}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func main() {
	cmd := NewCmdEventBrowser()
	cmd.Execute()
}
