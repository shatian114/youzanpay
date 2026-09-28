package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 查询银行支行联行号
func (s *Client) SearchBankBranchCode(ctx context.Context, req *youzan_model.BankBranchSearchReq) (res *[]youzan_model.BankBranchSearchRes, err error) {
	req.MchId = s.MchId
	resOrigin, err := s.Request[[]youzan_model.BankBranchSearchRes](ctx, req, "youzan.pay.open.bankcard.bankcode_search", "1.0.0")
	if err != nil {
		return
	}
	res = &resOrigin.Data
	return
}
