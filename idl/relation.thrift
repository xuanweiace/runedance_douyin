namespace go relation

struct RelationActionRequest {
    1: required string token (api.query="token")
    2: required i64 to_user_id (api.query="to_user_id")
    3: required i32 action_type (api.query="action_type")
}
struct RelationActionResponse {
    1: required i32 status_code;
    2: optional string status_msg;
}


struct User{
    1:required  i64     id
    2:required  string  name
    3:optional  i64     follow_count
    4:optional  i64     follower_count
    5:required  bool    is_follow
}

struct GetFollowListRequest{
    1:required  i64 user_id (api.query="user_id")
    2:required  string token (api.query="token")
}

struct GetFollowListResponse{
    1: required i32 status_code;
    2: optional string status_msg;
    3:list<User> user_list
}

struct GetFollowerListRequest{
    1:required  i64 user_id (api.query="user_id")
    2:required  string token (api.query="token")
}

struct GetFollowerListResponse{
    1: required i32 status_code;
    2: optional string status_msg;
    3:list<User> user_list
}

struct GetFriendListRequest{
    1:required  i64 user_id (api.query="user_id")
    2:required  string token (api.query="token")
}

struct GetFriendListResponse{
    1: required i32 status_code;
    2: optional string status_msg;
    3:list<User> user_list
}

service RelationService{
    RelationActionResponse   RelationAction(1:RelationActionRequest    req) (api.post="/douyin/relation/action/")
    GetFollowListResponse  GetFollowList(1:GetFollowListRequest  req) (api.get="/douyin/relation/follow/list/")
    GetFollowerListResponse  GetFollowerList(1:GetFollowerListRequest  req) (api.get="/douyin/relation/follower/list/")
    GetFriendListResponse    GetFriendList(1:GetFriendListRequest req) (api.get="/douyin/relation/friend/list/")
}


