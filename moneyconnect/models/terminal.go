package models

// GetMemberMerchantTerminalH5UrlParams http://47.110.246.50:6524/docs/qzt/qzt-1fbefr8ja7eup
type GetMemberMerchantTerminalH5UrlParams struct {
	MemberNo string `json:"member_no"` // 会员标识
	BackUrl  string `json:"back_url"`  // 后台通知地址
}

// GetMemberMerchantTerminalH5UrlResult http://47.110.246.50:6524/docs/qzt/qzt-1ep0llqtitpnu
type GetMemberMerchantTerminalH5UrlResult struct {
	Url string `json:"url"` // 在线地址
}
