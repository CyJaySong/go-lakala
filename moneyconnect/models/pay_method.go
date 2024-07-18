package models

type DepositParamsPayMethod struct {
	TransferPay *ConsumeParamsPayMethodByTRANSFERPAY `json:"TRANSFER_PAY,omitempty"` // 转账支付
}

type ConsumeParamsPayMethod struct {
	Scan        *ConsumeParamsPayMethodBySCAN        `json:"SCAN,omitempty"`         // 条码支付
	Jsapi       *ConsumeParamsPayMethodByJSAPI       `json:"JSAPI,omitempty"`        // 公众号/服务窗/小程序支付
	QrCode      *ConsumeParamsPayMethodByQRCode      `json:"QR_CODE,omitempty"`      // 创建支付宝付款码
	POS         *ConsumeParamsPayMethodByPOS         `json:"POS,omitempty"`          // POS支付
	Gateway     *ConsumeParamsPayMethodByGATEWAY     `json:"GATEWAY,omitempty"`      // 网关收银台支付
	TransferPay *ConsumeParamsPayMethodByTRANSFERPAY `json:"TRANSFER_PAY,omitempty"` // 转账支付
	BtsPay      *ConsumeParamsPayMethodByBTSPAY      `json:"BTS_PAY,omitempty"`      // 无卡分期支付
}

type ConsumeParamsPayMethodBySCAN struct {
	BarCode           string `json:"bar_code"`                       // 条码
	SubAppId          string `json:"app_id,omitempty"`               // 子商户公众账号ID的sub_appid
	Amount            int    `json:"amount"`                         // 支付金额,单位:分
	Subject           string `json:"subject,omitempty"`              // 订单标题(非必填)
	ChannelType       string `json:"channel_type,omitempty"`         // 渠道(非必填)
	MemberNo          string `json:"member_no,omitempty"`            // 会员标识(非必填),如上传则使用该会员进行收单
	MerchantNo        string `json:"merchant_no,omitempty"`          // 钱帐通商户号(非必填),如上传则使用该商户号进行收单且会员标识必填;如会员标识上传钱帐通商户号未传,则取该会员的第一个商户进行收单
	TermBaseStation   string `json:"term_base_station"`              // 商户终端基站信息,必填,上送格式为:MNC移动网络号码(2)＋LAC位置区域码(5)+CID基站编号(8)其中LAC不满五位,或CID不满足8位的均左补空格,如:00+LAC:6361+CID:58130
	TermLoc           string `json:"term_loc"`                       // 商户终端地理位置,必填,格式:纬度/经度,+表示北纬、东经,-表示南纬、西经,精度最长支持小数点后9位,举例:+37.123456789/-121.123456789
	HbFqNum           string `json:"hb_fq_num,omitempty"`            // 支付宝花呗分期期数,需花呗分期时必填,3:3期,6:6期,12:12期
	HbFqSellerPercent string `json:"hb_fq_seller_percent,omitempty"` // 支付宝花呗商家手续费承担比例,需花呗时必填,100:代表商家贴息,0:代表用户承担手续费
}

type ConsumeParamsPayMethodByJSAPI struct {
	Amount          int                                        `json:"amount"`                    // 支付金额,单位:分
	MchAppId        string                                     `json:"mch_appid"`                 // 必填,微信/支付宝/云闪付的appid
	OpenId          string                                     `json:"openid"`                    // 付款人ID,微信:openid,支付宝:buyer_id,云闪付:user_id
	Mode            string                                     `json:"mode"`                      // 支付宝:alipay,微信:wxpay,银联:unionpay
	Subject         string                                     `json:"subject,omitempty"`         // 订单标题(非必填)
	ChannelType     string                                     `json:"channel_type,omitempty"`    // 渠道(非必填)
	FrontURL        string                                     `json:"front_url,omitempty"`       // 银联前台通知地址
	FrontFailURL    string                                     `json:"front_fail_url,omitempty"`  // 银联失败交易前台通知地址
	GoodsDetail     []ConsumeParamsPayMethodByJsapiGoodsDetail `json:"goods_detail,omitempty"`    // 商品列表信息
	MemberNo        string                                     `json:"member_no,omitempty"`       // 会员标识(非必填),如果上传则使用该会员进行收单
	MerchantNo      string                                     `json:"merchant_no,omitempty"`     // 钱帐通商户号(非必填),如果上传则使用该商户号进行收单且会员标识必填;如果会员标识上传钱帐通商户号未传,则取该会员的第一个商户进行收单
	TermBaseStation string                                     `json:"term_base_station"`         // 商户终端基站信息,必填,上送格式为:MNC移动网络号码(2)＋LAC位置区域码(5)+CID基站编号(8)其中LAC不满五位,或CID不满足8位的均左补空格,如:00+LAC:6361+CID:58130
	TermLoc         string                                     `json:"term_loc"`                  // 商户终端地理位置,必填,格式:纬度/经度,+表示北纬、东经,-表示南纬、西经,精度最长支持小数点后9位,举例:+37.123456789/-121.123456789
	HbFqNum         string                                     `json:"hb_fq_num,omitempty"`       // 支付宝花呗分期期数,需花呗分期时必填,3:3期,6:6期,12:12期支付宝花呗商家手续费承担比例,需花呗时必填,100:代表商家贴息,0:代表用户承担手续费支付宝花呗商家手续费承担比例
	TimeoutExpress  string                                     `json:"timeout_express,omitempty"` // 交易有效时间(单位分),订单的有效时间,建议不超过15分钟,不传值则默认5分钟
}

type ConsumeParamsPayMethodByJsapiGoodsDetail struct {
	GoodsId       string  `json:"goods_id"`                 // 商品编号
	GoodsName     string  `json:"goods_name"`               // 商品名称
	Spec          string  `json:"spec,omitempty"`           // 商品规格
	Quantity      int     `json:"quantity"`                 // 数量
	Price         float64 `json:"price"`                    // 价格(元)
	GoodsCategory string  `json:"goods_category,omitempty"` // 商品类目
	Body          string  `json:"body,omitempty"`           // 商品描述信息
	ShowURL       string  `json:"show_url,omitempty"`       // 商品展示地址
}

type ConsumeParamsPayMethodByQRCode struct {
	Amount          int    `json:"amount"`                 // 支付金额,单位:分
	Subject         string `json:"subject,omitempty"`      // 订单标题(非必填)
	ChannelType     string `json:"channel_type,omitempty"` // 渠道(非必填)
	Mode            string `json:"mode"`                   // 支付宝:alipay
	MemberNo        string `json:"member_no,omitempty"`    // 会员标识(非必填),如果上传则使用该会员进行收单
	MerchantNo      string `json:"merchant_no,omitempty"`  // 钱帐通商户号(非必填),如果上传则使用该商户号进行收单且会员标识必填;如果会员标识上传钱帐通商户号未传,则取该会员的第一个商户进行收单
	TermBaseStation string `json:"term_base_station"`      // 商户终端基站信息,必填,上送格式为:MNC移动网络号码(2)＋LAC位置区域码(5)+CID基站编号(8)其中LAC不满五位,或CID不满足8位的均左补空格,如:00+LAC:6361+CID:58130
	TermLoc         string `json:"term_loc"`               // 商户终端地理位置,必填,格式:纬度/经度,+表示北纬、东经,-表示南纬、西经,精度最长支持小数点后9位,举例:+37.123456789/-121.123456789
	ValidTime       int    `json:"valid_time,omitempty"`   // 付款码有效时间,mode为lakala时有效,单位秒，可选上送，不上送默认60秒，区间为60秒-7200秒,小于60秒则设为60秒，大于7200秒设为7200秒
}

type ConsumeParamsPayMethodByPOS struct {
	Amount     int    `json:"amount"`                // 支付金额,单位:分
	MemberNo   string `json:"member_no,omitempty"`   // 会员标识(非必填),如上传则使用该会员进行收单
	MerchantNo string `json:"merchant_no,omitempty"` // 钱帐通商户号(非必填),如上传则使用该商户号进行收单且会员标识必填;如会员标识上传钱帐通商户号未传,则取该会员的第一个商户进行收单
}

type ConsumeParamsPayMethodByGATEWAY struct {
	Amount           int               `json:"amount"`                       // 支付金额,单位:分
	MerchantNo       string            `json:"merchant_no"`                  // 钱帐通商户号(必填)
	MemberNo         string            `json:"member_no"`                    // 会员标识(非必填),如上传则使用该会员进行收单
	ProductName      string            `json:"product_name"`                 // 商品名称（非必填）
	SupportRefund    int               `json:"support_refund,omitempty"`     // 是否支持退款(非必填),0:不支持,1:支持,默认1
	SupportCancel    int               `json:"support_cancel,omitempty"`     // 是否支持撤销(非必填),0:不支持,1:支持,默认1
	SupportRepeatPay int               `json:"support_repeat_pay,omitempty"` // 是否支持多次支付(非必填),0:不支持,1:支持,默认1
	ExpireTime       string            `json:"expire_time,omitempty"`        // 过期时间(非必填),格式yyyyMMddHHmmss,最大支持下单时间+2天
	CounterParam     map[string]string `json:"counter_param,omitempty"`      // 收银台展示参数
	IdentityInfo     map[string]string `json:"identity_info,omitempty"`      // 实名信息(非必填)
}

type ConsumeParamsPayMethodByTRANSFERPAY struct {
	Amount int `json:"amount"` // 支付金额,单位:分
}

type ConsumeParamsPayMethodByBTSPAY struct {
	Amount     int    `json:"amount"`                // 支付金额,单位:分
	MemberNo   string `json:"member_no,omitempty"`   // 会员标识(非必填),如上传则使用该会员进行收单
	MerchantNo string `json:"merchant_no,omitempty"` // 钱帐通商户号(非必填),如上传则使用该商户号进行收单且会员标识必填;如会员标识上传钱帐通商户号未传,则取该会员的第一个商户进行收单
}
