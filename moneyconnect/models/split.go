package models

// OrderConsumeCompleteParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcjji520jp
type OrderConsumeCompleteParams struct {
	MemberNo      string        `json:"member_no"`       // 会员标识,会员消费订单必填
	OrderNo       string        `json:"order_no"`        // 订单编号
	OutRequestNo  string        `json:"out_request_no"`  // 外部请求号,接口调用方上送,唯一
	SplitAmount   int           `json:"split_amount"`    // 分账金额
	SplitRuleData SplitRuleData `json:"split_rule_data"` // 分账规则,分账规则为空时默认全给商户,分账规则中分账方必须实名并且设置会员角色,分账记录总金额=订单金额
}

// OrderConsumeCompleteResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcjji520jp
type OrderConsumeCompleteResult struct {
	SplitRuleResult []SplitRuleResultItem `json:"split_rule_result"` // 分账会员列表
	OrderStatus     int                   `json:"order_status"`      // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
	SplitSeqNo      string                `json:"split_seq_no"`      // 分账流水号
}

// MerchantCommissionBalanceSplitParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcqhkp9v4c
type MerchantCommissionBalanceSplitParams struct {
	MemberNo      string        `json:"member_no"`            // 会员标识
	AccountType   string        `json:"account_type"`         // 账户类型编码,固定值:1000
	OutOrderNo    string        `json:"out_order_no"`         // 外部请求号
	Amount        int           `json:"amount"`               // 分账金额
	SplitRuleData SplitRuleData `json:"split_rule_data"`      // 分账规则,JSON格式
	OrderDesc     string        `json:"order_desc,omitempty"` // 订单描述,最多300个字符,可选
}

// MerchantCommissionBalanceSplitResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcqhkp9v4c
type MerchantCommissionBalanceSplitResult struct {
	OrderNo         string                `json:"order_no"`          // 订单编号
	SplitSeqNo      string                `json:"split_seq_no"`      // 分账流水号
	SplitRuleResult []SplitRuleResultItem `json:"split_rule_result"` // 分账会员列表
}

// RevokeSplitParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcr9mjaici
type RevokeSplitParams struct {
	OrderNo    string `json:"order_no"`     // 订单编号
	SplitSeqNo string `json:"split_seq_no"` // 分账流水号
	Amount     int    `json:"amount"`       // 金额(分),该金额和分账流水号对应的分账金额一致
}

// RevokeSplitResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcr9mjaici
type RevokeSplitResult struct {
	OrderNo    string `json:"order_no"`     // 订单编号
	SplitSeqNo string `json:"split_seq_no"` // 分账流水号
}
