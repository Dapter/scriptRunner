package cli

import (
	myAction "github.com/Dapter/scriptRunner/internal/adapters/left/cli/action"
	"github.com/Dapter/scriptRunner/internal/data/command"
	"github.com/Dapter/scriptRunner/internal/interfaces"
	"github.com/Dapter/scriptRunner/internal/interfaces/ports"
	"os"
)

type cli struct {
	app interfaces.App
}

func New(app interfaces.App) *cli {
	return &cli{app}
}

func (c *cli) Start() {
	left := ports.Left(c)
	action := myAction.Create(&left)

	for _, arg := range os.Args[2:] {
		if arg == "--help" {
			action.Help()
			os.Exit(0)
		}
	}

	action.Run()
}

func (c cli) Run(name string) command.Results {
	return c.app.Run(name)
}

func (c cli) RemoveCommand(name string) error {
	return c.app.RemoveCommand(name)
}

func (c cli) GetCommand(name string) (*command.Command, error) {
	return c.app.GetCommand(name)
}

func (c cli) GetList() (map[string]command.Command, error) {
	return c.app.GetList()
}

func (c cli) AddCommand(cmd command.Command) error {
	return c.app.AddCommand(cmd)
}

func (c cli) UpdateCommand(cmd command.Command) error {
	return c.app.UpdateCommand(cmd)
}
