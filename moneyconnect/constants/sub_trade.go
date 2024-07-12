package constants

const (
	TradeSubTypeRecharge             = 1001 // 充值
	TradeSubTypeWithdrawal           = 1002 // 提现
	TradeSubTypeTransfer             = 1003 // 转账
	TradeSubTypeRefund               = 1004 // 退款
	TradeSubTypeAutoWithdrawal       = 1005 // 自动提现
	TradeSubTypeFundAdvance          = 1006 // 资金垫付
	TradeSubTypeHangAccount          = 1007 // 挂账
	TradeSubTypeOfflineRecharge      = 1008 // 线下充值
	TradeSubTypeTaxTransfer          = 1009 // 税优转账
	TradeSubTypeTaxPreTransfer       = 1010 // 税优预转账
	TradeSubTypeMerchantTransfer     = 1011 // 商户间转账
	TradeSubTypeRefundPreRecharge    = 1012 // 退款预充值
	TradeSubTypeMemberConsumption    = 2009 // 会员消费
	TradeSubTypeMarketingConsumption = 2010 // 营销消费
	TradeSubTypeMarketingRefund      = 2011 // 营销退款
	TradeSubTypeSmallAmountAuth      = 2020 // 小额鉴权
	TradeSubTypeHandlingFee          = 3001 // 手续费
	TradeSubTypeProfitSharing        = 3002 // 分账
	TradeSubTypeBalanceProfitSharing = 3003 // 余额分账
	TradeSubTypeIncomeExpense        = 3004 // 收支两条线
	TradeSubTypeProfitSharingCancel  = 3005 // 分账撤销
	TradeSubTypeFreeze               = 4001 // 冻结
	TradeSubTypeUnfreeze             = 4002 // 解冻
	TradeSubTypeAdjustment           = 5001 // 调账
	TradeSubTypeAccountSettlement    = 5002 // 在途账户平账
	TradeSubTypeFeeSettlement        = 5003 // 手续费平账
)
