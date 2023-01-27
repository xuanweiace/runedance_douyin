package main

import (
	"context"
	message "runedance_douyin/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetMessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageChat(ctx context.Context, req *message.GetMessageChatRequest) (resp *message.GetMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}
