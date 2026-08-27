package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

type DB struct {
	sql *sql.DB
	path string
}

func Open(path string) (*DB, error) {
	sqldb, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	if _, err := sqldb.Exec("PRAGMA journal_mode=WAL"); err != nil {
		return nil, err
	}
	return &DB{sql: sqldb, path: path}, nil
}

func (d *DB) Close() error { return d.sql.Close() }
func (d *DB) Path() string { return d.path }
func (d *DB) Sql() *sql.DB { return d.sql }

func (d *DB) ListTables() ([]string, error) {
	rows, err := d.sql.Query("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%' ORDER BY name")
	if err != nil { return nil, err }
	defer rows.Close()
	var out []string
	for rows.Next() {
		var n string
		if err := rows.Scan(&n); err != nil { return nil, err }
		out = append(out, n)
	}
	return out, rows.Err()
}

func (d *DB) RowCount(table string) (int, error) {
	var c int
	// safe: table comes from ListTables, but also validate
	if table == "" || len(table) > 64 { return 0, fmt.Errorf("invalid table") }
	err := d.sql.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %q", table)).Scan(&c)
	return c, err
}

func (d *DB) Query(table string, limit int) ([]map[string]any, []string, error) {
	if limit <= 0 || limit > 100 { limit = 20 }
	rows, err := d.sql.Query(fmt.Sprintf("SELECT * FROM %q LIMIT ?", table), limit)
	if err != nil { return nil, nil, err }
	defer rows.Close()
	cols, _ := rows.Columns()
	var out []map[string]any
	for rows.Next() {
		vals := make([]any, len(cols))
		ptrs := make([]any, len(cols))
		for i := range vals { ptrs[i] = &vals[i] }
		if err := rows.Scan(ptrs...); err != nil { return nil, nil, err }
		m := map[string]any{}
		for i, c := range cols { m[c] = vals[i] }
		out = append(out, m)
	}
	return out, cols, rows.Err()
}

func (d *DB) WalSize() (int64, error) {
	fi, err := os.Stat(d.path + "-wal")
	if err != nil {
		if os.IsNotExist(err) { return 0, nil }
		return 0, err
	}
	return fi.Size(), nil
}
