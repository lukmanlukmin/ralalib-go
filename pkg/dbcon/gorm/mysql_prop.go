package dbcon

// MSQLProp :nodoc:
type MSQLProp struct {
	Name          string `yaml:"name"`
	User          string `yaml:"user"`
	Pass          string `yaml:"pass"`
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	MaxOpen       int    `yaml:"max_open"`
	MaxIdle       int    `yaml:"max_idle"`
	TimeoutSecond int    `yaml:"timeout_second"`
	MaxLifeTimeMS int    `yaml:"life_time_ms"`
	Charset       string `yaml:"charset"`
}
