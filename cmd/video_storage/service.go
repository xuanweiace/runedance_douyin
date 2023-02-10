package main

import (
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"time"
)

func UploadFileToDB(db *mongo.Database, bucketName string, dataFile *[]byte, sid *string, fileName *string) (err error) {
	opts := options.GridFSBucket().SetName(bucketName)
	bucket, err3 := gridfs.NewBucket(db, opts)
	if err3 != nil {
		return err3
	}
	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{"upload-time", time.Now()}})
	err3 = bucket.UploadFromStreamWithID(sid, *fileName, bytes.NewReader(*dataFile), uploadOpts)
	return err3
}
func DownloadFromCos(id string) (*[]byte, *[]byte) {

	vu := "_transcode_100030.mp4"
	cu := "_snapshotByOffset_10_0.jpg"
	pre := "dir02/" + id
	resp1, err1 := cosClient.Object.Get(context.Background(), pre+vu, nil)
	bs1, err2 := ioutil.ReadAll(resp1.Body)
	if len(bs1) == 0 {
		fmt.Println(err1)
		fmt.Println(err2)
		// log
		return nil, nil
	}
	resp2, err3 := cosClient.Object.Get(context.TODO(), pre+cu, nil)
	bs2, err4 := ioutil.ReadAll(resp2.Body)
	if len(bs2) == 0 {
		//TODO: 统一日志持久化格式与方式
		fmt.Println(err3)
		fmt.Println(err4)
		return nil, nil
	}
	return &bs1, &bs2 //video&cover;
}
