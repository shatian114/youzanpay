package youzan_model

// 查询账单
type RecordQueryReq struct { // 有赞订单查询
	SubMchId  string `json:"sub_mch_id"`         // 是否必填: 是;长度：0;二级子商户号;有赞支付商户号；又名特约商户
	OrderNo   string `json:"order_no,omitempty"` // 是否必填: 否;长度：64;订单号;外部订单号
	BizType   string `json:"biz_type,omitempty"` // 是否必填: 否;长度：64;业务类型;业务类型：交易手续费：TRADE_CHARGES订单入账：SETTLEMENT退款： REFUND开放平台-分账：PROFIT_SHARING开放平台-分账退回：PROFIT_SHARING_REFUND提现：WITHDRAW提现失败退回：CASH_FAIL_RETURN
	StartDate string `json:"start_date"`         // 是否必填: 是;长度：19;开始时间;按照标准格式格式yyyy-MM-dd HH:mm:ss
	EndDate   string `json:"end_date"`           // 是否必填: 是;长度：19;结束时间;按照标准格式格式yyyy-MM-dd HH:mm:ss
	Page      int64  `json:"page"`               // 是否必填: 是;长度：0;页数;当前页码
	PageSize  int64  `json:"page_size"`          // 是否必填: 是;长度：0;每页条数;每页条数
}

type RecordQueryRes struct {
	RecordResults []RecordQueryResult    `json:"record_results,omitempty"`
	PageInfo      RecordQueryResPageInfo `json:"page_info,omitempty"`
}

type RecordQueryResult struct { // 有赞订单查询响应
	SubMchId        string `json:"sub_mch_id"`        // 是否必填: 是;长度：0;二级子商户号;又名特约商户
	BizType         string `json:"biz_type"`          // 是否必填: 是;长度：0;业务类型;业务类型：交易手续费：TRADE_CHARGES交易手续费退回：TRADE_CHARGES_RETURN订单入账：SETTLEMENT退款： REFUND开放平台-分账：PROFIT_SHARING开放平台-分账退回：PROFIT_SHARING_REFUND提现：WITHDRAW提现失败退回：CASH_FAIL_RETURN
	BizTypeCode     string `json:"biz_type_code"`     // 是否必填: 是;长度：0;业务类型code;业务类型代码
	ChangeMoney     string `json:"change_money"`      // 是否必填: 是;长度：0;动账金额;动账金额
	SettleTime      string `json:"settle_time"`       // 是否必填: 是;长度：0;入账时间;入账时间
	Desc            string `json:"desc"`              // 是否必填: 是;长度：0;描述;交易描述
	OrderNo         string `json:"order_no"`          // 是否必填: 是;长度：0;订单号;订单号
	TradeNo         string `json:"trade_no"`          // 是否必填: 是;长度：0;支付收单号;支付收单号
	PayWaterNo      string `json:"pay_water_no"`      // 是否必填: 是;长度：0;支付流水号;支付流水号
	Amount          string `json:"amount"`            // 是否必填: 是;长度：0;金额;金额
	AccountBalance  string `json:"account_balance"`   // 是否必填: 是;长度：0;账户余额;账户余额
	Inout           int64  `json:"inout"`             // 是否必填: 是;长度：0;收入/支付;1:收入2:支出
	AccountTypeCode string `json:"account_type_code"` // 是否必填: 是;长度：0;变动账户类型;变动账户类型
	Remark          string `json:"remark"`            // 是否必填: 是;长度：0;备注;备注
}

type RecordQueryResPageInfo struct { // 查询有赞侧子商户列表结果页码信息
	Page     int64 `json:"page"`      // 是否必填: 是;长度：0;数据列表;页码
	PageSize int64 `json:"page_size"` // 是否必填: 是;长度：0;分页信息;页大小
	TotalNum int64 `json:"total_num"` // 是否必填: 是;长度：0;数据总数;数据总数
}
