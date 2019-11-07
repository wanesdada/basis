package util

const (
	SuccessCode    = 0
	ErrorLackCode  = 1
	ErrorSqlCode   = 2
	ErrorRidesCode = 3

	//参数签名秘钥
	DesKey = "fahus!dv*(fasfnjk!#s21mg"

	//redis key
	RedisKeyToken = "user:login:token:" //token 缓存

	//time
	FormatTime      = "15:04:05"            //时间格式
	FormatDate      = "2006-01-02"          //日期格式
	FormatDateTime  = "2006-01-02 15:04:05" //完整时间格式
	FormatDateTime2 = "2006-01-02 15:04"    //完整时间格式
)
