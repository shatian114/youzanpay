package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 查询有赞账单
func (s *Client) RecordQuery(ctx context.Context, req *youzan_model.RecordQueryReq) (res *youzan_model.RecordQueryRes, err error) {
	youzanRes, err := s.Request[youzan_model.RecordQueryRes](ctx, req, "youzan.pay.open.merchant.account.record_query", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}
