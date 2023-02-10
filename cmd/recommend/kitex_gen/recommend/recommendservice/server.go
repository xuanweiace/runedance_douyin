// Code generated by Kitex v0.4.4. DO NOT EDIT.
package recommendservice

import (
	server "github.com/cloudwego/kitex/server"
	recommend "runedance_douyin/cmd/recommend/kitex_gen/recommend"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler recommend.RecommendService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
