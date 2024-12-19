package result

import "time"

type ResponseSuccessBean struct {
	Code       uint32      `json:"code"`
	Mark       string      `json:"mark"`
	ServerTime int64       `json:"server_time"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{
		Code:       0,
		Mark:       "mxshop",
		ServerTime: time.Now().Unix(),
		Msg:        "success",
		Data:       data,
	}
}

type ResponseErrorBean struct {
	Code       uint32 `json:"code"`
	Mark       string `json:"mark"`
	ServerTime int64  `json:"server_time"`
	Msg        string `json:"msg"`
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{
		Code:       errCode,
		Mark:       "mxshop",
		ServerTime: time.Now().Unix(),
		Msg:        errMsg,
	}
}
