package moneyconnect

import (
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// GetMemberOtherChannelSubMerchant 查询渠道子商户
func (t *Client) GetMemberOtherChannelSubMerchant(params models.GetMemberOtherChannelSubMerchantParams) (body *models.BaseResponse[models.GetMemberOtherChannelSubMerchantResult], err error) {
	val := t.getBaseBody("member.other.channel.sub.merchant.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetMemberOtherChannelSubMerchantResult])
	err = t.reqPost(req, body)
	return
}

// GetMemberTransferPayAccount 查询转账支付打款账户
func (t *Client) GetMemberTransferPayAccount(params models.GetMemberTransferPayAccountParams) (body *models.BaseResponse[models.GetMemberTransferPayAccountResult], err error) {
	val := t.getBaseBody("member.transfer.pay.account.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetMemberTransferPayAccountResult])
	err = t.reqPost(req, body)
	return
}

// GetUnionPayUserId 获取银联UserId
func (t *Client) GetUnionPayUserId(params models.GetUnionPayUserIdParams) (body *models.BaseResponse[models.GetUnionPayUserIdResult], err error) {
	val := t.getBaseBody("other.openid.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetUnionPayUserIdResult])
	err = t.reqPost(req, body)
	return
}
