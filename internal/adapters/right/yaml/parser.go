package yaml

import (
	"fmt"
	"github.com/Dapter/scriptRunner/internal/data/command"
	"github.com/Dapter/scriptRunner/internal/data/config"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type parsedCommands struct {
	commands       command.Commands
	configFilePath string
	rawData        map[interface{}]map[interface{}]interface{}
}

func read(configFilePath string) map[interface{}]map[interface{}]interface{} {
	var data map[interface{}]map[interface{}]interface{}
	yfile, err := os.ReadFile(configFilePath)

	if err != nil {
		log.Fatalf("cannot read yaml config file: %v", err)
	}

	err = yaml.Unmarshal(yfile, &data)

	if err != nil {
		log.Fatalf("cannot parse yaml config file: %v", err)
	}

	return data
}

func New() *parsedCommands {
	configFilePath := config.GetConfigFilePath()
	commands := make(command.Commands)
	data := read(configFilePath)

	for k, v := range data {
		commandName := fmt.Sprintf("%s", k)
		commandsToRun, ok := v["command"]

		if !ok {
			log.Println("cannot load command for %s", commandName)
			continue
		}

		icon, ok := v["icon"]

		if !ok {
			icon = ""
		}

		cmd := command.Command{
			Name:     fmt.Sprintf("%s", k),
			Commands: getCommandsAsString(commandsToRun),
			Icon:     fmt.Sprintf("%s", icon),
		}

		commands[cmd.Name] = cmd
	}

	return &parsedCommands{commands, configFilePath, data}
}

func write(configFilePath string, data map[interface{}]map[interface{}]interface{}) {
	bData, err := yaml.Marshal(&data)

	if err != nil {
		log.Fatalf("cannot parse commands to write it to yaml config file: %v", err)
	}

	// write to file
	f, err := os.Create(configFilePath)

	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(configFilePath, bData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (yamlData *parsedCommands) RemoveCommand(cmd command.Command) error {
	delete(yamlData.rawData, cmd.Name)

	write(yamlData.configFilePath, yamlData.rawData)

	return nil
}

func getCommandsAsString(commands interface{}) []string {
	r := reflect.ValueOf(commands)
	switch r.Kind() {
	case reflect.String:
		return []string{r.String()}
	case reflect.Slice:
		var result = []string{}
		for i := 0; i < r.Len(); i++ {
			result = append(result, fmt.Sprintf("%s", r.Index(i)))
		}

		return result
	default:
		return []string{}
	}
}

func (yamlData *parsedCommands) GetCommand(name string) (*command.Command, error) {
	cmd, ok := yamlData.commands[name]

	if ok == false {
		return nil, fmt.Errorf("cannot find command named %s in loaded command list, check config file", name)
	}

	return &cmd, nil
}

func (yamlData *parsedCommands) GetList() (map[string]command.Command, error) {
	return yamlData.commands, nil
}

func (yamlData *parsedCommands) AddCommand(cmd command.Command) error {
	_, exists := yamlData.rawData[cmd.Name]
	if exists {
		return fmt.Errorf("cannot add command, command with such name already exists")
	}

	yamlData.rawData[cmd.Name] = make(map[interface{}]interface{})
	yamlData.rawData[cmd.Name]["command"] = cmd.Commands[0]

	write(yamlData.configFilePath, yamlData.rawData)

	return nil
}

func (yamlData *parsedCommands) UpdateCommand(cmd command.Command) error {
	_, exists := yamlData.rawData[cmd.Name]
	if !exists {
		return fmt.Errorf("cannot edit command, command with such name not exists")
	}

	yamlData.rawData[cmd.Name] = make(map[interface{}]interface{})
	yamlData.rawData[cmd.Name]["command"] = cmd.Commands

	write(yamlData.configFilePath, yamlData.rawData)

	return nil
}
