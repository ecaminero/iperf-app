package entities

type Config struct {
	App       AppConfig       `yaml:"app"`
	Execution ExecutionConfig `yaml:"execution"`
	Scripts   ScriptsConfig   `yaml:"scripts"`
	Iperf     IperfConfig     `yaml:"iperf"`
}

type AppConfig struct {
	Name     string `yaml:"name"`
	Version  string `yaml:"version"`
	LogLevel string `yaml:"log_level"`
}

type ExecutionConfig struct {
	Timeout          string   `yaml:"timeout"`
	WorkingDirectory string   `yaml:"working_directory"`
	Environment      []string `yaml:"environment_variables"`
}

type ScriptsConfig struct {
	Directory       string `yaml:"directory"`
	ConfigDirectory string `yaml:"config_directory"`
}

type IperfConfig struct {
	DefaultServers map[string]string `yaml:"default_servers"`
	DefaultTimeout int               `yaml:"default_timeout"`
	DefaultPath    string            `yaml:"default_path"`
}
