package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

func (s *Client) CreateOrder(ctx context.Context, req *youzan_model.CreateOrderReq) (res *youzan_model.CreateOrderRes, err error) {
	youzanRes, err := s.Request[youzan_model.CreateOrderRes](ctx, req, "youzan.pay.unified.partner.payopen_apply", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

func (s *Client) QueryOrder(ctx context.Context, req *youzan_model.QueryOrderReq) (res *youzan_model.QueryOrderRes, err error) {
	youzanRes, err := s.Request[youzan_model.QueryOrderRes](ctx, req, "youzan.pay.unified.partner.payopen_query", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}
