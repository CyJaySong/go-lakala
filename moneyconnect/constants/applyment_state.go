package constants

const (
	ApplymentStateFail   = "APPLYMENT_STATE_FAIL"   // 提交失败
	ApplymentStateCommit = "APPLYMENT_STATE_COMMIT" // 已提交
	/* 支付宝实名状态 */
	Auditing       = "AUDITING"        // 支付宝-审核中
	ContactConfirm = "CONTACT_CONFIRM" // 支付宝-待联系人确认
	LegalConfirm   = "LEGAL_CONFIRM"   // 支付宝-待法人确认
	AuditPass      = "AUDIT_PASS"      // 支付宝-审核通过
	AuditReject    = "AUDIT_REJECT"    // 支付宝-审核驳回
	AuditFreeze    = "AUDIT_FREEZE"    // 支付宝-已冻结
	Canceled       = "CANCELED"        // 支付宝-已撤回
	/* 微信实名状态 */
	ApplymentStateWaittingForAudit              = "APPLYMENT_STATE_WAITTING_FOR_AUDIT"               // 审核中
	ApplymentStateEditting                      = "APPLYMENT_STATE_EDITTING"                         // 编辑中
	ApplymentStateWaittingForConfirmContact     = "APPLYMENT_STATE_WAITTING_FOR_CONFIRM_CONTACT"     // 待确认联系信息
	ApplymentStateWaittingForConfirmLegalperson = "APPLYMENT_STATE_WAITTING_FOR_CONFIRM_LEGALPERSON" // 待账户验证
	ApplymentStatePassed                        = "APPLYMENT_STATE_PASSED"                           // 审核通过
	ApplymentStateRejected                      = "APPLYMENT_STATE_REJECTED"                         // 审核驳回
	ApplymentStateFreezed                       = "APPLYMENT_STATE_FREEZED"                          // 已冻结
	ApplymentStateCanceled                      = "APPLYMENT_STATE_CANCELED"                         // 已作废
)
