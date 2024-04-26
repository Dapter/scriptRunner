package app

import (
	"fmt"
	myCore "github.com/Dapter/scriptRunner/internal/core"
	"github.com/Dapter/scriptRunner/internal/data/command"
	"github.com/Dapter/scriptRunner/internal/interfaces"
	"github.com/Dapter/scriptRunner/internal/interfaces/ports"
	"log"
)

type App struct {
	rightAdapter ports.Right
	core         interfaces.Core
	commands     command.Commands
}

func New(rightAdapter ports.Right) *App {
	commands, err := rightAdapter.GetList()

	if err != nil {
		log.Fatalf("cannot create application layer: %v", err)
	}

	core := myCore.New()

	return &App{rightAdapter, core, commands}
}

func (api App) Run(name string) command.Results {
	command, ok := api.commands[name]

	if !ok {
		log.Fatalf("cannot find command named %s", name)
	}

	return api.core.Run(command)
}

func (api *App) LoadCommands() {
	commands, err := api.GetList()

	if err != nil {
		log.Fatalf("cannot load commands to app: %v", err)
	}

	api.commands = commands
}

func (api App) GetCommand(name string) (*command.Command, error) {
	command, ok := api.commands[name]

	if !ok {
		return nil, fmt.Errorf("command with name '%s' not found", name)
	}

	return &command, nil
}

func (api App) RemoveCommand(name string) error {
	command, ok := api.commands[name]

	if !ok {
		return fmt.Errorf("cannot remove command '%s', command not found", name)
	}

	return api.rightAdapter.RemoveCommand(command)
}

func (api App) GetList() (map[string]command.Command, error) {
	return api.commands, nil
}

func (api App) AddCommand(cmd command.Command) error {
	return api.rightAdapter.AddCommand(cmd)
}

func (api App) UpdateCommand(cmd command.Command) error {
	return api.rightAdapter.UpdateCommand(cmd)
}
