package constants

const (
	// BankCardTypeDebitCard 储蓄卡
	BankCardTypeDebitCard = 1
	// BankCardTypeCreditCard 信用卡
	BankCardTypeCreditCard = 2
)

const (
	// BankAccountTypePrivate 对私账户
	BankAccountTypePrivate = 0
	// BankAccountTypeCorporate 对公账户
	BankAccountTypeCorporate = 1
)

var BankCodes = []struct {
	Code int
	Name string
}{
	{102, "工商银行"},
	{103, "农业银行"},
	{104, "中国银行"},
	{105, "建设银行"},
	{301, "交通银行"},
	{302, "中信银行"},
	{303, "光大银行"},
	{304, "华夏银行"},
	{305, "民生银行"},
	{306, "广东发展银行"},
	{307, "深圳发展银行"},
	{308, "招商银行"},
	{309, "兴业银行"},
	{310, "上海浦东发展银行"},
	{313, "其它城市商业银行"},
	{315, "恒丰银行"},
	{316, "浙商银行"},
	{318, "渤海银行"},
	{319, "徽商银行"},
	{401, "城市信用社"},
	{403, "邮政储蓄银行"},
	{501, "汇丰银行"},
}
