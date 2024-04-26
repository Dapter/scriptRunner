package command

import (
	"bytes"
	"fmt"
	"strings"
)

type Command struct {
	Name     string
	Icon     string
	Commands []string
}

type Commands map[string]Command

type Results struct {
	CommandName string
	Items       map[string]Result
}

type Result struct {
	Output, ErrorOutput bytes.Buffer
	Err                 error
}

func (results Results) String() string {
	var output = fmt.Sprintf("-- %s --\n", results.CommandName)

	for cmd, result := range results.Items {
		subOutput := " successfully"
		if result.String() != "" {
			subOutput = fmt.Sprintf(", with output: \n%s", result)
		}
		output += fmt.Sprintf("Command '%s' runned%s\n", cmd, subOutput)
	}

	return output
}

func (result Result) String() string {
	var output = result.Output

	if result.Err != nil {
		output = result.ErrorOutput
	}

	return output.String()
}

func (cmd Command) String() string {
	return cmd.Name + ": " + strings.Join(cmd.Commands, "\n")
}
