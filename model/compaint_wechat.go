package youzan_model

type ComplaintWechatListReq struct { // 有赞微信投诉列表查询
	SubMchId  string `json:"sub_mch_id"` // 是否必填: 是;长度：0;平台商户号;有赞支付二级子商户号
	WxMchId   string `json:"wx_mch_id"`  // 是否必填: 是;长度：0;微信子商户号;微信被诉子商户号
	BeginDate string `json:"begin_date"` // 是否必填: 是;长度：10;开始日期;格式yyyy-MM-dd
	EndDate   string `json:"end_date"`   // 是否必填: 是;长度：10;结束日期;格式yyyy-MM-dd
	PageSize  int64  `json:"page_size"`  // 是否必填: 是;长度：0;分页大小;范围[1,50]
	Page      int64  `json:"page"`       // 是否必填: 是;长度：0;分页页码;范围[1,N]
}

type ComplaintWechatListRes struct {
	Data       []ComplaintWechatRes `json:"data"`
	Limit      int64                `json:"limit"`
	Offset     int64                `json:"offset"`
	TotalCount int64                `json:"total_count"`
}

type ComplaintWechatRes struct {
	ComplaintId           string                     `json:"complaint_id"`
	ComplaintTime         string                     `json:"complaint_time"`
	ComplaintDetail       string                     `json:"complaint_detail"`
	ComplaintState        string                     `json:"complaint_state"`
	PayerPhone            string                     `json:"payer_phone"`
	PayerOpenid           string                     `json:"payer_openid,omitempty"` // 详情里有，列表结果里没有
	ComplaintOrderInfo    []ComplaintOrderInfoEntity `json:"complaint_order_info"`
	ComplaintFullRefunded bool                       `json:"complaint_full_refunded"`
	IncomingUserResponse  bool                       `json:"incoming_user_response"`
	UserComplaintTimes    int64                      `json:"user_complaint_times"`
	ComplaintMediaList    []ComplaintMediaListEntity `json:"complaint_media_list"`
	ProblemDescription    string                     `json:"problem_description"`
	ProblemType           string                     `json:"problem_type"`
	ApplyRefundAmount     int64                      `json:"apply_refund_amount"`
	UserTagList           []string                   `json:"user_tag_list"`
	ServiceOrderInfo      []ServiceOrderInfoEntity   `json:"service_order_info"`
	AdditionalInfo        AdditionalInfoEntity       `json:"additional_info"`
	InPlatformService     bool                       `json:"in_platform_service"`
	NeedImmediateService  bool                       `json:"need_immediate_service"`
	IsAgentMode           bool                       `json:"is_agent_mode"`
}

type ComplaintOrderInfoEntity struct {
	TransactionId string `json:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no"`
	Amount        int64  `json:"amount"`
}

type ComplaintMediaListEntity struct {
	MediaType string   `json:"media_type"`
	MediaUrl  []string `json:"media_url"`
}

type ServiceOrderInfoEntity struct {
	OrderId    string `json:"order_id"`
	OutOrderNo string `json:"out_order_no"`
	State      string `json:"state"`
}

type AdditionalInfoEntity struct {
	Type           string               `json:"type"`
	SharePowerInfo SharePowerInfoEntity `json:"share_power_info"`
}

type SharePowerInfoEntity struct {
	ReturnTime              string                  `json:"return_time"`
	ReturnAddressInfo       ReturnAddressInfoEntity `json:"return_address_info"`
	IsReturnedToSameMachine bool                    `json:"is_returned_to_same_machine"`
}

type ReturnAddressInfoEntity struct {
	ReturnAddress string `json:"return_address"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
}

type ComplaintWechatDetailReq struct { // 微信投诉单个详情请求
	SubMchId    string `json:"sub_mch_id,omitempty"`   // 是否必填: 是;长度：0;平台商户号;有赞支付二级子商户号
	WxMchId     string `json:"wx_mch_id,omitempty"`    // 是否必填: 是;长度：0;微信子商户号;微信被诉子商户号
	ComplaintId string `json:"complaint_id,omitempty"` // 投诉订单号
}

type ComplaintWechatHistoryReq struct { // 微信投诉协商历史请求
	SubMchId    string `json:"sub_mch_id,omitempty"`   // 是否必填: 是;长度：0;平台商户号;有赞支付二级子商户号
	WxMchId     string `json:"wx_mch_id,omitempty"`    // 是否必填: 是;长度：0;微信子商户号;微信被诉子商户号
	ComplaintId string `json:"complaint_id,omitempty"` // 投诉订单号
	PageSize    int64  `json:"page_size"`              // 是否必填: 是;长度：0;分页大小;范围[1,50]
	Page        int64  `json:"page"`                   // 是否必填: 是;长度：0;分页页码;范围[1,N]
}

type ComplaintWechatHistoryRes struct {
	Data       []ComplaintWechatHistoryItem `json:"data"`
	Limit      int64                        `json:"limit"`
	Offset     int64                        `json:"offset"`
	TotalCount int64                        `json:"total_count"`
}

type ComplaintWechatHistoryItem struct {
	LogId                                    string                   `json:"log_id"`
	Operator                                 string                   `json:"operator"`
	OperateTime                              string                   `json:"operate_time"`
	OperateType                              string                   `json:"operate_type"`
	OperateDetails                           string                   `json:"operate_details"`
	ImageList                                []string                 `json:"image_list"`
	ComplaintMediaList                       ComplaintMediaListEntity `json:"complaint_media_list"`
	UserAppyPlatformServiceReason            string                   `json:"user_appy_platform_service_reason"`
	UserAppyPlatformServiceReasonDescription string                   `json:"user_appy_platform_service_reason_description"`
	NormalMessage                            NormalMessageEntity      `json:"normal_message"`
	ClickMessage                             ClickMessageEntity       `json:"click_message"`
}

type NormalMessageEntity struct {
	Blocks         []BlocksEntity `json:"blocks"`
	SenderIdentity string         `json:"sender_identity"`
	CustomData     string         `json:"custom_data"`
}

type BlocksEntity struct {
	Type        string            `json:"type"`
	Text        TextEntity        `json:"text"`
	Image       ImageEntity       `json:"image"`
	Link        LinkEntity        `json:"link"`
	FaqList     FaqListEntity     `json:"faq_list"`
	Button      ButtonEntity      `json:"button"`
	ButtonGroup ButtonGroupEntity `json:"button_group"`
}

type TextEntity struct {
	Text   string `json:"text"`
	Color  string `json:"color"`
	IsBold bool   `json:"is_bold"`
}

type ImageEntity struct {
	MediaId        string `json:"media_id"`
	ImageStyleType string `json:"image_style_type"`
}

type LinkEntity struct {
	Text        string            `json:"text"`
	Action      ActionEntity      `json:"action"`
	InvalidInfo InvalidInfoEntity `json:"invalid_info"`
}

type ActionEntity struct {
	ActionType          string                    `json:"action_type"`
	JumpUrl             string                    `json:"jump_url"`
	MiniProgramJumpInfo MiniProgramJumpInfoEntity `json:"mini_program_jump_info"`
	MessageInfo         MessageInfoEntity         `json:"message_info"`
	ActionId            string                    `json:"action_id"`
}

type MessageInfoEntity struct {
	Content    string `json:"content"`
	CustomData string `json:"custom_data"`
}

type InvalidInfoEntity struct {
	ExpiredTime    string `json:"expired_time"`
	MultiClickable bool   `json:"multi_clickable"`
}

type FaqListEntity struct {
	Faqs []FaqsEntity `json:"faqs"`
}

type FaqsEntity struct {
	FaqId    string       `json:"faq_id"`
	FaqTitle string       `json:"faq_title"`
	Action   ActionEntity `json:"action"`
}

type ButtonEntity struct {
	Text        string             `json:"text"`
	Action      ActionEntity1      `json:"action"`
	InvalidInfo InvalidInfoEntity2 `json:"invalid_info"`
}

type ActionEntity1 struct {
	ActionType          string                    `json:"action_type"`
	JumpUrl             string                    `json:"jump_url"`
	MiniProgramJumpInfo MiniProgramJumpInfoEntity `json:"mini_program_jump_info"`
	MessageInfo         MessageInfoEntity         `json:"message_info"`
	ActionId            string                    `json:"action_id"`
}

type InvalidInfoEntity2 struct {
	ExpiredTime    string `json:"expired_time"`
	MultiClickable bool   `json:"multi_clickable"`
}

type ButtonGroupEntity struct {
	Buttons      []ButtonsEntity    `json:"buttons"`
	ButtonLayout string             `json:"button_layout"`
	InvalidInfo  InvalidInfoEntity3 `json:"invalid_info"`
}

type ButtonsEntity struct {
	Text   string       `json:"text"`
	Action ActionEntity `json:"action"`
}

type MiniProgramJumpInfoEntity struct {
	Appid string `json:"appid"`
	Path  string `json:"path"`
}

type InvalidInfoEntity3 struct {
	ExpiredTime    string `json:"expired_time"`
	MultiClickable bool   `json:"multi_clickable"`
}

type ClickMessageEntity struct {
	MessageContent string `json:"message_content"`
	ActionId       string `json:"action_id"`
	ClickedLogId   string `json:"clicked_log_id"`
}

type ComplaintWechatReplyReq struct { // 有赞微信投诉回复用户
	SubMchId            string          `json:"sub_mch_id"`                       // 是否必填: 是;长度：0;平台商户号;有赞支付二级子商户号
	WxMchId             string          `json:"wx_mch_id"`                        // 是否必填: 是;长度：0;微信子商户号;微信被诉子商户号
	ComplaintId         string          `json:"complaint_id"`                     // 是否必填: 是;长度：0;投诉单号;投诉单对应的投诉单号
	ResponseContent     string          `json:"response_content"`                 // 是否必填: 是;长度：0;回复内容;具体的投诉处理方案，限制200个字符以内
	ResponseImages      []string        `json:"response_images,omitempty"`        // 是否必填: 否;长度：0;回复图片;最多上传4张图片凭证 // 暂不支持图片
	JumpUrl             string          `json:"jump_url,omitempty"`               // 是否必填: 否;长度：0;跳转链接;商户可在回复中附加跳转链接，引导用户跳转至商户客诉处理页面，链接需满足HTTPS格式
	JumpUrlText         string          `json:"jump_url_text,omitempty"`          // 是否必填: 否;长度：0;跳转链接文案;实际展示给用户的文案，附在回复内容之后。用户点击文案，即可进行跳转
	MiniProgramJumpInfo JumpMiniProgram `json:"mini_program_jump_info,omitempty"` // 是否必填: 否;长度：0;跳转小程序信息;商户可在回复中附加小程序信息，引导用户跳转至商户客诉处理小程序
}

type JumpMiniProgram struct { // 有赞微信投诉回复用户跳转小程序
	AppId string `json:"app_id"` // 是否必填: 是;长度：0;跳转小程序APPID;商户可在回复中附加小程序页面路径，引导用户跳转至商户服务工具页面。该字段为小程序APPID
	Path  string `json:"path"`   // 是否必填: 是;长度：0;跳转小程序页面PATH;商户可在回复中附加小程序页面路径，引导用户跳转至商户服务工具页面。该字段为小程序路径
	Text  string `json:"text"`   // 是否必填: 是;长度：0;跳转小程序页面名称;商户可在回复中附加小程序页面路径，引导用户跳转至商户服务工具页面。该字段为商户可自定义的页面名称
}

type ComplaintWechatCompleteReq struct { // 有赞微信投诉，给系统反馈已处理完成，系统给用户下发满意度评价
	SubMchId    string `json:"sub_mch_id"`   // 是否必填: 是;长度：0;平台商户号;有赞支付二级子商户号
	WxMchId     string `json:"wx_mch_id"`    // 是否必填: 是;长度：0;微信子商户号;微信被诉子商户号
	ComplaintId string `json:"complaint_id"` // 是否必填: 是;长度：0;投诉单号;投诉单对应的投诉单号
}
type ComplaintWechatRefundReq struct { // 有赞微信投诉处理退款
	SubMchId        string   `json:"sub_mch_id"`                  // 是否必填: 是;长度：0;平台商户号;有赞支付二级子商户号
	WxMchId         string   `json:"wx_mch_id"`                   // 是否必填: 是;长度：0;微信子商户号;微信被诉子商户号
	ComplaintId     string   `json:"complaint_id"`                // 是否必填: 是;长度：0;投诉单号;投诉单对应的投诉单号
	Action          string   `json:"action"`                      // 是否必填: 是;长度：0;审批动作;REJECT: 拒绝退款APPROVE: 同意退款
	LaunchRefundDay int64    `json:"launch_refund_day,omitempty"` // 是否必填: 否;长度：0;预计发起退款时间;在同意退款时返回，预计将在多少个工作日内能发起退款, 0代表当天
	RejectReason    string   `json:"reject_reason,omitempty"`     // 是否必填: 否;长度：0;拒绝退款原因;在拒绝退款时返回拒绝退款的原因
	RejectMediaList []string `json:"reject_media_list,omitempty"` // 是否必填: 否;长度：0;拒绝退款的举证图片列表;在拒绝退款时，如果有拒绝的图片举证，可以提供最多上传4张图片
	Remark          string   `json:"remark,omitempty"`            // 是否必填: 否;长度：0;备注;任何需要向微信支付客服反馈的信息
}
