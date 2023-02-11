package mysql

type Favorite struct {
	Id     string
	Uid    int64
	Vid    int64
	Action int32 //0为取消 1为点赞
}
