// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoprocessservice

import (
	server "github.com/cloudwego/kitex/server"
	"runedance_douyin/kitex_gen/videoProcess"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler videoprocess.VideoProcessService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}