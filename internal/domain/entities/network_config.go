package entities

type Script struct {
	Name          string              `yaml:"name"`
	Description   string              `yaml:"description"`
	Version       string              `yaml:"version"`
	Variables     []Variable          `yaml:"variables"`
	Commands      map[string][]string `yaml:"commands"` // OS -> commands
	Prerequisites []Prerequisite      `yaml:"prerequisites"`
}

type Variable struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
}

type Prerequisite struct {
	Command     string `yaml:"command"`
	Description string `yaml:"description"`
}
