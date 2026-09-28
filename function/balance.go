package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 查询商户余额
func (s *Client) BalanceQuery(ctx context.Context, req *youzan_model.BalanceReq) (res *youzan_model.BalanceRes, err error) {
	resOrigin, err := s.Request[youzan_model.BalanceRes](ctx, req, "youzan.pay.open.merchant.account.balance_query", "1.0.0")
	if err != nil {
		return
	}
	res = &resOrigin.Data
	return
}
