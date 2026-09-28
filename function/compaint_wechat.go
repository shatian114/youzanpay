package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 查询微信投诉列表
func (s *Client) ComplaintWechatList(ctx context.Context, req *youzan_model.ComplaintWechatListReq) (res *youzan_model.ComplaintWechatListRes, err error) {
	youzanRes, err := s.Request[youzan_model.ComplaintWechatListRes](ctx, req, "youzan.pay.merchant.wx.complaint.page_list", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

// 查询微信投诉详情
func (s *Client) ComplaintWechatDetail(ctx context.Context, req *youzan_model.ComplaintWechatDetailReq) (res *youzan_model.ComplaintWechatRes, err error) {
	youzanRes, err := s.Request[youzan_model.ComplaintWechatRes](ctx, req, "youzan.pay.merchant.wx.complaint.detail_get", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

// 查询微信投诉协商历史
func (s *Client) ComplaintWechatHistory(ctx context.Context, req *youzan_model.ComplaintWechatHistoryReq) (res *youzan_model.ComplaintWechatHistoryRes, err error) {
	youzanRes, err := s.Request[youzan_model.ComplaintWechatHistoryRes](ctx, req, "youzan.pay.merchant.wx.complaint.history_query", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

// 回复用户
func (s *Client) ComplaintWechatReply(ctx context.Context, req *youzan_model.ComplaintWechatReplyReq) (err error) {
	_, err = s.Request[any](ctx, req, "youzan.pay.merchant.wx.complaint.confirm_reply", "1.0.0")
	if err != nil {
		return
	}

	return
}

// 给系统反馈处理完成
func (s *Client) ComplaintWechatComplete(ctx context.Context, req *youzan_model.ComplaintWechatCompleteReq) (err error) {
	_, err = s.Request[any](ctx, req, "youzan.pay.merchant.wx.complaint.confirm_complete", "1.0.0")
	if err != nil {
		return
	}

	return
}

// 处理退款请求，是否退款
func (s *Client) ComplaintWechatRefund(ctx context.Context, req *youzan_model.ComplaintWechatRefundReq) (err error) {
	_, err = s.Request[any](ctx, req, "youzan.pay.merchant.wx.complaint.progress_update", "1.0.0")
	if err != nil {
		return
	}

	return
}
