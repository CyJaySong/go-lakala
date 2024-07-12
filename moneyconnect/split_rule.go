package moneyconnect

import (
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// CreateSplitRule 新建分账规则
func (t *Client) CreateSplitRule(params models.CreateSplitRuleParams) (body *models.BaseResponse[models.CreateSplitRuleResult], err error) {
	val := t.getBaseBody("order.splitrule.create")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.CreateSplitRuleResult])
	err = t.reqPost(req, body)
	return
}

// RemoveSplitRule 删除分账规则
func (t *Client) RemoveSplitRule(params models.RemoveSplitRuleParams) (body *models.BaseResponse[models.RemoveSplitRuleResult], err error) {
	val := t.getBaseBody("order.splitrule.remove")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.RemoveSplitRuleResult])
	err = t.reqPost(req, body)
	return
}

// StopSplitRule 停用分账规则
func (t *Client) StopSplitRule(params models.StopSplitRuleParams) (body *models.BaseResponse[models.StopSplitRuleResult], err error) {
	val := t.getBaseBody("order.splitrule.stop")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.StopSplitRuleResult])
	err = t.reqPost(req, body)
	return
}

// ResumeSplitRule 启用分账规则
func (t *Client) ResumeSplitRule(params models.ResumeSplitRuleParams) (body *models.BaseResponse[models.ResumeSplitRuleResult], err error) {
	val := t.getBaseBody("order.splitrule.resume")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.ResumeSplitRuleResult])
	err = t.reqPost(req, body)
	return
}

// GetSplitRule 获取分账规则
func (t *Client) GetSplitRule(params models.GetSplitRuleParams) (body *models.BaseResponse[models.GetSplitRuleResult], err error) {
	val := t.getBaseBody("order.splitrule.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetSplitRuleResult])
	err = t.reqPost(req, body)
	return
}

// EvaluateSplitRule 获取分账规则
func (t *Client) EvaluateSplitRule(params models.EvaluateSplitRuleParams) (body *models.BaseResponse[models.EvaluateSplitRuleResult], err error) {
	val := t.getBaseBody("order.splitrule.evaluate")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.EvaluateSplitRuleResult])
	err = t.reqPost(req, body)
	return
}
