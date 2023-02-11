package main

import (
	"bytes"
	"context"
	"strconv"
)

func up(id int64, file []byte) error {
	fname := "/dir01/" + strconv.FormatInt(id, 10) + ".mp4"
	_, err := cosClient.Object.Put(context.Background(), fname, bytes.NewReader(file), nil)
	return err
}
