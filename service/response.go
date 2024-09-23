package service

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

/*
可以根据自己需要去定义和使用
*/

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeCategoryExist
	CodeUserNotExist
	CodeErrUserPsw
	CodeServerBusy
	CodeAuthNull
	CodeAuthErrFormat
	CodeAuthInvalidToken
	CodeParamTypeErr
	CodeUserNotLogin
	CodeUserErrLogin
)

var CodeMsgText = map[int]string{
	CodeSuccess:          "success",
	CodeInvalidParams:    "请求参数错误 ",
	CodeUserExist:        "用户已经存在",
	CodeUserNotExist:     "用户不存在",
	CodeErrUserPsw:       "密码或者用户名输入有误",
	CodeServerBusy:       "服务器繁忙",
	CodeAuthNull:         "请求头中auth为空，需要登陆",
	CodeAuthErrFormat:    "请求头中auth格式有误，需要登陆",
	CodeAuthInvalidToken: "无效的Token",
	CodeParamTypeErr:     "参数类型错误",
	CodeUserNotLogin:     "用户没有登陆",
	CodeUserErrLogin:     "用户名或密码输入有误",
	CodeCategoryExist:    "分类已经存在",
}

func ResponseErr(ctx *app.RequestContext, code int) {
	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  CodeMsgText[code],
		Data: nil,
	})
}

func ResponseErrWithMsg(ctx *app.RequestContext, code int, msgErr string) {
	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  CodeMsgText[code] + msgErr,
		Data: nil,
	})
}

func ResponseSuccess(ctx *app.RequestContext, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code: CodeSuccess,
		Msg:  CodeMsgText[CodeSuccess],
		Data: data,
	})
}
