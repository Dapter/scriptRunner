package ports

import "github.com/Dapter/scriptRunner/internal/data/command"

type Right interface {
	GetCommand(name string) (*command.Command, error)
	GetList() (map[string]command.Command, error)
	AddCommand(cmd command.Command) error
	UpdateCommand(cmd command.Command) error
	RemoveCommand(cmd command.Command) error
}
