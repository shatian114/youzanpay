package youzan_model

type RefundReq struct { // 有赞退款请求
	OutRefundNo string `json:"out_refund_no"`    // 是否必填: 是;长度：0;商户退款单号;商户号内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的例：2305181535202621620585
	TradeNo     string `json:"trade_no"`         // 是否必填: 是;长度：0;支付收单号;有赞生分配，本次下单支付成功有赞支付单编号，用于在有赞支付域内标记订单的支付记录例：2305181616372621460130
	RefundFee   int64  `json:"refund_fee"`       // 是否必填: 是;长度：0;退款金额;退款金额 > 0，退款金额不大于支付金额单位：分
	SubMchId    string `json:"sub_mch_id"`       // 是否必填: 是;长度：18;收款商户号;有赞支付商户号, 有赞支付分配
	Memo        string `json:"memo,omitempty"`   // 是否必填: 否;长度：0;退款备注信息;退款备注信息
	NotifyUrl   string `json:"notify_url"`       // 是否必填: 是;长度：0;退款结果回调地址;用于退款成功结果通知，如果不传则不进行回调通知必须为https地址。不携带查询参数, 链接例：https://xxxx.cloud.com/service_api/pay/notif路径商家自定义设置，域名需要报备给有赞，否则无法正常接收回调请求！！！！！！
	Attach      string `json:"attach,omitempty"` // 是否必填: 否;长度：0;附加信息;附加数据，在退款响应原样返回，可作为自定义参数使用，例：xxxxxx
}

type RefundRes struct { // 有赞退款响应
	OutRefundNo  string `json:"out_refund_no"` // 是否必填: 是;长度：0;外部单号;商户号内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的例：81535202621620585
	TradeNo      string `json:"trade_no"`      // 是否必填: 是;长度：0;支付收单号;有赞生分配，本次下单支付成功有赞支付单编号，用于在有赞支付域内标记订单的支付记录例：2305181616372621460130
	RefundStatus string `json:"refund_status"` // 是否必填: 是;长度：0;退款状态;值：APPLYING : "退款申请中"APPLY_FAIL："退款申请失败" APPLY_SUCCESS : "退款申请成功"SUCCESS : "退款成功"ps：状态为APPLY_FAIL，明确业务失败，请换单重试；状态为退款申请中或接口调用报错请原单重试！！！！！若是申请成功请进行退款查询或退款结果回调拿到最终退款状态（SUCCESS）
	RefundFee    int64  `json:"refund_fee"`    // 是否必填: 是;长度：0;退款金额;退款金额 > 0单位：分
}

type RefundQueryReq struct { // 有赞退款查询请求
	OutRefundNo string `json:"out_refund_no"` // 是否必填: 是;长度：0;商户退款单号;商户号内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的例：2305181535202621620585
	TradeNo     string `json:"trade_no"`      // 是否必填: 是;长度：0;支付收单号;有赞生分配，本次下单支付成功有赞支付单编号，用于在有赞支付域内标记订单的支付记录例：2305181616372621460130
	SubMchId    string `json:"sub_mch_id"`    // 是否必填: 是;长度：18;收款商户号;有赞支付商户号, 有赞支付分配
}
type RefundQueryRes struct { // 有赞退款查询响应
	OutRefundNo         string `json:"out_refund_no"`                // 是否必填: 是;长度：0;商户退款单号;商户号内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的例：2305181535202621620585
	TradeNo             string `json:"trade_no"`                     // 是否必填: 是;长度：0;支付收单号;有赞生分配，本次下单支付成功有赞支付单编号，用于在有赞支付域内标记订单的支付记录例：2305181616372621460130
	ExternalOutRefundNo string `json:"external_out_refund_no"`       // 是否必填: 是;长度：0;三方退款商户单号;有赞向三方渠道(微信/支付宝)发起退款时，使用的商户退款单号，由有赞分配例：2305181616372621460130
	ExternalRefundNo    string `json:"external_refund_no,omitempty"` // 是否必填: 否;长度：0;三方退款流水号;微信/支付宝在处理这笔退款时生成并返回，由微信/支付宝分配，不由有赞分配例：2305181616372621460130
	RefundStatus        string `json:"refund_status"`                // 是否必填: 是;长度：0;退款状态;退款状态：APPLYING : "退款申请中"APPLY_FAIL："退款申请失败" APPLY_SUCCESS : "退款申请成功"SUCCESS : "退款成功"ps：状态为APPLY_FAIL，明确业务失败，请换单重试；状态为退款申请中或抛异常请原单重试；若是申请成功请进行退款查询或退款结果回调拿到最终退款状态（SUCCESS）
	RefundFee           int64  `json:"refund_fee"`                   // 是否必填: 是;长度：0;退款金额;退款金额 > 0单位：分
	CreateTime          int64  `json:"create_time"`                  // 是否必填: 是;长度：0;退款创建时间;退款受理时间
	SuccessTime         int64  `json:"success_time,omitempty"`       // 是否必填: 否;长度：0;退款成功时间;退款成功时间，当退款状态为退款成功时有返回。
	ErrorReason         string `json:"error_reason,omitempty"`       // 是否必填: 否;长度：0;退款失败原因;退款失败原因
	Attach              string `json:"attach,omitempty"`             // 是否必填: 否;长度：0;附加信息;附加数据，在退款响应原样返回，可作为自定义参数使用
}
