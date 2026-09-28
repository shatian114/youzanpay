package youzan_model

type CreateAccount struct { // 商户进件请求参数
	MchId                           string               `json:"mch_id"`                             // 是否必填: 是;长度：32;服务商商户号;有赞支付分配的服务商商户号，用于管理子商户号
	OutBizNo                        string               `json:"out_biz_no"`                         // 是否必填: 是;长度：32;业务关联账号;业务关联账号；服务商内唯一，由调用方（服务商）自己生成并传入，不是有赞分配的
	MchName                         string               `json:"mch_name"`                           // 是否必填: 是;长度：32;商户名称;商户名称
	Phone                           string               `json:"phone"`                              // 是否必填: 是;长度：128;联系电话;联系电话
	ContractName                    string               `json:"contract_name"`                      // 是否必填: 是;长度：128;联系人姓名;联系人姓名
	ContractEmail                   string               `json:"contract_email "`                    // 是否必填: 是;长度：128;联系人邮箱;联系人邮箱
	PaymentBusinessType             int64                `json:"payment_business_type"`              // 是否必填: 否;长度：0;是否收款;0-收款（包含仅收款，或既收款又分账）1-不收款，仅分账不填写时，默认为0
	CallbackUrl                     string               `json:"callback_url"`                       // 是否必填: 是;长度：512;结果通知地址;用于支付成功结果通知必须为https地址。不携带查询参数, 链接例：https://xxxx.cloud.com/service_api/pay/notify/路径商家自定义设置，域名需要报备给有赞，否则无法正常接收回调请求！！！！！！
	PrincipalCertType               int64                `json:"principal_cert_type"`                // 是否必填: 是;长度：0;主体认证类型;1: 个人 2: 企业 3: 个体工商户 注意，如果是收款用途需要用企业，如果是分账给个人，可以使用个人资质
	OpenAccountProtocol             string               `json:"open_account_protocol"`              // 是否必填: 否;长度：128;支付开户协议;支付开户协议 调用资源图片上传接口获取imgKey 并上传
	MainCategoryCode                string               `json:"main_category_code"`                 // 是否必填: 是;长度：128;主营类目;主营类目传入类目code 即可：类目文档：https://file.yzcdn.cn/upload_files/yz-file/2025/08/14/FqQvsCNBjrQp2jpN0JPAmT1NhQPr.csv?attname=有赞支付类目列表v4.csv建议用文本编辑器打开，excel 等工具会自动格式化类目编码为数字，导致丢失部分字符
	BusinessSpecialQualificationUrl string               `json:"business_special_qualification_url"` // 是否必填: 条件必传;长度：128;特殊资质照片;如果选中主营类目中，资质要求 不为空，请按照资质要求传入对应资质图片
	SupplementaryPaperUrlList       []string             `json:"supplementary_paper_url_list"`       // 是否必填: 条件必传;长度：0;补充文件;如果选中主营类目中，资质要求 不为空，按照资质要求传入，如果需要补充文件，传入补充信息即可
	OrganizationCertInfo            OrganizationCertInfo `json:"organization_cert_info"`             // 是否必填: 否;长度：0;组织信息;参数见 组织信息 非个人类型必传
	PersonCertInfo                  PersonCertInfo       `json:"person_cert_info"`                   // 是否必填: 是;长度：0;法人信息;参数见 法人代办人信息 个人认证类型，传个人主体信息企业/个体工商户，传法人/代办人信息
	BankCardInfo                    BankCardInfo         `json:"bank_card_info"`                     // 是否必填: 是;长度：0;结算卡信息;结算卡信息，用于后续资金提现 企业，只支持绑定企业对公同名卡个人   工商户，可支持绑定企业对公同名卡或法人对私卡  个人，只能绑定个人对私卡
	FinalBeneficiaries              FinalBeneficiaries   `json:"final_beneficiaries"`                // 是否必填: 是;长度：0;最终受益人信息;参数见 最终受益人信息最终受益人信息，当法人非最终受益人时，需要传入最终受益人信息，最多支持4个最终受益人信息
}

// 进件所需信息
type OrganizationCertInfo struct { // 组织信息
	OrganizationName     string `json:"organization_name"`       // 必填;长度： 128;企业名称
	ProvinceCode         string `json:"province_code"`           // 必填;长度：128;营业地址编码/省;营业地址码 见https://file.yzcdn.cn/upload_files/yz-file/2026/01/26/FkglVmQ9UXjzk_rVtwwZDIZ8bBtR.xls?attname=有赞支付地址码-20260126.xls
	CityCode             string `json:"city_code"`               // 必填;长度：128;营业地址编码;营业地址码 见https://file.yzcdn.cn/upload_files/yz-file/2026/01/26/FkglVmQ9UXjzk_rVtwwZDIZ8bBtR.xls?attname=有赞支付地址码-20260126.xls
	DistrictCode         string `json:"district_code"`           // 必填;长度：128;营业地址编码;营业地址码 见https://file.yzcdn.cn/upload_files/yz-file/2026/01/26/FkglVmQ9UXjzk_rVtwwZDIZ8bBtR.xls?attname=有赞支付地址码-20260126.xls
	Address              string `json:"address"`                 // 必填;长度：128;营业地址;营业地址
	BusinessScope        string `json:"business_scope"`          // 必填;长度：128;经营范围;经营范围，请和营业执照上的经营范围保持一致
	CreditNo             string `json:"credit_no"`               // 必填;长度：128;营业执照号;营业执照号
	CreditPhoto          string `json:"credit_photo"`            // 必填;长度：256;营业执照照片;营业执照照片调用资源图片上传接口获取imgKey 并上传
	CreditValidStartTime string `json:"credit_valid_start_time"` // 必填;长度：64;营业执照有效期起;营业执照有效期起应满足合法的YYYY-MM-DD格式
	CreditValidEndTime   string `json:"credit_valid_end_time"`   // 必填;长度：64;营业执照有效期止;营业执照有效期止应满足合法的YYYY-MM-DD格式若为“长期”，传值为：当前时间+100 年
	StoreHeadPhoto       string `json:"store_head_photo"`        // 可填;长度：64;门店门头照片;门店门头照片；收款商户必传，分账接收方可不传；若有线下门店，请传门店的门头照；若无线下门店，请传办公室前台照；调用资源图片上传接口获取imgKey 并上传
	StoreIndoorPhoto     string `json:"store_indoor_photo"`      // 可填;长度：64;门店内景照片;门店内景照片；收款商户必传，分账接收方可不传；若有线下门店，请传门店的内景照；若无线下门店，请传办公室照；调用资源图片上传接口获取imgKey 并上传
	CashierDeskPhoto     string `json:"cashier_desk_photo"`      // 可填;长度：256;收银台照片;收银台照片；收款商户必传，分账接收方可不传；若有线下门店，请传门店的收银台照；若无线下门店，请传办公室财务室照；调用资源图片上传接口获取imgKey 并上传
}

type PersonCertInfo struct { // 法人信息
	OpPersonType                int64  `json:"op_person_type"`                 // 是否必填 是;长度：0;人员类型;人员类型 1：代办人2：法人
	PapersType                  int64  `json:"papers_type"`                    // 是否必填 是;长度：0;法人身份证件类型;1 大陆身份证 2 港澳台（身份证、居民通行证 3 护照（限境外人士） 4 来往内地通行证（中国香港居民）5 来往内地通行证（中国澳门居民） 6 来往大陆通行证（中国台湾居民）
	PersonName                  string `json:"person_name"`                    // 是否必填 是;长度：128;法人姓名;法人姓名
	PapersNo                    string `json:"papers_no"`                      // 是否必填 是;长度：128;法人身份证件号;法人身份证件号
	PapersPhotoFront            string `json:"papers_photo_front"`             // 是否必填 是;长度：256;法人身份证正面;调用资源图片上传接口获取imgKey 并上传
	PapersPhotoReverse          string `json:"papers_photo_reverse"`           // 是否必填 是;长度：256;法人身份证反面;调用资源图片上传接口获取imgKey 并上传
	PapersPhotoHand             string `json:"papers_photo_hand"`              // 是否必填 条件;长度：256;法人手持身份证;法人手持身份证如果存在代办人 则不需要传法人手持身份证调用资源图片上传接口获取imgKey 并上传
	ValidStartTime              string `json:"valid_start_time"`               // 是否必填 否;长度：10;法人证件有效期开始时间;法人身份证件有效期开始时间日期格式应满足合法的YYYY-MM-DD格式
	ValidEndTime                string `json:"valid_end_time"`                 // 是否必填 是;长度：10;法人证件有效期结束时间;法人身份证件有效期结束时间日期格式应满足合法的YYYY-MM-DD格式若为“长期”，传值为：当前时间+100 年
	Birthday                    string `json:"birthday"`                       // 是否必填 是;长度：10;法人出生/年/月;法人出生日期日期格式应满足合法的YYYY-MM-DD格式
	ResidentialAddress          string `json:"residential_address"`            // 是否必填 是;长度：256;法人居住地址;法人居住地址传身份证上居住地址信息
	LegalPersonWillingnessPhoto string `json:"legal_person_willingness_photo"` // 是否必填 是;长度：128;法人开户意愿确认函;法人开户意愿确认函调用资源图片上传接口获取imgKey 并上传模版：开户意愿确认函.docx开户意愿确认函--星结.docx
	IsFinalBeneficiary          string `json:"is_final_beneficiary"`           // 是否必填 是;长度：1;法人是否最终受益人;法人是否最终受益人参数传入1 ：法人是最终受益人0 : 法人非最终受益人，需要额外传入最终受益人信息
	AgentPapersType             int64  `json:"agent_papers_type"`              // 是否必填 条件必传;长度：0;代理人身份证类型;op_person_type为 代办人时，必传；枚举如下：1 大陆身份证 2 港澳台（身份证、居民通行证 3 护照（限境外人士） 4 来往内地通行证（中国香港居民）5 来往内地通行证（中国澳门居民） 6 来往大陆通行证（中国台湾居民）
	AgentPersonName             string `json:"agent_person_name"`              // 是否必填 条件必传;长度：128;代理人姓名;op_person_type为 代办人时，必传；
	AgentPapersNo               string `json:"agent_papers_no"`                // 是否必填 条件必传;长度：128;代理人身份证件号;op_person_type为 代办人时，必传；
	AgentPapersPhotoFront       string `json:"agent_papers_photo_front"`       // 是否必填 条件必传;长度：256;代理人身份证件正面照片;op_person_type为 代办人时，必传；调用资源图片上传接口获取imgKey 并上传
	AgentPapersPhotoReverse     string `json:"agent_papers_photo_reverse"`     // 是否必填 条件必传;长度：256;代理人身份证件背面照片;op_person_type为 代办人时，必传；调用资源图片上传接口获取imgKey 并上传
	AgentPapersPhotoHand        string `json:"agent_papers_photo_hand"`        // 是否必填 条件必传;长度：256;代理人手持证件照片;op_person_type为 代办人时，必传；调用资源图片上传接口获取imgKey 并上传
	AgencyFile                  string `json:"agency_file"`                    // 是否必填 条件必传;长度：256;代理文件证明;op_person_type为 代办人时，必传；调用资源图片上传接口获取imgKey 并上传模版：https://b.yzcdn.cn/cert/www/doc/授权说明书 (认证-代办人) .docx
	AgentValidStartTime         string `json:"agent_valid_start_time"`         // 是否必填 条件必传;长度：16;代理人证件有效期开始时间;op_person_type为 代办人时，必传；代理人证件有效期开始时间日期格式应满足合法的YYYY-MM-DD格式
	AgentValidEndTime           string `json:"agent_valid_end_time"`           // 是否必填 条件必传;长度：16;代理人证件有效期结束时间;op_person_type为 代办人时，必传；代理人证件有效期结束时间日期格式应满足合法的YYYY-MM-DD格式
	AgentBirthday               string `json:"agent_birthday"`                 // 是否必填 条件必传;长度：16;代理人出生/年/月;op_person_type为 代办人时，必传；代理人出生/年/月日期格式应满足合法的YYYY-MM-DD格式
}

type BankCardInfo struct { // 结算卡信息
	CardBiz       string `json:"card_biz"`            // 是否必填: 是;长度：32;银行卡对公/对私类型;对私：TO_PRIVATE对公：TO_PUBLIC主体认证类型（principal_cert_type）为：个人（1） 、个体工商户（3 ）时，支持对私结算卡；企业（2）、个体工商户（3）、其他组织（4），支持对公结算卡；
	CardType      string `json:"card_type,omitempty"` // CREDIT_CARD 信用卡  DEBIT_CARD 储蓄卡 默认为 DEBIT_CARD 储蓄卡 ，进件时无需传
	CardNo        string `json:"card_no"`             // 是否必填: 是;长度：32;卡号;卡号
	CardHolder    string `json:"card_holder"`         // 是否必填: 是;长度：128;持卡人;持卡人名称； 对私（TO_PRIVATE）时，持卡人名称需 与 法人/代办人信息参数中的 法人姓名（person_name ）保持一致对公（TO_PUBLIC），持卡人名称 需与 组织信息参数中的 企业名称（organization_name）保持一致
	BankCode      string `json:"bank_code,omitempty"` // 是否必填: 是;长度：32;发行卡code码;银行code可参考：https://file.yzcdn.cn/upload_files/yz-file/2025/08/06/FoXiXawRCAJotjWhYuQrLa1hQkSa.json?attname=有赞支付支持银行名单.json
	BankName      string `json:"bank_name,omitempty"` // 是否必填: 是;长度：32;发行卡code码;银行code可参考：https://file.yzcdn.cn/upload_files/yz-file/2025/08/06/FoXiXawRCAJotjWhYuQrLa1hQkSa.json?attname=有赞支付支持银行名单.json
	BankPhone     string `json:"bank_phone"`          // 是否必填: 条件必传;长度：32;银行预留手机号;card_biz为TO_PRIVATE对私结算卡时，必传参数为银行卡在银行预留的手机号码
	InterBankCode string `json:"inter_bank_code"`     // 是否必填: 条件必传;长度：32;联行号;card_biz为TO_PRIVATE对私结算卡时，可不传。card_biz为TO_PUBLIC对公结算卡时，必传。
	Subbranch     string `json:"subbranch"`           // 是否必填: 条件必传;长度：128;支行名称;card_biz为TO_PRIVATE对私结算卡时，可不传。card_biz为TO_PUBLIC对公结算卡时，必传。
}

type FinalBeneficiaries struct { // 最终受益人信息
	PapersType         int64  `json:"papersType"`         // 是否必填: 是;长度：0;最终受益人身份证件类型;1 大陆身份证 2 港澳台（身份证、居民通行证 3 护照（限境外人士） 4 来往内地通行证（中国香港居民）5 来往内地通行证（中国澳门居民） 6 来往大陆通行证（中国台湾居民）
	PersonName         string `json:"personName"`         // 是否必填: 是;长度：128;最终受益人姓名;最终受益人姓名
	PapersNo           string `json:"papersNo"`           // 是否必填: 是;长度：128;最终受益人身份证件号;最终受益人身份证件号
	PapersPhotoFront   string `json:"papersPhotoFront"`   // 是否必填: 是;长度：256;最终受益人身份证正面;调用资源图片上传接口获取imgKey 并上传
	PapersPhotoReverse string `json:"papersPhotoReverse"` // 是否必填: 是;长度：256;最终受益人身份证反面;调用资源图片上传接口获取imgKey 并上传
	ValidStartTime     string `json:"validStartTime"`     // 是否必填: ;长度：10;最终受益人证件有效期开始时间;最终受益人身份证件有效期开始时间日期格式应满足合法的YYYY-MM-DD格式
	ValidEndTime       string `json:"validEndTime"`       // 是否必填: 是;长度：10;最终受益人证件有效期结束时间;最终受益人身份证件有效期结束时间日期格式应满足合法的YYYY-MM-DD格式若为“长期”，传值为：当前时间+100 年
	ResidentialAddress string `json:"residentialAddress"` // 是否必填: 是;长度：256;最终受益人居住地址;最终受益人居住地址取身份证件上的居住地址信息
}

type CreateAccountRes struct { // 进件结果
	AccountNo string `json:"account_no"`           // 进件申请凭证，用于查询进件开通进度
	SubMchId  string `json:"sub_mch_id,omitempty"` // 进件完成，生成的有赞支付商户号；又名特约商户、收款商户商户号
	Status    string `json:"status"`               // INIT :   初始化 IN_OPEN:  开户审核中 OPENED :  开户完成 REJECT： 开户驳回
}

type QueryCreateAccountStatusReq struct { // 查询上级进件结果请求
	OutBizNo string `json:"out_biz_no"`
	MchId    string `json:"mch_id"`
}

type QueryCreateAccountStatusRes struct { // 查询上级进件结果响应
	AccountNo      string            `json:"account_no"`
	SubMchId       string            `json:"sub_mch_id"`
	Status         string            `json:"status"`
	FailDesc       string            `json:"fail_desc"`
	FailDetailDesc map[string]string `json:"fail_detail_desc"`
}

type ChangeMchShortNameReq struct {
	SubMchId          string `json:"sub_mch_id,omitempty"`
	ChannelSubMchId   string `json:"channel_sub_mch_id,omitempty"`
	MerchantShortName string `json:"merchant_short_name,omitempty"`
}
