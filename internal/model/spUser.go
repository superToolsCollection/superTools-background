package model

/**
* @Author: super
* @Date: 2021-01-22 20:41
* @Description: 会员表
**/

type SpUser struct {
	UserEmail     string `gorm:"column:user_email" json:"user_email"`
	UserEmailCode string `gorm:"column:user_email_code" json:"user_email_code"`
	Username      string `gorm:"column:username" json:"username"`
	Password      string `gorm:"column:password" json:"password"`
	UserIntroduce string `gorm:"column:user_introduce" json:"user_introduce"`
	UpdateTime    int    `gorm:"column:update_time" json:"update_time"`
	UserTel       string `gorm:"column:user_tel" json:"user_tel"`
	UserHobby     string `gorm:"column:user_hobby" json:"user_hobby"`
	IsActive      string `gorm:"column:is_active" json:"is_active"`
	UserID        int    `gorm:"column:user_id;primary_key" json:"user_id"`
	QqOpenID      string `gorm:"column:qq_open_id" json:"qq_open_id"`
	UserXueli     string `gorm:"column:user_xueli" json:"user_xueli"`
	CreateTime    int    `gorm:"column:create_time" json:"create_time"`
	UserSex       string `gorm:"column:user_sex" json:"user_sex"`
	UserQq        string `gorm:"column:user_qq" json:"user_qq"`
}

// TableName sets the insert table name for this struct type
func (s *SpUser) TableName() string {
	return "sp_user"
}
