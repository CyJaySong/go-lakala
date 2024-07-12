package constants

const (
	EventTypeMemberOpen         = "member.open"          // 会员开户结果通知:开户流程中第一次验证成功后
	EventTypeEsignCreate        = "esign.create"         // 电子签约通知:开户流程中第二次验证成功后
	EventTypeMemberMerchantOpen = "member.merchant.open" // 商户开通结果通知:开户流程拉卡拉审核完成
	EventTypeMemberBankCard     = "member.bank.card"     // 银行卡页面相关操作结果通知:
	EventTypeMemberSplitScale   = "member.split.scale"   // 修改分账比例结果通知
	EventTypeMemberMerchantTerm = "member.merchant.term" // 会员商户增终结果通知
)

// 交易相关
const (
	EventTypeOrderDepositRequest       = "order.deposit.request"         // 充值结果通知
	EventTypeOrderWithdrawRequest      = "order.withdraw.request"        // 提现结果通知
	EventTypeOrderConsumePay           = "order.consume.pay"             // 消费支付结果通知
	EventTypeOrderConsumeBatchPay      = "order.consume.batch.pay"       // 批量支付结果通知
	EventTypeOrderConsumeBatchPaySplit = "order.consume.batch.pay.split" // 批量支付分账结果通知
	EventTypePosTradeNotify            = "pos.trade.notify"              // POS支付结果通知
	EventTypeOrderConsumeRefund        = "order.consume.refund"          // 退款结果通知
	EventTypeTransferPayMatchOrderFail = "transfer.pay.match.order.fail" // 转账支付匹配订单失败通知
	EventTypeTransferPayOrderTrade     = "transfer.pay.order.trade"      // 转账支付订单入账通知
	EventTypeOrderChannelFee           = "order.channel.fee"             // 通道手续费通知
	EventTypeOrderWithdrawChargeback   = "order.withdraw.chargeback"     // 提现退单通知
)
