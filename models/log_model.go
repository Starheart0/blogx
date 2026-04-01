package models

import "blogx_server/models/enum"

type LogModel struct {
	Model
	LogType     enum.LogType      `json:"logType"`
	Title       string            `gorm:"size:64" json:"title"`
	Content     string            `json:"content"`
	Level       enum.LogLevelType `json:"level"`
	UserID      uint              `json:"userID"`
	UserModel   UserModel         `gorm:"foreignKey:UserID" jsson:"-"`
	IP          string            `gorm:"size:32" json:"IP"`
	Addr        string            `gorm:"size:64" json:"addr"`
	IsRead      bool              `json:"isRead"`
	LoginStatus bool              `json:"loginStatus"`
	Username    string            `gorm:"size:32" json:"username"`
	Pwd         string            `gorm:"size:32" json:"pwd"`
	LoginType   enum.LoginType    `json:"loginType"`
	ServiceName string            `gorm:"size:32" json:"serviceName"`
}
