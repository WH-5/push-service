package service

import "github.com/go-kratos/kratos/v2/errors"

type ErrStruct struct {
	code int
	name string
}

// 推送消息相关错误
var (
	PushMessageSendError    = ErrStruct{code: 460, name: "PUSH_MESSAGE_SEND_ERROR"}
	PushMessageOfflineError = ErrStruct{code: 461, name: "PUSH_MESSAGE_OFFLINE_ERROR"}
	PushMessageFormatError  = ErrStruct{code: 462, name: "PUSH_MESSAGE_FORMAT_ERROR"}
)

func PushMessageError(err error) *errors.Error {
	e := PushMessageSendError
	return errors.New(e.code, e.name, err.Error())
}

func PushOfflineError(err error) *errors.Error {
	e := PushMessageOfflineError
	return errors.New(e.code, e.name, err.Error())
}

func PushFormatError(err error) *errors.Error {
	e := PushMessageFormatError
	return errors.New(e.code, e.name, err.Error())
}
