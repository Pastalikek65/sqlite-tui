# Contributing to sqlite-tui

```bash
git clone https://github.com/Pastalikek65/sqlite-tui.git
cd sqlite-tui
go vet ./...
go test ./... -count=1 -timeout 30s
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o sqlite-tui .
./sqlite-tui test.db
```

- Keep it phone-friendly: `CGO_ENABLED=0`, `8.5M`, `<40MB RAM`.
- TDD: failing test first.

PRs: fork → `feat/<name>` → `go test` → open PR.
