package interfaces

import (
	"github.com/Dapter/scriptRunner/internal/data/command"
)

type App interface {
	GetCommand(name string) (*command.Command, error)
	GetList() (map[string]command.Command, error)
	AddCommand(cmd command.Command) error
	UpdateCommand(cmd command.Command) error
	RemoveCommand(name string) error
	Run(name string) command.Results
	LoadCommands()
}
