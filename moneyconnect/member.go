package moneyconnect

import (
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// GetMemberOpenPartH5Url 获取开户进阶H5 Url
func (t *Client) GetMemberOpenPartH5Url(params models.GetMemberOpenPartH5UrlParams) (body *models.BaseResponse[models.GetMemberOpenPartH5UrlResult], err error) {
	val := t.getBaseBody("member.open.part.page.url.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetMemberOpenPartH5UrlResult])
	err = t.reqPost(req, body)
	return
}

// ChangeMemberInfo 变更会员信息
func (t *Client) ChangeMemberInfo(params models.ChangeMemberInfoParams) (body *models.BaseResponse[models.ChangeMemberInfoResult], err error) {
	val := t.getBaseBody("member.info.change")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.ChangeMemberInfoResult])
	err = t.reqPost(req, body)
	return
}

// QueryMemberInfo 查询会员信息
func (t *Client) QueryMemberInfo(params models.QueryMemberInfoParams) (body *models.BaseResponse[models.QueryMemberInfoResult], err error) {
	val := t.getBaseBody("member.info.query")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.QueryMemberInfoResult])
	err = t.reqPost(req, body)
	return
}

// ApplyMemberMobileModify 申请会员手机修改
func (t *Client) ApplyMemberMobileModify(params models.ApplyMemberMobileModifyParams) (body *models.BaseResponse[models.ApplyMemberMobileModifyResult], err error) {
	val := t.getBaseBody("member.mobile.modify.apply")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.ApplyMemberMobileModifyResult])
	err = t.reqPost(req, body)
	return
}

// ConfirmMemberMobileModify 确认会员手机修改
func (t *Client) ConfirmMemberMobileModify(params models.ConfirmMemberMobileModifyParams) (body *models.BaseResponse[models.ConfirmMemberMobileModifyResult], err error) {
	val := t.getBaseBody("member.mobile.modify.confirm")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.ConfirmMemberMobileModifyResult])
	err = t.reqPost(req, body)
	return
}

// QueryMemberChannelReal 查询支付宝微信实名认证
func (t *Client) QueryMemberChannelReal(params models.QueryMemberChannelRealParams) (body *models.BaseResponse[models.QueryMemberChannelRealResult], err error) {
	val := t.getBaseBody("member.merchant.query.channel.real")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.QueryMemberChannelRealResult])
	err = t.reqPost(req, body)
	return
}
