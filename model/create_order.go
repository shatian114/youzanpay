package youzan_model

// 根据不同的支持产品和环境进行搭配使用
// WX_JS: open_id(走有赞公共收银台不需要传) app_id(走有赞公共收银台不需要传) return_url(控制微信支付完成后，小票页面点击详情 跳转到指定链接地址。示例值：https://h5.xxxx.com/xxx.html)
// ALIPAY_JS: buyer_id(buyer_id 为支付宝环境下买家支付宝用户ID  示例值：208830******4490)
// WX_APPLET: open_id(open_id为支付环境app_id下 用户的open_id  示例值：oHIuujp************hxt3A) app_id()
type PayExt struct { // 下单的 Ext参数
	OpenId    string `json:"open_id"`    // open_id为支付环境app_id下 用户的open_id 示例值：oHIuujp************hxt3A  走有赞公共收银台不需要传
	AppId     string `json:"app_id"`     //app_id为支付环境所在公众号所属app_id 示例值：wx55****5e84478   走有赞公共收银台不需要传
	ReturnUrl string `json:"return_url"` // 控制微信支付完成后，小票页面点击详情 跳转到指定链接地址。 示例值：https://h5.xxxx.com/xxx.html
	BuyerId   string `json:"buyer_id"`   // buyer_id 为支付宝环境下买家支付宝用户ID  示例值：208830******4490
	Type      string `json:"type"`
}

type Payer struct { // 下单的买家信息
	Name        string `json:"name"`         // 可选，买家名称，对应身份证/有效证件的名称 示例值：王**
	NickName    string `json:"nick_name"`    // 可选， 买家昵称，对应支付环境下的的昵称信息，比如微信支付环境下，可传入微信 示例值：昵称* ～
	PhoneNumber string `json:"phone_number"` // 可选，买家手机号，买家的手机号  示例值：188******73
}

type CreateOrderReq struct { // 有赞下单请求
	OuterTradeNo string `json:"outer_trade_no"`       // 是否必填: 是;长度：0;外部订单号;服务商内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的例：2305181535202621620585
	TradeDesc    string `json:"trade_desc"`           // 是否必填: 是;长度：0;交易描述;交易描述信息，请确保真实、详尽例：有赞的店-玩具公仔
	PayAmount    int64  `json:"pay_amount"`           // 是否必填: 是;长度：0;支付金额;支付金额 > 0单位：分
	Currency     string `json:"currency,omitempty"`   // 是否必填: 否;长度：0;币种;默认：CNY
	SubMchId     string `json:"sub_mch_id"`           // 是否必填: 是;长度：18;收款商户号;有赞支付二级商户号（特约商户号），商户服务商身份开通二级子商户，并且完成支付能力配置；资金将结算至该商户号余额中。
	PayTool      string `json:"pay_tool"`             // 是否必填: 是;长度：0;支付工具;用于支付时指定支付工具，详见支付工具枚举例：WX_JS
	UserIp       string `json:"user_ip"`              // 是否必填: 是;长度：0;支付环境IP;支付环境IP，支持IPV6/IPV4指的是完成支付拉起三方收银台的设备环境IP，传入公网IP, 建议从前端直接获取例：8.8.8.8
	PayNotifyUrl string `json:"pay_notify_url"`       // 是否必填: 是;长度：0;支付结果回调地址;用于支付成功结果通知必须为https地址。不携带查询参数, 链接例：https://xxxx.cloud.com/service_api/pay/notify/路径商家自定义设置，域名需要报备给有赞，否则无法正常接收回调请求！！！！！！
	Attach       string `json:"attach,omitempty"`     // 是否必填: 否;长度：0;附加信息;附加数据，在支付通知中原样返回，可作为自定义参数使用，实际情况下只有支付完成状态才会返回该字段。例：xxxxxx
	ExpireTime   string `json:"expire_time"`          // 是否必填: 是;长度：0;支付结束时间;支付结束时间，超过此时间 支付单将失效，无法再完成支付，需要更换外部订单号重新支付东八区标准时间 yyyy-MM-dd hh:mm:ss例：2023-05-18 16:00:00
	Ext          PayExt `json:"ext,omitempty"`        // 是否必填: 是;长度：0;支付渠道请求扩展;支付扩展信息，用于传递不同支付工具下的支付参数，具体传参可参考支付渠道扩展参数说明
	PayerInfo    Payer  `json:"payer_info,omitempty"` // 是否必填: 是;长度：0;支付买家信息;支付买家信息，用于传递支付工具下的买家详细信息，具体传参可参考支付买家信息参数说明
}

type Deeplink struct { // 有赞公共收银台：对应微信和支付宝的h5支付
	AwakePayUrl *string `json:"awake_pay_url,omitempty"` // 有赞公共支付宝和微信h5收银台

	TradeNo string `json:"trade_no,omitempty"` // 商户自有支付宝收银台

	// 微信小程序：有赞公共收银台
	Path    string `json:"path,omitempty"`
	AppId   string `json:"appId,omitempty"` // 公共和自有都会返回
	Urllink string `json:"urllink,omitempty"`

	// 微信小程序：商户自有收银台
	TimeStamp string `json:"timeStamp,omitempty"`
	PaySign   string `json:"paySign,omitempty"`
	SignType  string `json:"signType,omitempty"`
	Package   string `json:"package,omitempty"`
	NonceStr  string `json:"nonceStr,omitempty"`
}

type CreateOrderRes struct { // 有赞下单响应
	OutBizNo  string   `json:"out_biz_no"`         // 是否必填: 是;长度：0;外部单号;和请求参数中传递的外部单号一致 ；服务商内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的例：81535202621620585
	AcquireNo string   `json:"acquire_no"`         // 是否必填: 是;长度：0;支付收单号;返回本次下单支付成功有赞支付单编号，用于在有赞支付域内标记订单的支付记录例：2608111847570000180807
	Status    string   `json:"status"`             // 是否必填: 是;长度：0;支付状态;CREATE : "支付创建"PROCESSING : "交易处理中SUCCESS : "支付成功"
	Deeplink  Deeplink `json:"deeplink,omitempty"` // 是否必填: 条件必填;长度：0;支付渠道返回字段;用于拉起第三方支付收银台参数，详情见支付渠道返回说明
}

// 查询订单
type QueryOrderReq struct { // 订单查询请求
	OuterTradeNo string `json:"outer_trade_no,omitempty"` // 是否必填: 支付收单号和外部订单号二选一;长度：0;外部订单号;务商内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的
	TradeNo      string `json:"trade_no,omitempty"`       // 是否必填: 支付收单号和外部订单号二选一;长度：0;收单号;本次下单支付成功有赞支付单编号，用于在有赞支付域内标记订单的支付记录
	SubMchId     string `json:"sub_mch_id"`               // 是否必填: 是;长度：18;商户号;有赞支付商户号，有赞支付分配
}

type QueryOrderRes struct { // 订单查询详情
	Status               string         `json:"status"`                          // 是否必填: 是;长度：0;支付状态;PROCESSING // 交易处理中， SUCCESS    // 支付成功， CLOSED     // 交易关闭， FINISHED   // 交易完成/全额退款完成
	StatusDesc           string         `json:"status_desc"`                     // 是否必填: 是;长度：0;支付描述;支付状态描述
	TradeNo              string         `json:"trade_no"`                        // 是否必填: 是;长度：0;支付收单号;有赞生分配，本次下单支付成功有赞支付单编号，用于在有赞支付域内标记订单的支付记录
	OuterTradeNo         string         `json:"outer_trade_no"`                  // 是否必填: 是;长度：0;业务单号;统一支付时传入；服务商内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的
	SubMchId             string         `json:"sub_mch_id"`                      // 是否必填: 是;长度：0;收款商户号;有赞支付商户号，有赞支付分配
	Currency             string         `json:"currency,omitempty"`              // 是否必填: 否;长度：0;支付币种;支付币种 默认CNY
	ExpireTime           int64          `json:"expire_time"`                     // 是否必填: 是;长度：0;支付过期时间;支付过期时间，统一支付传入的过期时间
	ExternalOutPaymentNo string         `json:"external_out_payment_no "`        // 是否必填: 是;长度：0;三方支付外部单号;支付商户单号，有赞分配，有赞发起支付时，调用三方支付（微信/支付宝）的商户单号
	ExternalPaymentNo    string         `json:"external_payment_no,omitempty"`   // 是否必填: 否;长度：0;三方支付单号;调用三方支付（微信/支付宝）返回支付单号
	ExternalCompletedAt  int64          `json:"external_completed_at,omitempty"` // 是否必填: 否;长度：0;三方支付完成时间;调用三方支付（微信/支付宝）完成时间， 毫秒级别时间戳
	PayAmount            int64          `json:"pay_amount"`                      // 是否必填: 是;长度：0;交易金额;单位：分
	PayTool              string         `json:"pay_tool,omitempty"`              // 是否必填: 否;长度：0;支付工具;用于完成支付时的支付工具
	SuccessTime          int64          `json:"success_time,omitempty"`          // 是否必填: 否;长度：0;支付完成时间;支付完成时间， 毫秒级别时间戳
	Attach               string         `json:"attach"`                          // 是否必填: 是;长度：0;支付附加信息;支付下单接口中，业务传递过来的信息，原样进行返回只有在支付成功情况下返回
	Extra                map[string]any `json:"extra,omitempty"`                 // 是否必填: 否;长度：0;渠道附加信息;返回支付渠道支付信息例：bank_type：'OTHERS' // 微信支付返回买家支付银行编码
}
