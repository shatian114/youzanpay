package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 设置商户费率
func (s *Client) FeeCreate(ctx context.Context, req *youzan_model.FeeCreateReq) (err error) {
	_, err = s.Request[any](ctx, req, "youzan.pay.open.fee.sign_create", "1.0.5")
	if err != nil {
		return
	}
	return
}

// 查询商户费率
func (s *Client) FeeQuery(ctx context.Context, req *youzan_model.FeeQueryReq) (res *youzan_model.FeeQueryRes, err error) {
	youzanRes, err := s.Request[youzan_model.FeeQueryRes](ctx, req, "youzan.pay.open.fee.sign_get", "1.0.3")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}
