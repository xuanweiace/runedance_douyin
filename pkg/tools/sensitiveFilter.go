package tools

import (
	"github.com/importcjj/sensitive"
	"runedance_douyin/pkg/errnos"
)

var Filter *sensitive.Filter

const WordDictPath = "sensitiveDict.txt"

func InitFilter() {
	Filter = sensitive.New()
	err := Filter.LoadWordDict(WordDictPath)
	if err != nil {
		errnos.Wrap(err, "初始化过滤器失败")
	}
}
func FilterSensitive(str string) string {
	return Filter.Replace(str, 42)
}
