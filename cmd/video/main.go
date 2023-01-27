package main

import (
	"log"
	videostorage "runedance_douyin/cmd/video/kitex_gen/videoStorage/videostorageservice"
)

func main() {
	svr := videostorage.NewServer(new(VideoStorageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
