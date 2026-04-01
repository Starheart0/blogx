package conf

type System struct {
	Ip      string `yaml:"ip"`
	Port    int    `yaml:"port"`
	GinMode string `yaml:"gin_mode"`
	Env     string `yaml:"env"`
}
