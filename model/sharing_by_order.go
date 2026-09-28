package youzan_model

// 基于订单分账
type ProfitSharingRecipientParamExt struct { //分账接收方参数里的Extra
	AttachmentMsg string `json:"attachmentMsg"`
}

type ProfitSharingRecipientParam struct { // 分账接收方参数
	SubOutBizNo     string                         `json:"sub_out_biz_no"`    // 是否必填: 是;长度：64;分账业务子单号;一次分账可以有多个子单号，需在商户号维度内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的ps：为了防止同一接收方重复分账，如果接口调用报错，或未明确拿到分账成功结果请原单重试！！！！！如果分账状态一直处于分账处理中状态，请联系有赞支付解决，请勿换单重试～
	RecipientUserNo string                         `json:"recipient_user_no"` // 是否必填: 是;长度：32;分账接收方商户号;分账接收方商户号，该商户号通过《进件申请接口》开出，详见技术研发阅读-01-管理特约商户开发文档
	RecipientType   string                         `json:"recipient_type"`    // 是否必填: 是;长度：32;分账接收方类型;子商户商户号，请使用该类型;   MCH_ID; 零钱商户号openId，定制，请勿使用;   YZ_OPEN_ID
	Amount          int64                          `json:"amount"`            // 是否必填: 是;长度：0;分账金额;分账金额（单位/分）
	Remark          string                         `json:"remark"`            // 是否必填: 是;长度：256;备注;备注
	Extra           ProfitSharingRecipientParamExt `json:"extra,omitempty"`   // 是否必填: 否;长度：0;扩展字段;attachmentMsg： 备注信息 用于一本帐备注信息展示 长度限制：128"extra": {;      "attachmentMsg":"活动时间2026-06-12 23:59:59 | 活动时间2026-06-12 23:59:59 |活动时间2026-06-12 23:59:59"; }
}

type ProfitSharingRecipientReq struct { // 分账请求
	SubMchId             string                        `json:"sub_mch_id"`             // 是否必填: 是;长度：0;子商户号;钱从这个商户号中分出
	OrderNo              string                        `json:"order_no"`               // 是否必填: 是;长度：0;业务订单号;请传《统一支付》接口中的outer_trade_no
	OutBizNo             string                        `json:"out_biz_no"`             // 是否必填: 是;长度：0;分账业务单号;在商户号维度内需要保证分账业务单号唯一，由调用方（服务商）自己生成并传入，不是有赞分配的ps：为了防止同一接收方重复分账，如果接口调用报错，或未明确拿到分账成功结果请原单重试！！！！！如果分账状态一直处于分账处理中状态，请联系有赞支付解决，请勿换单重试～
	Currency             string                        `json:"currency,omitempty"`     // 是否必填: 否;长度：0;币种;默认人民币：CNY
	TotalAmount          int64                         `json:"total_amount"`           // 是否必填: 是;长度：0;分账总金额;分账总金额 单位/分
	ProfitSharingDetails []ProfitSharingRecipientParam `json:"profit_sharing_details"` // 是否必填: 是;长度：0;分账接收方列表;数据类型详见 “分账接收方参数”
}

type ProfitSharingRecipientRes struct { // 分账响应
	SubMchId             string                      `json:"sub_mch_id"`             // 是否必填: 是;长度：0;商户号;分出方子商户号
	OrderNo              string                      `json:"order_no"`               // 是否必填: 是;长度：0;业务单号;对应《统一支付》接口中的outer_trade_no
	OutBizNo             string                      `json:"out_biz_no"`             // 是否必填: 是;长度：0;分账业务订单号;在商户号维度内需要保证分账业务单号唯一，由调用方（服务商）自己生成并传入，不是有赞分配的
	Currency             string                      `json:"currency"`               // 是否必填: 是;长度：0;币种;默认人民币：CNY
	TotalAmount          int64                       `json:"total_amount"`           // 是否必填: 是;长度：0;分账总金额;分账总金额 单位/分
	FeeType              string                      `json:"fee_type"`               // 是否必填: 是;长度：0;费用类型;费用类型默认为分账：PROFIT_SHARING
	SettleState          string                      `json:"settle_state"`           // 是否必填: 是;长度：0;分账状态;1：处理中2：成功3：受理成功99：失败
	ProfitSharingDetails ProfitSharingRecipientParam `json:"profit_sharing_details"` // 是否必填: 是;长度：0;分账结果明细;数据字段见 分账结果明细
}

type ProfitSharingQueryReq struct { // 分账查询请求
	SubMchId string `json:"sub_mch_id"` // 是否必填: 是;长度：0;子商户号;子商户号，服务商模式进件的子商户号
	OrderNo  string `json:"order_no"`   // 是否必填: 是;长度：0;业务订单号;对应《统一支付》接口中的outer_trade_no
	OutBizNo string `json:"out_biz_no"` // 是否必填: 是;长度：0;分账业务单号;在商户号维度内需要保证分账业务单号唯一，由调用方（服务商）自己生成并传入，不是有赞分配的
}

type ProfitSharingQueryRes struct { // 分账查询响应
	SubMchId                     string                      `json:"sub_mch_id"`                      // 是否必填: 是;长度：0;子商户号;分出方商户号，服务商模式进件的商户号
	OrderNo                      string                      `json:"order_no"`                        // 是否必填: 是;长度：0;业务订单号;对应《统一支付》接口中的outer_trade_no
	OutBizNo                     string                      `json:"out_biz_no"`                      // 是否必填: 是;长度：0;分账业务单号;在商户号维度内需要保证分账业务单号唯一，由调用方（服务商）自己生成并传入，不是有赞分配的
	Currency                     string                      `json:"currency"`                        // 是否必填: 是;长度：0;币种;默认 CNY 人民币
	TotalAmount                  int64                       `json:"total_amount"`                    // 是否必填: 是;长度：0;分账总金额;分账总金额 单位/分
	FeeType                      string                      `json:"fee_type"`                        // 是否必填: 是;长度：0;费用类型;费用类型分账默认为：PROFIT_SHARING
	SettleState                  int64                       `json:"settle_state"`                    // 是否必填: 是;长度：0;结算状态;1：处理中2：成功3：受理成功99：失败
	ProfitSharingDetailResponses ProfitSharingRecipientParam `json:"profit_sharing_detail_responses"` // 是否必填: 是;长度：0;分账明细结果;结果明细内容见 分账明细结果
}
