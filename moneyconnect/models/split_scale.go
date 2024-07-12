package models

// SplitScale 分账比例信息结构体 http://47.110.246.50:6524/docs/qzt/qzt-1d50jq19fkuop
type SplitScale struct {
	SeqNo          string  `json:"seq_no"`                     // 申请流水
	SellerScale    float64 `json:"seller_scale"`               // 商家分账比例(商家获得比列),20%传0.2,最长4位小数
	AgreementImgNo string  `json:"agreement_img_no,omitempty"` // 机构协议文件ID,请先调用图片上传接口上传
	Condition      string  `json:"condition"`                  // 条件,ge:大于等于(默认),eq:等于
	Status         int     `json:"status"`                     // 状态,1:创建,2:待初审,3:待复核,4:待进件,5:进件中,11:有效,10:无效
	CheckMessage   string  `json:"check_message,omitempty"`    // 初审信息
	RecheckMessage string  `json:"recheck_message,omitempty"`  // 复审信息
}

// ApplyMemberSplitScaleParams http://47.110.246.50:6524/docs/qzt/qzt-1d4rape2qm922
type ApplyMemberSplitScaleParams struct {
	MemberNo       string  `json:"member_no"`                  // 会员标识
	SellerScale    float64 `json:"seller_scale"`               // 商家分账比例(商家获得比列),20%传0.2,最长4位小数
	AgreementImgNo string  `json:"agreement_img_no,omitempty"` // 机构协议文件ID,请先调用图片上传接口上传
	Condition      string  `json:"condition"`                  // 条件,ge:大于等于(默认),eq:等于
	BackUrl        string  `json:"back_url"`                   // 后台通知地址
	Remark         string  `json:"remark,omitempty"`           // 备注
}

// ApplyMemberSplitScaleResult http://47.110.246.50:6524/docs/qzt/qzt-1d4rape2qm922
type ApplyMemberSplitScaleResult struct {
	MemberNo string `json:"member_no"` // 会员号
	SeqNo    string `json:"seq_no"`    // 申请流水
}

// ConfirmMemberSplitScaleParams http://47.110.246.50:6524/docs/qzt/qzt-1d4rapq54fh5i
type ConfirmMemberSplitScaleParams struct {
	MemberNo string `json:"member_no"` // 会员号
	SeqNo    string `json:"seq_no"`    // 申请流水
	Code     string `json:"code"`      // 验证码
}

// ConfirmMemberSplitScaleResult http://47.110.246.50:6524/docs/qzt/qzt-1d4rapq54fh5i
type ConfirmMemberSplitScaleResult struct {
	MemberNo string `json:"member_no"` // 会员号
}

// QueryMemberSplitScaleParams http://47.110.246.50:6524/docs/qzt/qzt-1d4raq703p3or
type QueryMemberSplitScaleParams struct {
	MemberNo string `json:"member_no"`        // 会员标识
	SeqNo    string `json:"seq_no,omitempty"` // 申请流水,可选
}

// QueryMemberSplitScaleResult http://47.110.246.50:6524/docs/qzt/qzt-1d4raq703p3or
type QueryMemberSplitScaleResult struct {
	SplitScaleList []SplitScale `json:"split_scale_list"`
}
