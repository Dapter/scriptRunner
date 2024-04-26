package config

import "os"

func GetConfigFilePath() string {
	configWorkDir := os.Getenv("CONFIG_WORK_DIR")

	if configWorkDir == "" {
		panic("Env 'CONFIG_WORK_DIR' required!")
	}

	configFilePath := os.Getenv("CONFIG_FILE_PATH")

	if configWorkDir == "" {
		panic("Env 'CONFIG_FILE_PATH' required!")
	}

	return configWorkDir + "/" + configFilePath
}
