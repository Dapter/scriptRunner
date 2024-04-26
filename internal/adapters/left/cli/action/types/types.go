package types

import "github.com/Dapter/scriptRunner/internal/interfaces/ports"

type ActionStruct struct {
	Cli        *ports.Left
	Parameters map[string]string
}

type Action interface {
	Run()
	Help()
}
