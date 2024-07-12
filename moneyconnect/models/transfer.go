package models

// TransferParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcohq06chg
type TransferParams struct {
	FromMemberNo         string         `json:"from_member_no"`          // 会员标识,存管系统指定会员号
	ToMemberNo           string         `json:"to_member_no"`            // 转账目标会员号
	OutOrderNo           string         `json:"out_order_no"`            // 应用平台订单号
	AccountType          string         `json:"account_type"`            // 账户类型编码，固定值:1000
	FromSpecialAccountNo string         `json:"from_special_account_no"` // 转出账户专用账户编号,营销专用账户编号:S004
	ToSpecialAccountNo   string         `json:"to_special_account_no"`   // 转入账户专用账户编号
	Amount               int            `json:"amount"`                  // 转账金额,单位:分
	OrderName            string         `json:"order_name"`              // 订单名称(通常是营销补贴活动名称)
	Remark               string         `json:"remark"`                  // 备注
	Exts                 map[string]any `json:"exts,omitempty"`          // 扩展信息(备注补贴的订单号，便于风控查单)
}

// TransferResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcohq06chg
type TransferResult struct {
	OrderNo      string `json:"order_no"`       // 存管系统订单号
	OutOrderNo   string `json:"out_order_no"`   // 应用平台订单号
	Amount       int    `json:"amount"`         // 金额
	FromMemberNo string `json:"from_member_no"` // 会员号
	ToMemberNo   string `json:"to_member_no"`   // 转账目标会员号
	EndTime      string `json:"end_time"`       // 交易时间,格式:YYYY-MM-DD HH:MI:SS
}

// RevokeTransferParams http://47.110.246.50:6524/docs/qzt/qzt-1f8f0i2r8bkt2
type RevokeTransferParams struct {
	OrderNo      string         `json:"order_no"`         // 订单编号
	OutRequestNo string         `json:"out_request_no"`   // 外部请求号
	Remark       string         `json:"remark,omitempty"` // 备注,可选
	Exts         map[string]any `json:"exts,omitempty"`   // 扩展信息,JSON,可选
}

// RevokeTransferResult http://47.110.246.50:6524/docs/qzt/qzt-1f8f0i2r8bkt2
type RevokeTransferResult struct {
	OrderNo     string `json:"order_no"`     // 订单编号
	OrderStatus int    `json:"order_status"` // 订单状态,http://47.110.246.50:6524/docs/qzt/qzt-1ekbdvs8bn5qi
}
