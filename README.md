
# 📝 todo-cli

A terminal-based todo list app built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and friends 🫖.

## 🚀 Features

- Add, toggle, remove, and update tasks.
- Navigate with keyboard.
- Pretty terminal UI using Bubble Tea and Bubbles.

## 📦 Requirements

- **Go 1.21 or later**
  This project was tested with **Go 1.24.4**.

### Install Go

If you haven't installed Go yet, download it from the official website:

🔗 [https://go.dev/dl](https://go.dev/dl)

Then verify installation:

```bash
go version
```

You should see something like:

```bash
go version go1.24.4 linux/amd64
```

## 📥 Installation

### Option 1: Install via `go install`

```bash
go install github.com/jimtrung/todo-cli@v0.1.0
```

This will install `todo-cli` into your `$GOBIN` (usually `$HOME/go/bin`). Make sure it’s in your `$PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Then run:

```bash
todo-cli
```

### Option 2: Clone and Run

```bash
git clone https://github.com/jimtrung/todo-cli.git
cd todo-cli
go run main.go
```

## 📦 Dependencies

This app uses:

- `bubbletea` for terminal UI
- `bubbles` for components
- `lipgloss` for styling
- And other helpful packages from the Charm ecosystem

All dependencies are managed in `go.mod`.

## 📄 License

MIT © 2025 [jimtrung](https://github.com/jimtrung)
