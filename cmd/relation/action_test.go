package main

import (
	"context"
	"fmt"
	"runedance_douyin/cmd/relation/dal"
	"runedance_douyin/cmd/relation/dal/db_mysql"

	"runedance_douyin/kitex_gen/relation"
	"testing"
)

func TestActionService_ExecuteAction(t *testing.T) {
	/*	test_token
		"username": "zxz",
		"user_id": 1,
	*/
	test_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Inp4eiIsInVzZXJfaWQiOjEsImlzcyI6IndlY2hhbiIsImV4cCI6MTY3NDkyMzc4Nn0.xNycv4CORlasadtYq0eYnczjdvcK0BrF-a9SKH-R3_g"
	r := new(RelationServiceImpl)

	resp, err := r.RelationAction(context.Background(), &relation.RelationActionRequest{
		Token:      test_token,
		ToUserId:   2,
		ActionType: 1,
	})
	fmt.Printf("resp:%+v, err:%+v", resp, err)
}

func TestQueryService_GetFollowList(t *testing.T) {
	dal.Init()
	queryRelation, err := db_mysql.QueryRelation(1, 3)
	fmt.Printf("relation:%+v, err=%v\n", queryRelation, err)
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
