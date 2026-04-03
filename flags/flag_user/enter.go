package flag_user

import (
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/utils/pwd"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

type FlagUser struct {
}

func (FlagUser) Create() {
	var role enum.RoleType
	fmt.Println("choose role   1 super admin   2 user  3 visitor")
	_, err := fmt.Scan(&role)
	if err != nil {
		logrus.Errorf("input err %s", err)
		return
	}
	if !(role == 1 || role == 2 || role == 3) {
		logrus.Errorf("role error %s", err)
		return
	}
	var username string
	fmt.Println("input username:")
	fmt.Scan(&username)

	var model models.UserModel
	err = global.DB.Take(&model, "username = ?", username).Error
	if err == nil {
		logrus.Errorf("user exist")
		return
	}
	fmt.Println("input password:")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		logrus.Errorf("password read error %s", err)
		return
	}
	fmt.Println("input password again:")
	rePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		logrus.Errorf("password read error %s", err)
		return
	}
	if string(password) != string(rePassword) {
		logrus.Errorf("password not same")
		return
	}
	hashPwd, _ := pwd.GenerateFromPassword(string(password))
	err = global.DB.Create(&models.UserModel{
		Username:       username,
		Nickname:       "user001",
		RegisterSource: enum.RegisterTerminalSourceType,
		Password:       hashPwd,
		Role:           role,
	}).Error
	if err != nil {
		logrus.Errorf("create user error %s", err)
		return
	}
	logrus.Infof("create user successful")
}
