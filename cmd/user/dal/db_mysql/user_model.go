package db_mysql

const TableNameUser string = "user"

type User struct {
	UserId        int64
	Username      string
	Password      string
	Avatar        string
	Salt          string
	FollowCount   int64
	FollowerCount int64
}

func (u *User) TableName() string {
	return TableNameUser
}
