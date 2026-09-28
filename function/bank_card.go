package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 查询结算卡信息
func (s *Client) BankCardQuery(ctx context.Context, req *youzan_model.BankCardReq) (res *youzan_model.BankCardInfo, err error) {
	resOrigin, err := s.Request[youzan_model.BankCardInfo](ctx, req, "youzan.pay.open.merchant.bank.account_query", "1.0.0")
	if err != nil {
		return
	}
	res = &resOrigin.Data
	return
}

// 修改结算卡信息
func (s *Client) BankCardUpdate(ctx context.Context, req *youzan_model.BankCardUpdateReq) (res *youzan_model.BankCardUpdateRes, err error) {
	resOrigin, err := s.Request[youzan_model.BankCardUpdateRes](ctx, req, "youzan.pay.open.merchant.bank.account_modify", "1.0.0")
	if err != nil {
		return
	}
	res = &resOrigin.Data
	return
}
