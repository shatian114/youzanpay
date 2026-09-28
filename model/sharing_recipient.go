package youzan_model

type SharingRecipientAddReq struct { // 添加分账关系请求

	SubMchId        string `json:"sub_mch_id"`        // 是否必填: 是;长度：32;商户号;二级子商户号为服务商在有赞支付服务商下二级子商户商户号
	RecipientType   string `json:"recipient_type"`    // 是否必填: 是;长度：32;分账接收方类型;枚举如下：商户号 - MCH_ID，请使用该类型有赞openId - YZ_OPEN_ID。定制，请勿使用。
	RecipientUserNo string `json:"recipient_user_no"` // 是否必填: 是;长度：32;分账接收方账号;如分账到二级子商户，分账接收方类型 recipient_type为MCH_ID, 本参数传递为有赞支付侧创建的子商户号 mchId, 如分账到个人零钱账户，分账接收方类型 recipient_type为YZ_OPEN_ID, 本参数传递用户的有赞 OpenId
	RecipientName   string `json:"recipient_name"`    // 是否必填: 是;长度：128;接收方名称;如分账到二级子商户，分账接收方类型 recipient_type为MCH_ID, 本参数传递为有赞支付侧创建的子商户名称, 如分账到个人零钱账户，分账接收方类型 recipient_type为YZ_OPEN_ID, 本参数传递用户在有赞的账户昵称
	Memo            string `json:"memo"`              // 是否必填: 是;长度：128;分账关系描述;分账关系描述
	RelationType    string `json:"relation_type"`     // 是否必填: 是;长度：12;分账关系类型;商户与分账接收方的关系；枚举如下：其他 - OTHERS，请使用该关系。 供应商 - SUPPLIER，定制，请勿使用。
}
