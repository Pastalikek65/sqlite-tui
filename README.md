# sqlite-tui

[![Go 1.25](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Termux](https://img.shields.io/badge/Termux-friendly-brightgreen?logo=android)](https://termux.dev)

Single-binary TUI viewer for SQLite — tables, rows, WAL size. No CGO, 8.5M, `<40MB RAM`.

I built it because `sqlite3` CLI is plain and `DB Browser` needs GUI. This shows your tables and rows in the terminal on your phone.

```bash
go install github.com/Pastalikek65/sqlite-tui@latest
sqlite-tui my.db
# j/k nav, r reload, q quit — shows WAL size and row preview
```

Works with `ceparsiv`'s `~/.cache/ceparsiv.db`, `pytermq`'s `queue.db`, or any SQLite.

## Install (Termux)

```bash
git clone https://github.com/Pastalikek65/sqlite-tui.git
cd sqlite-tui
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o sqlite-tui .
./sqlite-tui /path/to/db
```

## Why not DB Browser

DB Browser needs X11, 50MB+. This is 8.5M, pure Go `modernc.org/sqlite`, WAL aware, offline.

## Tech

Go 1.25, `bubbletea` 1.3.4, `lipgloss`, `modernc.org/sqlite` 1.38.2, `cobra`. 1 table `ListTables`, `RowCount`, `Query` with `LIMIT 20`.

## License

MIT
