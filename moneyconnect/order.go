package moneyconnect

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
	"strconv"
)

// Deposit 充值（仅用于平台营销专用账户及商户退货专用账户充值）
func (t *Client) Deposit(params models.DepositParams) (body *models.BaseResponse[models.DepositResult], err error) {
	val := t.getBaseBody("order.deposit.request")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.DepositResult])
	err = t.reqPost(req, body)
	return
}

// Withdraw 提现（手动结算）
func (t *Client) Withdraw(params models.WithdrawParams) (body *models.BaseResponse[models.WithdrawResult], err error) {
	val := t.getBaseBody("order.withdraw.request")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.WithdrawResult])
	err = t.reqPost(req, body)
	return
}

// OrderConsume 消费
func (t *Client) OrderConsume(params models.OrderConsumeParams) (body *models.BaseResponse[models.OrderConsumeResult], err error) {
	val := t.getBaseBody("order.consume.request")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.OrderConsumeResult])
	err = t.reqPost(req, body)
	return
}

// OrderConsumePay 支付消费（拆单支付）
func (t *Client) OrderConsumePay(params models.OrderConsumePayParams) (body *models.BaseResponse[models.OrderConsumePayResult], err error) {
	val := t.getBaseBody("order.consume.pay.request")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.OrderConsumePayResult])
	err = t.reqPost(req, body)
	return
}

// OrderConsumeBatchPay 批量支付消费订单（合单支付）
func (t *Client) OrderConsumeBatchPay(params models.OrderConsumeBatchPayParams) (body *models.BaseResponse[models.OrderConsumeBatchPayResult], err error) {
	val := t.getBaseBody("order.consume.batchpay.request")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.OrderConsumeBatchPayResult])
	err = t.reqPost(req, body)
	return
}

// OrderActionConfirm 订单确认（短信）
func (t *Client) OrderActionConfirm(params models.OrderActionConfirmParams) (body *models.BaseResponse[models.OrderActionConfirmResult], err error) {
	val := t.getBaseBody("order.action.confirm")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.OrderActionConfirmResult])
	err = t.reqPost(req, body)
	return
}

// OrderConsumeRefund 退款消费
func (t *Client) OrderConsumeRefund(params models.OrderConsumeRefundParams) (body *models.BaseResponse[models.OrderConsumeRefundResult], err error) {
	val := t.getBaseBody("order.consume.refund")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.OrderConsumeRefundResult])
	err = t.reqPost(req, body)
	return
}

// OrderConsumeClose 关闭订单交易,该功能只支持消费支付方式为JSAPI,说明:关单交易只能对主扫支付场景下,未完成的支付交易进行关单,NATIVE交易不可以关单。
func (t *Client) OrderConsumeClose(params models.OrderConsumeCloseParams) (body *models.BaseResponse[models.OrderConsumeCloseResult], err error) {
	val := t.getBaseBody("order.pay.request.close")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.OrderConsumeCloseResult])
	err = t.reqPost(req, body)
	return
}

// GetOrderStatus 查询订单状态
func (t *Client) GetOrderStatus(params models.GetOrderStatusParams) (body *models.BaseResponse[models.GetOrderStatusResult], err error) {
	val := t.getBaseBody("order.status.get")
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
	body = new(models.BaseResponse[models.GetOrderStatusResult])
	err = json.Unmarshal(bytes, body)
	return
}

// GetOrderPayStatus 查询订单支付状态
func (t *Client) GetOrderPayStatus(params models.GetOrderPayStatusParams) (body *models.BaseResponse[models.GetOrderPayStatusResult], err error) {
	val := t.getBaseBody("order.pay.status.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetOrderPayStatusResult])
	err = t.reqPost(req, body)
	return
}

// GetOrderRefundStatus 查询订单退款状态
func (t *Client) GetOrderRefundStatus(params models.GetOrderRefundStatusParams) (body *models.BaseResponse[models.GetOrderRefundStatusResult], err error) {
	val := t.getBaseBody("order.refund.status.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.GetOrderRefundStatusResult])
	err = t.reqPost(req, body)
	return
}

// GetOrderDetailInfo 查询订单详细信息
func (t *Client) GetOrderDetailInfo(params models.GetOrderDetailInfoParams) (body *models.BaseResponse[models.GetOrderDetailInfoResult], err error) {
	val := t.getBaseBody("order.detail.info.get")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	bytes, err := t.reqPostBytes(req)
	if err != nil {
		return nil, err
	} else if jsonObj, err := sonic.Get(bytes); err != nil {
		return nil, err
	} else if payListNode := jsonObj.GetByPath("result", "pay_list"); payListNode.Valid() {
		payList, err := payListNode.ArrayUseNode()
		if err != nil {
			return nil, err
		}
		for i, payNode := range payList {
			payMethodNode := payNode.Get("pay_method")
			if !payMethodNode.Valid() {
				break
			}
			payMethodStr, err := payMethodNode.String()
			if err != nil {
				return nil, err
			}
			payMethodJson, err := sonic.GetFromString(payMethodStr)
			if err != nil {
				return nil, err
			}
			if _, err = payNode.Set("pay_method", payMethodJson); err != nil {
				return nil, err
			}
			if _, err = payListNode.SetByIndex(i, payNode); err != nil {
				return nil, err
			}
		}
		if bytes, err = jsonObj.MarshalJSON(); err != nil {
			return nil, err
		}
	}
	body = new(models.BaseResponse[models.GetOrderDetailInfoResult])
	err = json.Unmarshal(bytes, body)
	return
}

// QueryOrderReceipt 查询电子回单
func (t *Client) QueryOrderReceipt(params models.QueryOrderReceiptParams) (body *models.BaseResponse[models.QueryOrderReceiptResult], err error) {
	val := t.getBaseBody("order.receipt.query")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.QueryOrderReceiptResult])
	err = t.reqPost(req, body)
	return
}
