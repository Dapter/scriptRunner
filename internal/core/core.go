package core

import (
	"github.com/Dapter/scriptRunner/internal/data/command"
	"os/exec"
	"strings"
)

type Core struct{}

func New() *Core {
	return &Core{}
}

func (c Core) Run(cmd command.Command) command.Results {
	results := make(map[string]command.Result)
	for _, command := range cmd.Commands {
		results[command] = runCommand(command)
	}

	return command.Results{
		CommandName: cmd.Name, Items: results,
	}
}

func runCommand(command string) (result command.Result) {
	var err error
	cmdParts := strings.Split(command, " ")
	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	cmd.Stdout = &result.Output
	cmd.Stderr = &result.ErrorOutput
	err = cmd.Start()
	if err == nil {
		err = cmd.Wait()
	}

	if err != nil || !cmd.ProcessState.Success() {
		result.Err = err
	}

	return
}
