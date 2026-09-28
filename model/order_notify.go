package youzan_model

type OrderNotify struct {
	EventType            string `json:"event_type"` // 时间类型，支付成功: PAY_SUCCESS
	Currency             string `json:"currency"`
	ExpireTime           string `json:"expire_time"`
	ExternalCompletedAt  string `json:"external_completed_at"`
	ExternalOutPaymentNo string `json:"external_out_payment_no"`
	ExternalPaymentNo    string `json:"external_payment_no"`
	Extra                string `json:"extra"`
	OuterTradeNo         string `json:"outer_trade_no"`
	PayAmount            string `json:"pay_amount"`
	PayTool              string `json:"pay_tool"`
	ServiceMchID         string `json:"service_mch_id"`
	Status               string `json:"status"`
	StatusDesc           string `json:"status_desc"`
	SubMchID             string `json:"sub_mch_id"`
	SuccessTime          string `json:"success_time"`
	TradeNo              string `json:"trade_no"`
}
