package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 修改商户简称 https://qima.feishu.cn/docx/ZPo4dm5WMofi90x5917c091LnOg
func (s *Client) ModityShortName(ctx context.Context, req *youzan_model.ModityShortNameReq) (err error) {
	_, err = s.Request[any](ctx, req, "youzan.pay.open.merchant.short.name_modify", "1.0.0")
	if err != nil {
		return
	}
	return
}
