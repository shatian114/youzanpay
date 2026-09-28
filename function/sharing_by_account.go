package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

func (s *Client) CreateSharingByAccount(ctx context.Context, req *youzan_model.SharingByAccountReq) (res *youzan_model.SharingByAccountRes, err error) {
	youzanRes, err := s.Request[youzan_model.SharingByAccountRes](ctx, req, "youzan.pay.settle.open.create_fundsharing", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

func (s *Client) QuerySharingByAccount(ctx context.Context, req *youzan_model.SharingByAccountQueryReq) (res *youzan_model.SharingByAccountRes, err error) {
	youzanRes, err := s.Request[youzan_model.SharingByAccountRes](ctx, req, "youzan.pay.settle.open.query_fundsharing", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

func (s *Client) RefundSharingByAccount(ctx context.Context, req *youzan_model.SharingByAccountRefundReq) (res *youzan_model.SharingByAccountRefundRes, err error) {
	youzanRes, err := s.Request[youzan_model.SharingByAccountRefundRes](ctx, req, "youzan.pay.settle.open.create_fundsharingrefund", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

func (s *Client) RefundSharingByAccountQuery(ctx context.Context, req *youzan_model.SharingByAccountRefundQueryReq) (res *youzan_model.SharingByAccountRefundQueryRes, err error) {
	youzanRes, err := s.Request[youzan_model.SharingByAccountRefundQueryRes](ctx, req, "youzan.pay.settle.open.query_fundsharingrefund", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}
