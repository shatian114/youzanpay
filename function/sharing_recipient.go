package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 添加分账关系，基于账户模式，https://qima.feishu.cn/docx/TC7Sdm20YopyMJxUziaczeDYnBg
func (s *Client) AddSharingRecipient(ctx context.Context, req *youzan_model.SharingRecipientAddReq) (err error) {
	_, err = s.Request[any](ctx, req, "youzan.pay.open.profitsharing.recipient_add", "1.0.0")
	if err != nil {
		return
	}
	return
}
