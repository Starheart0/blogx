package conf

import "fmt"

type ES struct {
	Addr     string `yaml:"addr"`
	Ishttps  bool   `yaml:"ishttps"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Enable   bool   `yaml:"enable"`
}

func (e ES) Url() string {
	if e.Ishttps {
		return fmt.Sprintf("https://%s", e.Addr)
	}
	return fmt.Sprintf("http://%s", e.Addr)
}
