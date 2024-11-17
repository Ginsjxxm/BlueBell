package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServeBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServeBusy:       "服务繁忙",
}

func (code ResCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeServeBusy]
	}
	return msg
}
