package handler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type Downloader struct {
	client   *mongo.Client
	database *mongo.Database
	bucket   *gridfs.Bucket
}

func (d *Downloader) downToStream(id *string, b *bytes.Buffer) error {
	_, err := d.bucket.DownloadToStream(*id, b)
	return err
}

func newDownloader(dbURI string, bucketName string) *Downloader {
	//url := "mongodb://user02:User02@qwq.bogo.ac.cn:23317/fdb02"
	client, err1 := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err1 != nil {
		panic(err1)
	}
	db := client.Database("fdb02")
	opts := options.GridFSBucket().SetName(bucketName)
	bucket, err2 := gridfs.NewBucket(db, opts)
	if err2 != nil {
		panic(err2)
	}
	return &Downloader{
		client:   client,
		database: db,
		bucket:   bucket,
	}
}
func disDownloader(downloader *Downloader) {
	if downloader.client != nil {
		err := downloader.client.Disconnect(context.TODO())
		if err != nil {
			//
		}
	}
	return
}
func getURI() string {
	return "mongodb://user02:User02@qwq.bogo.ac.cn:23317/fdb02"
}

func Play(ctx context.Context, c *app.RequestContext) {
	//a, _ := c.FormFile("123")
	fileType := c.Param("type")
	sid := c.Param("storageID")
	fmt.Println(fileType, sid)
	if sid == "" || fileType == "" ||
		(strings.EqualFold(fileType, "video") == false &&
			strings.EqualFold(fileType, "cover") == false) {
		c.String(consts.StatusBadRequest, "Invalid parameters")
		return
	}
	//sid = strings.ReplaceAll(sid, ".mp4", "")
	if strings.EqualFold(fileType, "video") {
		c.SetContentType("video/mp4")
	}
	downloader := newDownloader(getURI(), "fs_"+fileType)
	defer disDownloader(downloader)
	buff := bytes.NewBuffer(nil)
	c.SetBodyStream(buff, -1)
	var err error
	err = downloader.downToStream(&sid, buff)
	fmt.Println(buff.Len())
	if err != nil {
		fmt.Println(err)
		c.String(consts.StatusGatewayTimeout, "can't connect to database")
		//return
	} else {
		c.Finished()
	}
	return
}
