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

type TaskDelegate struct {
	styles TaskStyles
	keys   *TaskKeyMap
	editor Editor
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
	var (
		cmds    []tea.Cmd
		running = false
	)

	if v, ok := m.SelectedItem().(Task); ok {
		running = v.timer.Running()
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, t.keys.start, t.keys.stop):
			t.keys.start.SetEnabled(running)
			t.keys.stop.SetEnabled(!running)
			cmds = append(cmds, m.SelectedItem().(Task).timer.Toggle())
		case key.Matches(msg, t.keys.edit):
			cmds = append(cmds, t.Edit(m.SelectedItem().(Task)))
		default:
			t.keys.start.SetEnabled(!running)
			t.keys.stop.SetEnabled(running)
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

func (t TaskDelegate) Edit(task Task) tea.Cmd {
	t.editor.Load(task)
	return func() tea.Msg {
		return StateMsg{state: taskEdit}
	}
}
