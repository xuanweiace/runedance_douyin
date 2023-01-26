package tools

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func RandomStringUtil() string {
	l := 3                           //字符长度
	rand.Seed(time.Now().UnixNano()) //初始化种子
	b := make([]byte, l)             //随机生成len位字符数组
	rand.Read(b)                     //整合
	randStr := hex.EncodeToString(b) //转换为string
	return randStr                   //返回随机数
}
