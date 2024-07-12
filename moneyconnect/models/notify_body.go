package models

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
)

// NotifyBody http://47.110.246.50:6524/docs/qzt/qzt-1d50jciip1f9t
type NotifyBody struct {
	AppId     string `json:"app_id"`     // 应用ID
	Timestamp string `json:"timestamp"`  // UNIX时间戳
	Version   string `json:"version"`    // API版本
	Sign      string `json:"sign"`       // 签名
	EventType string `json:"event_type"` // 事件类型
	Event     string `json:"event"`      // 事件类型
}

// MemberOpenEvent 会员开户结果通知
func (nb *NotifyBody) MemberOpenEvent() (event *MemberOpenEvent, err error) {
	bytes := []byte(nb.Event)
	jsonObj, err := sonic.Get(bytes)
	if err != nil {
		return nil, err
	}
	if node := jsonObj.Get("member_state"); node.Valid() {
		num, err := node.String()
		if err != nil {
			return nil, err
		}
		if _, err = jsonObj.Set("member_state", ast.NewNumber(num)); err != nil {
			return nil, err
		}
	}
	if node := jsonObj.Get("member_type"); node.Valid() {
		num, err := node.String()
		if err != nil {
			return nil, err
		}
		if _, err = jsonObj.Set("member_type", ast.NewNumber(num)); err != nil {
			return nil, err
		}
	}
	if bytes, err = jsonObj.MarshalJSON(); err != nil {
		return nil, err
	}
	event = new(MemberOpenEvent)
	if err = json.Unmarshal(bytes, &event); err != nil {
		return nil, err
	}
	return
}

// EsignCreateEvent 电子签约通知
func (nb *NotifyBody) EsignCreateEvent() (event *EsignCreateEvent, err error) {
	event = new(EsignCreateEvent)
	if err = json.Unmarshal([]byte(nb.Event), &event); err != nil {
		return nil, err
	}
	return
}

// MemberMerchantOpenEvent 商户开通结果通知
func (nb *NotifyBody) MemberMerchantOpenEvent() (event *MemberMerchantOpenEvent, err error) {
	bytes := []byte(nb.Event)
	jsonObj, err := sonic.Get(bytes)
	if err != nil {
		return nil, err
	}
	if node := jsonObj.Get("merchant_status"); node.Valid() {
		num, err := node.String()
		if err != nil {
			return nil, err
		}
		if _, err = jsonObj.Set("merchant_status", ast.NewNumber(num)); err != nil {
			return nil, err
		}
	}
	if node := jsonObj.Get("split_scale_status"); node.Valid() {
		num, err := node.String()
		if err != nil {
			return nil, err
		}
		if _, err = jsonObj.Set("split_scale_status", ast.NewNumber(num)); err != nil {
			return nil, err
		}
	}
	if bytes, err = jsonObj.MarshalJSON(); err != nil {
		return nil, err
	}
	event = new(MemberMerchantOpenEvent)
	if err = json.Unmarshal(bytes, &event); err != nil {
		return nil, err
	}
	return
}

// MemberBankCardEvent 银行卡页面相关操作结果通知
func (nb *NotifyBody) MemberBankCardEvent() (event *MemberBankCardEvent, err error) {
	event = new(MemberBankCardEvent)
	if err = json.Unmarshal([]byte(nb.Event), &event); err != nil {
		return nil, err
	}
	return
}

// MemberSplitScaleEvent 修改分账比例结果通知
func (nb *NotifyBody) MemberSplitScaleEvent() (event *MemberSplitScaleEvent, err error) {
	bytes := []byte(nb.Event)
	jsonObj, err := sonic.Get(bytes)
	if err != nil {
		return nil, err
	} else if node := jsonObj.Get("status"); node.Valid() {
		num, err := node.String()
		if err != nil {
			return nil, err
		}
		if _, err = jsonObj.Set("status", ast.NewNumber(num)); err != nil {
			return nil, err
		}
		if bytes, err = jsonObj.MarshalJSON(); err != nil {
			return nil, err
		}
	}
	event = new(MemberSplitScaleEvent)
	if err = json.Unmarshal(bytes, &event); err != nil {
		return nil, err
	}
	return
}

// MemberMerchantTermEvent 会员商户增终结果通知
func (nb *NotifyBody) MemberMerchantTermEvent() (event *MemberMerchantTermEvent, err error) {
	bytes := []byte(nb.Event)
	jsonObj, err := sonic.Get(bytes)
	if err != nil {
		return nil, err
	} else if termsNode := jsonObj.Get("terms"); termsNode.Valid() {
		termList, err := termsNode.ArrayUseNode()
		if err != nil {
			return nil, err
		}
		for i, termNode := range termList {
			merchantStatusStr, err := termNode.Get("merchant_status").String()
			if err != nil {
				return nil, err
			}
			if _, err = termNode.Set("merchant_status", ast.NewNumber(merchantStatusStr)); err != nil {
				return nil, err
			}
			if _, err = termsNode.SetByIndex(i, termNode); err != nil {
				return nil, err
			}
		}
		if bytes, err = jsonObj.MarshalJSON(); err != nil {
			return nil, err
		}
	}
	event = new(MemberMerchantTermEvent)
	if err = json.Unmarshal(bytes, &event); err != nil {
		return nil, err
	}
	return
}

// OrderDepositRequestEvent 充值结果通知
func (nb *NotifyBody) OrderDepositRequestEvent() (event *OrderDepositRequestEvent, err error) {
	event = new(OrderDepositRequestEvent)
	if err = json.Unmarshal([]byte(nb.Event), &event); err != nil {
		return nil, err
	}
	return
}

// OrderWithdrawRequestEvent 提现结果通知
func (nb *NotifyBody) OrderWithdrawRequestEvent() (event *OrderWithdrawRequestEvent, err error) {
	event = new(OrderWithdrawRequestEvent)
	if err = json.Unmarshal([]byte(nb.Event), &event); err != nil {
		return nil, err
	}
	return
}

// OrderConsumePayEvent 消费支付结果通知
func (nb *NotifyBody) OrderConsumePayEvent() (event *OrderConsumePayEvent, err error) {
	bytes := []byte(nb.Event)
	if jsonObj, err := sonic.Get(bytes); err != nil {
		return nil, err
	} else if node := jsonObj.Get("pay_status"); node.Valid() {
		num, err := node.String()
		if err != nil {
			return nil, err
		}
		if _, err = jsonObj.Set("pay_status", ast.NewNumber(num)); err != nil {
			return nil, err
		}
		if bytes, err = jsonObj.MarshalJSON(); err != nil {
			return nil, err
		}
	}
	event = new(OrderConsumePayEvent)
	if err = json.Unmarshal(bytes, &event); err != nil {
		return nil, err
	}
	return
}

// OrderConsumeBatchPayEvent 批量支付结果通知
func (nb *NotifyBody) OrderConsumeBatchPayEvent() (event *OrderConsumeBatchPayEvent, err error) {
	bytes := []byte(nb.Event)
	jsonObj, err := sonic.Get(bytes)
	if err != nil {
		return nil, err
	}
	if node := jsonObj.Get("pay_status"); node.Valid() {
		num, err := node.String()
		if err != nil {
			return nil, err
		}
		if _, err = jsonObj.Set("pay_status", ast.NewNumber(num)); err != nil {
			return nil, err
		}
	}
	if subOrderListNode := jsonObj.Get("sub_order_list"); subOrderListNode.Valid() {
		subOrderList, err := subOrderListNode.ArrayUseNode()
		if err != nil {
			return nil, err
		}
		for i, subOrderNode := range subOrderList {
			subOrderStr, err := subOrderNode.String()
			if err != nil {
				return nil, err
			}
			subOrderJson, err := sonic.GetFromString(subOrderStr)
			if err != nil {
				return nil, err
			}
			if _, err = subOrderListNode.SetByIndex(i, subOrderJson); err != nil {
				return nil, err
			}
		}
	}
	if bytes, err = jsonObj.MarshalJSON(); err != nil {
		return nil, err
	}
	event = new(OrderConsumeBatchPayEvent)
	if err = json.Unmarshal(bytes, &event); err != nil {
		return nil, err
	}
	return
}

// OrderConsumeBatchPaySplitEvent 批量支付分账结果通知
func (nb *NotifyBody) OrderConsumeBatchPaySplitEvent() (event *OrderConsumeBatchPaySplitEvent, err error) {
	bytes := []byte(nb.Event)
	if jsonObj, err := sonic.Get(bytes); err != nil {
		return nil, err
	} else if subOrderSplitListNode := jsonObj.Get("sub_order_split_list"); subOrderSplitListNode.Valid() {
		subOrderSplitList, err := subOrderSplitListNode.ArrayUseNode()
		if err != nil {
			return nil, err
		}
		for i, subOrderSplitNode := range subOrderSplitList {
			subOrderSplitStr, err := subOrderSplitNode.String()
			if err != nil {
				return nil, err
			}
			subOrderSplitJson, err := sonic.GetFromString(subOrderSplitStr)
			if err != nil {
				return nil, err
			}
			if _, err = subOrderSplitListNode.SetByIndex(i, subOrderSplitJson); err != nil {
				return nil, err
			}
		}
		if bytes, err = jsonObj.MarshalJSON(); err != nil {
			return nil, err
		}
	}
	event = new(OrderConsumeBatchPaySplitEvent)
	if err = json.Unmarshal(bytes, &event); err != nil {
		return nil, err
	}
	return
}

// PosTradeNotifyEvent POS支付结果通知
func (nb *NotifyBody) PosTradeNotifyEvent() (event *PosTradeNotifyEvent, err error) {
	event = new(PosTradeNotifyEvent)
	if err = json.Unmarshal([]byte(nb.Event), &event); err != nil {
		return nil, err
	}
	return
}

// OrderConsumeRefundEvent 退款结果通知
func (nb *NotifyBody) OrderConsumeRefundEvent() (event *OrderConsumeRefundEvent, err error) {
	event = new(OrderConsumeRefundEvent)
	if err = json.Unmarshal([]byte(nb.Event), &event); err != nil {
		return nil, err
	}
	return
}

// TransferPayMatchOrderFailEvent 转账支付匹配订单失败通知
func (nb *NotifyBody) TransferPayMatchOrderFailEvent() (event *TransferPayMatchOrderFailEvent, err error) {
	event = new(TransferPayMatchOrderFailEvent)
	if err = json.Unmarshal([]byte(nb.Event), &event); err != nil {
		return nil, err
	}
	return
}

// TransferPayOrderTradeEvent 转账支付订单入账通知
func (nb *NotifyBody) TransferPayOrderTradeEvent() (event *TransferPayOrderTradeEvent, err error) {
	event = new(TransferPayOrderTradeEvent)
	if err = json.Unmarshal([]byte(nb.Event), &event); err != nil {
		return nil, err
	}
	return
}

// OrderChannelFeeEvent 通道手续费通知事件
func (nb *NotifyBody) OrderChannelFeeEvent() (event *OrderChannelFeeEvent, err error) {
	event = new(OrderChannelFeeEvent)
	if err = json.Unmarshal([]byte(nb.Event), &event); err != nil {
		return nil, err
	}
	return
}

// OrderWithdrawChargebackEvent 提现退单通知事件
func (nb *NotifyBody) OrderWithdrawChargebackEvent() (event *OrderWithdrawChargebackEvent, err error) {
	event = new(OrderWithdrawChargebackEvent)
	if err = json.Unmarshal([]byte(nb.Event), &event); err != nil {
		return nil, err
	}
	return
}
