package models

// MerchantInfo 商户列表中的单个商户信息的结构体 http://47.110.246.50:6524/docs/qzt/qzt-1d50jnpp82aqk
type MerchantInfo struct {
	MemberNo          string `json:"member_no"`                     // 会员编号
	OutMemberNo       string `json:"out_member_no"`                 // 外部会员号
	MerchantNo        string `json:"merchant_no"`                   // 钱账通商户号
	OutMerchantNo     string `json:"out_merchant_no,omitempty"`     // 外部商户号
	MerchantType      int    `json:"merchant_type"`                 // 商户类型
	MerchantStatus    int    `json:"merchant_status"`               // 商户状态,http://47.110.246.50:6524/docs/qzt/qzt-1d50j4fhhrekk
	ChannelNo         string `json:"channel_no"`                    // 渠道编号
	ChannelMerchantNo string `json:"channel_merchant_no,omitempty"` // 渠道商户号
	ChannelTerminalNo string `json:"channel_terminal_no,omitempty"` // 渠道终端号
	HandlerTime       string `json:"handler_time"`                  // 处理时间,格式为yyyy-MM-dd HH:mm:ss
	ErrMsg            string `json:"err_msg,omitempty"`             // 错误信息
	BuisTypeCode      string `json:"buis_type_code"`                // 终端业务码,BANK_CARD:银行卡,QR_CODE_CARD:扫码
	BuisTypeName      string `json:"buisTypeName"`                  // 终端业务名称,如银行卡、扫码
	BusiActiveNo      string `json:"busi_active_no"`                // 终端业务激活码
}

// BindMemberMerchantParams http://47.110.246.50:6524/docs/qzt/qzt-1d4raqu0cuofa
type BindMemberMerchantParams struct {
	MemberNo          string  `json:"member_no"`           // 会员标识
	ChannelMerchantNo string  `json:"channel_merchant_no"` // 渠道商户号
	ChannelTerminalNo string  `json:"channel_terminal_no"` // 渠道终端号,POS,BMCP必选
	MerchantType      string  `json:"merchant_type"`       // 商户类型,POS(POS机/码牌/收银台),BMCP(专业化扫码)
	SellerScale       float64 `json:"seller_scale"`        // 商家分账比例(商家获得比列),20%传0.2,最长4位小数
	Condition         string  `json:"condition"`           // 条件,ge:大于等于(默认),eq:等于
	BackUrl           string  `json:"back_url"`            // 后台通知地址
}

// BindMemberMerchantResult http://47.110.246.50:6524/docs/qzt/qzt-1d4raqu0cuofa
type BindMemberMerchantResult struct {
	MemberNo       string `json:"member_no"`       // 会员号
	MerchantNo     string `json:"merchant_no"`     // 钱账通商户号
	MerchantStatus int    `json:"merchant_status"` // 商户状态,1:审核中,2:进件中,3:有效,4:无效
}

// UnbindMemberMerchantParams http://47.110.246.50:6524/docs/qzt/qzt-1d4rar93ujpf3
type UnbindMemberMerchantParams struct {
	MemberNo          string `json:"member_no"`           // 会员标识
	ChannelMerchantNo string `json:"channel_merchant_no"` // 渠道商户号
	ChannelTerminalNo string `json:"channel_terminal_no"` // 渠道终端号,POS必选
	MerchantType      string `json:"merchant_type"`       // 商户类型,POS(POS终端),BMCP(开户终端),BTS_PAY(无卡分期支付)
}

// UnbindMemberMerchantResult http://47.110.246.50:6524/docs/qzt/qzt-1d4rar93ujpf3
type UnbindMemberMerchantResult struct {
	MemberNo   string `json:"member_no"`   // 会员号
	MerchantNo string `json:"merchant_no"` // 钱账通商户号
}

// QueryMemberMerchantParams http://47.110.246.50:6524/docs/qzt/qzt-1d4raol7h83o8
type QueryMemberMerchantParams struct {
	MemberNo      string `json:"member_no,omitempty"`       // 会员标识,二选一
	OutMemberNo   string `json:"out_member_no,omitempty"`   // 外部会员号,二选一
	OutMerchantNo string `json:"out_merchant_no,omitempty"` // 外部商户号,可选
	MerchantNo    string `json:"merchant_no,omitempty"`     // 钱账通商户号,可选
}

// QueryMemberMerchantResult http://47.110.246.50:6524/docs/qzt/qzt-1d4raol7h83o8
type QueryMemberMerchantResult struct {
	MerchantList []MerchantInfo `json:"merchant_list"`
}

// ReplenishMemberMerchantAttachmentParams http://47.110.246.50:6524/docs/qzt/qzt-1d8506t36jvvv
type ReplenishMemberMerchantAttachmentParams struct {
	MemberNo         string `json:"member_no"`          // 会员标识,必选
	MerchantNo       string `json:"merchant_no"`        // 钱账通商户号,必选
	AttachmentType   int    `json:"attachment_type"`    // 外部会员号,0:身份证正面照,1:身份证反面照,2:银行卡/对公账户,3:营业执照,4:商户门头照,5:电子签约,6:商铺内部照片,7:合作资质证明,8:食品经营相关资质,9:线下纸质协议,10:租赁合同,100:其他
	AttachmentFileNo string `json:"attachment_file_no"` // 附件文件编号
}

// ReplenishMemberMerchantAttachmentResult http://47.110.246.50:6524/docs/qzt/qzt-1d8506t36jvvv
type ReplenishMemberMerchantAttachmentResult struct {
	MerchantNo string `json:"merchant_no"` // 钱账通商户号
}

// QueryMemberMerchantQuotaParams http://47.110.246.50:6524/docs/qzt/qzt-1d4rarlm3eftm
type QueryMemberMerchantQuotaParams struct {
	MemberNo   string `json:"member_no"`   // 会员标识
	MerchantNo string `json:"merchant_no"` // 钱账通商户号
}

// QueryMemberMerchantQuotaResult http://47.110.246.50:6524/docs/qzt/qzt-1d4rarlm3eftm
type QueryMemberMerchantQuotaResult struct {
	DebitPerLimit    int `json:"debit_per_limit,omitempty"`    // 借记卡单笔交易限额,分
	DebitDayLimit    int `json:"debit_day_limit,omitempty"`    // 借记卡单日交易限额,分
	DebitMonthLimit  int `json:"debit_month_limit,omitempty"`  // 借记卡单月交易限额,分
	CreditPerLimit   int `json:"credit_per_limit,omitempty"`   // 贷记卡单笔交易限额,分
	CreditDayLimit   int `json:"credit_day_limit,omitempty"`   // 贷记卡单日交易限额,分
	CreditMonthLimit int `json:"credit_month_limit,omitempty"` // 贷记卡单月交易限额,分
	WechatPerLimit   int `json:"wechat_per_limit,omitempty"`   // 扫码单笔交易限额,分
	WechatDayLimit   int `json:"wechat_day_limit,omitempty"`   // 扫码单日交易限额,分
	WechatMonthLimit int `json:"wechat_month_limit,omitempty"` // 扫码单月交易限额,分
	D0PerLimit       int `json:"d0_per_limit,omitempty"`       // D0单笔限额,分
	D0DayLimit       int `json:"d0_day_limit,omitempty"`       // D0单日限额,分
	WechatMinAmt     int `json:"wechat_min_amt,omitempty"`     // 扫码手续费保底值,分
	WechatMaxAmt     int `json:"wechat_max_amt,omitempty"`     // 扫码手续费封顶值,分
	DebitMixAmt      int `json:"debit_mix_amt,omitempty"`      // 借记卡手续费保底值,分
	DebitMaxAmt      int `json:"debit_max_amt,omitempty"`      // 借记卡手续费封顶值,分
	CreditMixAmt     int `json:"credit_mix_amt,omitempty"`     // 贷记卡手续费保底值,分
	CreditMaxAmt     int `json:"credit_max_amt,omitempty"`     // 贷记卡手续费封顶值,分
}

// SubmitMemberMerchantReconsiderParams http://47.110.246.50:6524/docs/qzt/qzt-1ffo3mtv45fbm
type SubmitMemberMerchantReconsiderParams struct {
	MemberNo string `json:"member_no"` // 会员标识
}

// SubmitMemberMerchantReconsiderResult http://47.110.246.50:6524/docs/qzt/qzt-1ffo3mtv45fbm
type SubmitMemberMerchantReconsiderResult struct {
	Message string `json:"message"` // 复议备注
}
