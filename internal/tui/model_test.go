package tui

import (
	"path/filepath"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/Pastalikek65/sqlite-tui/internal/db"
)

func TestNewAndView(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.db")
	d, _ := db.Open(path)
	defer d.Close()
	d.Sql().Exec("CREATE TABLE foo(id INT, name TEXT)")
	d.Sql().Exec("INSERT INTO foo VALUES(1,'a'),(2,'b')")
	m := New(d)
	if len(m.Tables()) == 0 {
		t.Fatal("no tables")
	}
	v := m.View()
	if v == "" {
		t.Fatal("empty view")
	}
}

func TestUpdateNav(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.db")
	d, _ := db.Open(path)
	defer d.Close()
	d.Sql().Exec("CREATE TABLE a(id INT)")
	d.Sql().Exec("CREATE TABLE b(id INT)")
	m := New(d)
	// j down
	m2, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
	m = m2.(Model)
	// k up
	m2, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}})
	m = m2.(Model)
	// r reload
	m2, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'r'}})
	_ = m2
	// q quit
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	if cmd == nil {
		t.Fatal("q should quit")
	}
}
