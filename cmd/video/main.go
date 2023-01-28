package main

import (
	"crypto/sha256"
	"hash"
	"log"
	videostorage "runedance_douyin/kitex_gen/videoStorage/videostorageservice"
)

var hhh hash.Hash

func main() {
	hhh = sha256.New()
	svr := videostorage.NewServer(new(VideoStorageServiceImpl))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
