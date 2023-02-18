package tools

import(
	"strconv"
)
// generation keyname by user_id and to_user_id
func GenerateKeyname(userId int64, toUserId int64) (string){
	var str string
	if(userId < toUserId){
		str = strconv.FormatInt(userId, 10) + "-" + strconv.FormatInt(toUserId, 10)
		return str
	}
	str = strconv.FormatInt(toUserId, 10) + "-" + strconv.FormatInt(userId, 10)
	return str
}