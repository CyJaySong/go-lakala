package moneyconnect

import (
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// GetMemberBindBankCardH5Url 会员绑卡相关页面
func (t *Client) GetMemberBindBankCardH5Url(params models.GetMemberBindBankCardH5UrlParams) (body *models.BaseResponse[models.GetMemberBindBankCardH5UrlResult], err error) {
	val := t.getBaseBody("member.bank.card.page.url.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetMemberBindBankCardH5UrlResult])
	err = t.reqPost(req, body)
	return
}

// QueryMemberBindBankCard 查询绑定银行卡列表
func (t *Client) QueryMemberBindBankCard(params models.QueryMemberBindBankCardParams) (body *models.BaseResponse[models.QueryMemberBindBankCardResult], err error) {
	val := t.getBaseBody("member.bank.card.bind.query")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.QueryMemberBindBankCardResult])
	err = t.reqPost(req, body)
	return
}

// UpdateMemberBankCardContact 修改银行卡联行号
func (t *Client) UpdateMemberBankCardContact(params models.UpdateMemberBankCardContactParams) (body *models.BaseResponse[models.UpdateMemberBankCardContactResult], err error) {
	val := t.getBaseBody("member.bank.card.contact.update")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.UpdateMemberBankCardContactResult])
	err = t.reqPost(req, body)
	return
}
