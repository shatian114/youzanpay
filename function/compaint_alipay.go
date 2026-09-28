package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 查询支付宝投诉详情
func (s *Client) ComplaintAlipayDetail(ctx context.Context, req *youzan_model.ComplaintAlipayDetailReq) (res *youzan_model.ComplaintAlipayDetailRes, err error) {
	youzanRes, err := s.Request[youzan_model.ComplaintAlipayDetailRes](ctx, req, "youzan.pay.open.alipay.complaint_get", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

// 回复支付宝投诉
func (s *Client) ComplaintAlipayReply(ctx context.Context, req *youzan_model.ComplaintAlipayReplyReq) (err error) {
	_, err = s.Request[any](ctx, req, "youzan.pay.open.alipay.complaint_create", "1.0.0")
	if err != nil {
		return
	}

	return
}

// 完结支付宝投诉
func (s *Client) ComplaintAlipayComplete(ctx context.Context, req *youzan_model.ComplaintAlipayCompleteReq) (err error) {
	_, err = s.Request[any](ctx, req, "youzan.pay.open.alipay.complaint_operate", "1.0.1")
	if err != nil {
		return
	}

	return
}
