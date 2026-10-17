package youzan_func

import (
	"context"

	youzan_model "github.com/shatian114/youzanpay/model"
)

// 发起进件
func (s *Client) CreateAccount(ctx context.Context, req *youzan_model.CreateAccountReq) (res *youzan_model.CreateAccountRes, err error) {
	youzanRes, err := s.Request[youzan_model.CreateAccountRes](ctx, req, "youzan.pay.open.user.account_create", "1.0.3")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}

// 查询进件结果
func (s *Client) QueryCreateAccountStatus(ctx context.Context, req *youzan_model.QueryCreateAccountStatusReq) (res *youzan_model.QueryCreateAccountStatusRes, err error) {
	youzanRes, err := s.Request[youzan_model.QueryCreateAccountStatusRes](ctx, req, "youzan.pay.open.user.account.status_query", "1.0.0")
	if err != nil {
		return
	}
	res = &youzanRes.Data
	return
}
