// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package trvael

const VersionNo = "20131202"


/* 旅游商品结构。 */
type TravelItems struct {
	ApproveStatus string `json:"approve_status"`
	AuctionPoint int `json:"auction_point"`
	Cid int `json:"cid"`
	Created string `json:"created"`
	DelistTime string `json:"delist_time"`
	Desc string `json:"desc"`
	DetailUrl string `json:"detail_url"`
	Duration int `json:"duration"`
	FeeInclude string `json:"fee_include"`
	FeeNotInclude string `json:"fee_not_include"`
	FlightInfo string `json:"flight_info"`
	FreightPayer string `json:"freight_payer"`
	GatheringPlace string `json:"gathering_place"`
	HasDiscount bool `json:"has_discount"`
	HasInvoice bool `json:"has_invoice"`
	HasShowcase bool `json:"has_showcase"`
	HotelInfo string `json:"hotel_info"`
	InputPids string `json:"input_pids"`
	InputStr string `json:"input_str"`
	IsTdcy bool `json:"is_tdcy"`
	IsTiming bool `json:"is_timing"`
	ItemId int `json:"item_id"`
	ItemImgs *TravelItemsImgObject `json:"item_imgs"`
	ListTime string `json:"list_time"`
	LocalityLife *TravelItemsLocalityLife `json:"locality_life"`
	Location string `json:"location"`
	MaxShowcase int `json:"max_showcase"`
	Modified string `json:"modified"`
	Nick string `json:"nick"`
	Num int `json:"num"`
	OrderInfo string `json:"order_info"`
	OuterId string `json:"outer_id"`
	OwnExpense string `json:"own_expense"`
	PicUrl string `json:"pic_url"`
	Price int `json:"price"`
	Props string `json:"props"`
	PropsName string `json:"props_name"`
	RefundRegulation string `json:"refund_regulation"`
	SecondKill string `json:"second_kill"`
	SellerCids string `json:"seller_cids"`
	ShopingInfo string `json:"shoping_info"`
	Skus *TravelItemsSkuObject `json:"skus"`
	Start string `json:"start"`
	SubStock int `json:"sub_stock"`
	TicketInfo string `json:"ticket_info"`
	Title string `json:"title"`
	TravelItemsCombos *TravelItemsComboObject `json:"travel_items_combos"`
	Type string `json:"type"`
	UsedShowcase int `json:"used_showcase"`
	Violation bool `json:"violation"`
	WapDesc string `json:"wap_desc"`
	WapDetailUrl string `json:"wap_detail_url"`

}

/* 旅游商品结构。 */
type TravelItemsImgObject struct {
	TravelItemsImgs []*TravelItemsImg `json:"travel_items_img"`

}

/* 旅游商品图片结构。 */
type TravelItemsImg struct {
	Created string `json:"created"`
	Id int `json:"id"`
	Position int `json:"position"`
	Url string `json:"url"`

}

/* 旅游度假线路电子凭证（本地生活）结构。 */
type TravelItemsLocalityLife struct {
	ChooseLogis int `json:"choose_logis"`
	Expirydate string `json:"expirydate"`
	Merchant string `json:"merchant"`
	NetworkId string `json:"network_id"`
	OnsaleAutoRefundRatio int `json:"onsale_auto_refund_ratio"`
	RefundRatio int `json:"refund_ratio"`
	Verification string `json:"verification"`

}

/* 旅游商品结构。 */
type TravelItemsSkuObject struct {
	TravelItemsSkus []*TravelItemsSku `json:"travel_items_sku"`

}

/* 旅游商品销售属性SKU结构。 */
type TravelItemsSku struct {
	Created string `json:"created"`
	ItemId int `json:"item_id"`
	Modified string `json:"modified"`
	OuterId string `json:"outer_id"`
	Price int `json:"price"`
	Properties string `json:"properties"`
	PropertiesName string `json:"properties_name"`
	Quantity int `json:"quantity"`
	SkuId int `json:"sku_id"`

}

/* 旅游商品结构。 */
type TravelItemsComboObject struct {
	TravelItemsCombos []*TravelItemsCombo `json:"travel_items_combo"`

}

/* 旅游度假线路套餐价格日历结构。 */
type TravelItemsCombo struct {
	Combo *TravelItemsPropValue `json:"combo"`
	ComboPriceCalendars *TravelItemsPriceCalendarObject `json:"combo_price_calendars"`

}

/* 旅游商品类目属性值结构 */
type TravelItemsPropValue struct {
	Cid int `json:"cid"`
	Name string `json:"name"`
	Pid int `json:"pid"`
	PropName string `json:"prop_name"`
	SortOrder int `json:"sort_order"`
	Vid int `json:"vid"`

}

/* 旅游度假线路套餐价格日历结构。 */
type TravelItemsPriceCalendarObject struct {
	TravelItemsPriceCalendars []*TravelItemsPriceCalendar `json:"travel_items_price_calendar"`

}

/* 旅游度假线路价格日历结构。 */
type TravelItemsPriceCalendar struct {
	ChildNum int `json:"child_num"`
	ChildPrice int `json:"child_price"`
	Date string `json:"date"`
	DiffPrice int `json:"diff_price"`
	ManNum int `json:"man_num"`
	ManPrice int `json:"man_price"`

}

/* 旅游度假商品地区结构。 */
type TravelItemsAreaNode struct {
	SubPropValues *TravelItemsPropValueObject `json:"sub_prop_values"`
	TravelItemsPropValue *TravelItemsPropValue `json:"travel_items_prop_value"`

}

/* 旅游商品类目属性结构 */
type TravelItemsPropValueObject struct {
	TravelItemsPropValues []*TravelItemsPropValue `json:"travel_items_prop_value"`

}

/* 旅游商品类目属性结构 */
type TravelItemsProp struct {
	Cid int `json:"cid"`
	IsEnums bool `json:"is_enums"`
	IsInput bool `json:"is_input"`
	IsSaleProp bool `json:"is_sale_prop"`
	Multi bool `json:"multi"`
	Must bool `json:"must"`
	Name string `json:"name"`
	Pid int `json:"pid"`
	SortOrder int `json:"sort_order"`
	TravelItemsPropValues *TravelItemsPropValueObject `json:"travel_items_prop_values"`

}

