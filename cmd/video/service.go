package main

import (
	"bytes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
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
