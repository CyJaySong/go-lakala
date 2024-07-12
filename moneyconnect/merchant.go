package moneyconnect

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// BindMemberMerchant 商户绑定
func (t *Client) BindMemberMerchant(params models.BindMemberMerchantParams) (body *models.BaseResponse[models.BindMemberMerchantResult], err error) {
	val := t.getBaseBody("member.merchant.bind")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	bytes, err := t.reqPostBytes(req)
	if err != nil {
		return nil, err
	} else if jsonObj, err := sonic.Get(bytes); err != nil {
		return nil, err
	} else if merchantStatusNode := jsonObj.GetByPath("result", "merchant_status"); merchantStatusNode.Valid() {
		merchantStatusStr, err := merchantStatusNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = jsonObj.Get("result").Set("merchant_status", ast.NewNumber(merchantStatusStr)); err != nil {
			return nil, err
		}
		if bytes, err = jsonObj.MarshalJSON(); err != nil {
			return nil, err
		}
	}
	body = new(models.BaseResponse[models.BindMemberMerchantResult])
	err = json.Unmarshal(bytes, body)
	return
}

// UnbindMemberMerchant 商户解绑
func (t *Client) UnbindMemberMerchant(params models.UnbindMemberMerchantParams) (body *models.BaseResponse[models.UnbindMemberMerchantResult], err error) {
	val := t.getBaseBody("member.merchant.unbind")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.UnbindMemberMerchantResult])
	err = t.reqPost(req, body)
	return
}

// QueryMemberMerchant 查询会员商户
func (t *Client) QueryMemberMerchant(params models.QueryMemberMerchantParams) (body *models.QueryMemberMerchantResult, err error) {
	val := t.getBaseBody("member.merchant.query")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.QueryMemberMerchantResult)
	err = t.reqPost(req, body)
	return
}

// ReplenishMemberMerchantAttachment 补录商户附件
func (t *Client) ReplenishMemberMerchantAttachment(params models.ReplenishMemberMerchantAttachmentParams) (body *models.BaseResponse[models.ReplenishMemberMerchantAttachmentResult], err error) {
	val := t.getBaseBody("member.merchant.attachment.replenish")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.ReplenishMemberMerchantAttachmentResult])
	err = t.reqPost(req, body)
	return
}

// QueryMemberMerchantQuota 查询商户支付收单限额
func (t *Client) QueryMemberMerchantQuota(params models.QueryMemberMerchantQuotaParams) (body *models.BaseResponse[models.QueryMemberMerchantQuotaResult], err error) {
	val := t.getBaseBody("member.merchant.quota.query")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	bytes, err := t.reqPostBytes(req)
	if err != nil {
		return nil, err
	}
	jsonObj, err := sonic.Get(bytes)
	if err != nil {
		return nil, err
	}
	resultNode := jsonObj.Get("result")
	if numNode := resultNode.GetByPath("debit_per_limit"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("debit_per_limit", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("debit_day_limit"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("debit_day_limit", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("debit_month_limit"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("debit_month_limit", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("credit_per_limit"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("credit_per_limit", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("credit_day_limit"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("credit_day_limit", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("credit_month_limit"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("credit_month_limit", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("wechat_per_limit"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("wechat_per_limit", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("wechat_day_limit"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("wechat_day_limit", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("wechat_month_limit"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("wechat_month_limit", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("wechat_min_amt"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("wechat_min_amt", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("wechat_max_amt"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("wechat_max_amt", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("debit_mix_amt"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("debit_mix_amt", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("debit_max_amt"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("debit_max_amt", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("credit_mix_amt"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("credit_mix_amt", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if numNode := resultNode.GetByPath("credit_max_amt"); numNode.Valid() {
		numStr, err := numNode.String()
		if err != nil {
			return nil, err
		}
		if _, err = resultNode.Set("credit_max_amt", ast.NewNumber(numStr)); err != nil {
			return nil, err
		}
	}
	if bytes, err = jsonObj.MarshalJSON(); err != nil {
		return nil, err
	}
	body = new(models.BaseResponse[models.QueryMemberMerchantQuotaResult])
	err = json.Unmarshal(bytes, body)
	return
}

// SubmitMemberMerchantReconsider 进件会员商户复议
func (t *Client) SubmitMemberMerchantReconsider(params models.SubmitMemberMerchantReconsiderParams) (body *models.BaseResponse[models.SubmitMemberMerchantReconsiderResult], err error) {
	val := t.getBaseBody("member.merchant.reconsider.submit")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.SubmitMemberMerchantReconsiderResult])
	err = t.reqPost(req, body)
	return
}
