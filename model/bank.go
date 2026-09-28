package youzan_model

type BankBranchSearchReq struct { // 查询支行名称和联行号
	MchId      string `json:"mch_id"`
	BankName   string `json:"bank_name"`   // 不支持模糊匹配
	BranchName string `json:"branch_name"` // 支持模糊匹配
}

type BankBranchSearchRes struct {
	BankName   string `json:"bank_name"`
	BranchName string `json:"branch_name"`
	BankCode   string `json:"bank_code"` // 联行号
}
