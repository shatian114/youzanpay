package youzan_model

// 开通支付能力
type OpenPayConfig struct { // 使用商户自有收银台的配置
	WxWoaAppIdList    []string `json:"wx_woa_app_id_list,omitempty"`    // 微信公众号AppId列表
	WxAppletAppIdList []string `json:"wx_applet_app_id_list,omitempty"` // 微信小程序AppId列表
	WxWoaPathList     []string `json:"wx_woa_path_list,omitempty"`      // 微信支付授权目录列表
}

type OpenPaySupplement struct { // 补充资料，不填则沿用商户的
	StoreHeadPhoto   string `json:"store_head_photo,omitempty"`
	StoreIndoorPhoto string `json:"store_indoor_photo,omitempty"`
	CashierDeskPhoto string `json:"cashier_desk_photo,omitempty"`
}

type OpenPayReq struct { // 开通支付能力请求
	SubMchId               string            `json:"sub_mch_id"`               //
	OutBizNo               string            `json:"out_biz_no"`               // 此次请求订单号
	BizSceneType           string            `json:"biz_scene_type"`           // 场景 ONLINE(线上) OFFLINE(线下) 虚拟行业需联系产品宝贝开通
	SignType               string            `json:"sign_type"`                // 收款场景 WX_PUBLIC_ACCOUNT(微信公众号) WX_APPLET(微信小程序) ALIPAY_JS(支付宝js)
	PaymentConfig          OpenPayConfig     `json:"payment_config,omitempty"` //
	SupplementaryMaterials OpenPaySupplement `json:"supplementary_materials,omitempty"`
}

type OpenPayQueryReq struct {
	SubMchId string `json:"sub_mch_id"`           //
	OutBizNo string `json:"out_biz_no,omitempty"` // 此次请求订单号，不填查询全部
}
type OpenPayQueryRes struct {
	SubMchId          string                   `json:"sub_mch_id,omitempty"`
	ApplicationResult OpenPayApplicationResult `json:"application_result,omitempty"`
	Channels          []OpenPayChannel         `json:"channels,omitempty"`
}

type OpenPayChannel struct {
	Reason          string `json:"reason,omitempty"`
	PayConfigStatus int64  `json:"pay_config_status,omitempty"`
	QrcodeUrl       string `json:"qrcode_url,omitempty"`
	SubMerchantNo   string `json:"sub_merchant_no,omitempty"`
	State           string `json:"state,omitempty"`
	OutBizNo        string `json:"out_biz_no,omitempty"`
	AppId           string `json:"app_id,omitempty"`
	SignType        string `json:"sign_type,omitempty"`
	PayChannelType  string `json:"pay_channel_type,omitempty"`
}

type OpenPayApplicationResult struct {
	OutBizNo         string `json:"out_biz_no,omitempty"`
	SubmissionStatus string `json:"submission_status,omitempty"`
}
