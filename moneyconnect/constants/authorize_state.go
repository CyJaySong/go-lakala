package constants

const (
	/* 支付宝认证状态 */
	Unauthorized = "UNAUTHORIZED"   // 支付宝-未授权
	Authorized   = "AUTHORIZED"     // 支付宝-已授权
	Closed       = "CLOSED"         // 支付宝-已销户
	SmidNotExist = "SMID_NOT_EXIST" // 支付宝-SMID不存在
	/* 微信认证状态 */
	AuthorizeStateUnauthorized = "AUTHORIZE_STATE_UNAUTHORIZED" // 未授权
	AuthorizeStateAuthorized   = "AUTHORIZE_STATE_AUTHORIZED"   // 已授权
)
