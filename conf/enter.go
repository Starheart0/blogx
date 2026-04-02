package conf

import "fmt"

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     DB     `yaml:"db"`  //read
	DB1    DB     `yaml:"db1"` //write
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
