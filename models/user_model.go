package models

import (
	"blogx_server/models/enum"
	"math"
	"time"
)

type UserModel struct {
	Model
	Username       string                  `gorm:"size:32" json:"username"`
	Nickname       string                  `gorm:"size:32" json:"nickname"`
	Avatar         string                  `gorm:"size:256" json:"avatar"`
	Abstract       string                  `gorm:"size:256" json:"abstract"`
	RegisterSource enum.RegisterSourceType `json:"registerSource"`
	Password       string                  `gorm:"size:64" json:"-"`
	Email          string                  `gorm:"size:256" json:"email"`
	OpenID         string                  `gorm:"size:64" json:"openID"` //Third-party login unique value
	Role           enum.RoleType           `json:"role"`                  //1 Administrator 2 normal user 3 visitor
	UserConfModel  *UserConfModel          `gorm:"foreignKey:UserID" json:"-"`
	Ip             string                  `json:"ip"`
	Addr           string                  `json:"addr"`
}

type UserConfModel struct {
	UserID             uint       `gorm:"PrimaryKey;unique" json:"userID"`
	UserModel          UserModel  `gorm:"foreignKey:UserID" json:"-"`
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"`
	UpdateUsernameDate *time.Time `json:"updateUsernameDate"`
	OpenCollect        bool       `json:"openCollect"` //open user collect
	OpenFollow         bool       `json:"openFollow"`
	OpenFans           bool       `json:"openFans"`
	HomeStyleID        uint       `json:"homeStyleID"`
}

func (u *UserModel) CodeAge() int {
	sub := time.Now().Sub(u.CreatedAt)
	return int(math.Ceil(sub.Hours() / 24 / 365))
}
