package models

type TransferRes struct {
	MemberNo     string `json:"member_no"`                // 会员编号
	OutRequestNo string `json:"out_request_no,omitempty"` // 商户请求流水
	RequestNo    string `json:"request_no"`               // 钱账通转账通流水号
	Amount       int    `json:"amount"`                   // 金额
	PayStatus    int    `json:"pay_status"`               // 支付状态
}
