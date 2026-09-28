package youzan_model

type ComplaintAlipayDetailReq struct {
	SubMchId       string `json:"sub_mch_id,omitempty"`
	OutComplaintId string `json:"out_complaint_id"`
}

type ComplaintAlipayDetailRes struct { // 有赞支付宝投诉详情响应
	ComplaintId          string `json:"complaint_id"`            // 是否必填: 是;长度：0;投诉单号;merchant 内部投诉单号。
	OutComplaintId       string `json:"out_complaint_id"`        // 是否必填: 是;长度：0;支付宝投诉单号;机构侧投诉单号
	OutComplaintRecordId string `json:"out_complaint_record_id"` // 是否必填: 是;长度：0;支付宝投诉记录主键;机构侧投诉记录主键，对应支付宝 record_id/id。
	ChannelSource        string `json:"channel_source"`          // 是否必填: 是;长度：0;投诉来源渠道;投诉来源渠道。
	OutComplaintMchId    string `json:"out_complaint_mch_id"`    // 是否必填: 是;长度：0;支付宝侧商户号;支付宝侧被投诉商户号。
	ComplaintDetail      string `json:"complaint_detail"`        // 是否必填: 是;长度：0;投诉详情;投诉详情文本。
	OutState             string `json:"out_state"`               // 是否必填: 是;长度：0;投诉状态;投诉状态 用户撤诉: DROP_COMPLAIN超时后用户撤诉: DROP_OVERDUE_COMPLAIN超时处理完成用户撤诉: DROP_OVERDUE_PROCESSED处理完成用户撤诉: DROP_PROCESSED超时未处理: OVERDUE超时处理完成: OVERDUE_PROCESSED部分超时未处理: PART_OVERDUE处理完成: PROCESSED处理中: PROCESSING待处理: WAIT_PROCESS
	OutStateDesc         string `json:"out_state_desc"`          // 是否必填: 是;长度：0;状态说明;投诉状态说明。
	ComplaintTime        int64  `json:"complaint_time"`          // 是否必填: 是;长度：0;投诉发生时间;投诉发生时间 ms
	PayerPhone           string `json:"payer_phone"`             // 是否必填: 是;长度：0;消费者联系方式;消费者手机号
	ProcessTime          int64  `json:"process_time"`            // 是否必填: 是;长度：0;处理时间;处理时间 ms
	ProcessResult        string `json:"process_result"`          // 是否必填: 是;长度：0;处理结果;处理结果。
	ProcessRemark        string `json:"process_remark"`          // 是否必填: 是;长度：0;处理备注;处理备注。
}

type ComplaintAlipayReplyReq struct { //回复请求
	SubMchId       string `json:"sub_mch_id,omitempty"`
	OutComplaintId string `json:"out_complaint_id,omitempty"`
	ReplyContent   string `json:"reply_content,omitempty"`
}

type ComplaintAlipayCompleteReq struct { // 有赞支付宝投诉完结请求
	SubMchId       string         `json:"sub_mch_id"`              // 是否必填: 是;长度：0;商户号;merchant 内部商户号。
	OutComplaintId string         `json:"out_complaint_id"`        // 是否必填: 是;长度：0;支付宝投诉单号;支付宝投诉单号，即通知中的 complaint_id。
	ProcessCode    string         `json:"process_code"`            // 是否必填: 是;长度：0;处理结果编码;支付宝处理结果编码枚举值已联系到用户，协商一致，无异议: CONSENSUS_WITH_CLIENT其他: ORTHER不涉及退款，已针对投诉内容进行整改: RECTIFICATION_NO_REFUND已退款，用户无异议: REFUND不认可举报的证明材料: SUBMIT_PROOF_NOT_CONTACTED其他渠道已退款: OTHER_CHANNEL_REFUND
	ReplyContent   string         `json:"reply_content"`           // 是否必填: 是;长度：0;回复内容;商家完结留言内容
	Remark         string         `json:"remark,omitempty"`        // 是否必填: 否;长度：0;商户处理备注;商户处理备注
	ImgFileList    []AlipayImgUrl `json:"img_file_list,omitempty"` // 是否必填: 否;长度：0;完结举证图片列表;图片先通过投诉文件上传接口取得 file_url、file_key，再按支付宝要求传入。
}
type AlipayImgUrl struct {
	ImgUrl    string `json:"img_url,omitempty"`
	ImgUrlKey string `json:"img_url_key,omitempty"`
}
