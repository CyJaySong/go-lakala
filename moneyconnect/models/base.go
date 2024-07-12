package models

type BaseResponse[T any] struct {
	Status    string `json:"status"`     // 服务调用是否成功;"OK"表示成功;"ERROR"表示失败
	ErrorCode string `json:"error_code"` // 错误代码;仅当status=error时有效
	Message   string `json:"message"`    // 错误信息;仅当status=error时有效
	Result    T      `json:"result"`     // JSON Object类型的返回值
}
