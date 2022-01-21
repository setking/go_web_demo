package controller

type CodeType int64

const (
	CodeSuccess CodeType = 1000 + iota
	CodeInvalidParm
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
)

var CodeMsgMap = map[CodeType]string{
	CodeSuccess:         "success",
	CodeInvalidParm:     "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效的Token",
}

func (c CodeType) Msg() string {
	msg, ok := CodeMsgMap[c]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
