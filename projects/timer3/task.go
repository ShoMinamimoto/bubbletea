package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"time"
)

type Task struct {
	label       string
	description string
	timer       stopwatch.Model
}

func (t Task) FilterValue() string {
	return ""
}

func (t Task) Update(msg tea.Msg) (Task, tea.Cmd) {
	var cmd tea.Cmd
	t.timer, cmd = t.timer.Update(msg)
	return t, cmd
}

func newTask(i int) Task {
	return Task{
		label:       fmt.Sprintf("Task %d", i),
		description: "",
		timer:       stopwatch.NewWithInterval(time.Second),
	}
}

type TaskStyles struct {
	NormalLabel   lipgloss.Style
	NormalTimer   lipgloss.Style
	SelectedLabel lipgloss.Style
	SelectedTimer lipgloss.Style
	ActiveLabel   lipgloss.Style
	ActiveTimer   lipgloss.Style
}

func NewTaskStyles() (s TaskStyles) {
	s.NormalLabel = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
		Padding(0, 0, 0, 2)

	s.NormalTimer = s.NormalLabel.Copy().Padding(0, 2, 0, 0)

	s.SelectedLabel = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#f792ff", Dark: "#ad58b4"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#ee6ff8", Dark: "#ee6ff8"}).
		Padding(0, 0, 0, 1)

	s.SelectedTimer = s.SelectedLabel.Copy().Padding(0, 1, 0, 0).
		Border(lipgloss.NormalBorder(), false, true, false, false)

	s.ActiveLabel = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#6F73F8", Dark: "#6F73F8"}).
		Padding(0, 0, 0, 2)

	s.ActiveTimer = s.ActiveLabel.Copy().Padding(0, 2, 0, 0)

	return s
}

type TaskKeyMap struct {
	start key.Binding
	stop  key.Binding
}

func NewTaskKeyMap() *TaskKeyMap {
	return &TaskKeyMap{
		start: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "start timer"),
		),
		stop: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "stop timer"),
			key.WithDisabled(),
		),
	}
}

type TaskDelegate struct {
	styles TaskStyles
	keys   *TaskKeyMap
}

func (t TaskDelegate) ShortHelp() []key.Binding {
	return []key.Binding{
		t.keys.start,
		t.keys.stop,
	}
}

func (t TaskDelegate) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		t.ShortHelp(),
	}
}

func newTaskDelegate() TaskDelegate {
	return TaskDelegate{
		styles: NewTaskStyles(),
		keys:   NewTaskKeyMap(),
	}
}

func (t TaskDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var (
		label  = item.(Task).label
		timer  = item.(Task).timer.Elapsed().String()
		styles = &t.styles
	)

	if m.Width() <= 0 {
		return
	}

	var (
		isSelected = index == m.Index()
		isActive   = item.(Task).timer.Running()
	)

	switch {
	case isSelected:
		label = styles.SelectedLabel.Render(label)
		timer = styles.SelectedTimer.Render(timer)
	case isActive:
		label = styles.ActiveLabel.Render(label)
		timer = styles.ActiveTimer.Render(timer)
	default:
		label = styles.NormalLabel.Render(label)
		timer = styles.NormalTimer.Render(timer)
	}

	timer = lipgloss.PlaceHorizontal(m.Width()-lipgloss.Width(label), lipgloss.Right, timer)
	line := lipgloss.JoinHorizontal(lipgloss.Center, label, timer)

	fmt.Fprintf(w, "%s", line)
}

func (t TaskDelegate) Height() int {
	return 1
}

func (t TaskDelegate) Spacing() int {
	return 1
}

func (t TaskDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, t.keys.start, t.keys.stop):
			t.keys.start.SetEnabled(m.SelectedItem().(Task).timer.Running())
			t.keys.stop.SetEnabled(!m.SelectedItem().(Task).timer.Running())
			cmds = append(cmds, m.SelectedItem().(Task).timer.Toggle())
		default:
			t.keys.start.SetEnabled(!m.SelectedItem().(Task).timer.Running())
			t.keys.stop.SetEnabled(m.SelectedItem().(Task).timer.Running())
		}
	}

	for index, task := range m.Items() {
		if task.(Task).timer.Running() || index == m.Index() {
			newTask, cmd := task.(Task).Update(msg)
			cmds = append(cmds, m.SetItem(index, newTask), cmd)
		}
	}

	return tea.Batch(cmds...)
}
