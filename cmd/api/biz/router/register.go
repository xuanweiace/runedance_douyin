// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
<<<<<<< HEAD
	relation "runedance_douyin/cmd/api/biz/router/relation"
=======
	douyin "runedance_douyin/cmd/api/biz/router/douyin"
>>>>>>> main
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
<<<<<<< HEAD
	relation.Register(r)
=======
	douyin.Register(r)
>>>>>>> main
}
