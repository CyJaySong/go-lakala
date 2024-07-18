package models

// DepositParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbccp043dlg
type DepositParams struct {
	MemberNo         string                 `json:"member_no"`            // 会员标识
	OutOrderNo       string                 `json:"out_order_no"`         // 应用平台订单号,接口调用方上送、唯一
	AccountType      string                 `json:"account_type"`         // 账户类型编码,固定值:1000
	SpecialAccountNo string                 `json:"special_account_no"`   // 专用账户,营销专用账户编号:S004,退货专用账户编号:S010
	Amount           int                    `json:"amount"`               // 充值金额,单位:分
	FrontUrl         string                 `json:"front_url,omitempty"`  // 前台通知地址
	BackUrl          string                 `json:"back_url,omitempty"`   // 后台通知地址
	Expire           string                 `json:"expire,omitempty"`     // 订单过期时间,默认过期时间24小时,yyyy-MM-dd HH:mm:ss
	PayMethod        ConsumeParamsPayMethod `json:"pay_method"`           // 支付方式,TRANSFER_PAY(转账支付)
	TerminalIP       string                 `json:"terminal_ip"`          // 终端ip
	OrderName        string                 `json:"order_name,omitempty"` // 订单名称
	Remark           string                 `json:"remark,omitempty"`     // 备注
	Exts             string                 `json:"exts,omitempty"`       // 扩展信息,最多32个字符
}

// DepositResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbccp043dlg
type DepositResult struct {
	OrderNo                string  `json:"order_no"`                             // 存管系统订单号
	OrderStatus            int     `json:"order_status"`                         // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	ErrorMessage           string  `json:"error_message,omitempty"`              // 错误信息
	PayInfo                PayInfo `json:"pay_info"`                             // 支付信息,pay_no:(转账支付),原生支付时返回
	ChannelTradeNo         string  `json:"channel_trade_no,omitempty"`           // 渠道交易流水号(收单)
	QztChannelPayRequestNo string  `json:"qzt_channel_pay_request_no,omitempty"` // 钱帐通请求通道的流水号
	ChannelSeqNo           string  `json:"channel_seq_no,omitempty"`             // 渠道支付流水号(收单)
	PayChannelTradeNo      string  `json:"pay_channel_trade_no,omitempty"`       // 支付通道交易流水号(支付宝、微信)
	ThirdPartyPayment      string  `json:"third_party_payment,omitempty"`        // 三方支付渠道
	EndTime                string  `json:"end_time,omitempty"`                   // 完成时间,yyyy-MM-dd HH:mm:ss
}

// WithdrawParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcdpaq4jns
type WithdrawParams struct {
	MemberNo           string `json:"member_no"`               // 会员标识
	OutOrderNo         string `json:"out_order_no"`            // 应用平台订单号
	AccountType        string `json:"account_type"`            // 账户类型编码,固定值:1000
	SpecialAccountType string `json:"special_account_type"`    // 专用账户类型编码
	Amount             int    `json:"amount"`                  // 提现金额(含平台手续费),单位:分
	BillingFee         int    `json:"billing_fee"`             // 平台手续费,单位:分
	BankCardId         string `json:"bank_card_no"`            // 钱账通绑定银行卡的编号
	WithdrawType       string `json:"withdraw_type,omitempty"` // 提现方式,D0:D0提现,T1:T1提现,D1:D1提现,默认T1提现
	FrontUrl           string `json:"front_url,omitempty"`     // 前台通知地址
	BackUrl            string `json:"back_url"`                // 后台通知地址
	Expire             string `json:"expire,omitempty"`        // 订单过期时间,格式:yyyy-MM-ddHH:mm:ss,默认过期时间24小时
	OrderName          string `json:"order_name,omitempty"`    // 订单名称
	Exts               string `json:"exts,omitempty"`          // 扩展信息,最多32个字符
}

// WithdrawResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcdpaq4jns
type WithdrawResult struct {
	OrderNo        string `json:"order_no"`                   // 存管系统订单号
	OrderStatus    int    `json:"order_status"`               // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	ErrorMessage   string `json:"error_message,omitempty"`    // 错误信息
	ChannelTradeNo string `json:"channel_trade_no,omitempty"` // 渠道交易流水号
}

// OrderConsumeParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbceebjccgf
type OrderConsumeParams struct {
	BuyerMemberNo            string                  `json:"buyer_member_no,omitempty"`              // 买家会员标识
	SellerMemberNo           string                  `json:"seller_member_no"`                       // 卖家会员标识
	OutOrderNo               string                  `json:"out_order_no"`                           // 应用平台订单号
	Amount                   int                     `json:"amount"`                                 // 订单金额,单位:分
	PayAmount                int                     `json:"pay_amount"`                             // 首次支付金额,单位:分,支付金额<=订单金额,本次不支付传0
	OutRequestNo             string                  `json:"out_request_no,omitempty"`               // 平台支付请求号
	PayMethod                *ConsumeParamsPayMethod `json:"pay_method,omitempty"`                   // 支付方式,pay_amount=0可不传
	TerminalIP               string                  `json:"terminal_ip"`                            // 终端IP,支付方式为SCAN时上传支付者IP,其他支付方式上传商户支付相关设备的ip
	SplitRuleData            *SplitRuleData          `json:"split_rule_data,omitempty"`              // 分账规则,分账规则为空时默认全给商户,分账规则中分账方必须实名并且设置会员角色,分账记录总金额=订单金额
	FrontUrl                 string                  `json:"front_url,omitempty"`                    // 前台通知地址
	BackUrl                  string                  `json:"back_url,omitempty"`                     // 后台通知地址
	OrderName                string                  `json:"order_name"`                             // 订单名称,最多100个字符
	Exts                     string                  `json:"exts,omitempty"`                         // 扩展参数,最多32个字符
	SplitMode                int                     `json:"split_mode"`                             // 分账模式,0:即时分账,1:确权分账
	AutoSettleTypeAfterSplit int                     `json:"auto_settle_type_after_split,omitempty"` // 分账后自动提现模式,11:分账后D1提现(订单笔笔结算专用,项目需要特殊申请,正常不需要传该字段)
}

// OrderConsumeResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbceebjccgf
type OrderConsumeResult struct {
	OrderNo                string                `json:"order_no"`                             // 订单号
	PayStatus              int                   `json:"pay_status,omitempty"`                 // 支付状态,pay_amount>0时返回,http://47.110.246.50:6524/docs/qzt/qzt-1ekbe0j5a5ard
	OrderStatus            int                   `json:"order_status"`                         // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	ErrorMessage           string                `json:"error_message,omitempty"`              // 错误信息
	PaySeqNo               string                `json:"pay_seq_no,omitempty"`                 // 支付流水号,pay_amount>0时返回
	SplitRuleResult        []SplitRuleResultItem `json:"split_rule_result"`                    // 分账会员列表
	SplitSeqNo             string                `json:"split_seq_no,omitempty"`               // 分账流水号
	BuyerMemberNo          string                `json:"buyer_member_no,omitempty"`            // 买家会员标识,有买家会员时返回
	PayInfo                PayInfo               `json:"pay_info,omitempty"`                   // 支付信息,原生支付时返回
	QztChannelPayRequestNo string                `json:"qzt_channel_pay_request_no,omitempty"` // 钱帐通请求通道的流水号,支付完成返回
	ChannelTradeNo         string                `json:"channel_trade_no,omitempty"`           // 渠道交易流水号(收单),支付完成返回
	ChannelSeqNo           string                `json:"channel_seq_no,omitempty"`             // 渠道支付流水号(收单),支付完成返回
	PayChannelTradeNo      string                `json:"pay_channel_trade_no,omitempty"`       // 支付通道交易流水号(支付宝、微信),支付完成返回
	ThirdPartyPayment      string                `json:"third_party_payment,omitempty"`        // 三方支付渠道,交易完成返回,微信:WECHAT,支付宝:ALIPAY
	OpenId                 string                `json:"open_id,omitempty"`                    // 微信/支付宝/云闪付的openid
	SubOpenId              string                `json:"sub_open_id,omitempty"`                // 微信子appid的openid
	PayTime                string                `json:"pay_time,omitempty"`                   // 支付时间,yyyy-MM-dd HH:mm:ss
}

// OrderConsumePayParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcfiauedbj
type OrderConsumePayParams struct {
	OrderNo      string                 `json:"order_no"`                 // 钱账通订单号
	OutRequestNo string                 `json:"out_request_no,omitempty"` // 平台支付请求号
	PayAmount    int                    `json:"pay_amount"`               // 支付金额,单位:分
	PayMethod    ConsumeParamsPayMethod `json:"pay_method"`               // 支付方式
	TerminalIP   string                 `json:"terminal_ip"`              // 终端IP,支付方式为SCAN时上传支付者IP,其他支付方式上传商户支付相关设备的ip
	Exts         string                 `json:"exts,omitempty"`           // 扩展参数,最多32个字符
}

// OrderConsumePayResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcfiauedbj
type OrderConsumePayResult struct {
	OrderNo                string  `json:"order_no"`                             // 存管系统订单号
	OrderStatus            int     `json:"order_status"`                         // 支付状态,pay_amount>0时返回,http://47.110.246.50:6524/docs/qzt/qzt-1ekbe0j5a5ard
	PayStatus              int     `json:"pay_status"`                           // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	ErrorMessage           string  `json:"error_message,omitempty"`              // 错误信息
	OutRequestNo           string  `json:"out_request_no,omitempty"`             // 外部支付请求号
	PaySeqNo               string  `json:"pay_seq_no"`                           // 支付流水号
	PayInfo                PayInfo `json:"pay_info,omitempty"`                   // 支付信息,原生支付时返回
	QztChannelPayRequestNo string  `json:"qzt_channel_pay_request_no,omitempty"` // 钱帐通请求通道的流水号,支付完成返回
	ChannelTradeNo         string  `json:"channel_trade_no,omitempty"`           // 渠道交易流水号(收单),支付完成返回
	ChannelSeqNo           string  `json:"channel_seq_no,omitempty"`             // 渠道支付流水号(收单),支付完成返回
	PayChannelTradeNo      string  `json:"pay_channel_trade_no,omitempty"`       // 支付通道交易流水号(支付宝、微信),支付完成返回
	ThirdPartyPayment      string  `json:"third_party_payment,omitempty"`        // 三方支付渠道,交易完成返回,微信:WECHAT,支付宝:ALIPAY
	OpenId                 string  `json:"open_id,omitempty"`                    // 微信/支付宝/云闪付的openid
	SubOpenId              string  `json:"sub_open_id,omitempty"`                // 微信子appId的openid
	PayTime                string  `json:"pay_time,omitempty"`                   // 支付时间,yyyy-MM-dd HH:mm:ss
}

// OrderConsumeBatchPayParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbchps1ckgp
type OrderConsumeBatchPayParams struct {
	SubOrderNoList  []string               `json:"sub_order_no_list"`           // 子订单号列表,order_no以String传输
	AcquireMemberNo string                 `json:"acquire_member_no,omitempty"` // 收单商户会员,可选,若不传,如果子订单中卖家会员一致则使用卖家会员,如卖家会员不一致则使用平台会员;若上传,上传会员必须是子订单中的卖家
	OutOrderNo      string                 `json:"out_order_no"`                // 应用平台订单号
	OutRequestNo    string                 `json:"out_request_no,omitempty"`    // 外部请求号
	Amount          int                    `json:"amount"`                      // 订单总金额
	PayMethod       ConsumeParamsPayMethod `json:"pay_method"`                  // 支付方式
	TerminalIP      string                 `json:"terminal_ip"`                 // 终端IP,支付方式为SCAN时上传支付者IP,其他支付方式上传商户支付相关设备的ip
	Exts            string                 `json:"exts,omitempty"`              // 扩展参数,最多32个字符
	FrontUrl        string                 `json:"front_url,omitempty"`         // 前台通知地址
	BackUrl         string                 `json:"back_url,omitempty"`          // 后台通知地址
}

// OrderConsumeBatchPayResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbchps1ckgp
type OrderConsumeBatchPayResult struct {
	BatchOrderNo           string  `json:"batch_order_no"`                       // 批量支付订单号
	OrderStatus            int     `json:"order_status"`                         // 支付状态,pay_amount>0时返回,http://47.110.246.50:6524/docs/qzt/qzt-1ekbe0j5a5ard
	PayStatus              int     `json:"pay_status"`                           // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	ErrorMessage           string  `json:"error_message,omitempty"`              // 错误信息
	OutRequestNo           string  `json:"out_request_no,omitempty"`             // 外部支付请求号
	PaySeqNo               string  `json:"pay_seq_no"`                           // 支付流水号,pay_amount>0时返回
	PayInfo                PayInfo `json:"pay_info,omitempty"`                   // 支付信息,原生支付时返回
	QztChannelPayRequestNo string  `json:"qzt_channel_pay_request_no,omitempty"` // 钱帐通请求通道的流水号,支付完成返回
	ChannelTradeNo         string  `json:"channel_trade_no,omitempty"`           // 渠道交易流水号(收单),支付完成返回
	ChannelSeqNo           string  `json:"channel_seq_no,omitempty"`             // 渠道支付流水号(收单),支付完成返回
	PayChannelTradeNo      string  `json:"pay_channel_trade_no,omitempty"`       // 支付通道交易流水号(支付宝、微信),支付完成返回
	ThirdPartyPayment      string  `json:"third_party_payment,omitempty"`        // 三方支付渠道,交易完成返回,微信:WECHAT,支付宝:ALIPAY
	OpenId                 string  `json:"open_id,omitempty"`                    // 微信/支付宝/云闪付的openid
	SubOpenId              string  `json:"sub_open_id,omitempty"`                // 微信子appId的openid
	PayTime                string  `json:"pay_time,omitempty"`                   // 支付时间,yyyy-MM-dd HH:mm:ss
}

// OrderActionConfirmParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcpiggemfj
type OrderActionConfirmParams struct {
	MemberNo string `json:"member_no,omitempty"` // 会员标识,可选
	OrderNo  string `json:"order_no"`            // 订单编号,必选
	SeqNo    string `json:"seq_no,omitempty"`    // 支付流水号,3、4、6时必填,可选
	Action   int    `json:"action"`              // 订单动作,3:完成订单,4:支付,6:转账(平安模式),必选
	Code     string `json:"code"`                // 验证码,必选
	PayNo    string `json:"pay_no,omitempty"`    // 支付编号,快捷支付需上传返回的pay_token,String类型,可选
	IP       string `json:"ip,omitempty"`        // IP地址,String类型,可选
}

// OrderActionConfirmResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcpiggemfj
type OrderActionConfirmResult struct {
	OrderNo                string                `json:"order_no"`                             // 存管系统订单号,必选
	PayStatus              int                   `json:"pay_status,omitempty"`                 // 支付状态,pay_amount>0时返回,可选
	OrderStatus            int                   `json:"order_status"`                         // 订单状态,必选
	ErrorMessage           string                `json:"error_message,omitempty"`              // 错误信息,可选
	PaySeqNo               string                `json:"pay_seq_no,omitempty"`                 // 支付流水号,pay_amount>0时返回,必选
	SplitRuleResult        []SplitRuleResultItem `json:"split_rule_result"`                    // 分账会员列表,必选
	SplitSeqNo             string                `json:"split_seq_no,omitempty"`               // 分账流水号,可选
	QztChannelPayRequestNo string                `json:"qzt_channel_pay_request_no,omitempty"` // 钱帐通请求通道的流水号,支付完成返回,可选
	ChannelTradeNo         string                `json:"channel_trade_no,omitempty"`           // 渠道交易流水号(收单),支付成功返回,可选
	ChannelSeqNo           string                `json:"channel_seq_no,omitempty"`             // 渠道支付流水号(收单),支付成功返回,可选
	PayChannelTradeNo      string                `json:"pay_channel_trade_no,omitempty"`       // 支付通道交易流水号(支付宝、微信),支付成功返回,可选
	PayTime                string                `json:"pay_time,omitempty"`                   // 支付时间,格式为yyyy-MM-dd HH:mm:ss,可选
}

// OrderConsumeRefundParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcin2rs400
type OrderConsumeRefundParams struct {
	OrderNo      string                         `json:"order_no"`               // 原订单（需要退款的订单编号）
	OutRequestNo string                         `json:"out_request_no"`         // 应用平台请求号
	MemberNo     string                         `json:"member_no,omitempty"`    // 会员标识
	RefundAmount int                            `json:"refund_amount"`          // 退款总额,退款金额<=原订单未退款金额
	PaySeqNo     string                         `json:"pay_seq_no"`             // 支付流水号
	RefundRule   []OrderConsumeRefundParamsRule `json:"refund_rule,omitempty"`  // 退款规则，交易已分账时必填
	RefundParam  string                         `json:"refund_param,omitempty"` // 退款参数，只有转账支付退款时需要
	Remark       string                         `json:"remark,omitempty"`       // 备注
}

type OrderConsumeRefundParamsRule struct {
	MemberNo         string `json:"member_no"`                    // 会员标识
	AccountTypeNo    string `json:"account_type_no"`              // 账户类型编号,固定值:1000
	Amount           int    `json:"amount"`                       // 退款金额
	SpecialAccountNo string `json:"special_account_no,omitempty"` // 专用账户
}

// OrderConsumeRefundResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcin2rs400
type OrderConsumeRefundResult struct {
	OrderNo                string `json:"order_no"`                             // 交易订单号
	OrderStatus            int    `json:"order_status"`                         // 订单状态
	RefundStatus           int    `json:"refund_status"`                        // 退款状态
	ErrorMessage           string `json:"error_message,omitempty"`              // 错误信息
	OutRefundRequestNo     string `json:"out_refund_request_no"`                // 外部退款请求号
	RefundSeqNo            string `json:"refund_seq_no"`                        // 退款流水号,有退款时必选
	QztChannelPayRequestNo string `json:"qzt_channel_pay_request_no,omitempty"` // 钱账通请求通道的流水号,支付完成返回
	ChannelTradeNo         string `json:"channel_trade_no,omitempty"`           // 渠道交易流水号(收单),退款完成可选
	ChannelSeqNo           string `json:"channel_seq_no,omitempty"`             // 渠道支付流水号(收单),退款完成可选
	PayChannelTradeNo      string `json:"pay_channel_trade_no,omitempty"`       // 支付通道交易流水号(支付宝、微信),退款完成可选
	RefundTime             string `json:"refund_time,omitempty"`                // 退款时间,yyyy-MM-dd HH:mm:ss
}

// OrderConsumeCloseParams http://47.110.246.50:6524/docs/qzt/qzt-1f00qp3efdvso
type OrderConsumeCloseParams struct {
	PaySeqNo     string                              `json:"pay_seq_no"`    // 支付请求号
	LocationInfo OrderConsumeCloseParamsLocationInfo `json:"location_info"` // 地址位置信息
}

type OrderConsumeCloseParamsLocationInfo struct {
	RequestIP   string `json:"request_ip"`   // 请求方IP地址
	BaseStation string `json:"base_station"` // 基站信息
	Location    string `json:"location"`     // 维度,经度
}

// OrderConsumeCloseResult http://47.110.246.50:6524/docs/qzt/qzt-1f00qp3efdvso
type OrderConsumeCloseResult struct{}

// GetOrderStatusParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcl771rsu6
type GetOrderStatusParams struct {
	OrderNo    string `json:"order_no,omitempty"`     // 订单编号
	OutOrderNo string `json:"out_order_no,omitempty"` // 平台订单号
}

// GetOrderStatusResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcl771rsu6
type GetOrderStatusResult struct {
	OrderNo            string        `json:"order_no"`                       // 订单编号
	OrderStatus        int           `json:"order_status"`                   // 订单状态
	OutOrderNo         string        `json:"out_order_no"`                   // 商户订单号
	Amount             int           `json:"amount"`                         // 金额
	TransferResultList []TransferRes `json:"transfer_result_list,omitempty"` // 转账结果,税优转账订单有
}

// GetOrderPayStatusParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcm6ee3m1n
type GetOrderPayStatusParams struct {
	OrderNo  string `json:"order_no"`   // 订单编号
	PaySeqNo string `json:"pay_seq_no"` // 支付流水号
}

// GetOrderPayStatusResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcm6ee3m1n
type GetOrderPayStatusResult struct {
	PayStatus              int                    `json:"pay_status"`                           // 支付状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbe0j5a5ard
	OutOrderNo             string                 `json:"out_order_no"`                         // 商户订单号
	PayMethod              ConsumeParamsPayMethod `json:"pay_method"`                           // 支付方式,
	Amount                 int                    `json:"amount"`                               // 金额
	FeeAmount              int                    `json:"fee_amount,omitempty"`                 // 渠道手续费
	QztChannelPayRequestNo string                 `json:"qzt_channel_pay_request_no,omitempty"` // 钱帐通请求通道的流水号
	ChannelTradeNo         string                 `json:"channel_trade_no,omitempty"`           // 渠道交易流水号(收单)
	ChannelSeqNo           string                 `json:"channel_seq_no,omitempty"`             // 渠道支付流水号(收单)
	PayChannelTradeNo      string                 `json:"pay_channel_trade_no,omitempty"`       // 支付通道交易流水号(支付宝、微信)
	PayTime                string                 `json:"pay_time,omitempty"`                   // 支付时间,yyyy-MM-dd HH:mm:ss
	ThirdPartyPayment      string                 `json:"third_party_payment,omitempty"`        // 三方支付渠道,微信:WECHAT,支付宝:ALIPAY,借记卡:DEBIT_CARD,贷记卡:CREDIT_CARD
}

// GetOrderRefundStatusParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcn7n81jcq
type GetOrderRefundStatusParams struct {
	OrderNo     string `json:"order_no"`      // 订单编号
	RefundSeqNo string `json:"refund_seq_no"` // 退款流水号
}

// GetOrderRefundStatusResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcn7n81jcq
type GetOrderRefundStatusResult struct {
	RefundStatus           int    `json:"refund_status"`              // 退款状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbe0vhq9dkh
	OutRefundRequestNo     string `json:"out_refund_request_no"`      // 外部退款请求号
	RefundSeqNo            string `json:"refund_seq_no"`              // 退款流水号,有退款时必选
	QztChannelPayRequestNo string `json:"qzt_channel_pay_request_no"` // 钱帐通请求通道的流水号
	ChannelTradeNo         string `json:"channel_trade_no"`           // 渠道交易流水号(收单),退款完成可选
	ChannelSeqNo           string `json:"channel_seq_no"`             // 渠道支付流水号(收单),退款完成可选
	PayChannelTradeNo      string `json:"pay_channel_trade_no"`       // 支付通道交易流水号(支付宝、微信),退款完成可选
	RefundTime             string `json:"refund_time"`                // 退款时间,yyyy-MM-dd HH:mm:ss
	RefundFeeAmount        int    `json:"refund_fee_amount"`          // 渠道退款手续费
	RefundAmount           int    `json:"refund_amount"`              // 退款金额
}

// GetOrderDetailInfoParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcnqllsto1
type GetOrderDetailInfoParams struct {
	OrderNo    string `json:"order_no"`     // 订单编号
	OutOrderNo string `json:"out_order_no"` // 平台订单号
}

// GetOrderDetailInfoResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcnqllsto1
type GetOrderDetailInfoResult struct {
	OrderNo            string                       `json:"order_no"`                       // 订单编号
	OrderStatus        int                          `json:"order_status"`                   // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	OutOrderNo         string                       `json:"out_order_no"`                   // 商户订单号
	OrderAmount        int                          `json:"order_amount"`                   // 订单金额
	PayAmount          int                          `json:"pay_amount"`                     // 已支付金额
	SplitAmount        int                          `json:"split_amount"`                   // 已分账金额
	RefundAmount       int                          `json:"refund_amount"`                  // 已退款金额
	PayList            []OrderDetailInfoPaymentInfo `json:"pay_list,omitempty"`             // 支付列表
	RefundList         []OrderDetailInfoRefundInfo  `json:"refund_list,omitempty"`          // 退款列表
	SplitList          []OrderDetailInfoSplitInfo   `json:"split_list,omitempty"`           // 分账列表
	TransferResultList []TransferRes                `json:"transfer_result_list,omitempty"` // 转账结果,税优转账订单有
}

type OrderDetailInfoPaymentInfo struct {
	OutRequestNo           string                 `json:"out_request_no,omitempty"`             // 平台支付请求号
	PaySeqNo               string                 `json:"pay_seq_no"`                           // 支付流水号
	PayStatus              int                    `json:"pay_status"`                           // 支付状态
	PayMethod              ConsumeParamsPayMethod `json:"pay_method"`                           // 支付方式
	Amount                 int                    `json:"amount"`                               // 金额
	QztChannelPayRequestNo string                 `json:"qzt_channel_pay_request_no,omitempty"` // 钱帐通请求通道的流水号,支付完成返回
	ChannelTradeNo         string                 `json:"channel_trade_no,omitempty"`           // 渠道交易流水号(收单)
	ChannelSeqNo           string                 `json:"channel_seq_no,omitempty"`             // 渠道支付流水号(收单)
	PayChannelTradeNo      string                 `json:"pay_channel_trade_no,omitempty"`       // 支付通道交易流水号(支付宝、微信)
	PayTime                string                 `json:"pay_time,omitempty"`                   // 支付时间,格式 yyyy-MM-dd HH:mm:ss
	ThirdPartyPayment      string                 `json:"third_party_payment,omitempty"`        // 三方支付渠道,微信:WECHAT,支付宝:ALIPAY,借记卡:DEBIT_CARD,贷记卡:CREDIT_CARD
}

type OrderDetailInfoRefundInfo struct {
	RefundStatus           int    `json:"refund_status"`                        // 退款状态
	OutRequestNo           string `json:"out_request_no"`                       // 应用平台退款请求号
	Amount                 int    `json:"amount"`                               // 金额
	RefundSeqNo            string `json:"refund_seq_no"`                        // 退款流水号,有退款时必选
	QztChannelPayRequestNo string `json:"qzt_channel_pay_request_no,omitempty"` // 钱帐通请求通道的流水号,支付完成返回
	ChannelTradeNo         string `json:"channel_trade_no,omitempty"`           // 渠道交易流水号(收单)
	ChannelSeqNo           string `json:"channel_seq_no,omitempty"`             // 渠道支付流水号(收单)
	PayChannelTradeNo      string `json:"pay_channel_trade_no,omitempty"`       // 支付通道交易流水号(支付宝、微信)
	RefundTime             string `json:"refund_time,omitempty"`                // 退款时间,格式 yyyy-MM-dd HH:mm:ss
}

type OrderDetailInfoSplitInfo struct {
	SplitStatus  int    `json:"split_status"`             // 分账状态
	OutRequestNo string `json:"out_request_no,omitempty"` // 应用平台退款请求号
	Amount       int    `json:"amount"`                   // 金额
	SplitSeqNo   string `json:"split_seq_no"`             // 分账流水号
	MemberIdNo   string `json:"member_id_no"`             // 分账会员编号
	SplitTime    string `json:"split_time,omitempty"`     // 分账时间,格式 yyyy-MM-dd HH:mm:ss
}

// QueryOrderReceiptParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcs393drqg
type QueryOrderReceiptParams struct {
	OrderNo string `json:"order_no"` // 订单编号
}

// QueryOrderReceiptResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcs393drqg
type QueryOrderReceiptResult struct {
	IncomeFileNo       string `json:"income_file_no,omitempty"` // 入金电子回单文件编号
	ExpensesFileNoList []struct {
		MemberNo string `json:"member_no"` // 分账的会员ID
		Amount   int    `json:"amount"`    // 分账的金额
		FileNo   string `json:"file_no"`   // 回单文件编号
	} `json:"expenses_file_no_list,omitempty"` // 出金电子回单文件编号列表
}
