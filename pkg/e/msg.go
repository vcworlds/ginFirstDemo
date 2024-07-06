package e

var MsgFlags = map[int]string{
	Success:      "成功",
	Error:        "失败",
	InvilidError: "参数不正确",
}

func GetMgsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
