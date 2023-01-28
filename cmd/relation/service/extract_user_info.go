package service

import (
	"runedance_douyin/pkg/tools"
)

// jwt_secret_key是一样的，已经写在constants里。
func extract_user_id_from_jwt_token(token string) (int64, error) {
	//todo 调用jwt得到用户id
	claims, err := tools.ParseToken(token)

	if err != nil {
		return 0, err
	}
	return claims.User_id, nil
}
