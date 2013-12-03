// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jipiao

const VersionNo = "20131202"


/* 全量、增量的添加政策压缩包的json对应对象 */
type BatchPolicy struct {
	Airline string `json:"airline"`
	ArrAirports string `json:"arr_airports"`
	Attributes string `json:"attributes"`
	AutoHkFlag bool `json:"auto_hk_flag"`
	AutoTicketFlag bool `json:"auto_ticket_flag"`
	CabinRules string `json:"cabin_rules"`
	ChangeRule string `json:"change_rule"`
	DayOfWeeks string `json:"day_of_weeks"`
	DepAirports string `json:"dep_airports"`
	Ei string `json:"ei"`
	ExcludeDate string `json:"exclude_date"`
	FirstSaleAdvanceDay int `json:"first_sale_advance_day"`
	Flags int `json:"flags"`
	FlightInfo string `json:"flight_info"`
	LastSaleAdvanceDay int `json:"last_sale_advance_day"`
	Memo string `json:"memo"`
	OfficeId string `json:"office_id"`
	OutProductId string `json:"out_product_id"`
	PolicyType int `json:"policy_type"`
	RefundRule string `json:"refund_rule"`
	ReissueRule string `json:"reissue_rule"`
	SaleEndDate string `json:"sale_end_date"`
	SaleStartDate string `json:"sale_start_date"`
	SeatInfo string `json:"seat_info"`
	ShareSupport bool `json:"share_support"`
	Source string `json:"source"`
	TravelEndDate string `json:"travel_end_date"`
	TravelStartDate string `json:"travel_start_date"`

}

/* 机票产品政策数据对象 */
type Policy struct {
	AgentId int `json:"agent_id"`
	Airline string `json:"airline"`
	ArrAirports string `json:"arr_airports"`
	Attributes string `json:"attributes"`
	AutoHkFlag bool `json:"auto_hk_flag"`
	AutoTicketFlag bool `json:"auto_ticket_flag"`
	CabinRules string `json:"cabin_rules"`
	DepAirports string `json:"dep_airports"`
	FirstSaleAdvanceDay int `json:"first_sale_advance_day"`
	Flags int `json:"flags"`
	FlightInfo string `json:"flight_info"`
	Id int `json:"id"`
	LastSaleAdvanceDay int `json:"last_sale_advance_day"`
	OutProductId string `json:"out_product_id"`
	PolicyDetail *PolicyDetail `json:"policy_detail"`
	PolicyType int `json:"policy_type"`
	SaleEndDate string `json:"sale_end_date"`
	SaleStartDate string `json:"sale_start_date"`
	SeatInfo string `json:"seat_info"`
	ShareSupport bool `json:"share_support"`
	Status int `json:"status"`
	TravelEndDate string `json:"travel_end_date"`
	TravelStartDate string `json:"travel_start_date"`
	TripType int `json:"trip_type"`

}

/* 淘宝机票政策详情信息 */
type PolicyDetail struct {
	ChangeRule string `json:"change_rule"`
	DayOfWeeks string `json:"day_of_weeks"`
	Ei string `json:"ei"`
	ExcludeDate string `json:"exclude_date"`
	Memo string `json:"memo"`
	OfficeId string `json:"office_id"`
	RefundRule string `json:"refund_rule"`
	ReissueRule string `json:"reissue_rule"`
	SpecialRule string `json:"special_rule"`

}

/* 代理商订单搜索接口返回数据对象【订单优化】 */
type SearchOrderResult struct {
	HasNext bool `json:"has_next"`
	OrderIds []int `json:"order_ids"`
	PageSize int `json:"page_size"`

}

/* 国内机票订单数据结构【top订单优化】 */
type AtOrder struct {
	BaseInfo *BaseInfo `json:"base_info"`
	CorpInfo *CorpInfo `json:"corp_info"`
	Extra string `json:"extra"`
	Itinerary *Itinerary `json:"itinerary"`
	SegmentInfos *SegmentInfoObject `json:"segment_infos"`

}

/* 国内机票，订单基本信息数据结构【top订单优化】 */
type BaseInfo struct {
	AccountNo string `json:"account_no"`
	AlipayTradeNo string `json:"alipay_trade_no"`
	BookWay int `json:"book_way"`
	Commission string `json:"commission"`
	CommissionDiscount string `json:"commission_discount"`
	CreateTime string `json:"create_time"`
	Extra string `json:"extra"`
	ForceInsure bool `json:"force_insure"`
	InsurePromotion bool `json:"insure_promotion"`
	ModifyTime string `json:"modify_time"`
	OrderId int `json:"order_id"`
	PayLatestTime string `json:"pay_latest_time"`
	PayStatus int `json:"pay_status"`
	RelationEmail string `json:"relation_email"`
	RelationMobile string `json:"relation_mobile"`
	RelationName string `json:"relation_name"`
	RelationPhoneBak string `json:"relation_phone_bak"`
	RelativeOrderId int `json:"relative_order_id"`
	Status int `json:"status"`
	TotalPrice int `json:"total_price"`
	TripType int `json:"trip_type"`

}

/* 国内机票订单行程购票数据结构录入【top订单优化】 */
type CorpInfo struct {
	ApplyName string `json:"apply_name"`
	ApplyNo string `json:"apply_no"`
	ApplyTime string `json:"apply_time"`
	CorprationId string `json:"corpration_id"`
	CostCenter string `json:"cost_center"`
	CostCenterCode string `json:"cost_center_code"`
	CostOu string `json:"cost_ou"`
	Extra string `json:"extra"`
	FormNo string `json:"form_no"`
	FormStatus string `json:"form_status"`
	ReceiptsStatus string `json:"receipts_status"`
	TripPersonName string `json:"trip_person_name"`
	TripPersonNo string `json:"trip_person_no"`
	WorkSpace string `json:"work_space"`

}

/* 国内机票行程单数据结构定义【top订单优化】 */
type Itinerary struct {
	Address string `json:"address"`
	AlipayTradeNo string `json:"alipay_trade_no"`
	CompanyCode string `json:"company_code"`
	ExpressCode string `json:"express_code"`
	Extra string `json:"extra"`
	Id int `json:"id"`
	ItineraryNo string `json:"itinerary_no"`
	Mobile string `json:"mobile"`
	MobileBak string `json:"mobile_bak"`
	Name string `json:"name"`
	Price string `json:"price"`
	SendDate string `json:"send_date"`
	Status int `json:"status"`
	Type int `json:"type"`

}

/* 国内机票订单数据结构【top订单优化】 */
type SegmentInfoObject struct {
	SegmentInfos []*SegmentInfo `json:"segment_info"`

}

/* 国内机票航段信息数据结构，【订单优化】 */
type SegmentInfo struct {
	AirlineCode string `json:"airline_code"`
	ArrAirportCode string `json:"arr_airport_code"`
	ArrCityCode string `json:"arr_city_code"`
	ArrTime string `json:"arr_time"`
	BookStatus int `json:"book_status"`
	CabinClass int `json:"cabin_class"`
	CabinCode string `json:"cabin_code"`
	CabinId int `json:"cabin_id"`
	Carrier string `json:"carrier"`
	ChildFee int `json:"child_fee"`
	ChildInsurePromotionPrice int `json:"child_insure_promotion_price"`
	ChildPrice int `json:"child_price"`
	ChildTax int `json:"child_tax"`
	DepAirportCode string `json:"dep_airport_code"`
	DepCityCode string `json:"dep_city_code"`
	DepTime string `json:"dep_time"`
	Extra string `json:"extra"`
	Fee int `json:"fee"`
	FlightId int `json:"flight_id"`
	FlightNo string `json:"flight_no"`
	FlightType string `json:"flight_type"`
	InsurePromotionPrice int `json:"insure_promotion_price"`
	Memo string `json:"memo"`
	Passengers *PassergerObject `json:"passengers"`
	PolicyId int `json:"policy_id"`
	PolicyType int `json:"policy_type"`
	Price int `json:"price"`
	SegmentType int `json:"segment_type"`
	SpecialRule string `json:"special_rule"`
	Tax int `json:"tax"`
	TicketPrice int `json:"ticket_price"`

}

/* 国内机票航段信息数据结构，【订单优化】 */
type PassergerObject struct {
	Passergers []*Passerger `json:"passerger"`

}

/* 国内机票乘机人信息数据结构【top订单优化】 */
type Passerger struct {
	Birthday string `json:"birthday"`
	CertNo string `json:"cert_no"`
	CertType int `json:"cert_type"`
	Ei string `json:"ei"`
	Extra string `json:"extra"`
	ForceInsurePrice int `json:"force_insure_price"`
	InsurePromotionPrice int `json:"insure_promotion_price"`
	Name string `json:"name"`
	PassengerType int `json:"passenger_type"`
	Pnr string `json:"pnr"`
	TicketNo string `json:"ticket_no"`
	TripCardNo string `json:"trip_card_no"`
	Tuigaiqian string `json:"tuigaiqian"`

}

