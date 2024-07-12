package moneyconnect

import (
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// OrderConsumeComplete 完成订单（确权分账）
func (t *Client) OrderConsumeComplete(params models.OrderConsumeCompleteParams) (body *models.BaseResponse[models.OrderConsumeCompleteResult], err error) {
	val := t.getBaseBody("order.complete")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.OrderConsumeCompleteResult])
	err = t.reqPost(req, body)
	return
}

// MerchantCommissionBalanceSplit 商户佣金专用账户分账
func (t *Client) MerchantCommissionBalanceSplit(params models.MerchantCommissionBalanceSplitParams) (body *models.BaseResponse[models.MerchantCommissionBalanceSplitResult], err error) {
	val := t.getBaseBody("order.balance.split")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.MerchantCommissionBalanceSplitResult])
	err = t.reqPost(req, body)
	return
}

// RevokeSplit 撤销分账，只支持分账全额撤销
func (t *Client) RevokeSplit(params models.RevokeSplitParams) (body *models.BaseResponse[models.RevokeSplitResult], err error) {
	val := t.getBaseBody("order.split.revoke")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.RevokeSplitResult])
	err = t.reqPost(req, body)
	return
}
