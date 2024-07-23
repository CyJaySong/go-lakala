package constants

const (
	OrderStatusUnpaid   = 1001 // 未支付
	OrderStatusPartPaid = 1002 // 部分支付
	OrderStatusOngoing  = 1003 // 进行中
	OrderStatusPaid     = 1004 // 已付款
	OrderStatusDealDone = 1006 // 交易完成
	OrderStatusPartDone = 1007 // 交易完成
	OrderStatusClose    = 2002 // 关闭(失败)
	OrderStatusRefunded = 2003 // 已退款
)
