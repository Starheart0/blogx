package models

import "time"

type UserModel struct {
	Model
	Username       string `gorm:"size:32" json:"username"`
	Nickname       string `gorm:"size:32" json:"nickname"`
	Avatar         string `gorm:"size:256" json:"avatar"`
	Abstract       string `gorm:"size:256" json:"abstract"`
	RegisterSource int8   `json:"registerSource"`
	CodeAge        int    `json:"codeAge"`
	Password       string `gorm:"size:64" json:"-"`
	Email          string `gorm:"size:256" json:"email"`
	OpenID         string `gorm:"size:64" json:"openID"` //Third-party login unique value
	Role           int8   `json:"role"`                  //1 Administrator 2 normal user 3 visitor
}

type UserConfModel struct {
	UserID             uint      `gorm:"unique" json:"userID"`
	UserModel          UserModel `gorm:"foreignKey:UserID" json:"-"`
	LikeTags           []string  `gorm:"type:longtext;serializer:json" json:"likeTags"`
	UpdataUserNameDate time.Time `json:"updataUserNameDate"`
	OpenCollect        bool      `json:"openCollect"` //open user collect
	OpenFollow         bool      `json:"openFollow"`
	OpenFans           bool      `json:"openFans"`
	HomeStyleID        uint      `json:"homeStyleID"`
}
