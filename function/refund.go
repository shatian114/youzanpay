package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

func (s *Client) CreateRefund(ctx context.Context, req *youzan_model.RefundReq) (res *youzan_model.RefundRes, err error) {
	youzanRes, err := s.Request[youzan_model.RefundRes](ctx, req, "youzan.pay.unified.partner.refundopen_apply", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

func (s *Client) QueryRefund(ctx context.Context, req *youzan_model.RefundQueryReq) (res *youzan_model.RefundQueryRes, err error) {
	youzanRes, err := s.Request[youzan_model.RefundQueryRes](ctx, req, "youzan.pay.unified.partner.refundopen_query", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}
