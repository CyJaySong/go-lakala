package models

type FeeRate struct {
	Wechat            string `json:"wechat,omitempty"`               // 微信手续费费率,大于0.001，小于100
	Alipay            string `json:"alipay,omitempty"`               // 支付宝手续费费率,大于0.001，小于100
	UnionScanDebit    string `json:"union_scan_debit,omitempty"`     // 专业扫码-借记卡费率,大于0.001，小于100专
	UnionScanCredit   string `json:"union_scan_credit,omitempty"`    // 专业扫码-贷记卡费率,大于0.001，小于100
	UnionScanDebitPct string `json:"union_scan_debit_pct,omitempty"` // 专业扫码-借记卡封顶,元
}
