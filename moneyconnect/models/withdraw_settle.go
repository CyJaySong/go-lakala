package models

// MemberSettleRecord 会员结算记录 http://47.110.246.50:6524/docs/qzt/qzt-1ep13d35nlbk8
type MemberSettleRecord struct {
	MemberNo      string `json:"member_no"`      // 会员号
	SettleStatus  int    `json:"settle_status"`  // 结算状态,1:待处理;2:处理中;3:处理成功;4:处理失败
	SettleAmount  int    `json:"settle_amount"`  // 结算金额,单位:分
	SettleFee     int    `json:"settle_fee"`     // 结算手续费,单位:分
	SettleDate    string `json:"settle_date"`    // 结算日期,格式为yyyyMMdd
	SettleMessage string `json:"settle_message"` // 结算错误信息,在处理失败时有值
}

// SetMemberWithdrawSettleCardParams http://47.110.246.50:6524/docs/qzt/qzt-1ep0llqtitpnu
type SetMemberWithdrawSettleCardParams struct {
	MemberNo string `json:"member_no"` // 会员标识
	CardId   string `json:"card_no"`   // 银行卡编号
}

// SetMemberWithdrawSettleCardResult http://47.110.246.50:6524/docs/qzt/qzt-1ep0llqtitpnu
type SetMemberWithdrawSettleCardResult struct{}

// GetMemberWithdrawSettleRecordParams http://47.110.246.50:6524/docs/qzt/qzt-1ep13d35nlbk8
type GetMemberWithdrawSettleRecordParams struct {
	MemberNo  string `json:"member_no"`  // 会员标识
	StartDate string `json:"start_date"` // 开始日期,格式为yyyy-MM-dd
	EndDate   string `json:"end_date"`   // 结束日期,格式为yyyy-MM-dd,最大不能超过开始日期七天
}

// GetMemberWithdrawSettleRecordResult http://47.110.246.50:6524/docs/qzt/qzt-1ep13d35nlbk8
type GetMemberWithdrawSettleRecordResult struct {
	Record []MemberSettleRecord `json:"record"`
}
