package moneyconnect

import (
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// GetMemberMerchantTerminalH5Url 获取会员商户在线增终H5Url
func (t *Client) GetMemberMerchantTerminalH5Url(params models.GetMemberMerchantTerminalH5UrlParams) (body *models.BaseResponse[models.GetMemberMerchantTerminalH5UrlResult], err error) {
	val := t.getBaseBody("member.merchant.term.page.url.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetMemberMerchantTerminalH5UrlResult])
	err = t.reqPost(req, body)
	return
}
