package models

// MemberOpenEvent 会员开户结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1d50jddv5culm
type MemberOpenEvent struct {
	MemberNo    string `json:"member_no"`              // 会员编号
	OutMemberNo string `json:"out_member_no"`          // 外部会员号
	Mobile      string `json:"mobile"`                 // 手机号
	MemberState int    `json:"member_state,omitempty"` // 会员状态,http://47.110.246.50:6524/docs/qzt/qzt-1d50j2f4efqcs
	MemberType  int    `json:"member_type,omitempty"`  // 会员类型,http://47.110.246.50:6524/docs/qzt/qzt-1d50iu3jl96p8
	CardId      string `json:"card_no,omitempty"`      // 银行卡编号
	ErrMsg      string `json:"err_msg,omitempty"`      // 错误信息
}

// EsignCreateEvent 电子签约通知 http://47.110.246.50:6524/docs/qzt/qzt-1d50ji2qik3ui
type EsignCreateEvent struct {
	FileNo             string `json:"file_no"`                         // 合同文件编号
	EsignProjectCode   string `json:"esign_project_code"`              // 电子合同签约授权号
	ContractNo         string `json:"contract_no"`                     // 合同编号
	EsignStatus        int    `json:"esign_status"`                    // 合同状态,0:签署中,1:生效,2:失效
	PbMemberSignStatus int    `json:"pb_member_sign_status"`           // 会员签署状态(乙方),0:未签署,1:已签署
	PaMemberSignStatus int    `json:"pa_member_sign_status,omitempty"` // 会员签署状态(甲方),0:未签署,1:已签署,可选
	LklSignStatus      int    `json:"lkl_sign_status,omitempty"`       // 拉卡拉签署状态,0:未签署,1:已签署,可选,合同类型中有此签署方时返回
	MemberList         []struct {
		MemberNo      string `json:"member_no"`       // 会员编号
		EsignUserType int    `json:"esign_user_type"` // 签约用户类型
	} `json:"member_list"` // 签约会员列表
	FailMsg string `json:"fail_msg,omitempty"` // 失败原因,可选
}

// MemberMerchantOpenEvent 商户开通结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1d50jg4rmf7tf
type MemberMerchantOpenEvent struct {
	MemberNo          string  `json:"member_no"`                     // 会员编号
	OutMemberNo       string  `json:"out_member_no"`                 // 外部会员号
	MerchantNo        string  `json:"merchant_no"`                   // 钱账通商户号
	OutMerchantNo     string  `json:"out_merchant_no"`               // 外部商户号
	ChannelMerchantNo string  `json:"channel_merchant_no,omitempty"` // 渠道商户号
	ChannelTerminalNo string  `json:"channel_terminal_no,omitempty"` // 渠道终端号
	HandlerTime       string  `json:"handler_time,omitempty"`        // 处理时间,yyyy-MM-dd HH:mm:ss
	SeqNo             string  `json:"seq_no"`                        // 申请流水
	SellerScale       float64 `json:"seller_scale"`                  // 商家分账比例
	Condition         string  `json:"condition"`                     // 条件
	MerchantStatus    int     `json:"merchant_status,omitempty"`     // 商户状态,http://47.110.246.50:6524/docs/qzt/qzt-1d50j4fhhrekk
	SplitScaleStatus  int     `json:"split_scale_status"`            // 分账比例审核状态,http://47.110.246.50:6524/docs/qzt/qzt-1d50jq19fkuop
	ErrMsg            string  `json:"err_msg,omitempty"`             // 错误信息
	ChannelType       string  `json:"channel_type"`                  // 通道类型
}

// MemberBankCardEvent 银行卡页面相关操作结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1d50jf8urle5h
type MemberBankCardEvent struct {
	MemberNo        string `json:"member_no"`                   // 会员编号
	OperateType     int    `json:"operate_type"`                // 操作类型,1:绑卡,2:解绑,3:设置结算卡
	CardId          string `json:"card_no"`                     // 银行卡编号
	IdentityNo      string `json:"identity_no,omitempty"`       // 证件号码,RSA加密
	BankCardNo      string `json:"bank_card_no,omitempty"`      // 银行卡号,掩码
	Mobile          string `json:"mobile,omitempty"`            // 银行预留手机号
	ProvinceCode    string `json:"province_code,omitempty"`     // 省编号
	CityCode        string `json:"city_code,omitempty"`         // 市编号
	AreaCode        string `json:"area_code,omitempty"`         // 区/县编码
	BankContactName string `json:"bank_contact_name,omitempty"` // 联行号名称
	BankContactNo   string `json:"bank_contact_no,omitempty"`   // 联行号
	IsSettleCard    bool   `json:"is_settle_card"`              // 是否结算卡
}

// MemberSplitScaleEvent 修改分账比例结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1d50jh0pqolhr
type MemberSplitScaleEvent struct {
	MemberNo    string  `json:"member_no"`           // 会员编号
	SellerScale float64 `json:"seller_scale"`        // 商家分账比例
	Condition   string  `json:"condition"`           // 条件
	Remark      string  `json:"remark,omitempty"`    // 备注(可选)
	Status      int     `json:"status"`              // 审核状态,http://47.110.246.50:6524/docs/qzt/qzt-1d50j93ri5atp
	SeqNo       string  `json:"seq_no"`              // 申请流水
	ErrorMsg    string  `json:"error_msg,omitempty"` // 错误信息(可选)
}

// MemberMerchantTermEvent 会员商户增终结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1fbeg2djvup8e
type MemberMerchantTermEvent struct {
	MemberNo    string `json:"member_no"`    // 会员编号
	ChannelType string `json:"channel_type"` // 通道类型
	PosType     string `json:"pos_type"`     // 进件类型
	HandlerTime string `json:"handler_time"` // 处理时间
	Terms       []struct {
		MerchantNo        string `json:"merchant_no"`         // 商户编号
		MerchantStatus    int    `json:"merchant_status"`     // 商户状态,1:审核中,2:进件中,3:有效,4:无效
		ChannelMerchantNo string `json:"channel_merchant_no"` // 渠道商户号
		ChannelTerminalNo string `json:"channel_terminal_no"` // 渠道终端号
		BuisTypeCode      string `json:"buis_type_code"`      // 终端业务码,BANK_CARD:银行卡,QR_CODE_CARD:扫码
		BuisTypeName      string `json:"buis_type_name"`      // 终端业务名称,银行卡\扫码
		BusiActiveNo      string `json:"busi_active_no"`      // 终端业务激活码
	} `json:"terms"` // 终端列表
	ErrMsg string `json:"err_msg,omitempty"` // 错误信息(可选)
}

// OrderDepositRequestEvent 充值结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1ekbd2vjjgaai
type OrderDepositRequestEvent struct {
	MemberNo          string                 `json:"member_no"`                      // 会员标识
	OutOrderNo        string                 `json:"out_order_no"`                   // 应用平台订单号
	OrderNo           string                 `json:"order_no"`                       // 存管订单号
	AccountType       string                 `json:"account_type"`                   // 账户类型编码,固定值:1000
	Amount            int                    `json:"amount"`                         // 充值金额,单位:分
	PayMethod         ConsumeParamsPayMethod `json:"pay_method"`                     // 支付方式,TRANSFER_PAY(转账支付)
	ChannelTradeNo    string                 `json:"channel_trade_no,omitempty"`     // 渠道交易流水号(收单)
	ChannelSeqNo      string                 `json:"channel_seq_no,omitempty"`       // 渠道支付流水号(收单)
	PayChannelTradeNo string                 `json:"pay_channel_trade_no,omitempty"` // 支付通道交易流水号(支付宝、微信)
	Refernumber       string                 `json:"refernumber,omitempty"`          // 系统参考号,POS支付时返回
	ThirdPartyPayment string                 `json:"third_party_payment,omitempty"`  // 三方支付渠道
	Remark            string                 `json:"remark,omitempty,omitempty"`     // 备注
	Exts              string                 `json:"exts,omitempty"`                 // 扩展信息
	OrderStatus       int                    `json:"order_status"`                   // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	EndTime           string                 `json:"end_time"`                       // 交易完成时间,yyyy-MM-dd	HH:mm:ss
	Status            string                 `json:"status"`                         // 状态,"OK"表示成功,"ERROR"表示失败
	ErrorCode         string                 `json:"error_code,omitempty"`           // 错误代码
	ErrorMsg          string                 `json:"error_msg,omitempty"`            // 错误信息
}

// OrderWithdrawRequestEvent 提现结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1ekbd3e1stea7
// 仅调用提现接口发起提现才会有提现结果通知，自动提现不会有通知
type OrderWithdrawRequestEvent struct {
	MemberNo       string `json:"member_no"`                  // 会员标识
	OutOrderNo     string `json:"out_order_no"`               // 应用平台订单号
	OrderNo        string `json:"order_no"`                   // 平台订单号
	AccountType    string `json:"account_type"`               // 账户类型编码,固定值:1000
	Amount         int    `json:"amount"`                     // 提现金额(含平台手续费),单位:分
	BillingFee     int    `json:"billing_fee"`                // 平台手续费,单位:分
	ChannelFee     int    `json:"channel_fee,omitempty"`      // 渠道手续费,单位:分
	BankCardID     string `json:"bank_card_id"`               // 绑定银行卡的编号
	ChannelTradeNo string `json:"channel_trade_no,omitempty"` // 渠道交易流水号
	WithdrawType   string `json:"withdraw_type"`              // 提现方式,T0:T+0到账,T1:T+1到账
	Remark         string `json:"remark,omitempty"`           // 备注
	Exts           string `json:"exts,omitempty"`             // 扩展信息,最多32个字符
	OrderStatus    int    `json:"order_status"`               // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	EndTime        string `json:"end_time"`                   // 交易完成时间,yyyy-MM-dd HH:mm:ss
	Status         string `json:"status"`                     // 状态,OK:成功,PROCESSING:进行中,ERROR:失败
	ErrorCode      string `json:"error_code,omitempty"`       // 错误代码
	ErrorMsg       string `json:"error_msg,omitempty"`        // 错误信息
}

// OrderConsumePayEvent 消费支付结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1ekbd486q44av
type OrderConsumePayEvent struct {
	OrderNo           string                  `json:"order_no"`                       // 订单号
	OutOrderNo        string                  `json:"out_order_no"`                   // 应用平台订单号
	BuyerMemberNo     string                  `json:"buyer_member_no,omitempty"`      // 买家会员号
	SellerMemberNo    string                  `json:"seller_member_no"`               // 卖家会员号
	Amount            int                     `json:"amount"`                         // 订单金额,单位:分
	PayAmount         int                     `json:"pay_amount"`                     // 支付金额,单位:分
	PayMethod         *ConsumeParamsPayMethod `json:"pay_method,omitempty"`           // 支付方式
	OrderName         string                  `json:"order_name"`                     // 订单名称
	OrderDesc         string                  `json:"order_desc"`                     // 订单描述
	OrderStatus       int                     `json:"order_status"`                   // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	PayStatus         int                     `json:"pay_status"`                     // 支付状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbe0j5a5ard
	PaySeqNo          string                  `json:"pay_seq_no"`                     // 支付流水号(钱账通)
	ChannelTradeNo    string                  `json:"channel_trade_no,omitempty"`     // 渠道交易流水号(收单)
	ChannelSeqNo      string                  `json:"channel_seq_no,omitempty"`       // 渠道支付流水号(收单)
	PayChannelTradeNo string                  `json:"pay_channel_trade_no,omitempty"` // 支付通道交易流水号(支付宝、微信)
	ReferNumber       string                  `json:"refernumber,omitempty"`          // 系统参考号,POS支付时返回
	ThirdPartyPayment string                  `json:"third_party_payment,omitempty"`  // 三方支付渠道,交易完成返回,微信:WECHAT,支付宝:ALIPAY,借记卡:DEBIT_CARD,贷记卡:CREDIT_CARD,借记卡和贷记卡只在POS支付时返回
	BankType          string                  `json:"bank_type,omitempty"`            // 付款银行,00:借记,01:贷记,03:支付宝花呗,04:支付宝其他,05:数字货币,06:拉卡拉支付账户,99:未知sha
	CardType          string                  `json:"card_type,omitempty"`            // 银行卡类型
	SplitRuleResult   []SplitRuleResultItem   `json:"split_rule_result,omitempty"`    // 分账规则结果
	SplitSeqNo        string                  `json:"split_seq_no,omitempty"`         // 分账流水,订单全部支付完成并即时分账时返回
	PayTime           string                  `json:"pay_time"`                       // 支付时间,yyyy-MM-dd HH:mm:ss
	PayRequestNo      string                  `json:"pay_request_no,omitempty"`       // 我方渠道订单号
	Status            string                  `json:"status"`                         // 状态,OK:成功,ERROR:失败
	Exts              string                  `json:"exts,omitempty"`                 // 扩展信息,最多32个字符
	ErrorCode         string                  `json:"error_code,omitempty"`           // 错误代码
	ErrorMsg          string                  `json:"error_msg,omitempty"`            // 错误信息
}

// OrderConsumeBatchPayEvent 批量支付结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1ekbd4reg3el5
type OrderConsumeBatchPayEvent struct {
	BatchOrderNo      string                  `json:"batch_order_no"`                 // 订单号
	OutOrderNo        string                  `json:"out_order_no"`                   // 应用平台订单号
	Amount            int                     `json:"amount"`                         // 订单金额
	PayAmount         int                     `json:"pay_amount"`                     // 支付金额
	PayMethod         *ConsumeParamsPayMethod `json:"pay_method,omitempty"`           // 支付方式
	OrderStatus       int                     `json:"order_status"`                   // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	PayStatus         int                     `json:"pay_status"`                     // 支付状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbe0j5a5ard
	PaySeqNo          string                  `json:"pay_seq_no"`                     // 支付流水号(钱账通)
	ChannelTradeNo    string                  `json:"channel_trade_no,omitempty"`     // 渠道交易流水号(收单)
	ChannelSeqNo      string                  `json:"channel_seq_no,omitempty"`       // 渠道支付流水号(收单)
	PayChannelTradeNo string                  `json:"pay_channel_trade_no,omitempty"` // 支付通道交易流水号(支付宝、微信),
	ThirdPartyPayment string                  `json:"third_party_payment,omitempty"`  // 三方支付渠道,交易完成返回,微信:WECHAT,支付宝:ALIPAY,借记卡:DEBIT_CARD,贷记卡:CREDIT_CARD,借记卡和贷记卡只在POS支付时返回
	Refernumber       string                  `json:"refernumber,omitempty"`          // 系统参考号,POS支付时返回
	SubOrderList      []SubOrder              `json:"sub_order_list"`                 // 子订单列表
	Status            string                  `json:"status"`                         // 状态,OK:成功,ERROR:失败
	PayTime           string                  `json:"pay_time"`                       // 支付时间,yyyy-MM-dd HH:mm:ss
	Exts              string                  `json:"exts,omitempty"`                 // 扩展信息,最多32个字符
	ErrorCode         string                  `json:"error_code,omitempty"`           // 错误代码
	ErrorMsg          string                  `json:"error_msg,omitempty"`            // 错误信息
}

// SubOrder 结构体表示子订单列表中的每个子订单信息
type SubOrder struct {
	OrderNo        string `json:"order_no"`         // 订单号
	BuyerMemberNo  string `json:"buyer_member_no"`  // 买家会员号
	SellerMemberNo string `json:"seller_member_no"` // 卖家会员号
	Amount         int    `json:"amount"`           // 订单金额
}

// OrderConsumeBatchPaySplitEvent 批量支付分账结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1ekbd5rodomi6
type OrderConsumeBatchPaySplitEvent struct {
	BatchOrderNo      string              `json:"batch_order_no"`       // 订单号
	OutOrderNo        string              `json:"out_order_no"`         // 应用平台订单号
	SubOrderSplitList []SubOrderSplitItem `json:"sub_order_split_list"` // 子订单分账列表
	Status            string              `json:"status"`               // 状态,OK:成功,ERROR:失败
	ErrorCode         string              `json:"error_code,omitempty"` // 错误代码
	ErrorMsg          string              `json:"error_msg,omitempty"`  // 错误信息
}

type SubOrderSplitItem struct {
	OrderNo         string                `json:"order_no"`                    // 订单号
	OutOrderNo      string                `json:"out_order_no"`                // 应用平台订单号
	Amount          int                   `json:"amount"`                      // 订单金额
	OrderStatus     int                   `json:"order_status"`                // 订单状态
	SplitRuleResult []SplitRuleResultItem `json:"split_rule_result,omitempty"` // 分账规则结果
	SplitSeqNo      string                `json:"split_seq_no,omitempty"`      // 分账流水
}

// PosTradeNotifyEvent POS支付结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1ekbd6ieijaul
type PosTradeNotifyEvent struct {
	OrderNo           string `json:"order_no"`                      // 钱账通订单号
	OutOrderNo        string `json:"out_order_no"`                  // 机构订单号(没有则和order_no一致)
	PaySeqNo          string `json:"pay_seq_no,omitempty"`          // 支付流水号,pay_amount>0时返回
	TradeNo           string `json:"trade_no,omitempty"`            // POS交易请求流水号
	TradeType         int    `json:"trade_type"`                    // 交易类型:1-支付,2-退款,3-冲正
	Amount            int    `json:"amount"`                        // 交易金额,单位:分
	MerchantNo        string `json:"merchant_no"`                   // 商户号
	TerminalNo        string `json:"terminal_no"`                   // 终端号
	TransactionID     string `json:"transaction_id"`                // POS交易通道返回流水号
	ThirdPartyPayment string `json:"third_party_payment,omitempty"` // 三方支付渠道,微信:WECHAT,支付宝:ALIPAY,借记卡:DEBIT_CARD,贷记卡:CREDIT_CARD,银联二维码借记:UNION_DEBIT_CARD,银联二维码贷记:UNION_CREDIT_CARD
	Batchbillno       string `json:"batchbillno"`                   // POS批次号
	Systraceno        string `json:"systraceno"`                    // POS凭证号
	Refernumber       string `json:"refernumber"`                   // POS系统参考号
	PayTime           string `json:"pay_time"`                      // 交易时间,格式为yyyy-MM-dd HH:mm:ss
	IsOutCard         bool   `json:"is_out_card,omitempty"`         // 是否外卡,true:是,false/空:否
	Status            string `json:"status"`                        // 状态:OK:成功,ERROR:失败
	ErrorCode         string `json:"error_code,omitempty"`          // 错误代码
	ErrorMsg          string `json:"error_msg,omitempty"`           // 错误信息
}

// OrderConsumeRefundEvent 退款结果通知 http://47.110.246.50:6524/docs/qzt/qzt-1ekbd7bgi1dab
type OrderConsumeRefundEvent struct {
	OrderNo                string                `json:"order_no"`                             // 订单编号
	OutRequestNo           string                `json:"out_request_no"`                       // 外部退款请求号
	RefundSeqNo            string                `json:"refund_seq_no"`                        // 退款流水号
	OrderStatus            int                   `json:"order_status"`                         // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	RefundStatus           int                   `json:"refund_status"`                        // 退款状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbe0vhq9dkh
	SplitRuleResult        []SplitRuleResultItem `json:"split_rule_result,omitempty"`          // 分账会员列表
	QztChannelPayRequestNo string                `json:"qzt_channel_pay_request_no,omitempty"` // 钱账通请求通道的流水号
	ChannelTradeNo         string                `json:"channel_trade_no,omitempty"`           // 渠道交易流水号(收单)
	ChannelSeqNo           string                `json:"channel_seq_no,omitempty"`             // 渠道支付流水号(收单)
	PayChannelTradeNo      string                `json:"pay_channel_trade_no,omitempty"`       // 支付通道交易流水号(支付宝、微信)
	RefundTime             string                `json:"refund_time,omitempty"`                // 退款时间,格式为 yyyy-MM-dd HH:mm:ss
	ErrorCode              string                `json:"error_code,omitempty"`                 // 错误代码
	ErrorMsg               string                `json:"error_msg,omitempty"`                  // 错误信息,代付失败的退款会返回"退款代付失败"；代付成功，退款失败时返回通道返回信息
}

// TransferPayMatchOrderFailEvent 转账支付匹配订单失败通知 http://47.110.246.50:6524/docs/qzt/qzt-1ekbd89s1i0ch
type TransferPayMatchOrderFailEvent struct {
	PayerAccountNo   string `json:"payer_account_no"`       // 付款方账号,掩码
	PayerAccountName string `json:"payer_account_name"`     // 付款方账户名称
	PayeeAccountNo   string `json:"payee_account_no"`       // 收款方账号,掩码
	PayeeAccountName string `json:"payee_account_name"`     // 收款方账户名称
	TradeTime        string `json:"trade_time"`             // 交易时间,格式为 yyyy-MM-dd HH:mm:ss
	Amount           int    `json:"amount"`                 // 打款金额,单位:分
	TradeRemark      string `json:"trade_remark,omitempty"` // 交易备注,即打款方打款时的备注
	Message          string `json:"message,omitempty"`      // 失败原因
}

// TransferPayOrderTradeEvent 转账支付订单入账通知 http://47.110.246.50:6524/docs/qzt/qzt-1ekbd95g51vs4
type TransferPayOrderTradeEvent struct {
	PayerAccountNo   string `json:"payer_account_no"`       // 付款方账号,掩码处理
	PayerAccountName string `json:"payer_account_name"`     // 付款方账户名称
	PayeeAccountNo   string `json:"payee_account_no"`       // 收款方账号,掩码处理
	PayeeAccountName string `json:"payee_account_name"`     // 收款方账户名称
	TradeTime        string `json:"trade_time"`             // 交易时间,格式为:yyyy-MM-dd HH:mm:ss
	Amount           int    `json:"amount"`                 // 打款金额,单位:分
	TradeRemark      string `json:"trade_remark,omitempty"` // 交易备注,即打款方打款时的备注
	OrderNo          string `json:"order_no"`               // 匹配的钱账通订单号
	PaySeqNo         string `json:"pay_seq_no"`             // 匹配的钱账通支付流水号
	SellerMemberNo   string `json:"seller_member_no"`       // 卖家会员编号
	SellerMemberName string `json:"seller_member_name"`     // 卖家会员名称
}

// OrderChannelFeeEvent 通道手续费通知事件 http://47.110.246.50:6524/docs/qzt/qzt-1fansp4c4acoh
type OrderChannelFeeEvent struct {
	OrderNo       string `json:"order_no"`                 // 订单编号
	OutOrderNo    string `json:"out_order_no,omitempty"`   // 外部订单编号
	RequestNo     string `json:"request_no,omitempty"`     // 请求编号
	OutRequestNo  string `json:"out_request_no,omitempty"` // 外部请求编号
	RequestType   int    `json:"request_type"`             // 请求类型,2:消费,4:退款
	OrderStatus   int    `json:"order_status"`             // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	TradeAmount   int    `json:"trade_amount"`             // 交易金额,单位:分
	RequestStatus int    `json:"request_status,omitempty"` // 请求状态,1006:交易完成,2002:关闭
	ChannelFee    int    `json:"channel_fee"`              // 通道手续费
}

// OrderWithdrawChargebackEvent 提现退单通知事件 http://47.110.246.50:6524/docs/qzt/qzt-1faoobfh3sd8f
type OrderWithdrawChargebackEvent struct {
	OrderNo        string `json:"order_no"`        // 订单编号
	OutOrderNo     string `json:"out_order_no"`    // 外部订单编号
	MemberNo       string `json:"member_no"`       // 会员编号
	WithdrawType   string `json:"withdraw_type"`   // 提现类型,如D1,T1
	OrderStatus    int    `json:"order_status"`    // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	Amount         int    `json:"amount"`          // 提现金额
	ChargebackDate string `json:"chargeback_date"` // 退单日期,格式为YYYYMMDD
	ErrorMsg       string `json:"error_msg"`       // 错误信息
}
