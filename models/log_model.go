package models

type LogModel struct {
	Model
	LogType   int8      `json:"logType"`
	Title     string    `gorm:"size:64" json:"title"`
	Content   string    `json:"content"`
	Level     int8      `json:"level"`
	UserID    uint      `json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" jsson:"-"`
	IP        string    `gorm:"size:32" json:"IP"`
	Addr      string    `gorm:"size:64" json:"addr"`
	IsRead    bool      `json:"isRead"`
}
