namespace go douyin

struct User{
    1:required  i64     id
    2:required  string  name
    3:optional  i64     follow_count
    4:optional  i64     follower_count
    5:required  bool    is_follow
}
struct Message {
    1: required i64 id
    2: required string content
    3: optional string create_time
}
struct Video {
  1: required i64 id // 视频唯一标识
  2: required User author // 视频作者信息
  3: required string play_url // 视频播放地址
  4: required string cover_url // 视频封面地址
  5: required i64 favorite_count // 视频的点赞总数
  6: required i64 comment_count // 视频的评论总数
  7: required bool is_favorite // true-已点赞，false-未点赞
  8: required string title // 视频标题
}

struct Comment{
    1: required i64 id
    2: required User user
    3: string content
    4: required string create_date
}
///////////////////////////////////////////////////

struct RelationActionRequest {
    1: required string token (api.query="token")
    2: required i64 to_user_id (api.query="to_user_id")
    3: required i32 action_type (api.query="action_type" )
}

struct RelationActionResponse {
    1: required i32 status_code;
    2: optional string status_msg;
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

///////////////////////////////////////////////////


struct GetMessageChatRequest {
    1: required string token (api.query="token")
    2: required i64 to_user_id (api.query="to_user_id ")
}

struct GetMessageChatResponse {
    1: required i32 status_code
    2: optional string status_msg 
    3: required list<Message> msg_list
}

struct MessageActionRequest {
    1: required string token (api.query="token")
    2: required i64 to_user_id (api.query="to_user_id ")
    3: required i32 action_type = 1 (api.query="action_type ")
    4: required string content (api.query="content ")
}

struct MessageActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}
///////////////////////////////////////////////////
//用户注册接口
struct DouyinUserRegisterRequest {
    1: required string username (api.query="username")
    2: required string password (api.query="password")
}

struct DouyinUserRegisterResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

//用户登录接口
struct DouyinUserLoginRequest{
    1: required string username (api.query="username")
    2: required string password (api.query="password")
}

struct DouyinUserLoginResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

//用户信息接口
struct DouyinUserRequest{
    1: required i64 user_id (api.query="user_id")
    2: required string token (api.query="token")
}

struct DouyinUserResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: required User user
}

///////////////////////////////////////////////////

struct DouyinFeedRequest {
  1: optional i64 latest_time (api.query="latest_time")// 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
  2: optional string token (api.query="token") // 可选参数，登录用户设置
}

struct DouyinFeedResponse {
  1: required i32 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: list<Video>  video_list// 视频列表
  4: optional i64 next_time// 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct DouyinPublishActionRequest {
  1: required string token  (api.query="token")// 用户鉴权token
  2: required byte data  (api.form="data")// 视频数据
  3: required string title  (api.query="title")// 视频标题
}

struct DouyinPublishActionResponse {
  1: required i32 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
}
struct DouyinPublishListRequest {
  1: required i64 user_id  (api.query="user_id")// 用户id
  2: required string token (api.query="token")// 用户鉴权token
}

struct DouyinPublishListResponse {
  1: required i32 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: list<Video> video_list // 用户发布的视频列表
}

///////////////////////////////////////////////////

struct FavoriteRequest {
    1: required string token (api.query="token")
    2: required i64 video_id (api.query="video_id")
    3: required i32 action_type (api.query="action_type")
}

struct FavoriteResponse {
    1: required i32 status_code
    2: optional string status_msg
}

struct GetFavoriteListRequest{
    1:required  i64 user_id (api.query="user_id")
    2:required  string token (api.query="token")
}

struct GetFavoriteListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: list<Video> vedio_list
}

struct CommentRequest {
    1: required string token (api.query="token")
    2: required i64 video_id (api.query="video_id")
    3: required i32 action_type (api.query="action_type")
    4: optional string comment_text (api.query="comment_text")
    5: optional string comment_id (api.query="comment_id")
}

struct CommentResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: optional Comment comment
}

struct GetCommentListRequest{
    1:required  i64 video_id (api.query="video_id")
    2:required  string token (api.query="token")
}

struct GetCommentListResponse{
    1: required i32 status_code
    2: optional string status_msg
    3: list<Comment> comment_list
}

service DouyinService {
    RelationActionResponse   RelationAction(1:RelationActionRequest    req) (api.post="/douyin/relation/action/")
    GetFollowListResponse  GetFollowList(1:GetFollowListRequest  req) (api.get="/douyin/relation/follow/list/")
    GetFollowerListResponse  GetFollowerList(1:GetFollowerListRequest  req) (api.get="/douyin/relation/follower/list/")
    GetFriendListResponse    GetFriendList(1:GetFriendListRequest req) (api.get="/douyin/relation/friend/list/")
    
    MessageActionResponse MessageAction(1: MessageActionRequest req) (api.post="/douyin/message/action/")
    GetMessageChatResponse GetMessageChat(1: GetMessageChatRequest req) (api.post="/douyin/message/chat/")
    
    DouyinUserRegisterResponse UserRegister (1:DouyinUserRegisterRequest req) (api.post="/douyin/user/register/")
    DouyinUserLoginResponse UserLogin (1: DouyinUserLoginRequest req) (api.post="/douyin/user/login/")
    DouyinUserResponse GetUser (1: DouyinUserRequest req) (api.get="/douyin/user/")
    
    DouyinFeedResponse Feed(1: DouyinFeedRequest req) (api.get="/douyin/feed/")
    DouyinPublishActionResponse PublishAction(1: DouyinPublishActionRequest req) (api.post="/douyin/publish/action/")
    DouyinPublishListResponse PublishList(1: DouyinPublishListRequest req) (api.get="/douyin/publish/list/")

    FavoriteResponse FavoriteAction(1: FavoriteRequest req) (api.post="/douyin/favorite/action/")
    GetFavoriteListResponse GetFavoriteList(1: GetFavoriteListRequest req)(api.get="/douyin/favorite/list/")
    CommentResponse CommentAction(1: CommentRequest req) (api.post="/douyin/comment/action/")
    GetCommentListResponse GetCommentList(1: GetCommentListRequest req)(api.get="/douyin/comment/list/")
}