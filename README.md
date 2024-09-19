# GO TIME TUI

This is a simple time project using ![golang](https://go.dev),
![bubbletea](https://github.com/charmbracelet/bubbletea)
and ![lipgloss](https://github.com/charmbracelet/lipgloss)

## How to build

```bash
git clone https://github.com/zakwanhisham/go-time.git
make build
mv gotime ~/.local/bin/
export PATH=PATH:~/.local/bin
```

## How to use

- Normal clock

```bash
gotime

```

- Countdown mode

```bash
gotime -countdown=<time in seconds>
```

- Timer mode

```bash
gotime -timer`
```
