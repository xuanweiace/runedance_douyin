package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"runedance_douyin/kitex_gen/interaction"
	"runedance_douyin/kitex_gen/interaction/interactionservice"
	constants "runedance_douyin/pkg/consts"
	"runedance_douyin/pkg/errnos"
	"time"
)

var interactionClient interactionservice.Client

func initInteraction() {

	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := interactionservice.NewClient(constants.InteractionServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(500*time.Microsecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	log.Println("[api服务 initInteraction] Interaction服务启动成功")
	interactionClient = c
}

func FavoriteAction(ctx context.Context, user_id, video_id int64, action_type int32) error {
	req := interaction.FavoriteRequest{
		UserId:     user_id,
		VideoId:    video_id,
		ActionType: action_type,
	}
	resp, err := interactionClient.FavoriteAction(ctx, &req)
	if err != nil {
		return err
	}
	if resp.StatusCode != errnos.CodeSuccess {
		return fmt.Errorf("[rpc.InteractionAction] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

func GetFavoriteList(ctx context.Context, user_id int64) ([]*interaction.Video, error) {
	req := interaction.GetFavoriteListRequest{
		UserId: user_id,
	}
	resp, err := interactionClient.GetFavoriteList(ctx, &req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errnos.CodeSuccess {
		return nil, fmt.Errorf("[rpc.GetFavoriteList] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}
	return resp.VedioList, nil
}

func CommentAction(ctx context.Context, user_id, video_id int64, action_type int32, comment_text string, comment_id string) (*interaction.Comment, error) {
	req := interaction.CommentRequest{
		UserId:      user_id,
		VideoId:     video_id,
		ActionType:  action_type,
		CommentText: &comment_text,
		CommentId:   &comment_id,
	}
	resp, err := interactionClient.CommentAction(ctx, &req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errnos.CodeSuccess {
		return nil, fmt.Errorf("[rpc.CommentAction] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}
	return resp.Comment, nil
}

func GetCommentList(ctx context.Context, video_id int64) ([]*interaction.Comment, error) {
	req := interaction.GetCommentListRequest{
		VideoId: video_id,
	}
	resp, err := interactionClient.GetCommentList(ctx, &req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errnos.CodeSuccess {
		return nil, fmt.Errorf("[rpc.GetCommentList] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}
	return resp.CommentList, nil
}
