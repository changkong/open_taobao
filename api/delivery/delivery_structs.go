// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package delivery

const VersionNo = "20131207"


/*  */
type AreaListObject struct {
	Area []*Area `json:"area"`

}

/* 地址区域结构 */
type Area struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ParentId int `json:"parent_id"`
	Type int `json:"type"`
	Zip string `json:"zip"`

}

/* 运费模板对象 */
type DeliveryTemplate struct {
	Address string `json:"address"`
	Assumer int `json:"assumer"`
	ConsignAreaId int `json:"consign_area_id"`
	Created string `json:"created"`
	FeeList *TopFeeListObject `json:"fee_list"`
	Modified string `json:"modified"`
	Name string `json:"name"`
	Supports string `json:"supports"`
	TemplateId int `json:"template_id"`
	Valuation int `json:"valuation"`

}

/*  */
type TopFeeListObject struct {
	TopFee []*TopFee `json:"top_fee"`

}

/* 运费模板中运费信息对象 */
type TopFee struct {
	AddFee string `json:"add_fee"`
	AddStandard string `json:"add_standard"`
	Destination string `json:"destination"`
	ServiceType string `json:"service_type"`
	StartFee string `json:"start_fee"`
	StartStandard string `json:"start_standard"`

}

/*  */
type DeliveryTemplateListObject struct {
	DeliveryTemplate []*DeliveryTemplate `json:"delivery_template"`

}

/* 地址库返回数据信息 */
type AddressResult struct {
	Addr string `json:"addr"`
	AreaId int `json:"area_id"`
	CancelDef bool `json:"cancel_def"`
	City string `json:"city"`
	ContactId int `json:"contact_id"`
	ContactName string `json:"contact_name"`
	Country string `json:"country"`
	GetDef bool `json:"get_def"`
	Memo string `json:"memo"`
	MobilePhone string `json:"mobile_phone"`
	ModifyDate string `json:"modify_date"`
	Phone string `json:"phone"`
	Province string `json:"province"`
	SellerCompany string `json:"seller_company"`
	SendDef bool `json:"send_def"`
	ZipCode string `json:"zip_code"`

}

/*  */
type AddressResultListObject struct {
	AddressResult []*AddressResult `json:"address_result"`

}

/*  */
type LogisticsCompanyListObject struct {
	LogisticsCompany []*LogisticsCompany `json:"logistics_company"`

}

/* 物流公司基础数据结构 */
type LogisticsCompany struct {
	Code string `json:"code"`
	Id int `json:"id"`
	Name string `json:"name"`
	RegMailNo string `json:"reg_mail_no"`

}

/* 获取的物流订单详情列表 返回的Shipping包含的具体信息为入参fields请求的字段信息 */
type Shipping struct {
	BuyerNick string `json:"buyer_nick"`
	CompanyName string `json:"company_name"`
	Created string `json:"created"`
	DeliveryEnd string `json:"delivery_end"`
	DeliveryStart string `json:"delivery_start"`
	FreightPayer string `json:"freight_payer"`
	IsQuickCodOrder bool `json:"is_quick_cod_order"`
	IsSpilt int `json:"is_spilt"`
	IsSuccess bool `json:"is_success"`
	ItemTitle string `json:"item_title"`
	Location *Location `json:"location"`
	Modified string `json:"modified"`
	OrderCode string `json:"order_code"`
	OutSid string `json:"out_sid"`
	ReceiverMobile string `json:"receiver_mobile"`
	ReceiverName string `json:"receiver_name"`
	ReceiverPhone string `json:"receiver_phone"`
	SellerConfirm string `json:"seller_confirm"`
	SellerNick string `json:"seller_nick"`
	Status string `json:"status"`
	SubTids []int `json:"sub_tids"`
	Tid int `json:"tid"`
	Type string `json:"type"`

}

/* 用户地址 */
type Location struct {
	Address string `json:"address"`
	City string `json:"city"`
	Country string `json:"country"`
	District string `json:"district"`
	State string `json:"state"`
	Zip string `json:"zip"`

}

/*  */
type ShippingListObject struct {
	Shipping []*Shipping `json:"shipping"`

}

/*  */
type LogisticsPartnerListObject struct {
	LogisticsPartner []*LogisticsPartner `json:"logistics_partner"`

}

/* 查询揽送范围之内的物流公司信息 */
type LogisticsPartner struct {
	Carriage *CarriageDetail `json:"carriage"`
	CoverRemark string `json:"cover_remark"`
	Partner *PartnerDetail `json:"partner"`
	UncoverRemark string `json:"uncover_remark"`

}

/* 物流公司资费相关信息 */
type CarriageDetail struct {
	AddFee int `json:"add_fee"`
	AddWeight int `json:"add_weight"`
	DamagePayment string `json:"damage_payment"`
	GotTime string `json:"got_time"`
	InitialFee int `json:"initial_fee"`
	InitialWeight int `json:"initial_weight"`
	LostPayment string `json:"lost_payment"`
	WayDay string `json:"way_day"`

}

/* 物流公司详细信息 */
type PartnerDetail struct {
	AccountNo string `json:"account_no"`
	CompanyCode string `json:"company_code"`
	CompanyId int `json:"company_id"`
	CompanyName string `json:"company_name"`
	FullName string `json:"full_name"`
	RegMailNo string `json:"reg_mail_no"`
	WangwangId string `json:"wangwang_id"`

}

/*  */
type TransitStepInfoListObject struct {
	TransitStepInfo []*TransitStepInfo `json:"transit_step_info"`

}

/* 物流跟踪信息的一条 */
type TransitStepInfo struct {
	Action string `json:"action"`
	Desc string `json:"desc"`
	NodeDescription string `json:"node_description"`
	StatusDesc string `json:"status_desc"`
	StatusTime string `json:"status_time"`
	Time string `json:"time"`

}

/*  */
type AddressReachableResultListObject struct {
	AddressReachableResult []*AddressReachableResult `json:"address_reachable_result"`

}

/* 判定服务是否可达的返回结果 */
type AddressReachableResult struct {
	ErrorCode string `json:"error_code"`
	PartnerId int `json:"partner_id"`
	Reachable int `json:"reachable"`
	ServiceType int `json:"service_type"`

}

