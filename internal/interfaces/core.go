package interfaces

import "github.com/Dapter/scriptRunner/internal/data/command"

type Core interface {
	Run(command command.Command) command.Results
}
