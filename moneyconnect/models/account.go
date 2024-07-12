package models

// GetAccountBalanceParams http://47.110.246.50:6524/docs/qzt/qzt-1che8b7438hhi
type GetAccountBalanceParams struct {
	MemberNo         string `json:"member_no"`                    // 会员标识
	AccountType      string `json:"account_type"`                 // 账户类型编码,固定值:1000
	SpecialAccountNo string `json:"special_account_no,omitempty"` // 专用账户,可选
}

// GetAccountBalanceResult http://47.110.246.50:6524/docs/qzt/qzt-1che8b7438hhi
type GetAccountBalanceResult struct {
	TotalAmount     int `json:"total_amount"`     // 所有余额(单位:分),包括可用余额、冻结金额
	AvailableAmount int `json:"available_amount"` // 可用余额(单位:分)
	FreezeAmount    int `json:"freeze_amount"`    // 已冻结金额(单位:分)
}

// GetTrustAccountBalanceParams http://47.110.246.50:6524/docs/qzt/qzt-1d4rhsg1pflhr
type GetTrustAccountBalanceParams struct {
	MemberNo  string `json:"member_no"`  // 会员标识
	TrustType string `json:"trust_type"` // 托管类型;pingan:平安银行
}

// GetTrustAccountBalanceResult http://47.110.246.50:6524/docs/qzt/qzt-1d4rhsg1pflhr
type GetTrustAccountBalanceResult struct {
	TotalAmount      int `json:"total_amount"`      // 所有余额(单位:分)
	WithdrawalAmount int `json:"withdrawal_amount"` // 可提现余额(单位:分)
}

// GetAccountIncomeCostParams http://47.110.246.50:6524/docs/qzt/qzt-1che8cosbrq64
type GetAccountIncomeCostParams struct {
	MemberNo         string `json:"member_no"`                    // 会员标识
	AccountType      string `json:"account_type"`                 // 账户类型编码,固定值:1000
	SpecialAccountNo string `json:"special_account_no,omitempty"` // 专用账户,可选
	StartTime        string `json:"start_time"`                   // 开始时间,格式为yyyy-MM-dd
	EndTime          string `json:"end_time"`                     // 结束时间,格式为yyyy-MM-dd,时间不能超过一个月
}

// GetAccountIncomeCostResult http://47.110.246.50:6524/docs/qzt/qzt-1che8cosbrq64
type GetAccountIncomeCostResult struct {
	IncomeAmount int `json:"income_amount"` // 收入金额(单位:分)
	CostAmount   int `json:"cost_amount"`   // 支出金额(单位:分)
}

// FreezeAccountBalanceParams http://47.110.246.50:6524/docs/qzt/qzt-1che8dslbifrc
type FreezeAccountBalanceParams struct {
	MemberNo         string `json:"member_no"`                    // 会员标识
	OutOrderNo       string `json:"out_order_no"`                 // 平台订单编号
	AccountType      string `json:"account_type"`                 // 账户类型编码,固定值:1000
	SpecialAccountNo string `json:"special_account_no,omitempty"` // 专用账户,可选
	Amount           int    `json:"amount"`                       // 冻结金额,单位:分
	FreezeNo         string `json:"freeze_no,omitempty"`          // 冻结号,可选
	Remark           string `json:"remark,omitempty"`             // 冻结说明,可选
}

// FreezeAccountBalanceResult http://47.110.246.50:6524/docs/qzt/qzt-1che8dslbifrc
type FreezeAccountBalanceResult struct {
	OrderNo  string `json:"order_no"`  // 订单号
	FreezeNo string `json:"freeze_no"` // 冻结编号
}

// UnfreezeAccountBalanceParams http://47.110.246.50:6524/docs/qzt/qzt-1che8e73223k6
type UnfreezeAccountBalanceParams struct {
	MemberNo   string `json:"member_no"`           // 会员标识
	OutOrderNo string `json:"out_order_no"`        // 平台订单编号
	Amount     int    `json:"amount"`              // 冻结金额,单位:分
	FreezeNo   string `json:"freeze_no,omitempty"` // 冻结号,可选
	Remark     string `json:"remark,omitempty"`    // 冻结说明,可选
}

// UnfreezeAccountBalanceResult http://47.110.246.50:6524/docs/qzt/qzt-1che8e73223k6
type UnfreezeAccountBalanceResult struct {
	FreezeNo string `json:"freeze_no"` // 冻结编号
}

// GetAccountHistoryListParams http://47.110.246.50:6524/docs/qzt/qzt-1che8gj3vp37p
type GetAccountHistoryListParams struct {
	MemberNo         string `json:"member_no"`                    // 会员标识
	AccountType      string `json:"account_type"`                 // 账户类型编码,固定值为1000
	SpecialAccountNo string `json:"special_account_no,omitempty"` // 专用账户,可选
	StartTime        string `json:"start_time"`                   // 开始时间,格式为yyyy-MM-dd
	EndTime          string `json:"end_time"`                     // 结束时间,格式为yyyy-MM-dd,时间不能超过一个月
	StartIndex       int    `json:"start_index"`                  // 起始行数,从0开始
	EndIndex         int    `json:"end_index"`                    // 结束行数,最多不能超过20,end_index-start_index不能超过20
	Sort             string `json:"sort,omitempty"`               // 排序,可选值为asc,desc,默认为asc
}

// GetAccountHistoryListResult http://47.110.246.50:6524/docs/qzt/qzt-1che8gj3vp37p
type GetAccountHistoryListResult struct {
	TotalCount   int                    `json:"total_count"`   // 总记录数,当start_index等于0时返回
	CurrentCount int                    `json:"current_count"` // 当前返回记录数
	List         []AccountHistoryDetail `json:"list"`          // 账户变更明细列表,http://47.110.246.50:6524/docs/qzt/qzt-1d50jkreaoign
}

// AccountHistoryDetail 账户变更历史明细 http://47.110.246.50:6524/docs/qzt/qzt-1d50jkreaoign
type AccountHistoryDetail struct {
	AccountNo      string `json:"account_no"`       // 交易流水号
	OrderNo        string `json:"order_no"`         // 订单号
	AccountType    string `json:"account_type"`     // 账户类型编号
	TradeTime      string `json:"trade_time"`       // 变更时间,格式为yyyy-MM-dd HH:mm:ss
	CurrentAmount  int    `json:"cur_amount"`       // 现有金额
	OriginalAmount int    `json:"ori_amount"`       // 原始金额
	ChangeAmount   int    `json:"chg_amount"`       // 变更金额
	LogType        string `json:"log_type"`         // 借贷标记,OUT:转出,IN:转入
	Remark         string `json:"remark"`           // 备注
	SubTradeType   int    `json:"sub_trade_type"`   // 交易子类型,详见 http://47.110.246.50:6524/docs/qzt/qzt-1dcqq6e7k7m3v
	TargetMemberNo string `json:"target_member_no"` // 对方会员ID
}

// QueryAccountFeeUnCollectParams http://47.110.246.50:6524/docs/qzt/qzt-1che8hmupoqmb
type QueryAccountFeeUnCollectParams struct {
	MemberNo string `json:"member_no"` // 会员标识
}

// QueryAccountFeeUnCollectResult http://47.110.246.50:6524/docs/qzt/qzt-1che8hmupoqmb
type QueryAccountFeeUnCollectResult struct {
	UnCollectFee int `json:"uncollect_fee"` // 未收取手续费金额,当start_index等于0时返回
}

// QueryAccountFeeListParams http://47.110.246.50:6524/docs/qzt/qzt-1che8ik1p2jun
type QueryAccountFeeListParams struct {
	MemberNo   string `json:"member_no"`   // 会员标识
	StartTime  string `json:"start_time"`  // 手续费产生开始时间,yyyy-MM-dd
	EndTime    string `json:"end_time"`    // 手续费产生结束时间,yyyy-MM-dd
	StartIndex int    `json:"start_index"` // 起始行数,从0开始
	EndIndex   int    `json:"end_index"`   // 结束行数,最多不能超过20,end_index-start_index不能超20
}

// QueryAccountFeeListResult http://47.110.246.50:6524/docs/qzt/qzt-1che8ik1p2jun
type QueryAccountFeeListResult struct {
	TotalCount   int                `json:"total_count"`   // 总记录数,当start_index等于0时返回
	CurrentCount int                `json:"current_count"` // 当前返回记录数
	List         []AccountFeeDetail `json:"list"`          // 账户手续费记录列表,http://47.110.246.50:6524/docs/qzt/qzt-1d50jpaa7rkb8
}

// AccountFeeDetail 账户手续费明细 http://47.110.246.50:6524/docs/qzt/qzt-1d50jpaa7rkb8
type AccountFeeDetail struct {
	Fee              int    `json:"fee"`                          // 手续费金额
	FeeType          int    `json:"fee_type"`                     // 手续费产生类型,1:收单,4:提现,5:转账,6:分账,8:税筹,9:返利
	Remark           string `json:"remark,omitempty"`             // 备注
	FeeDate          string `json:"fee_date"`                     // 手续费产生日期,yyyy-MM-dd
	Status           int    `json:"status"`                       // 收取状态,1:欠款,2:已收取,4:无需收取
	CollectType      string `json:"collect_type,omitempty"`       // 收取方式,1:自动收取,2:结算收取,3:定时收取
	CollectTime      string `json:"collect_time,omitempty"`       // 收取时间,yyyy-MM-dd HH:mm:ss
	AccountType      string `json:"account_type,omitempty"`       // 收取账户
	SpecialAccountNo string `json:"special_account_no,omitempty"` // 收取专用账户
}
