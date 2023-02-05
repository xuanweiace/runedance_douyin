package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"log"
	"runedance_douyin/kitex_gen/interaction"
	"runedance_douyin/kitex_gen/interaction/messageservice"
	"runedance_douyin/pkg/tools"
)

func main() {
	client, _ := messageservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	/*
		if err != nil {
			log.Fatal(err)
		}

		token, err := tools.GenToken("qw", 1)
		req := &interaction.FavoriteRequest{Token: token, VideoId: 1, ActionType: 1}
		resp, err := client.FavoriteAction(context.Background(), req)
		client.FavoriteAction(context.Background(), &interaction.FavoriteRequest{Token: token, VideoId: 2, ActionType: 1})
		client.FavoriteAction(context.Background(), &interaction.FavoriteRequest{Token: token, VideoId: 3, ActionType: 1})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)*/
	//req:=&interaction.
	token, _ := tools.GenToken("qw", 1)
	resp1, err := client.GetFavoriteList(context.Background(), &interaction.GetFavoriteListRequest{UserId: 1, Token: token})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp1)
	var c *string
	c = new(string)
	*c = "123"
	resp2, err := client.CommentAction(context.Background(), &interaction.CommentRequest{token, 3, 1, c, nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp2)
	resp3, err := client.GetCommentList(context.Background(), &interaction.GetCommentListRequest{1, token})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp3)
}
