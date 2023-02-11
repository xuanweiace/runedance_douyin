// Code generated by Kitex v0.4.4. DO NOT EDIT.

package recommendservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	recommend "runedance_douyin/cmd/recommend/kitex_gen/recommend"
)

func serviceInfo() *kitex.ServiceInfo {
	return recommendServiceServiceInfo
}

var recommendServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RecommendService"
	handlerType := (*recommend.RecommendService)(nil)
	methods := map[string]kitex.MethodInfo{
		"getRecommended": kitex.NewMethodInfo(getRecommendedHandler, newRecommendServiceGetRecommendedArgs, newRecommendServiceGetRecommendedResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "recommend",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func getRecommendedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*recommend.RecommendServiceGetRecommendedArgs)
	realResult := result.(*recommend.RecommendServiceGetRecommendedResult)
	success, err := handler.(recommend.RecommendService).GetRecommended(ctx, realArg.User)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRecommendServiceGetRecommendedArgs() interface{} {
	return recommend.NewRecommendServiceGetRecommendedArgs()
}

func newRecommendServiceGetRecommendedResult() interface{} {
	return recommend.NewRecommendServiceGetRecommendedResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetRecommended(ctx context.Context, user int64) (r *recommend.RecommendResponse, err error) {
	var _args recommend.RecommendServiceGetRecommendedArgs
	_args.User = user
	var _result recommend.RecommendServiceGetRecommendedResult
	if err = p.c.Call(ctx, "getRecommended", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}