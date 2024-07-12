package models

type PayInfo struct {
	PrepayId    string `json:"prepay_id"`    // 支付宝
	Package     string `json:"package"`      // 微信
	Timestamp   string `json:"timestamp"`    // 微信
	PaySign     string `json:"pay_sign"`     // 微信
	NonceStr    string `json:"nonce_str"`    // 微信
	AppId       string `json:"appId"`        // 微信
	TradeNo     string `json:"trade_no"`     // POS
	QrCode      string `json:"qr_code"`      // 支付宝qrcode返回
	PayToken    string `json:"pay_token"`    // 支付令牌(线下APP支付、快捷支付)
	RedirectUrl string `json:"redirect_url"` // 银联
	PayNo       string `json:"pay_no"`       // 转账支付、快捷支付
	BtsUrl      string `json:"bts_url"`      // 无卡分期支付url
	CounterUrl  string `json:"counter_url"`  // 网关收银台支付url
	PayOrderNo  string `json:"pay_order_no"` // 网关收银台支付订单号
}
