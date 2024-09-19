# GO TIME TUI

This is a simple time project using ![golang](https://go.dev),
![bubbletea](https://github.com/charmbracelet/bubbletea)
and ![lipgloss](https://github.com/charmbracelet/lipgloss)

## How to build

```bash
git clone https://github.com/zakwanhisham/go-time.git
make build
```

## How to install

```bash
git clone https://github.com/zakwanhisham/go-time.git
make install
# Put this on your rc file (.bashrc, .zshrc)
export PATH=$PATH:~/.local/bin
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

## TODO

- [ ] Fix the text styling
- [ ] Make the font larger
- [ ] Make it more interactive
- [ ] Add more feature
  - [ ] Alarm
  - [ ] Reminder
  - [ ] Timer Alarm
  - [ ] Countdown Alarm
