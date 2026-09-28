package youzan_model

import "encoding/json"

type PreErrInner struct { // 有赞请求前置错误响应
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

type PreErr struct {
	GwErrResp PreErrInner `json:"gw_err_resp"`
}

type YouZanResCommon[T any] struct { // 有赞统一的响应
	Success   bool        `json:"success"`
	Code      json.Number `json:"code"` // 成功为200
	Message   string      `json:"message"`
	RequestId string      `json:"request_id"` // 成功为200
	Data      T           `json:"data,omitempty"`
}
