package yamlconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Theme    string    `yaml:"theme"`
	Title    string    `yaml:"title"`
	Subtitle string    `yaml:"subtitle"`
	Sections []Section `yaml:"sections"`
}

type Section struct {
	Name    string  `yaml:"name"`
	Icon    string  `yaml:"icon"`
	Entries []Entry `yaml:"entries"`
}

type Entry struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type Colors struct {
	Title           string `yaml:"title"`
	Background      string `yaml:"background"`
	Text            string `yaml:"text"`
	BackgroundLight string `yaml:"background-light"`
}

func Parse() (Config, Colors, error) {
	config := Config{}
	colors := Colors{}

	data, err := os.ReadFile("config.yml")
	if err != nil {
		return config, colors, err
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return config, colors, err
	}

	data, err = os.ReadFile("themes/" + config.Theme + ".yml")
	if err != nil {
		return config, colors, err
	}

	err = yaml.Unmarshal([]byte(data), &colors)
	if err != nil {
		return config, colors, err
	}

	return config, colors, nil
}
