package youzan_model

type AcctransQueryReq struct { // 查询商户可结算余额请求
	SubMchId     string   `json:"sub_mch_id"`    // 又名特约商户、收款商户
	AccountTypes []string `json:"account_types"` //账户类型固定值：["WAIT_TO_SETTLE"]
}

type AcctransQueryRes struct { // 查询商户可结算余额响应
	Balance       int64          `json:"balance"`        //账户余额，金额单位：分
	Exists        bool           `json:"exists"`         //查询的账户是否存在
	AccountDetail map[string]any `json:"account_detail"` // 按账户类型返回账户明细，当前账户类型为 WAIT_TO_SETTLE
}
