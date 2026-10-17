package youzan_model

type ModityShortNameReq struct { // 修改商家简称请求
	SubMchId          string `json:"sub_mch_id,omitempty"`          // 有赞发放的商户号
	ChannelSubMchId   string `json:"channel_sub_mch_id,omitempty"`  // 对应的子商户号
	MerchantShortName string `json:"merchant_short_name,omitempty"` // 修改的简称
}
