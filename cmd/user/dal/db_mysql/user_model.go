package db_mysql

const TableNameUser string = "user"

type User struct {
	UserId        int64  `gorm:"column:user_id;PRIMARY_KEY"`
	Username      string `gorm:"column:username;UNIQUE"`
	Password      string `gorm:"column:password;NOT NULL"`
	Avatar        string `gorm:"column:avatar"`
	Salt          string `gorm:"column:Salt"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (u *User) TableName() string {
	return TableNameUser
}
