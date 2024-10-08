package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

type model struct {
	time    time.Time
	timer   *time.Duration
	countup *time.Duration
	width   int
	height  int
}

var (
	timeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#c4b28a")).
			Background(lipgloss.Color("#181616")).
			Padding(1, 4).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#c5c9c5")).
			Align(lipgloss.Center)

	emptyStyle = lipgloss.NewStyle()
)

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.time = time.Time(msg)
		if m.timer != nil && *m.timer > 0 {
			*m.timer -= time.Second
		}
		if m.countup != nil && *m.countup > 0 {
			*m.countup += time.Second
		}
		return m, tick()
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "esc" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m model) View() string {
	var timeView, textView string
	if m.timer != nil && *m.timer > 0 {
		remainingTime := int64(time.Duration(*m.timer).Seconds())
		minutes := remainingTime / 60
		seconds := remainingTime % 60
		timeView = timeStyle.Render(fmt.Sprintf("%02d:%02d", minutes, seconds))
	} else if m.timer != nil && *m.timer <= 0 {
		timeView = timeStyle.Render("00:00")
	} else if m.countup != nil && *m.countup >= 0 {
		countTime := int64(time.Duration(*m.countup).Seconds())
		minutes := countTime / 60
		seconds := countTime % 60
		timeView = timeStyle.Render(fmt.Sprintf("%02d:%02d", minutes, seconds))
	} else {
		timeView = timeStyle.Render(m.time.Format("Mon, Jan 2 2006 \n 15:04:05"))
	}
	combinedView := fmt.Sprintf("%s\n\n%s", textView, timeView)

	verticalMargin := (m.height - lipgloss.Height(combinedView)) / 2
	horizontalMargin := (m.width - lipgloss.Width(combinedView)) / 2

	return emptyStyle.
		Margin(verticalMargin, horizontalMargin).
		Render(combinedView)
}

func main() {
	var timerArg string
	var countupFlag bool

	flag.StringVar(
		&timerArg,
		"timer",
		"",
		"Countdown timer duration in seconds (e.g. -timer=120)",
	)
	flag.BoolVar(&countupFlag, "countup", false, "Countup mode")
	flag.Parse()

	var countdown *time.Duration
	if timerArg != "" {
		countdownSec, err := strconv.Atoi(timerArg)
		if err != nil {
			fmt.Println("Invalid countdown value. Please provide a number in seconds.")
			os.Exit(1)
		}
		duration := time.Duration(countdownSec) * time.Second
		countdown = &duration
	}

	var timer *time.Duration
	if countupFlag {
		defaultTimerDuration := time.Second
		timer = &defaultTimerDuration
	}

	m := model{time: time.Now(), timer: countdown, countup: timer}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v", err)
		return
	}
}
