// Package wallet command line operations
package wallet

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	wallet "github.com/trevatk/go-wallet"
)

const (
	defaultDir                 = "$PWD"
	useHighPerformanceRenderer = false
)

var (
	// output file flag
	outputFile string

	createWalletCmd = &cobra.Command{
		Use: "create",
		RunE: func(_ *cobra.Command, _ []string) error {

			w := wallet.New()

			if outputFile == defaultDir {

				wd, err := os.Getwd()
				if err != nil {
					return fmt.Errorf("failed to get current working directory %v", err)
				}
				outputFile = filepath.Join(wd, "wallet.json")
			}

			err := w.MarshalToFile(outputFile)
			if err != nil {
				return fmt.Errorf("failed to marshal wallet %v", err)
			}

			content, err := os.ReadFile(filepath.Clean(outputFile))
			if err != nil {
				return fmt.Errorf("failed to read wallet file %v", err)
			}

			s := string(content)
			split := strings.Split(s, ",")
			var b strings.Builder
			for _, sp := range split {
				b.WriteString(sp + "\n")
			}

			p := tea.NewProgram(
				CreateWalletModel{content: b.String()},
				tea.WithAltScreen(),
				tea.WithMouseCellMotion(),
			)

			_, err = p.Run()
			if err != nil {
				return fmt.Errorf("failed to run bubbletea program %v", err)
			}

			return nil
		},
	}

	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

// CreateWalletModel bubbletea create wallet command model
type CreateWalletModel struct {
	content  string
	ready    bool
	viewport viewport.Model
}

// Init model
func (m CreateWalletModel) Init() tea.Cmd {
	return nil
}

// Update model from user input
func (m CreateWalletModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.content)
			m.ready = true

			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View returns bubbletea rendered view
func (m CreateWalletModel) View() string {
	if !m.ready {
		return "\n Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m CreateWalletModel) headerView() string {
	title := titleStyle.Render(outputFile)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m CreateWalletModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func init() {

	createWalletCmd.Flags().StringVarP(&outputFile, "output-file", "o", "$PWD", "set output file location")

	walletCmd.AddCommand(createWalletCmd)
}
