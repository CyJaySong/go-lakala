package constants

const (

	// MemberTypePerson 个人会员
	MemberTypePerson = 2
	// MemberTypeEnterprise 企业会员
	MemberTypeEnterprise = 3
	// MemberTypeIndividualBusiness 个体户
	MemberTypeIndividualBusiness = 4

	// MemberCategoryGeneral 普通会员
	MemberCategoryGeneral = 0
	// MemberCategoryMerchant 收单商户
	MemberCategoryMerchant = 1
)

const (
	// MemberStateEffective 有效的
	MemberStateEffective = 1
	// MemberStateLock 锁定
	MemberStateLock = 10
	// MemberStateClose 关闭
	MemberStateClose = 20
	// MemberStatePendingAudit 待审核
	MemberStatePendingAudit = 99
)

const (
	// MemberReadStateNot 未实名
	MemberReadStateNot = 0
	// MemberReadStateDone 已实名
	MemberReadStateDone = 1
)
