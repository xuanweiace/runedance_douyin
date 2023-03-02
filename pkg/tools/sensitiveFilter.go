package tools

import (
	"github.com/importcjj/sensitive"
	"log"
	"runedance_douyin/pkg/errnos"
)

var Filter *sensitive.Filter

const WordDictPath = "pkg/tools/SensitiveDic.txt"

func InitFilter() {
	log.Println("init filter")
	Filter = sensitive.New()
	err := Filter.LoadWordDict(WordDictPath)
	Filter.AddWord("党国")
	if err != nil {
		errnos.Wrap(err, "初始化过滤器失败")
	}
}
func FilterSensitive(str string) string {
	return Filter.Replace(str, '#')
}
