package constants

type ErrorCode struct {
	Code string
	Msg  string
}

var (
	OrderAlreadyExist = ErrorCode{"OD032010", "订单已存在"}
)
