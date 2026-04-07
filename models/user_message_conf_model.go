package models

type UserMessageConfModel struct {
	UserID             uint      `gorm:"primaryKey;unique" json:"userID"`
	UserModel          UserModel `gorm:"foreignKey:UserID" json:"-"`
	OpenCommentMessage bool      `json:"openCommentMessage"`
	OpenDiggMessage    bool      `json:"openDiggMessage"`
	OpenPrivateChat    bool      `json:"openPrivateChat"`
}
