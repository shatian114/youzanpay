package youzan_model

type FeeCreateReq struct {
	SubMchId string `json:"sub_mch_id"` // 有赞商户号。不能为空。
	FeeRate  int64  `json:"fee_rate"`   // 费率整数值。例如 1.2% 传 120，0.6% 传 60，0.65% 传 65；目前支持的费率：从60（既0.6%）到400（既4%），支持百分号内2位小数，例如0.65%，不支持0.651%。
}

type FeeQueryReq struct {
	SubMchId string `json:"sub_mch_id"` // 有赞商户号。不能为空。
}

type FeeQueryRes struct {
	FeeRate         int    `json:"fee_rate"`          // 费率整数值。例如 1.2% 传 120，0.6% 传 60，0.65% 传 65；目前支持的费率：从60（既0.6%）到400（既4%），支持百分号内2位小数，例如0.65%，不支持0.651%。
	AgreeEffectDate string `json:"agree_effect_date"` // 协议生效日期，格式为 yyyy-MM-dd。
	AgreeExpireDate string `json:"agree_expire_date"` // 协议失效日期，格式为 yyyy-MM-dd。
}
