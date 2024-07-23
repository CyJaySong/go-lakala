package models

// BankCardInfo 单个银行卡信息的结构体 http://47.110.246.50:6524/docs/qzt/qzt-1d50jk2k3qsoi
type BankCardInfo struct {
	CardId          string `json:"card_id"`           // 银行卡标识
	CardNo          string `json:"card_no"`           // 银行卡号,带掩码
	IsSettleCard    bool   `json:"is_settle_card"`    // 是否是结算卡,true/false
	BankCode        string `json:"bank_code"`         // 银行代码,http://47.110.246.50:6524/attach_files/qzt/277
	BankCodeEn      string `json:"bank_code_en"`      // 银行简称
	BankName        string `json:"bank_name"`         // 银行名称
	BindTime        string `json:"bind_time"`         // 绑定时间,格式为yyyy-MM-dd HH:mm:ss
	CardType        int    `json:"card_type"`         // 银行卡类型,1:储蓄卡,2:信用卡
	BankAccountProp int    `json:"bank_account_prop"` // 银行卡账户属性,0:对私,1:对公
	BankProvince    string `json:"bank_province"`     // 银行卡省编号,http://47.110.246.50:6524/attach_files/qzt/278
	BankCity        string `json:"bank_city"`         // 银行卡市编号,http://47.110.246.50:6524/attach_files/qzt/278
	BankContactNo   string `json:"bank_contact_no"`   // 开户支行编号
	BankContactName string `json:"bank_contact_name"` // 开户支行名称
	Mobile          string `json:"mobile"`            // 银行预留手机
	BindChannelCode string `json:"bind_channel_code"` // 绑卡通道，如BMCP/BOSS
}

// GetMemberBindBankCardH5UrlParams http://47.110.246.50:6524/docs/qzt/qzt-1cj3r5ovfo3ek
type GetMemberBindBankCardH5UrlParams struct {
	MemberNo string `json:"member_no"` // 会员标识
	BackUrl  string `json:"back_url"`  // 后台通知地址
}

// GetMemberBindBankCardH5UrlResult http://47.110.246.50:6524/docs/qzt/qzt-1cj3r5ovfo3ek
type GetMemberBindBankCardH5UrlResult struct {
	Url string `json:"url"`
}

// QueryMemberBindBankCardParams http://47.110.246.50:6524/docs/qzt/qzt-1d4rap36c09s0
type QueryMemberBindBankCardParams struct {
	MemberNo string `json:"member_no"` // 会员标识
}

// QueryMemberBindBankCardResult http://47.110.246.50:6524/docs/qzt/qzt-1d4rap36c09s0
type QueryMemberBindBankCardResult struct {
	BankCardList []BankCardInfo `json:"bank_card_list"`
}

// UpdateMemberBankCardContactParams http://47.110.246.50:6524/docs/qzt/qzt-1d4raqih9j0lc
type UpdateMemberBankCardContactParams struct {
	MemberNo        string `json:"member_no"`         // 会员标识
	CardId          string `json:"card_no"`           // 银行卡编号
	BankContactName string `json:"bank_contact_name"` // 联行号名称
	BankContactNo   string `json:"bank_contact_no"`   // 联行号
	BankProvince    string `json:"bank_province"`     // 支行所在省
	BankCity        string `json:"bank_city"`         // 支行所在市
}

// UpdateMemberBankCardContactResult http://47.110.246.50:6524/docs/qzt/qzt-1d4raqih9j0lc
type UpdateMemberBankCardContactResult struct {
	MemberNo string `json:"member_no"` // 会员号
	CardId   string `json:"card_no"`   // 银行卡编号
}
