package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 开通支付能力
func (s *Client) OpenPay(ctx context.Context, req *youzan_model.OpenPayReq) (err error) {
	_, err = s.Request[any](ctx, req, "youzan.pay.open.merchant.unified_create", "1.0.0")
	if err != nil {
		return
	}
	return
}

// 查询支付能力开通进度
func (s *Client) OpenPayQuery(ctx context.Context, req *youzan_model.OpenPayQueryReq) (res *youzan_model.OpenPayQueryRes, err error) {
	youzanRes, err := s.Request[youzan_model.OpenPayQueryRes](ctx, req, "youzan.pay.open.merchant.unified_query", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}
