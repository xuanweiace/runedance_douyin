package db_redis

import (
	"context"
	"fmt"
	"runedance_douyin/cmd/relation/dal/db_mysql"
	"strconv"
)

// 默认成功
func Sadd(ctx context.Context, key string, values ...string) {
	Rdb.SAdd(ctx, key, values)
}

func Srem(ctx context.Context, key string, values ...string) {
	Rdb.SRem(ctx, key, values)
}

func ListFolloweridsByUserid(ctx context.Context, userId int64) (ids []int64, err error) {
	res := Rdb.SMembers(ctx, gen_fans_key(userId))
	for _, id := range res.Val() {
		idd, _ := strconv.ParseInt(id, 10, 64)
		ids = append(ids, idd)
	}
	return ids, res.Err()
}
func ListFollowidsByUserid(ctx context.Context, userId int64) (ids []int64, err error) {
	res := Rdb.SMembers(ctx, gen_follow_key(userId))
	for _, id := range res.Val() {
		idd, _ := strconv.ParseInt(id, 10, 64)
		ids = append(ids, idd)
	}
	return ids, res.Err()
}

func CreateRelation(ctx context.Context, fansId, userId int64) error {
	Rdb.SAdd(ctx, gen_follow_key(fansId), userId)
	Rdb.SAdd(ctx, gen_fans_key(userId), fansId)
	return nil // todo
}

func DeleteRelation(ctx context.Context, fansId, userId int64) error {
	cnt1 := Rdb.SRem(ctx, gen_follow_key(fansId), userId)

	cnt2 := Rdb.SRem(ctx, gen_fans_key(userId), fansId)
	if cnt1.Val() == 0 || cnt2.Val() == 0 { // todo 是|| 还是 &&？ 两次rem会不会有并发问题
		return fmt.Errorf("[db_redis.DeleteRelation] %d %d not exist", cnt1.Val(), cnt2.Val())
	}
	return nil
}

func QueryRelation(ctx context.Context, fans_id, user_id int64) (*db_mysql.Relation, error) {
	bc := Rdb.SIsMember(ctx, gen_follow_key(fans_id), user_id)
	//为了兼容mysql的写法 其实不太优雅
	if bc.Val() == true {
		return &db_mysql.Relation{FansID: fans_id, UserID: user_id}, nil
	}
	return nil, nil
}

// userId的关注列表
func gen_follow_key(userId int64) string {
	return strconv.Itoa(int(userId)) + "-follow"
}

// userId的粉丝列表
func gen_fans_key(userId int64) string {
	return strconv.Itoa(int(userId)) + "-fans"
}
