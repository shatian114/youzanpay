package youzan_model

type WithDrawReq struct { // 申请提现请求
	SubMchId       string `json:"sub_mch_id"`             // 是否必填: 是;长度：32;二级子商户号;又名特约商户；
	OutBizNo       string `json:"out_biz_no"`             // 是否必填: 是;长度：32;提现业务单号;提现业务单号，服务商内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的
	WithdrawAmount int64  `json:"withdraw_amount"`        // 是否必填: 是;长度：0;提现金额;提现金额 单位/分
	CardNo         string `json:"card_no"`                // 是否必填: 是;长度：32;卡号;提现卡号
	Remark         string `json:"remark"`                 // 是否必填: 是;长度：32;备注;提现备注
	TerminalType   string `json:"terminal_type"`          // 是否必填: 是;长度：32;终端类型;枚举如下：电脑终端：PC苹果手机：IOS平板电脑：PAD安卓手机：ANDROID
	AccountType    string `json:"account_type,omitempty"` // 是否必填: 否;长度：32;提现账户类型;10：余额户提现80：分账账户提现（查询余额接口返回值中 余额在：PROFIT_SHARING_ACCT 账户时才需要传此值）
}

type WithDrawRes struct { // 提现申请返回结果
	TradeNo    string `json:"trade_no"`    // 是否必填: 是;长度：0;提现收单号;提现支付收单号
	OutBizNo   string `json:"out_biz_no"`  // 是否必填: 是;长度：0;业务单号;提现外部业务单号
	Status     int    `json:"status"`      // 是否必填: 是;长度：0;状态码;提现状态200:处理成功,1001:未知异常,1002:参数异常,1003:收单不存在,1005:提现金额非法,1006:收单状态非法,1007:获取卡信息异常,2001:店铺提现功能已锁定，详情请登录商家后台通知中心查看,2002:账务出账失败,2003:账务入账失败,2004:出金提现受理失败,2005:店铺认证校验失败,2006:商户信息未开户完成,2007:多门店提现权限校验失败,1008:提现申请中,1009:提现申请失败,1010:提现处理中,1011:提现成功,1012:提现失败,1013:银行退票
	StatusDesc string `json:"status_desc"` // 是否必填: 是;长度：0;状态描述;状态描述
	Success    bool   `json:"success"`     // 是否必填: 是;长度：0;是否成功;是否提现成功
	ErrorMsg   string `json:"errorMsg"`    // 是否必填: 是;长度：0;错误信息;提现错误原因
}

type WithDrawQueryReq struct { // 申请提现请求
	SubMchId string `json:"sub_mch_id"` // 是否必填: 是;长度：32;二级子商户号;又名特约商户；
	OutBizNo string `json:"out_biz_no"` // 是否必填: 是;长度：32;提现业务单号;提现业务单号，服务商内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的
}

type WithDrawQueryRes struct { // 提现申请查询结果
	SubMchId        string `json:"sub_mch_id"`        // 是否必填: 是;长度：0;二级子商户号;二级子商户号；又名特约商户
	OutBizNo        string `json:"out_biz_no"`        // 是否必填: 是;长度：0;提现申请单号;提现申请单号
	TradeNo         string `json:"trade_no"`          // 是否必填: 是;长度：0;提现收单号;提现收单号
	Status          int64  `json:"status"`            // 是否必填: 是;长度：0;提现状态代码;提现状态0, "申请中"，1, "申请成功"，2, "申请失败"，3, "提现处理中"，4, "风控拦截"，5, "提现异常转人工处理"，6, "提现失败"，7, "提现成功"，8,"银行退票"
	StatusDesc      string `json:"status_desc"`       // 是否必填: 是;长度：0;提现状态描述;见提现状态代码描述
	ApplyTime       int64  `json:"apply_time"`        // 是否必填: 是;长度：0;提现申请时间;提现申请时间
	ProcessTime     int64  `json:"process_time"`      // 是否必填: 是;长度：0;银行处理时间;银行处理时间
	FinishTime      int64  `json:"finish_time"`       // 是否必填: 是;长度：0;到账时间;到账时间
	FailDesc        string `json:"fail_desc"`         // 是否必填: 是;长度：0;失败原因;失败原因
	InstructionId   string `json:"instruction_id"`    // 是否必填: 是;长度：0;提现明细号;提现明细号
	WithdrawAmount  int64  `json:"withdraw_amount"`   // 是否必填: 是;长度：0;提现金额;提现金额
	InstAccountNo   string `json:"inst_account_no"`   // 是否必填: 是;长度：0;银行卡号;银行卡号
	InstName        string `json:"inst_name"`         // 是否必填: 是;长度：0;银行编码;银行编码
	InstBranchName  string `json:"inst_branch_name"`  // 是否必填: 是;长度：0;银行名称;银行名称
	InstAccountName string `json:"inst_account_name"` // 是否必填: 是;长度：0;持卡人名称;持卡人名称
	InstAccountType string `json:"inst_account_type"` // 是否必填: 是;长度：0;银行账户类型;银行账户类型
}
