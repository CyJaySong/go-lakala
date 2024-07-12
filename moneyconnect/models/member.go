package models

type MemberInfo struct {
	MemberStatus               int    `json:"member_status"`                            // 会员状态,1:有效,10:锁定,20:关闭,99:待审核
	RealStatus                 int    `json:"real_status"`                              // 实名状态,0:未实名,1:已实名
	MemberCategory             int    `json:"member_category"`                          // 会员类别,0:普通,1:商户
	MobileBind                 bool   `json:"mobile_bind"`                              // 手机号是否绑定
	MemberType                 int    `json:"member_type,omitempty"`                    // 会员类型,登记后不允许修改,3:企业会员,2:个人会员,4:个人工商户
	NickName                   string `json:"nick_name,omitempty"`                      // 会员昵称
	Name                       string `json:"name,omitempty"`                           // 会员名称/企业名称,实名后不可修改
	IdentityType               int    `json:"identity_type,omitempty"`                  // 个人或法人证件类型,实名后不可修改,1:身份证,2:护照,3:港澳通行证,4:台胞证
	IdentityName               string `json:"identity_name,omitempty"`                  // 个人或法人证件姓名,实名后不可修改
	IdentityNo                 string `json:"identity_no,omitempty"`                    // 个人或法人证件号码,RSA加密,实名后不可修改
	IdentityNoStartDate        string `json:"identity_no_start_date,omitempty"`         // 身份证生效日期,实名后不可修改,yyyy-MM-dd
	IdentityNoValidDate        string `json:"identity_no_valid_date,omitempty"`         // 身份证失效日期,实名后不可修改,yyyy-MM-dd
	IdentityFrontImgNo         string `json:"identity_front_img_no,omitempty"`          // 证件正面照图片ID,实名后不可修改
	IdentityBackImgNo          string `json:"identity_back_img_no,omitempty"`           // 证件反面照图片ID,实名后不可修改
	Address                    string `json:"address,omitempty"`                        // 地址,企业信息,实名后不可修改
	BusinessLicense            string `json:"business_license,omitempty"`               // 营业执照编号,实名后不可修改
	BusinessLicenseStartDate   string `json:"business_license_start_date,omitempty"`    // 营业执照生效日期,yyyy-MM-dd
	BusinessLicenseValidDate   string `json:"business_license_valid_date,omitempty"`    // 营业执照失效日期,yyyy-MM-dd
	BusinessLicenseImgNo       string `json:"business_license_img_no,omitempty"`        // 营业执照图片ID,实名后不可修改
	AccountOpeningLicenseImgNo string `json:"account_opening_license_img_no,omitempty"` // 开户许可证图片ID,实名后不可修改
	LinkmanName                string `json:"linkman_name,omitempty"`                   // 企业联系人
	LinkmanPhone               string `json:"linkman_phone,omitempty"`                  // 企业联系人手机号码
	LinkmanIdentityNo          string `json:"linkman_identity_no,omitempty"`            // 企业联系人身份证ID,RSA加密
	LinkmanFrontImgNo          string `json:"linkman_front_img_no,omitempty"`           // 企业联系人身份证正面照ID
	LinkmanBackImgNo           string `json:"linkman_back_img_no,omitempty"`            // 企业联系人身份证反面照ID
	ProvinceCode               string `json:"province_code,omitempty"`                  // 所在地省编号
	CityCode                   string `json:"city_code,omitempty"`                      // 所在地市编号
	AreaCode                   string `json:"area_code,omitempty"`                      // 所在地区编号
	AgreementImgNo             string `json:"agreement_img_no,omitempty"`               // 平台协议文件ID
	FacadeNo                   string `json:"facade_no,omitempty"`                      // 门头照片ID
	Mobile                     string `json:"mobile,omitempty"`                         // 绑定手机号码
}

// GetMemberOpenPartH5UrlParams http://47.110.246.50:6524/docs/qzt/qzt-1dt3fhr2456as
type GetMemberOpenPartH5UrlParams struct {
	OutMemberNo                string     `json:"out_member_no"`                            // 必填,外部会员标识,唯一
	OutMerchantNo              string     `json:"out_merchant_no,omitempty"`                // 选填,外部会员标识,收单商户角色必填
	OpenPart                   int        `json:"open_part"`                                // 必填,开通角色,1:普通会员;2:收单商户
	BackUrl                    string     `json:"back_url"`                                 // 必填,后台通知地址
	FrontUrl                   string     `json:"front_url,omitempty"`                      // 选填,前台跳转地址
	MemberCategory             int        `json:"member_category"`                          // 必填,会员类别,固定值:0
	MemberRole                 string     `json:"member_role"`                              // 必填,会员角色,由运营人员在管理后台创建，生产投产时请向运营人员索取, 拉卡拉方面告知生产环境固定R001
	MemberType                 int        `json:"member_type"`                              // 必填,会员类型,2:个人会员,3:企业会员,4:个人工商户,http://47.110.246.50:6524/docs/qzt/qzt-1d50iu3jl96p8
	Phone                      string     `json:"phone,omitempty"`                          // 选填,电话号码
	Name                       string     `json:"name,omitempty"`                           // 选填,会员名称/企业名称,会员类型如果是企业或者个体工商户则必填
	IdentityType               int        `json:"identity_type,omitempty"`                  // 选填,证件类型,1:身份证,2:护照,3:港澳通行证,4:台胞证
	IdentityName               string     `json:"identity_name,omitempty"`                  // 选填,证件姓名,个人姓名或法人姓名
	IdentityNo                 string     `json:"identity_no,omitempty"`                    // 选填,证件号码,个人或法人证件号,RSA加密
	BusinessLicense            string     `json:"business_license,omitempty"`               // 选填,统一社会信用代码
	BankProvinceCode           string     `json:"bank_province_code,omitempty"`             // 选填,银行卡开户行-省编号,http://47.110.246.50:6524/attach_files/qzt/278
	BankCityCode               string     `json:"bank_city_code,omitempty"`                 // 选填,银行卡开户行-市编号,http://47.110.246.50:6524/attach_files/qzt/278
	BankAreaCode               string     `json:"bank_area_code,omitempty"`                 // 选填,银行卡开户行-区/县编码,见商户对接资料,http://47.110.246.50:6524/attach_files/qzt/278
	BankCode                   string     `json:"bank_code,omitempty"`                      // 选填,银行编码,http://47.110.246.50:6524/attach_files/qzt/277
	BankCardNo                 string     `json:"bank_card_no,omitempty"`                   // 选填,银行卡号,RSA加密
	MerchantName               string     `json:"merchant_name,omitempty"`                  // 选填,商户名称,4-20个字
	MccCode                    string     `json:"mcc_code,omitempty"`                       // 选填,行业代码,见商户对接资料
	MerchantProvinceCode       string     `json:"merchant_province_code,omitempty"`         // 选填,所在地省编号,http://47.110.246.50:6524/attach_files/qzt/278
	MerchantCityCode           string     `json:"merchant_city_code,omitempty"`             // 选填,所在地市编号,http://47.110.246.50:6524/attach_files/qzt/278
	MerchantAreaCode           string     `json:"merchant_area_code,omitempty"`             // 选填,所在地区编号,http://47.110.246.50:6524/attach_files/qzt/278
	ManageAddress              string     `json:"manage_address,omitempty"`                 // 选填,经营地址,大于等于3个字
	BizContent                 string     `json:"biz_content,omitempty"`                    // 选填,商户经营内容,见商户对接资料
	SellerScale                float64    `json:"seller_scale,omitempty"`                   // 选填,商家分账比例(商家获得比列),20%传0.2,最长4位小数
	Condition                  string     `json:"condition,omitempty"`                      // 选填,条件,ge:大于等于(默认);eq:等于
	IdentityNoStartDate        string     `json:"identity_no_start_date,omitempty"`         // 选填,身份证生效日期,yyyy-MM-dd
	IdentityNoValidDate        string     `json:"identity_no_valid_date,omitempty"`         // 选填,身份证失效日期,yyyy-MM-dd,长期:9999-12-31
	BusinessLicenseStartDate   string     `json:"business_license_start_date,omitempty"`    // 选填,营业执照生效日期,yyyy-MM-dd
	BusinessLicenseValidDate   string     `json:"business_license_valid_date,omitempty"`    // 选填,营业执照失效日期,yyyy-MM-dd,长期:9999-12-31
	IdentityFrontImgNo         string     `json:"identity_front_img_no,omitempty"`          // 选填,证件正面照图片ID,请先调用图片上传接口上传
	IdentityBackImgNo          string     `json:"identity_back_img_no,omitempty"`           // 选填,证件反面照图片ID,请先调用图片上传接口上传
	BusinessLicenseImgNo       string     `json:"business_license_img_no,omitempty"`        // 选填,营业执照图片ID,请先调用图片上传接口上传
	AccountOpeningLicenseImgNo string     `json:"account_opening_license_img_no,omitempty"` // 选填,开户许可证或银行卡图片ID,请先调用图片上传接口上传
	FacadeNo                   string     `json:"facade_no,omitempty"`                      // 选填,门头照片ID,请先调用图片上传接口上传
	MerchantNature             int        `json:"merchant_nature,omitempty"`                // 选填,商户性质,1:国营,2:股份制,3:集体,4:中外合资,5:外商独资,6:私营合伙,7:私营独资,8:个体,9:事业单位,10:党政及国家机关,11:社会团体,12:其他,13:小微商户
	FeeRate                    *FeeRate   `json:"fee_rate,omitempty"`                       // 选填,手续费费率,http://47.110.246.50:6524/docs/qzt/qzt-1d50jmfo7lire
	EsignInfo                  *EsignInfo `json:"esign_info,omitempty"`                     // 选填,收单商户必填,http://47.110.246.50:6524/docs/qzt/qzt-1dt3ij0i64btg
}

// GetMemberOpenPartH5UrlResult http://47.110.246.50:6524/docs/qzt/qzt-1dt3fhr2456as
type GetMemberOpenPartH5UrlResult struct {
	Url string `json:"url"` // 开通商户页面的URL
}

// ChangeMemberInfoParams http://47.110.246.50:6524/docs/qzt/qzt-1dldqnlkitr8j
type ChangeMemberInfoParams struct {
	MemberNo                 string `json:"member_no"`                              // 应用平台会员标识
	IdentityNoStartDate      string `json:"identity_no_start_date,omitempty"`       // 选填,证件生效日期,yyyy-MM-dd,长期:9999-12-31
	IdentityNoValidDate      string `json:"identity_no_valid_date,omitempty"`       // 选填,证件失效日期,yyyy-MM-dd,长期:9999-12-31
	IdentityFrontImgNo       string `json:"identity_front_img_no,omitempty"`        // 选填,证件正面照图片ID,请先调用图片上传接口上传
	IdentityBackImgNo        string `json:"identity_back_img_no,omitempty"`         // 选填,证件反面照图片ID,请先调用图片上传接口上传
	ProvinceCode             string `json:"province_code,omitempty"`                // 选填,所在地省编号,http://47.110.246.50:6524/attach_files/qzt/278
	CityCode                 string `json:"city_code,omitempty"`                    // 选填,所在地市编号,http://47.110.246.50:6524/attach_files/qzt/278
	AreaCode                 string `json:"area_code,omitempty"`                    // 选填,所在地区编号,http://47.110.246.50:6524/attach_files/qzt/278
	Address                  string `json:"address,omitempty"`                      // 选填,联系地址/经营地址
	CompanyName              string `json:"company_name,omitempty"`                 // 选填,企业名称
	BusinessLicense          string `json:"business_license,omitempty"`             // 选填,营业执照编号
	BusinessLicenseStartDate string `json:"business_license_start_date,omitempty"`  // 选填,营业执照生效日期,yyyy-MM-dd
	BusinessLicenseValidDate string `json:"business_license_valid_date,omitempty"`  // 选填,营业执照失效日期,yyyy-MM-dd,长期:9999-12-31
	BusinessLicenseImgNo     string `json:"business_license_img_no,omitempty"`      // 选填,营业执照图片ID,请先调用图片上传接口上传
	LegalName                string `json:"legal_name,omitempty"`                   // 选填,个人或法人证件姓名
	LegalIdentityNo          string `json:"legal_identity_no,omitempty"`            // 选填,个人或法人证件号码,RSA加密
	LegalIdentityType        int    `json:"legal_identity_type,omitempty"`          // 选填,个人或法人证件类型,1:身份证,2:护照,3:港澳通行证,4:台胞证
	LegalIdentityNoStartDate string `json:"legal_identity_no_start_date,omitempty"` // 选填,个人或法人证件生效日期,yyyy-MM-dd,长期:9999-12-31
	LegalIdentityNoValidDate string `json:"legal_identity_no_valid_date,omitempty"` // 选填,个人或法人证件失效日期,yyyy-MM-dd,长期:9999-12-31
	LegalIdentityFrontImgNo  string `json:"legal_identity_front_img_no,omitempty"`  // 选填,个人或法人证件正面照图片ID,请先调用图片上传接口上传
	LegalIdentityBackImgNo   string `json:"legal_identity_back_img_no,omitempty"`   // 选填,个人或法人证件反面照图片ID,请先调用图片上传接口上传
	BizContent               string `json:"biz_content,omitempty"`                  // 选填,经营内容
	MccCode                  string `json:"mcc_code,omitempty"`                     // 选填,行业代码
}

// ChangeMemberInfoResult http://47.110.246.50:6524/docs/qzt/qzt-1dldqnlkitr8j
type ChangeMemberInfoResult struct{}

// QueryMemberInfoParams http://47.110.246.50:6524/docs/qzt/qzt-1d4ranrhtk1r4
type QueryMemberInfoParams struct {
	MemberNo    string `json:"member_no,omitempty"`     // 会员标识,二选一
	OutMemberNo string `json:"out_member_no,omitempty"` // 外部会员号,二选一
}

// QueryMemberInfoResult http://47.110.246.50:6524/docs/qzt/qzt-1d4ranrhtk1r4
type QueryMemberInfoResult = MemberInfo

// ApplyMemberMobileModifyParams http://47.110.246.50:6524/docs/qzt/qzt-1dg1fpra22vv8
type ApplyMemberMobileModifyParams struct {
	MemberNo string `json:"member_no"` // 会员标识
	Mobile   string `json:"mobile"`    // 手机号
}

// ApplyMemberMobileModifyResult http://47.110.246.50:6524/docs/qzt/qzt-1dg1fpra22vv8
type ApplyMemberMobileModifyResult struct {
	MemberNo string `json:"member_no"` // 会员号
}

// ConfirmMemberMobileModifyParams http://47.110.246.50:6524/docs/qzt/qzt-1dg1fqe4ne8bk
type ConfirmMemberMobileModifyParams struct {
	MemberNo string `json:"member_no"` // 会员标识
	Code     string `json:"code"`      // 验证码
}

// ConfirmMemberMobileModifyResult http://47.110.246.50:6524/docs/qzt/qzt-1dg1fqe4ne8bk
type ConfirmMemberMobileModifyResult struct {
	MemberNo string `json:"member_no"` // 会员号
}

// QueryMemberChannelRealParams http://47.110.246.50:6524/docs/qzt/qzt-1em3jgq8h8pqc
type QueryMemberChannelRealParams struct {
	MemberNo   string `json:"member_no"`            // 会员标识
	MerchantNo string `json:"merchant_no"`          // 会员商户标识
	RealType   string `json:"real_type"`            // 实名类型;WXZF:代表微信,ZFBZF:代表支付宝
	SubMchID   string `json:"sub_mch_id"`           // 子商户号
	ChannelID  string `json:"channel_id,omitempty"` // 渠道ID,可选
}

// QueryMemberChannelRealResult http://47.110.246.50:6524/docs/qzt/qzt-1em3jgq8h8pqc
type QueryMemberChannelRealResult struct {
	ApplymentID     string `json:"applymentId"`               // 申请编号
	ApplymentState  string `json:"applymentState"`            // 申请状态
	AuthorizeState  string `json:"authorizeState"`            // 认证状态
	RegisterChannel string `json:"registerChannel"`           // 报备通道
	RealNameType    string `json:"realNameType"`              // 实名认证类型;WXZF:代表微信,ZFBZF:代表支付宝
	QrcodeData      string `json:"qrcodeData,omitempty"`      // 小程序码图片;当申请单状态为CONTACT_CONFIRM、LEGAL_CONFIRM、AUDIT_PASS、AUDIT_FREEZE时会返回此链接。
	RejectParameter string `json:"rejectParameter,omitempty"` // 驳回参数
	RejectReason    string `json:"rejectReason,omitempty"`    // 驳回原因
}
