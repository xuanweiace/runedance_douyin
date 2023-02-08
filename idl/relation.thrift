namespace go relation

struct User{
    1:required  i64     id
    2:required  string  name
    3:optional  i64     follow_count
    4:optional  i64     follower_count
    5:required  bool    is_follow
}

struct RelationActionRequest {
    1: required i64 from_user_id 
    2: required i64 to_user_id (api.query="to_user_id")
    3: required i32 action_type (api.query="action_type")
}
struct RelationActionResponse {
    1: required i32 status_code;
    2: optional string status_msg;
}

struct GetFollowListRequest{
    1:required  i64 user_id (api.query="user_id")
}

struct GetFollowListResponse{
    1: required i32 status_code;
    2: optional string status_msg;
    3:list<User> user_list
}

struct GetFollowerListRequest{
    1:required  i64 user_id (api.query="user_id")
}

struct GetFollowerListResponse{
    1: required i32 status_code;
    2: optional string status_msg;
    3:list<User> user_list
}

struct GetFriendListRequest{
    1:required  i64 user_id (api.query="user_id")
}

struct GetFriendListResponse{
    1: required i32 status_code;
    2: optional string status_msg;
    3:list<User> user_list
}

struct ExistRelationRequest {
    1: required i64 from_user_id 
    2: required i64 to_user_id 
}
struct ExistRelationResponse{
    1: required i32 status_code;
    2: optional string status_msg;
    3: bool existed
}

service RelationService{
    RelationActionResponse   RelationAction(1:RelationActionRequest    req) (api.post="/douyin/relation/action/")
    GetFollowListResponse  GetFollowList(1:GetFollowListRequest  req) (api.get="/douyin/relation/follow/list/")
    GetFollowerListResponse  GetFollowerList(1:GetFollowerListRequest  req) (api.get="/douyin/relation/follower/list/")
    GetFriendListResponse    GetFriendList(1:GetFriendListRequest req) (api.get="/douyin/relation/friend/list/")
    ExistRelationResponse   ExistRelation(1:ExistRelationRequest req)
}


