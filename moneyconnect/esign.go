package moneyconnect

import (
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// GetMemberEsignCreateH5Url 获取独立电子签约H5 Url
func (t *Client) GetMemberEsignCreateH5Url(params models.GetMemberEsignCreateH5UrlParams) (body *models.BaseResponse[models.GetMemberEsignCreateH5UrlResult], err error) {
	val := t.getBaseBody("member.esign.create.page.url.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetMemberEsignCreateH5UrlResult])
	err = t.reqPost(req, body)
	return
}

// CreateMemberEsign 创建电子签约
func (t *Client) CreateMemberEsign(params models.CreateMemberEsignParams) (body *models.BaseResponse[models.CreateMemberEsignResult], err error) {
	val := t.getBaseBody("esign.create")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.CreateMemberEsignResult])
	err = t.reqPost(req, body)
	return
}

// QueryMemberEsignByProjectCode 根据电子合同签约授权号查询电子合同
func (t *Client) QueryMemberEsignByProjectCode(params models.QueryMemberEsignParams) (body *models.BaseResponse[models.QueryMemberEsignByProjectCodeResult], err error) {
	val := t.getBaseBody("esign.query.by.project.code")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.QueryMemberEsignByProjectCodeResult])
	err = t.reqPost(req, body)
	return
}

// QueryMemberEsignByContractNo 根据合同号查询电子合同
func (t *Client) QueryMemberEsignByContractNo(params models.QueryMemberEsignParams) (body *models.BaseResponse[models.QueryMemberEsignByContractNoResult], err error) {
	val := t.getBaseBody("esign.query.by.contract.no")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.QueryMemberEsignByContractNoResult])
	err = t.reqPost(req, body)
	return
}
