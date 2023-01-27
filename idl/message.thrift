namespace go message

struct GetMessageChatRequest {
    1: required string token;
    2: required i64 to_user_id;
}

struct Message {
    1: required i64 id;
    2: required string content;
    3: optional string create_time;
}

struct GetMessageChatResponse {
    1: required i32 status_code;
    2: optional string status_msg; 
    3: required list<Message> msg_list;
}

struct MessageActionRequest {
    1: required string token;
    2: required i64 to_user_id;
    3: required i32 action_type = 1;
    4: required string content;
}

struct MessageActionResponse {
    1: required i32 status_code;
    2: optional string status_msg;
}


service MessageService {
    MessageActionResponse MessageAction(1: MessageActionRequest req)
    GetMessageChatResponse GetMessageChat(1: GetMessageChatRequest req)
}