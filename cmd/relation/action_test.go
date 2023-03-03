package main

import (
	"context"
	"fmt"
	"runedance_douyin/cmd/relation/dal"
	"runedance_douyin/cmd/relation/dal/db_mysql"
	"runedance_douyin/cmd/relation/rpc"
	"runedance_douyin/pkg/tools"

	"runedance_douyin/kitex_gen/relation"
	"testing"
)

var token_1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Inp4eiIsInVzZXJfaWQiOjEsImlzcyI6IndlY2hhbiIsImV4cCI6MTY3NTI0NTk0OH0.FE8HtdX4IekS0R89_hV5K5a-7k8-f9F7TfKebMIHiN0"

func InitEnv() {
	dal.Init()
	rpc.Init()
}

func TestRelationServiceImpl_RelationAction(t *testing.T) {
	InitEnv()
	/*	test_token
		"username": "zxz",
		"user_id": 1,
	*/
	//test_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Inp4eiIsInVzZXJfaWQiOjEsImlzcyI6IndlY2hhbiIsImV4cCI6MTY3NDkyMzc4Nn0.xNycv4CORlasadtYq0eYnczjdvcK0BrF-a9SKH-R3_g"
	r := new(RelationServiceImpl)

	//插入关系(1,2)，预期插入成功且err=nil
	{
		resp, err := r.RelationAction(context.Background(), &relation.RelationActionRequest{
			FromUserId: 1,
			ToUserId:   2,
			ActionType: 1,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}
	//再次插入关系(1,2)，预期插入成功且err=nil
	{
		resp, err := r.RelationAction(context.Background(), &relation.RelationActionRequest{
			FromUserId: 1,
			ToUserId:   2,
			ActionType: 1,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}

	// ActionType 参数错误
	{
		resp, err := r.RelationAction(context.Background(), &relation.RelationActionRequest{
			FromUserId: 1,
			ToUserId:   2,
			ActionType: 0,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	}

	//ToUserId 参数错误 用户不存在
	// todo 细化err
	{
		resp, err := r.RelationAction(context.Background(), &relation.RelationActionRequest{
			FromUserId: 1,
			ToUserId:   -1,
			ActionType: 0,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	}

	//取消关系(1,2)，预期删除成功且err=nil
	{
		resp, err := r.RelationAction(context.Background(), &relation.RelationActionRequest{
			FromUserId: 1,
			ToUserId:   2,
			ActionType: 2,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}

	//再次取消关系(1,2)，预期删除成功且err=nil
	{
		resp, err := r.RelationAction(context.Background(), &relation.RelationActionRequest{
			FromUserId: 1,
			ToUserId:   2,
			ActionType: 2,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}
}

func TestQueryService_GetFollowList(t *testing.T) {
	InitEnv()
	queryRelation, err := db_mysql.QueryRelation(1, 3)
	fmt.Printf("relation:%+v, err=%v\n", queryRelation, err)
}

func TestRelationServiceImpl_GetFollowList(t *testing.T) {
	InitEnv()
	r := new(RelationServiceImpl)

	//插入关系查询userid1的关注列表，预期返回数据且err=nil
	{
		resp, err := r.GetFollowList(context.Background(), &relation.GetFollowListRequest{
			UserId: 1,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}

	//参数错误
	{
		resp, err := r.GetFollowList(context.Background(), &relation.GetFollowListRequest{
			UserId: 2,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	}

}

func TestRelationServiceImpl_GetFollowerList(t *testing.T) {
	InitEnv()
	r := new(RelationServiceImpl)

	//插入关系查询userid1的粉丝列表，预期返回数据且err=nil
	{
		resp, err := r.GetFollowerList(context.Background(), &relation.GetFollowerListRequest{
			UserId: 1,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}

	//插入关系查询userid1的粉丝列表，预期返回数据且err=nil
	{
		resp, err := r.GetFollowerList(context.Background(), &relation.GetFollowerListRequest{
			UserId: 1,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}

}

func TestRelationServiceImpl_GetFriendList(t *testing.T) {
	InitEnv()
	r := new(RelationServiceImpl)

	//插入关系查询userid1的粉丝列表，预期返回数据且err=nil
	{
		resp, err := r.GetFriendList(context.Background(), &relation.GetFriendListRequest{
			UserId: 1,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}

	//插入关系查询userid1的粉丝列表，预期返回数据且err=nil
	{
		resp, err := r.GetFriendList(context.Background(), &relation.GetFriendListRequest{
			UserId: 1,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}

}
func intersection_of_id(arr1, arr2 []int64) (ret []int64) {
	sets := make(map[int64]struct{})
	for _, id := range arr1 {
		sets[id] = struct{}{}
	}
	for _, id := range arr2 {
		if _, ok := sets[id]; ok {
			ret = append(ret, id)
		}
	}
	return ret
}
func TestQ_abc(t *testing.T) {
	a1 := []int64{1, 2, 4}
	a2 := []int64{2, 4, 6}
	fmt.Println(intersection_of_id(a1, a2))

}

func Test_a(t *testing.T) {
	a := 1
	defer func(a int) {
		x := a
		println(x)
	}(a)
	a = 2
}

func Test_jwt(t *testing.T) {
	token, _ := tools.GenToken("zxz", 2)
	println(token)
	//过期时间一秒，则 err=token is expired by 3m8.0721174s
	s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Inp4eiIsInVzZXJfaWQiOjEsImlzcyI6IndlY2hhbiIsImV4cCI6MTY3NDkxMjg2MH0.vpY2d_pt3hDJs2erO42z94DrizoswhotjWpwE0xk81c"
	//s2 := token
	parseToken, err := tools.ParseToken(s)
	fmt.Println(parseToken)

	fmt.Println(err)

}

func Test_ExistRelation(t *testing.T) {
	InitEnv()
	r := new(RelationServiceImpl)

	//查询已存在数据，预期返回True且err=nil
	{
		resp, err := r.ExistRelation(context.Background(), &relation.ExistRelationRequest{
			FromUserId: 1,
			ToUserId:   2,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}

	//查询不存在数据，预期返回False且err=nil
	{
		resp, err := r.ExistRelation(context.Background(), &relation.ExistRelationRequest{
			FromUserId: 1,
			ToUserId:   2000,
		})
		fmt.Printf("resp:%+v, err:%+v\n", resp, err)
	}
}

func TestBatchInsertDelete(t *testing.T) {
	// relation := []*db_mysql.Relation{{5, 6}, {7, 8}, {5, 10}}
	db_mysql.Init()
	// err := db_mysql.BatchInsertIgnore(relation)
	// if err != nil {
	// 	fmt.Println("err1:", err)
	// 	t.Fail()
	// }
	relation2 := []*db_mysql.Relation{{5, 6}, {7, 8}, {5, 10}, {100, 100}}
	err := db_mysql.BatchDelete(relation2)
	if err != nil {
		fmt.Println("err2:", err)
		t.Fail()
	}
}
