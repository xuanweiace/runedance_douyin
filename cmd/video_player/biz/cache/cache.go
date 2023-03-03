package cache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"time"
)

var Cache map[string]*bigcache.BigCache
var BCache1, BCache2 *bigcache.BigCache

func Init() error {
	cache1, err1 := bigcache.New(context.TODO(), bigcache.DefaultConfig(10*time.Minute))
	if err1 != nil {
		return err1
	}
	BCache1 = cache1
	cache2, err2 := bigcache.New(context.TODO(), bigcache.DefaultConfig(10*time.Minute))
	if err2 != nil {
		return err2
	}
	BCache2 = cache2
	Cache = make(map[string]*bigcache.BigCache)
	Cache["video"] = BCache1
	Cache["cover"] = BCache2
	return nil
}
func GetFile(tp string, sid string) (*[]byte, int) {
	b, e := Cache[tp].Get(sid)
	if e != nil {
		return nil, 0
	}
	if b == nil {
		return nil, 1
	}
	return &b, 2
}
func SetFile(tp string, sid string, data []byte) error {
	err := Cache[tp].Set(sid, data)
	return err
}
