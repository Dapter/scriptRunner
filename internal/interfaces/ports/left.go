package ports

import (
	"github.com/Dapter/scriptRunner/internal/data/command"
)

type Left interface {
	GetList() (map[string]command.Command, error)
	AddCommand(cmd command.Command) error
	UpdateCommand(cmd command.Command) error
	RemoveCommand(name string) error
	GetCommand(name string) (*command.Command, error)
	Run(name string) command.Results
	Start()
}
