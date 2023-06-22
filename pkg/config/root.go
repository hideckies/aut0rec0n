package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ApiKeys struct {
	GitHub     string `yaml:"github"`
	Shodan     string `yaml:"shodan"`
	VirusTotal string `yaml:"virusTotal"`
}

type Config struct {
	ApiKeys ApiKeys `yaml:"apikeys"`
}

// Create config files if not exist
func createConfigFiles() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Check if the .config directory exists under $HOME
	configDir := homeDir + "/.config"
	_, err = os.Stat(configDir)
	if os.IsNotExist(err) {
		// Create new .config directory
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// Check if the aut0rec0n directory exists under $HOME/.config/
	aut0rec0nDir := configDir + "/aut0rec0n"
	_, err = os.Stat(aut0rec0nDir)
	if os.IsNotExist(err) {
		// Create new aut0rec0n directory
		err = os.MkdirAll(aut0rec0nDir, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// Check if the config.yaml exists under $HOME/.config/aut0rec0n/
	configFile := aut0rec0nDir + "/config.yaml"
	_, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		// Create new config.yaml
		file, err := os.Create(configFile)
		if err != nil {
			return err
		}
		defer file.Close()

		// Write the YAML file
		conf := Config{}
		yamlData, err := yaml.Marshal(&conf)
		if err != nil {
			return err
		}

		err = os.WriteFile(configFile, yamlData, 0644)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func loadConfigFiles() (Config, error) {
	conf := Config{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return conf, err
	}

	aut0rec0nDir := homeDir + "/.config/aut0rec0n"

	// Load config
	configFile := aut0rec0nDir + "/config.yaml"
	loadedConfig, err := os.ReadFile(configFile)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(loadedConfig, &conf)
	if err != nil {
		return conf, err
	}

	// Update the config file
	updatedConfig, err := yaml.Marshal(&conf)
	if err != nil {
		return conf, err
	}

	err = os.WriteFile(configFile, updatedConfig, 0644)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

// The function to execute
func Execute() (Config, error) {
	conf := Config{}

	err := createConfigFiles()
	if err != nil {
		return conf, err
	}

	conf, err = loadConfigFiles()
	if err != nil {
		return conf, err
	}

	return conf, nil
}
