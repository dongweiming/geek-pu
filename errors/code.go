package errors

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
)

var MsgFlags = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "发生错误",
	INVALID_PARAMS: "请求参数错误",
}

func GetMsg(code int) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}

	return MsgFlags[ERROR]
}
