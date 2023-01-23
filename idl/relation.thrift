namespace go relation

struct BaseResponse {
    1: required i32 status_code;
    2: optional string status_msg;
}
struct RelationActionRequest {
    1: required string token;
    2: required i64 to_user_id;
    3: required i32 action_type;
}
struct RelationActionResponse {
    1: required BaseResponse base_resp
}
