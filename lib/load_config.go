package lib

import (
	"github.com/pelletier/go-toml/v2"
)

type ConfigData struct {
	DistroName *string
}

func LoadConfig(filename string) ConfigData {
	var Config ConfigData
	fileBin := Load_file(filename)
	tomlUnmarshalErr := toml.Unmarshal(fileBin, &Config)
	if tomlUnmarshalErr != nil {
		panic(tomlUnmarshalErr)
	}
	return Config
}

func (config ConfigData) Init() ConfigData {
	if config.DistroName == nil {
		*config.DistroName = "Arch Linux"
	}
	return config
}
