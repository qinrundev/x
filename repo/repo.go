package repo

import (
	"fmt"
	"github.com/qinrundev/x/consts"
)

const (
	OK = "001"

	ErrParametersValidation = "101"
	ErrInternalServer       = "102"
	ErrUnauthorized         = "103"

	ErrInvalidClient      = "201"
	ErrUserNotExists      = "202"
	ErrInvalidPassword    = "203"
	ErrInvalidCaptcha     = "204"
	ErrDuplicatedUsername = "205"
	ErrInvalidDomainType  = "206"
	ErrInvalidDomain      = "207"

	ErrRecordNotFound  = "301"
	ErrStockIsOpened   = "302"
	ErrRepeatOperation = "303"
	ErrCanNotOperation = "304"

	ErrUnknown = "999"
)

var msg = map[string]string{
	OK:                      "成功",
	ErrUnauthorized:         "未授权",
	ErrParametersValidation: "请求参数错误",
	ErrInternalServer:       "内部错误",
	ErrInvalidClient:        "无效的客户端",
	ErrUserNotExists:        "用户不存在",
	ErrInvalidPassword:      "密码错误",
	ErrInvalidCaptcha:       "验证码错误",
	ErrDuplicatedUsername:   "用户名已存在",
	ErrInvalidDomainType:    "无效的域类型",
	ErrInvalidDomain:        "无效的域",
	ErrRecordNotFound:       "操作的数据不存在",
	ErrStockIsOpened:        "股票已经开盘",
	ErrRepeatOperation:      "请不要重复操作",
	ErrCanNotOperation:      "当前状态不可操作数据",
	ErrUnknown:              "未知错误",
}

func GetMsg(code string) string {
	message, ok := msg[code]
	if ok {
		return message
	}
	return msg[ErrUnknown]
}

func GetCode(code string, serviceName string) string {
	switch serviceName {
	case consts.IAMServiceName:
		return fmt.Sprintf("%s%s", consts.IAMServiceName, code)
	case consts.PredictionServiceName:
		return fmt.Sprintf("%s%s", consts.PredictionServiceName, code)
	default:
		return ErrUnknown
	}
}
