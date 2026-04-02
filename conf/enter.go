package conf

import "fmt"

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	DB     DB     `yaml:"db"`  //read
	DB1    DB     `yaml:"db1"` //write
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
	Site   Site   `yaml:"site"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	QiNiu  QiNiu  `yaml:"qiNiu"`
	Ai     Ai     `yaml:"ai"`
	Upload Upload `yaml:"upload"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
