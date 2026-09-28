package youzan_model

type ListSubMchReq struct { // 查询有赞侧子商户列表
	MchId    string `json:"mch_id"`    // 是否必填: 是;长度：0;服务商商户号;服务商商户号
	Page     int64  `json:"page"`      // 是否必填: 是;长度：0;页码;页从1开始
	PageSize int64  `json:"page_size"` // 是否必填: 是;长度：0;页大小;页大小 最大50
}

type ListSubMchDetail struct { // 查询有赞侧子商户列表结果单个详情
	SubMchId  string `json:"sub_mch_id"` // 是否必填: 是;长度：0;二级子商户号;特约商户商户号；例：2044xxxxxxx011
	MchName   string `json:"mch_name"`   // 是否必填: 是;长度：0;商户名称;商户名称例：xxx科技公司
	ShortName string `json:"short_name"` // 是否必填: 是;长度：0;商户简称;商户简称例：xxx科技公司简称
	Type      int64  `json:"type"`       // 是否必填: 是;长度：0;商户类型;商户类型例：2 - 特约商户
	Status    string `json:"status"`     // 是否必填: 是;长度：0;商户状态;INIT 初始化 PRE_OPEN 预开户 APPLYED 认证申请 SIGNED 已签约 SIGN_FAIL 签约失败 SIGN_SUCCESS 签约成功 FINISH 开户成功注意：特约商户进件时，当审核通过后，会给商家发送“激活确认邮件”，邮件内容为商家登录高汇通后台的登录密码，需商家人工在邮件内点击【激活】后，商户号方可使用。可联系有赞去掉邮件发送及激活环节。; 若走发送邮件及激活环节：认证审核通过后，商户状态变为“SIGN_SUCCESS”；商家邮件激活后，商户状态变为“FINISH “，至此，商户开户成功，可进行特约商户分账等资金操作 ; 若跳过邮件及激活环节：认证审核通过后，商户状态直接变为“FINISH “，至此，商户开户成功，可进行特约商户分账等资金操作
}

type ListSubMchResPageInfo struct { // 查询有赞侧子商户列表结果页码信息
	Page     int64 `json:"page"`      // 是否必填: 是;长度：0;数据列表;页码
	PageSize int64 `json:"page_size"` // 是否必填: 是;长度：0;分页信息;页大小
	TotalNum int64 `json:"total_num"` // 是否必填: 是;长度：0;数据总数;数据总数
}

type ListSubMchRes struct { // 查询有赞侧子商户列表结果
	List     []ListSubMchDetail    `json:"list"`      // 是否必填: 是;长度：0;数据列表;数据列表信息见：数据信息
	PageInfo ListSubMchResPageInfo `json:"page_info"` // 是否必填: 是;长度：0;分页信息;分页信息见：分页信息
}
