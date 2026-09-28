package youzan_model

// 基于余额分账
type SharingByAccountReqRecipient struct { // 基于账户分账请求的接收方
	SubOutBizNo     string `json:"sub_out_biz_no"`    // 是否必填: 是;长度：64;分账业务子单号;分账业务子单号，每一笔分账的业务单号需要保持唯一如果未获取到分账结果，请原单重试ps：为了防止同一接收方重复分账，如果接口调用报错，或未明确拿到分账成功结果请原单重试！！！！！如果分账状态一直处于分账处理中状态，请联系有赞支付解决，请勿换单重试～
	RecipientUserNo string `json:"recipient_user_no"` // 是否必填: 是;长度：32;分账接收方商户号;分账接收方商户号 MCH_ID（默认）; 零钱商户号openId YZ_OPEN_ID（定制专用）
	RecipientType   string `json:"recipient_type"`    // 是否必填: 是;长度：32;分账接收方类型;子商户商户号;   MCH_ID（默认）; 有赞零钱;   YZ_OPEN_ID（定制专用）
	Amount          int64  `json:"amount"`            // 是否必填: 是;长度：0;分账金额;分账金额（单位/分）
	Remark          string `json:"remark,omitempty"`  // 是否必填: 否;长度：256;备注;备注
}

type SharingByAccountReq struct { // 基于账户分账请求
	SubMchId      string                         `json:"sub_mch_id"`             // 是否必填: 是;长度：0;子商户号;分账商户号; 子商户号，服务商模式进件的子商户号
	OrderNo       string                         `json:"order_no"`               // 是否必填: 是;长度：0;业务订单号;业务订单号外部传入，需要保证商户号内唯一
	OutBizNo      string                         `json:"out_biz_no"`             // 是否必填: 是;长度：0;分账业务单号;分账业务订单号外部传入，需要保证商户号内唯一如果未获取到分账结果，请原单重试ps：为了防止同一接收方重复分账，如果接口调用报错，或未明确拿到分账成功结果请原单重试！！！！！如果分账状态一直处于分账处理中状态，请联系有赞支付解决，请勿换单重试～
	Currency      string                         `json:"currency,omitempty"`     // 是否必填: 否;长度：0;币种;默认人民币：CNY
	TotalAmount   int64                          `json:"total_amount"`           // 是否必填: 是;长度：0;分账总金额;分账总金额（单位/分）
	RecipientList []SharingByAccountReqRecipient `json:"profit_sharing_details"` // 是否必填: 是;长度：0;分账接收方列表;数据类型详见 “分账接收方参数”
}

type SharingByAccountResDetails struct { // 基于账户分账请求响应结果明细
	SubOutBizNo       string `json:"sub_out_biz_no"`      // 是否必填: 是;长度：0;分账业务子单号;分账业务子单号
	RecipientUserNo   string `json:"recipient_user_no"`   // 是否必填: 是;长度：32;分账接收方账号;分账接收方账号
	RecipientUserType string `json:"recipient_user_type"` // 是否必填: 是;长度：32;分账接收方类型;子商户商户号; MCH_ID
	Amount            int64  `json:"amount"`              // 是否必填: 是;长度：0;分账金额;分账金额（单位/分）
}
type SharingByAccountRes struct { // 基于账户分账请求响应
	SubMchId    string                       `json:"sub_mch_id"`                      // 是否必填: 是;长度：0;商户号;子商户号，服务商模式进件的子商户号
	OrderNo     string                       `json:"order_no,omitempty"`              // 是否必填: 是;长度：0;业务单号;业务单号
	OutBizNo    string                       `json:"out_biz_no"`                      // 是否必填: 是;长度：0;分账业务订单号;分账业务订单号
	Currency    string                       `json:"currency"`                        // 是否必填: 是;长度：0;币种;默认人民币：CNY
	TotalAmount int64                        `json:"total_amount"`                    // 是否必填: 是;长度：0;分账总金额;分账总金额（单位/分）
	FeeType     string                       `json:"fee_type"`                        // 是否必填: 是;长度：0;费用类型;费用类型默认为资金分账：FUND_SHARING
	SettleState int64                        `json:"settle_state"`                    // 是否必填: 是;长度：0;分账状态;1：处理中2：成功3：受理成功99：失败
	Details     []SharingByAccountResDetails `json:"profit_sharing_detail_responses"` // 是否必填: 是;长度：0;分账结果明细;数据字段见 分账结果明细
}
type SharingByAccountQueryReq struct { // 查询基于账户分账请求
	SubMchId string `json:"sub_mch_id"` // 是否必填: 是;长度：0;商户号;子商户号，服务商模式进件的子商户号
	OrderNo  string `json:"order_no"`   // 是否必填: 是;长度：0;业务单号;业务订单号，此次业务订单号；请和分账时闪送的order_no一致
	OutBizNo string `json:"out_biz_no"` // 是否必填: 是;长度：0;分账业务订单号;分账业务单号，此次分账业务单号；请和分账时闪送的out_biz_no一致
}

type SharingByAccountRefundReq struct { // 有赞基于账户的分账退回请求
	OrderNo             string `json:"order_no"`                // 是否必填: 是;长度：64;业务订单号;业务订单号
	OutBizNo            string `json:"out_biz_no"`              // 是否必填: 是;长度：64;退分账业务单号;退分账业务单号如果未获取到分账结果，请原单重试
	RelationOutBizNo    string `json:"relation_out_biz_no"`     // 是否必填: 是;长度：64;退分账关联的分账单号;关联分账单号
	RelationSubOutBizNo string `json:"relation_sub_out_biz_no"` // 是否必填: 是;长度：64;退分账关联的分账子单号;关联分账子单号
	SubMchId            string `json:"sub_mch_id"`              // 是否必填: 是;长度：32;退分账接收方商户号;即分账出款方二级子商户号为服务商在有赞支付服务商下二级子商户商户号
	UserNo              string `json:"user_no"`                 // 是否必填: 是;长度：32;退分账出款方账号;退分账出款方账号用户类型为商户 为商户号
	UserType            string `json:"user_type"`               // 是否必填: 是;长度：32;退分账出款方用户类型;用户类型; 商户号MCH_ID
	Currency            string `json:"currency"`                // 是否必填: 是;长度：32;币种;默认币种 CNY
	Amount              int64  `json:"amount"`                  // 是否必填: 是;长度：0;退分账金额;退分账金额（单位/分）不能大于 分账金额
	Remark              string `json:"remark,omitempty"`        // 是否必填: 否;长度：32;备注;备注字段用于对帐单展示交易名称
}

type SharingByAccountRefundRes struct { // 有赞基于账户的分账退回响应
	SettleState int64 `json:"settle_state"` // 是否必填: 是;长度：0;分账状态;1：处理中2：成功3：受理成功99：失败
}

type SharingByAccountRefundQueryReq struct { // 查询基于账户分账退回查询请求
	SubMchId string `json:"sub_mch_id"` // 是否必填: 是;长度：0;商户号;子商户号，服务商模式进件的子商户号
	OrderNo  string `json:"order_no"`   // 是否必填: 是;长度：0;业务单号;业务订单号，此次业务订单号；请和分账时闪送的order_no一致
	OutBizNo string `json:"out_biz_no"` // 是否必填: 是;长度：0;分账业务订单号;分账业务单号，此次分账业务单号；请和分账时闪送的out_biz_no一致
}

type SharingByAccountRefundQueryRes struct { // 有赞基于账户的分账退回响应
	SettleState int64 `json:"settle_state"` // 是否必填: 是;长度：0;分账状态;1：处理中2：成功3：受理成功99：失败
}
