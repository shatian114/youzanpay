package youzan_model

type BankCardReq struct {
	SubMchId string `json:"sub_mch_id"` // 子商户号
}

type BankCardUpdateReq struct {
	SubMchId string `json:"sub_mch_id"` // 子商户号
	BankCardInfo
}

type BankCardUpdateRes struct {
	Succ bool `json:"succ"` // 是否绑定
}
