package main

import (
	"context"
	message "runedance_douyin/kitex_gen/message"
	service "runedance_douyin/cmd/message/service"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionRequest) (*message.MessageActionResponse, error) {
	// TODO: Your code here...
	resp := message.NewMessageActionResponse()
	resp.StatusCode = 1
	var msg string

	// parse token
	err2 := service.NewMessageActionService(ctx).MessageAction(ctx, req.UserId, req.ToUserId, req.ActionType, req.Content)
	
	if(err2 != nil){
		msg = err2.Error()
		resp.StatusMsg = msg
		return resp, err2
	}
	msg = "success"
	resp.StatusMsg = msg
	resp.StatusCode = 0
	return resp, nil
}


// GetMessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageChat(ctx context.Context, req *message.GetMessageChatRequest) (*message.GetMessageChatResponse, error) {
	// TODO: Your code here...
	resp := message.NewGetMessageChatResponse()
	resp.StatusCode = 1
	var msg string

	messageList, err2 := service.NewGetMessageChatService(ctx).GetMessageChat(ctx, req.UserId, req.ToUserId)
	if(err2 != nil){
		msg = err2.Error()
		resp.StatusMsg = msg
		return resp, err2
	}
	msg = "success"
	resp.StatusCode = 0
	resp.StatusMsg = msg
	resp.MsgList = messageList
	return resp, nil
}




