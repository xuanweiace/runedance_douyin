namespace go interaction

struct FavoriteRequest {
    1: required string token;
    2: required i64 video_id;
    3: required i32 action_type;
}

struct FavoriteResponse {
    1: required i32 status_code;
    2: optional string status_msg;
}

struct GetFavoriteListRequest{
    1:required  i64 user_id
    2:required  string token
}

struct GetFavoriteListResponse{
    1: required i32 status_code;
    2: optional string status_msg;
    3: list<Video> vedio_list
}

struct Video{
   1:required i64 id;
   2:required User author;
   3:required string play_url;
   4:required string cover_url;
   5:required i64 favorite_count;
   6:required i64 comment_count;
   7:required bool is_favorite;
   8:required string title;
}

struct User{
    1:required i64 user_id;
    2:required string username;
    3:optional i64 follow_count;
    4:optional i64 follower_count;
    5:required bool is_follow;
}

struct CommentRequest {
    1: required string token;
    2: required i64 video_id;
    3: required i32 action_type;
    4: optional string comment_text;
}

struct CommentResponse {
    1: required i32 status_code;
    2: optional string status_msg;
    3: optional Comment comment;
}
struct Comment{
    1: required i64 id;
    2: required User user;
    3: string content;
    4: required string create_date;
}

struct GetCommentListRequest{
    1:required  i64 video_id;
    2:required  string token;
}

struct GetCommentListResponse{
    1: required i32 status_code;
    2: optional string status_msg;
    3: list<Comment> comment_list;
}

service InteractionService {
    FavoriteResponse FavoriteAction(1: FavoriteRequest req)
    GetFavoriteListResponse GetFavoriteList(1: GetFavoriteListRequest req)
    CommentResponse CommentAction(1: CommentRequest req)
    GetCommentListResponse GetCommentList(1: GetCommentListRequest req)
}