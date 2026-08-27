# Security Policy

## Supported Versions

| Version | Supported |
| ------- | --------- |
| 0.1.x   | :white_check_mark: |

## Reporting a Vulnerability

Use GitHub private advisory: https://github.com/Pastalikek65/sqlite-tui/security/advisories/new

Response within 48h, fix within 14 days.

## Notes

- Opens local SQLite files only — no network, no `shell` injection (table names are quoted via `%q`).
- WAL size is read via `os.Stat(path + "-wal")` — no exec.
