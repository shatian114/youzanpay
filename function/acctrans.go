package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

func (s *Client) QueryAcctrans(ctx context.Context, req *youzan_model.AcctransQueryReq) (res *youzan_model.AcctransQueryRes, err error) {
	resOrigin, err := s.Request[youzan_model.AcctransQueryRes](ctx, req, "youzan.pay.merchant.paycloud.acctrans_query", "1.0.2")
	if err != nil {
		return
	}
	res = &resOrigin.Data
	return
}
