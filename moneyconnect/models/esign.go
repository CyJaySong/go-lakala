package models

// EsignInfo http://47.110.246.50:6524/docs/qzt/qzt-1dt3ij0i64btg
type EsignInfo struct {
	A7  string `json:"A7,omitempty"`
	A8  string `json:"A8,omitempty"`
	A10 string `json:"A10,omitempty"`
	A19 string `json:"A19,omitempty"`
	A21 string `json:"A21,omitempty"`
	A22 string `json:"A22,omitempty"`
	A24 string `json:"A24,omitempty"`
	A25 string `json:"A25,omitempty"`
	A27 string `json:"A27,omitempty"`
	A28 string `json:"A28,omitempty"`
	A30 string `json:"A30,omitempty"`
	A31 string `json:"A31,omitempty"`
	A32 string `json:"A32,omitempty"`
	A33 string `json:"A33,omitempty"`
	A35 string `json:"A35,omitempty"`
	A37 string `json:"A37,omitempty"`
	A39 string `json:"A39,omitempty"`
	A41 string `json:"A41,omitempty"`
	A42 string `json:"A42,omitempty"`
	A44 string `json:"A44,omitempty"`
	A58 string `json:"A58,omitempty"`
	B12 string `json:"B12,omitempty"`
	B14 string `json:"B14,omitempty"`
	B15 string `json:"B15,omitempty"`
	B16 string `json:"B16,omitempty"`
	C1  string `json:"C1,omitempty"`
	C2  string `json:"C2,omitempty"`
	C3  string `json:"C3,omitempty"`
	E1  string `json:"E1,omitempty"`
	E2  string `json:"E2,omitempty"`
	E3  string `json:"E3,omitempty"`
	E4  string `json:"E4,omitempty"`
	E5  string `json:"E5,omitempty"`
	E6  string `json:"E6,omitempty"`
	E7  string `json:"E7,omitempty"`
	E8  string `json:"E8,omitempty"`
	E9  string `json:"E9,omitempty"`
	E10 string `json:"E10,omitempty"`
	E11 string `json:"E11,omitempty"`
	E12 string `json:"E12,omitempty"`
	E13 string `json:"E13,omitempty"`
	E14 string `json:"E14,omitempty"`
	E15 string `json:"E15,omitempty"`
}

// EsignSignInfo http://47.110.246.50:6524/docs/qzt/qzt-1d50joj8scim1
type EsignSignInfo struct {
	EsignUserType int    `json:"esign_user_type"` // 合同用户类型,1:甲方,2:乙方
	MemberNo      string `json:"member_no"`       // 钱账通会员
	SignLocation  string `json:"sign_location"`   // 签名域标签值,签署人可以签多个,以";"分割
}

// GetMemberEsignCreateH5UrlParams http://47.110.246.50:6524/docs/qzt/qzt-1d4ranfbg4k6j
type GetMemberEsignCreateH5UrlParams struct {
	OutMemberNo      string `json:"out_member_no"`       // 必填,外部会员标识
	EsignProjectCode string `json:"esign_project_code"`  // 必填,电子合同签约授权号,如果需要电子签约则必传
	FrontUrl         string `json:"front_url,omitempty"` // 前台跳转地址
}

// GetMemberEsignCreateH5UrlResult http://47.110.246.50:6524/docs/qzt/qzt-1d4ranfbg4k6j
type GetMemberEsignCreateH5UrlResult struct {
	Url string `json:"url"` // 开户页面的URL
}

// CreateMemberEsignParams http://47.110.246.50:6524/docs/qzt/qzt-1chlfcis1hkef
type CreateMemberEsignParams struct {
	EsignType       int             `json:"esign_type"`        // 电子签约类型,1:特约商户支付服务合作协议三合一,2:机构业务协议
	EsignTemplateNo string          `json:"esign_template_no"` // 电子合同签约的合同模板编号
	EsignInfo       EsignInfo       `json:"esign_info"`        // 电子合同签约的合同内容,http://47.110.246.50:6524/docs/qzt/qzt-1dt3ij0i64btg
	EsignSignInfo   []EsignSignInfo `json:"esign_sign_info"`   // 电子合同签约的签名域内容(公章的区域),http://47.110.246.50:6524/docs/qzt/qzt-1d50joj8scim1
	BackUrl         string          `json:"back_url"`          // 后台通知地址
}

// CreateMemberEsignResult http://47.110.246.50:6524/docs/qzt/qzt-1chlfcis1hkef
type CreateMemberEsignResult struct {
	ContractNo       string `json:"contract_no"`                  // 合同编号
	EsignProjectCode string `json:"esign_project_code,omitempty"` // 如果电子合同签约，则必选
}

// QueryMemberEsignParams
// http://47.110.246.50:6524/docs/qzt/qzt-1cj4ekl02kho2
// http://47.110.246.50:6524/docs/qzt/qzt-1d4rkl8qckpe7
type QueryMemberEsignParams struct {
	EsignProjectCode string `json:"esign_project_code,omitempty"` // 电子合同签约授权号
	ContractNo       string `json:"contract_no,omitempty"`        // 合同编号
}

// QueryMemberEsignByProjectCodeResult http://47.110.246.50:6524/docs/qzt/qzt-1cj4ekl02kho2
type QueryMemberEsignByProjectCodeResult struct {
	ContractInfo []ContractInfo `json:"contract_info"` // 合同信息
}

// QueryMemberEsignByContractNoResult http://47.110.246.50:6524/docs/qzt/qzt-1d4rkl8qckpe7
type QueryMemberEsignByContractNoResult = ContractInfo

// http://47.110.246.50:6524/docs/qzt/qzt-1cj4ekl02kho2
// http://47.110.246.50:6524/docs/qzt/qzt-1d4rkl8qckpe7
// ContractInfo 合同信息的结构体
type ContractInfo struct {
	EsignStatus        int    `json:"esign_status"`                    // 合同状态,0:签署中,1:生效,2:失效
	EsignType          int    `json:"esign_type"`                      // 电子签约类型,1:特约商户支付服务合作协议,2:特约商户信息授权,3:商户结算授权委托书
	EsignProjectCode   string `json:"esign_project_code"`              // 电子合同签约授权号
	PbMemberSignStatus int    `json:"pb_member_sign_status"`           // 会员签署状态(乙方会员),0:未签署,1:已签署
	PaMemberSignStatus int    `json:"pa_member_sign_status,omitempty"` // 会员签署状态(甲方会员),合同类型中有此签署方时返回,0:未签署,1:已签署
	LklSignStatus      int    `json:"lkl_sign_status,omitempty"`       // 拉卡拉签署状态,合同类型中有此签署方时返回,0:未签署,1:已签署
	FileNo             int    `json:"file_no,omitempty"`               // 文件编号,生成后产生
	ContractNo         string `json:"contract_no"`                     // 合同编号
}
