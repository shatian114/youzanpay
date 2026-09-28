package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 查询有赞侧子商户列表
func (s *Client) SubMchList(ctx context.Context, req *youzan_model.ListSubMchReq) (res *youzan_model.ListSubMchRes, err error) {
	req.MchId = s.MchId
	resOrigin, err := s.Request[youzan_model.ListSubMchRes](ctx, req, "youzan.pay.open.service.list_submchinfo", "1.0.0")
	if err != nil {
		return
	}
	res = &resOrigin.Data
	return
}
