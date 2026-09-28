package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 发起进件
// TODO: 待完成
func (s *Client) CreateAccount(ctx context.Context, req *youzan_model.CreateOrderReq) (res *youzan_model.CreateOrderRes, err error) {
	youzanRes, err := s.Request[youzan_model.CreateOrderRes](ctx, req, "youzan.pay.unified.partner.payopen_apply", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

// 查询进件结果
func (s *Client) QueryCreateAccountStatus(ctx context.Context, req *youzan_model.QueryCreateAccountStatusReq) (res *youzan_model.QueryCreateAccountStatusRes, err error) {
	youzanRes, err := s.Request[youzan_model.QueryCreateAccountStatusRes](ctx, req, "youzan.pay.open.user.account.status_query", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}
