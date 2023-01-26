package service

import "strconv"

// jwt_secret_key是一样的，已经写在constants里。
func extract_user_id_from_jwt_token(token string) (int64, error) {
	//todo 调用jwt得到用户id
	user_id, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return nil, err
	}
	return user_id, nil
}
