# Changelog

All notable changes to `sqlite-tui` will be documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-08-25
### Added
- SQLite WAL `ListTables`, `RowCount`, `Query(limit 20)`, `WalSize` at any `*.db`
- TUI `bubbletea` `j/k` nav, `r` reload, `q` quit, `WalSize` header, `lipgloss` styles
- Single binary 8.5M, `CGO_ENABLED=0`, `modernc.org/sqlite` 1.38.2, `cobra` CLI
- 1 test `TestOpenAndListTables` <0.1s
