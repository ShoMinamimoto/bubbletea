package main

// A simple program that counts down from 5 and then exits.

import (
	"errors"
	"fmt"
	"log"

	"github.com/charmbracelet/boba"
	"github.com/charmbracelet/teaparty/input"
)

type Model struct {
	textInput input.Model
	err       error
}

type tickMsg struct{}
type errMsg error

func main() {
	p := boba.NewProgram(
		initialize,
		update,
		view,
		subscriptions,
	)

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func initialize() (boba.Model, boba.Cmd) {
	inputModel := input.DefaultModel()
	inputModel.Placeholder = "Pikachu"

	return Model{
		textInput: inputModel,
		err:       nil,
	}, nil
}

func update(msg boba.Msg, model boba.Model) (boba.Model, boba.Cmd) {
	var cmd boba.Cmd
	m, ok := model.(Model)
	if !ok {
		// When we encounter errors in Update we simply add the error to the
		// model so we can handle it in the view. We could also return a command
		// that does something else with the error, like logs it via IO.
		return Model{
			err: errors.New("could not perform assertion on model in update"),
		}, nil
	}

	switch msg := msg.(type) {
	case boba.KeyMsg:
		switch msg.Type {
		case boba.KeyCtrlC:
			fallthrough
		case boba.KeyEsc:
			return m, boba.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = input.Update(msg, m.textInput)
	return m, cmd
}

func subscriptions(model boba.Model) boba.Subs {
	return boba.Subs{
		// We just hand off the subscription to the input component, giving
		// it the model it expects.
		"input": func(model boba.Model) boba.Msg {
			m, _ := model.(Model)
			return input.Blink(m.textInput)
		},
	}
}

func view(model boba.Model) string {
	m, ok := model.(Model)
	if !ok {
		return "Oh no: could not perform assertion on model."
	} else if m.err != nil {
		return fmt.Sprintf("Uh oh: %s", m.err)
	}
	return fmt.Sprintf(
		"What’s your favorite Pokémon?\n\n%s\n\n%s",
		input.View(m.textInput),
		"(esc to quit)",
	)
}
