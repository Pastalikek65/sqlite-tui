package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/Pastalikek65/sqlite-tui/internal/db"
	"github.com/Pastalikek65/sqlite-tui/internal/tui"
)

var version = "0.1.0"

func main() {
	root := &cobra.Command{
		Use:   "sqlite-tui [dbfile]",
		Short: "TUI viewer for SQLite — tables, rows, WAL size",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "test.db"
			if len(args)==1 { path = args[0] }
			if _, err := os.Stat(path); err != nil {
				return fmt.Errorf("db not found %s: %w", path, err)
			}
			d, err := db.Open(path)
			if err != nil { return err }
			defer d.Close()
			m := tui.New(d)
			p := tea.NewProgram(m, tea.WithAltScreen())
			_, err = p.Run()
			return err
		},
	}
	root.Version = version
	if err := root.Execute(); err != nil { os.Exit(1) }
}
