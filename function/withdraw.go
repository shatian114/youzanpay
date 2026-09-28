package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 发起提现
func (s *Client) WithdrawCreate(ctx context.Context, req *youzan_model.WithDrawReq) (res *youzan_model.WithDrawRes, err error) {
	resOrigin, err := s.Request[youzan_model.WithDrawRes](ctx, req, "youzan.pay.open.merchant.withdraw_apply", "1.0.2")
	if err != nil {
		return
	}
	res = &resOrigin.Data
	return
}

// 修改结算卡信息
func (s *Client) WithdrawQuery(ctx context.Context, req *youzan_model.WithDrawQueryReq) (res *youzan_model.WithDrawQueryRes, err error) {
	resOrigin, err := s.Request[youzan_model.WithDrawQueryRes](ctx, req, "youzan.pay.open.merchant.withdraw_query", "1.0.0")
	if err != nil {
		return
	}
	res = &resOrigin.Data
	return
}
