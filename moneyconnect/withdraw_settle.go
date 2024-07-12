package moneyconnect

import (
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// SetMemberWithdrawSettleCard 设置会员结算卡
func (t *Client) SetMemberWithdrawSettleCard(params models.SetMemberWithdrawSettleCardParams) (body *models.BaseResponse[models.SetMemberWithdrawSettleCardResult], err error) {
	val := t.getBaseBody("member.withdraw.settle.card.set")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.SetMemberWithdrawSettleCardResult])
	err = t.reqPost(req, body)
	return
}

// GetMemberWithdrawSettleRecord 查询会员结算记录
func (t *Client) GetMemberWithdrawSettleRecord(params models.GetMemberWithdrawSettleRecordParams) (body *models.BaseResponse[models.GetMemberWithdrawSettleRecordResult], err error) {
	val := t.getBaseBody("member.settle.record.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetMemberWithdrawSettleRecordResult])
	err = t.reqPost(req, body)
	return
}
