package models

type SplitRuleData struct {
	RuleId    string                `json:"rule_id,omitempty"` // 规则ID
	Params    map[string]string     `json:"params,omitempty"`  // 规则中的参数值:{"参数1":"1000","参数2":"1000"}
	Members   map[string]string     `json:"members,omitempty"` // 分账参与方在规则中的角色 ，会员角色唯一:{"会员1ID":"角色1","会员2ID":"角色2"}
	SplitList []CustomSplitRuleData `json:"split_list"`        // 自定义模板
}

type CustomSplitRuleData struct {
	MemberNo         string `json:"member_no"`                    // 会员标示
	Amount           int    `json:"amount"`                       // 分账金额,分
	SpecialAccountNo string `json:"special_account_no,omitempty"` // 可选,不填则进入余额账户,选择了则分到商户佣金专用账户,仅商户会员可选
}

type SplitRuleResultItem struct {
	MemberNo string `json:"member_no"` // 分账参与方在规则中的角色
	Amount   int    `json:"amount"`    // 分账金额
}

type SplitRules struct {
	Rules  []SplitRuleItem   `json:"rules"`
	Params map[string]string `json:"params"`
}

type SplitRuleItem struct {
	Expression       string `json:"expression"`
	RoleAlias        string `json:"role_alias"`
	RoleCode         string `json:"role_code"`
	Remark           string `json:"remark"`
	SpecialAccountNo string `json:"special_account_no,omitempty"`
}

// CreateSplitRuleParams http://47.110.246.50:6524/docs/qzt/qzt-1che8kcosoahl
type CreateSplitRuleParams struct {
	Name    string     `json:"name"`             // 分账规则名称
	Rules   SplitRules `json:"rules"`            // 规则
	Remark  string     `json:"remark,omitempty"` // 备注
	BackUrl string     `json:"back_url"`         // 后台通知地址
}

// CreateSplitRuleResult http://47.110.246.50:6524/docs/qzt/qzt-1che8kcosoahl
type CreateSplitRuleResult struct {
	RuleNo string `json:"rule_no"` // 分账规则编号
	Status int    `json:"status"`  // 分账规则状态,0:待审核,1:有效,2:停用,3:审核失败,9:删除
}

// RemoveSplitRuleParams http://47.110.246.50:6524/docs/qzt/qzt-1che8laiklcr0
type RemoveSplitRuleParams struct {
	RuleNo string `json:"rule_no"`          // 分账规则编号
	Remark string `json:"remark,omitempty"` // 备注
}

// RemoveSplitRuleResult http://47.110.246.50:6524/docs/qzt/qzt-1che8laiklcr0
type RemoveSplitRuleResult struct {
	RuleNo string `json:"rule_no"` // 分账规则编号
}

// StopSplitRuleParams http://47.110.246.50:6524/docs/qzt/qzt-1che8m6mkbe0e
type StopSplitRuleParams struct {
	RuleNo string `json:"rule_no"`          // 分账规则编号
	Remark string `json:"remark,omitempty"` // 备注
}

// StopSplitRuleResult http://47.110.246.50:6524/docs/qzt/qzt-1che8m6mkbe0e
type StopSplitRuleResult struct {
	RuleNo string `json:"rule_no"` // 分账规则编号
}

// ResumeSplitRuleParams http://47.110.246.50:6524/docs/qzt/qzt-1che8n27q090j
type ResumeSplitRuleParams struct {
	RuleNo string `json:"rule_no"`          // 分账规则编号
	Remark string `json:"remark,omitempty"` // 备注
}

// ResumeSplitRuleResult http://47.110.246.50:6524/docs/qzt/qzt-1che8n27q090j
type ResumeSplitRuleResult struct {
	RuleNo string `json:"rule_no"` // 分账规则编号
}

// GetSplitRuleParams http://47.110.246.50:6524/docs/qzt/qzt-1che8nr3dpnul
type GetSplitRuleParams struct {
	RuleNo string `json:"rule_no"` // 分账规则编号
}

// GetSplitRuleResult http://47.110.246.50:6524/docs/qzt/qzt-1che8nr3dpnul
type GetSplitRuleResult struct {
	Rules  SplitRules `json:"rules"`  // 规则
	Status int        `json:"status"` // 分账规则状态,0:待审核,1:有效,2:停用,3:审核失败,9:删除
}

type EvaluateSplitRules struct {
	RuleId  string            `json:"rule_id"` // 规则ID
	Params  map[string]int    `json:"params"`  // 规则中的参数值 // TODO
	Members map[string]string `json:"members"` // 分账参与方在规则中的角色
}

// EvaluateSplitRuleParams http://47.110.246.50:6524/docs/qzt/qzt-1che8otbrihmi
type EvaluateSplitRuleParams struct {
	RuleData []EvaluateSplitRules `json:"rule_data"` // 规则数据,可包含多个规则
}

// EvaluateSplitRuleResult http://47.110.246.50:6524/docs/qzt/qzt-1che8otbrihmi
type EvaluateSplitRuleResult struct {
	Results []struct {
		MemberNo    string `json:"member_no"`    // 分账的会员ID
		TotalAmount int    `json:"total_amount"` // 分账的金额
	} `json:"results"` // 分账结果
}
