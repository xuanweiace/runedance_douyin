package main

import (
	"log"
	"runedance_douyin/pkg/tools"
)

func main() {
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

	log.Println(token)

}
