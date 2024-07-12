package moneyconnect

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

// ApplyMemberSplitScale 申请修改分账比例
func (t *Client) ApplyMemberSplitScale(params models.ApplyMemberSplitScaleParams) (body *models.BaseResponse[models.ApplyMemberSplitScaleResult], err error) {
	val := t.getBaseBody("member.split.scale.apply")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.ApplyMemberSplitScaleResult])
	err = t.reqPost(req, body)
	return
}

// ConfirmMemberSplitScale 确认修改分账比例
func (t *Client) ConfirmMemberSplitScale(params models.ConfirmMemberSplitScaleParams) (body *models.BaseResponse[models.ConfirmMemberSplitScaleResult], err error) {
	val := t.getBaseBody("member.split.scale.confirm")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.ConfirmMemberSplitScaleResult])
	err = t.reqPost(req, body)
	return
}

// QueryMemberSplitScale 查询分账比例
func (t *Client) QueryMemberSplitScale(params models.QueryMemberSplitScaleParams) (body *models.BaseResponse[models.QueryMemberSplitScaleResult], err error) {
	val := t.getBaseBody("member.split.scale.query")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	bytes, err := t.reqPostBytes(req)
	if err != nil {
		return nil, err
	} else if jsonObj, err := sonic.Get(bytes); err != nil {
		return nil, err
	} else if splitScaleListNode := jsonObj.GetByPath("result", "split_scale_list"); splitScaleListNode.Valid() {
		splitScaleList, err := splitScaleListNode.ArrayUseNode()
		if err != nil {
			return nil, err
		}
		for i, splitScaleNode := range splitScaleList {
			statusStr, err := splitScaleNode.Get("status").String()
			if err != nil {
				return nil, err
			}
			if _, err = splitScaleNode.Set("status", ast.NewNumber(statusStr)); err != nil {
				return nil, err
			}
			if _, err = splitScaleListNode.SetByIndex(i, splitScaleNode); err != nil {
				return nil, err
			}
		}
		if bytes, err = jsonObj.MarshalJSON(); err != nil {
			return nil, err
		}
	}
	body = new(models.BaseResponse[models.QueryMemberSplitScaleResult])
	err = json.Unmarshal(bytes, body)
	return
}
