package moneyconnect

import (
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// GetAccountBalance 查询账户余额
func (t *Client) GetAccountBalance(params models.GetAccountBalanceParams) (body *models.BaseResponse[models.GetAccountBalanceResult], err error) {
	val := t.getBaseBody("account.balance.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetAccountBalanceResult])
	err = t.reqPost(req, body)
	return
}

// GetTrustAccountBalance 查询账户余额
func (t *Client) GetTrustAccountBalance(params models.GetTrustAccountBalanceParams) (body *models.BaseResponse[models.GetTrustAccountBalanceResult], err error) {
	val := t.getBaseBody("account.trust.balance.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetTrustAccountBalanceResult])
	err = t.reqPost(req, body)
	return
}

// GetAccountIncomeCost 查询收入及支出
func (t *Client) GetAccountIncomeCost(params models.GetAccountIncomeCostParams) (body *models.BaseResponse[models.GetAccountIncomeCostResult], err error) {
	val := t.getBaseBody("account.income.cost.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetAccountIncomeCostResult])
	err = t.reqPost(req, body)
	return
}

// FreezeAccountBalance 冻结账户余额
func (t *Client) FreezeAccountBalance(params models.FreezeAccountBalanceParams) (body *models.BaseResponse[models.FreezeAccountBalanceResult], err error) {
	val := t.getBaseBody("account.balance.freeze")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.FreezeAccountBalanceResult])
	err = t.reqPost(req, body)
	return
}

// UnfreezeAccountBalance 解冻账户余额
func (t *Client) UnfreezeAccountBalance(params models.UnfreezeAccountBalanceParams) (body *models.BaseResponse[models.UnfreezeAccountBalanceResult], err error) {
	val := t.getBaseBody("account.balance.unfreeze")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.UnfreezeAccountBalanceResult])
	err = t.reqPost(req, body)
	return
}

// GetAccountHistoryList 查询账户变更明细
func (t *Client) GetAccountHistoryList(params models.GetAccountHistoryListParams) (body *models.BaseResponse[models.GetAccountHistoryListResult], err error) {
	val := t.getBaseBody("account.history.list")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetAccountHistoryListResult])
	err = t.reqPost(req, body)
	return
}

// QueryAccountFeeUnCollect 查询账户手续费欠款
func (t *Client) QueryAccountFeeUnCollect(params models.QueryAccountFeeUnCollectParams) (body *models.BaseResponse[models.QueryAccountFeeUnCollectResult], err error) {
	val := t.getBaseBody("account.fee.uncollect.query")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.QueryAccountFeeUnCollectResult])
	err = t.reqPost(req, body)
	return
}

// QueryAccountFeeList 查询账户手续费记录
func (t *Client) QueryAccountFeeList(params models.QueryAccountFeeListParams) (body *models.BaseResponse[models.QueryAccountFeeListResult], err error) {
	val := t.getBaseBody("account.fee.list")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.QueryAccountFeeListResult])
	err = t.reqPost(req, body)
	return
}
