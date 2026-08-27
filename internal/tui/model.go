package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/Pastalikek65/sqlite-tui/internal/db"
)

var (
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("99")).Padding(0,1)
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("240"))
	selStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	normStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	faintStyle = lipgloss.NewStyle().Faint(true)
)

type Model struct {
	db *db.DB
	tables []string
	cursor int
	rows []map[string]any
	cols []string
	wal int64
	err string
}

func New(d *db.DB) Model {
	tables, _ := d.ListTables()
	m := Model{db: d, tables: tables}
	if len(tables)>0 {
		m.loadTable(tables[0])
	}
	return m
}

func (m *Model) loadTable(name string) {
	rows, cols, err := m.db.Query(name, 20)
	if err != nil { m.err = err.Error(); return }
	m.rows = rows
	m.cols = cols
	m.wal, _ = m.db.WalSize()
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q","ctrl+c": return m, tea.Quit
		case "up","k": if m.cursor>0 { m.cursor-- }
		case "down","j": if m.cursor < len(m.tables)-1 { m.cursor++; m.loadTable(m.tables[m.cursor]) }
		case "r": if len(m.tables)>0 { m.loadTable(m.tables[m.cursor]) }
		}
	}
	return m, nil
}

func (m Model) Tables() []string { return m.tables }
func (m Model) View() string {
	s := titleStyle.Render(" sqlite-tui ") + faintStyle.Render(fmt.Sprintf(" %s WAL:%dB", m.dbPath(), m.wal)) + "\n"
	s += headerStyle.Render(" TABLES ") + "\n"
	for i, t := range m.tables {
		line := fmt.Sprintf(" %-28s", t)
		if i==m.cursor { s+= selStyle.Render("▶"+line)+"\n" } else { s+= normStyle.Render(" "+line)+"\n" }
	}
	if m.err != "" { s+= "\n"+lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("err: "+m.err)+"\n" }
	s += "\n" + headerStyle.Render(fmt.Sprintf(" %s (%d rows) ", m.currentTable(), len(m.rows))) + "\n"
	if len(m.cols)>0 {
		s+= strings.Join(m.cols, " | ") + "\n"
		s+= strings.Repeat("─", 60)+"\n"
		for _, r := range m.rows {
			var vals []string
			for _, c := range m.cols { vals = append(vals, fmt.Sprint(r[c])) }
			line := strings.Join(vals, " | ")
			if len(line)>70 { line=line[:70]+"..." }
			s+= normStyle.Render(" "+line)+"\n"
		}
	}
	s += "\n" + faintStyle.Render("j/k:nav  r:reload  q:quit") + "\n"
	return s
}

func (m Model) dbPath() string { return m.db.Path() }
func (m Model) currentTable() string {
	if len(m.tables)==0 { return "(no tables)" }
	if m.cursor < len(m.tables) { return m.tables[m.cursor] }
	return m.tables[0]
}
