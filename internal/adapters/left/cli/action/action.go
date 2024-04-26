package action

import (
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/add"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/delete"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/display"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/help"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/list"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/run"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/types"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/update"
	"github.com/Dapter/scriptRunner/internal/interfaces/ports"
	"log"
	"os"
)

// actions
const (
	listAction    = "list"
	addAction     = "add"
	updateAction  = "edit"
	deleteAction  = "delete"
	displayAction = "display"
	runAction     = "run"
	helpAction    = "help"
)

var availableActions = []string{
	listAction,
	addAction,
	updateAction,
	deleteAction,
	displayAction,
	runAction,
	helpAction,
}

func init() {
	if len(os.Args) == 1 {
		log.Fatalf("No command provided, use 'help' to display availible commands")
	}
}

func Create(c *ports.Left) (action types.Action) {
	command := os.Args[1]

	switch command {
	case helpAction:
		help := *help.Create()

		action = types.Action(help)
	case listAction:
		list := *list.Create(c)

		action = types.Action(list)
	case runAction:
		run := *run.Create(c)

		action = types.Action(run)
	case addAction:
		add := *add.Create(c)

		action = types.Action(add)
	case updateAction:
		update := *update.Create(c)

		action = types.Action(update)
	case deleteAction:
		delete := *delete.Create(c)

		action = types.Action(delete)
	case displayAction:
		display := *display.Create(c)

		action = types.Action(display)
	default:
		log.Fatalf("Command '%s' not found, run 'help' to check available commands", command)
	}

	return
}
