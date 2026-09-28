package youzan_model

type BalanceReq struct {
	SubMchId string `json:"sub_mch_id"`
}

type BalanceRes struct {
	// BalanceMap map[string]int64 `json:"balance_map"`
	BalanceDetail BalanceDetail `json:"balance_map"`
}

type BalanceDetail struct {
	AVAILABLE           int64 `json:"AVAILABLE"`
	FREEZE              int64 `json:"FREEZE"`
	PROFIT_SHARING_ACCT int64 `json:"PROFIT_SHARING_ACCT"`
}
