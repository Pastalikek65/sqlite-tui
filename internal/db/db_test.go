package db

import (
	"path/filepath"
	"testing"
)

func TestOpenAndListTables(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.db")
	db, err := Open(path)
	if err != nil { t.Fatalf("open failed: %v", err) }
	defer db.Close()
	// create table
	if _, err := db.sql.Exec("CREATE TABLE foo(id INT)"); err != nil { t.Fatal(err) }
	tables, err := db.ListTables()
	if err != nil { t.Fatal(err) }
	if len(tables) != 1 || tables[0] != "foo" {
		t.Fatalf("want [foo] got %v", tables)
	}
	// row count
	c, err := db.RowCount("foo")
	if err != nil { t.Fatal(err) }
	if c != 0 { t.Fatalf("want 0 got %d", c) }
	// query
	rows, cols, err := db.Query("foo", 10)
	if err != nil { t.Fatal(err) }
	if len(cols) != 1 || cols[0] != "id" { t.Fatalf("cols %v", cols) }
	if len(rows) != 0 { t.Fatalf("want 0 rows got %d", len(rows)) }
}
