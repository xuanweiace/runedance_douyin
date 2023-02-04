namespace go user

//用户注册接口
struct douyin_user_register_request {
    1: required string username;
    2: required string password;
}

struct douyin_user_register_response {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required i64 user_id;
    4: required string token;
}


//用户登录接口
struct douyin_user_login_request{
    1: required string username;
    2: required string password;
}

struct douyin_user_login_response {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required i64 user_id;
    4: required string token;
}


//用户信息接口
struct douyin_user_request{
    1: required i64 user_id;
    2: required i64 my_user_id;
}

struct douyin_user_response{
    1: required i32 status_code;
    2: optional string status_msg;
    3: required User user;
}

struct User{
    1:required i64 user_id;
    2:required string username;
    3:optional i64 follow_count;
    4:optional i64 follower_count;
    5:required bool is_follow;
}

service UserService {
     douyin_user_register_response UserRegister (1:douyin_user_register_request req)
     douyin_user_login_response UserLogin (1: douyin_user_login_request req)
     douyin_user_response GetUser (1: douyin_user_request req)
}

