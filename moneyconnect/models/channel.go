package models

// GetMemberOtherChannelSubMerchantParams http://47.110.246.50:6524/docs/qzt/qzt-1d50jau3leaor
type GetMemberOtherChannelSubMerchantParams struct {
	MemberNo          string `json:"member_no"`           // 钱账通会员标识
	MerchantNo        string `json:"merchant_no"`         // 钱账通商户号
	ChannelNo         string `json:"channel_no"`          // 渠道号
	ThirdPartyPayment string `json:"third_party_payment"` // 三方渠道,WECHAT:微信,ALIPAY:支付宝,UNION:银联
	Channel           string `json:"channel"`             // 报备渠道,UNIONPAY:银联,NETUNION:网联,UNIONPAY_NEIMENG:内蒙银联(默认),ALIPAY_FLOWER:支付宝健康分
}

// GetMemberOtherChannelSubMerchantResult http://47.110.246.50:6524/docs/qzt/qzt-1d50jau3leaor
type GetMemberOtherChannelSubMerchantResult struct {
	ChannelSubMerchantNo string `json:"channel_sub_merchant_no,omitempty"` // 渠道子商户号
}

// GetMemberTransferPayAccountParams http://47.110.246.50:6524/docs/qzt/qzt-1d50jblfquu2i
type GetMemberTransferPayAccountParams struct {
	MemberNo         string `json:"member_no"`                    // 会员标识
	SpecialAccountNo string `json:"special_account_no,omitempty"` // 账户标识,如退货专用账户:S010
}

// GetMemberTransferPayAccountResult http://47.110.246.50:6524/docs/qzt/qzt-1d50jblfquu2i
type GetMemberTransferPayAccountResult struct {
	ReceiveAccountNo   string `json:"receive_account_no"`   // 收款账号
	ReceiveAccountName string `json:"receive_account_name"` // 收款户名
	ReceiveBankName    string `json:"receive_bank_name"`    // 收款银行
}

// GetUnionPayUserIdParams http://47.110.246.50:6524/docs/qzt/qzt-1ekbcuoggq19p
type GetUnionPayUserIdParams struct {
	MemberMerchantId  string `json:"member_merchant_id"`  // 钱账通商户编号,必选
	ThirdPartyPayment string `json:"third_party_payment"` // 三方渠道,如"UNIONPAY",必选
	UserAuthCode      string `json:"user_auth_code"`      // 用户授权码,必选
	AppUpIdentifier   string `json:"app_up_identifier"`   // 支付标识,可选
}

// GetUnionPayUserIdResult http://47.110.246.50:6524/docs/qzt/qzt-1ekbcuoggq19p
type GetUnionPayUserIdResult struct {
	OpenId string `json:"openid"` // openid,必选
}
