package db_mysql

const TableNameUser string = "usermess"

type Usermess struct {
	UserId        int64
	Username      string
	Password      string
	Avatar        string
	Salt          string
	FollowCount   int64
	FollowerCount int64
}

func (*Usermess) TableName() string {
	return TableNameUser
}
