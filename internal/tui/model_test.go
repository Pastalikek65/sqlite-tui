package tui

import (
	"path/filepath"
	"testing"

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
