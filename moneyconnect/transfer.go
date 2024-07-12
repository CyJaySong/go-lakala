package moneyconnect

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
	"strconv"
)

// Transfer 平台补贴商户会员
func (t *Client) Transfer(params models.TransferParams) (body *models.BaseResponse[models.TransferResult], err error) {
	val := t.getBaseBody("app.order.transfer")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.TransferResult])
	err = t.reqPost(req, body)
	return
}

// RevokeTransfer 平台转账撤销
func (t *Client) RevokeTransfer(params models.TransferParams) (body *models.BaseResponse[models.RevokeTransferResult], err error) {
	val := t.getBaseBody("app.order.transfer.revoke")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	bytes, err := t.reqPostBytes(req)
	if err != nil {
		return nil, err
	} else if jsonObj, err := sonic.Get(bytes); err != nil {
		return nil, err
	} else if node := jsonObj.GetByPath("result", "order_no"); node.Valid() {
		num, err := node.Int64()
		if err != nil {
			return nil, err
		}
		orderNo := strconv.FormatInt(num, 10)
		if _, err = jsonObj.Get("result").Set("order_no", ast.NewString(orderNo)); err != nil {
			return nil, err
		}
		if bytes, err = jsonObj.MarshalJSON(); err != nil {
			return nil, err
		}
	}
	body = new(models.BaseResponse[models.RevokeTransferResult])
	err = json.Unmarshal(bytes, body)
	return
}
