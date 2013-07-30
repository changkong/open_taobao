// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

const VersionNo = "20130729"

/* 淘宝客商品 */
type TaobaokeItem struct {
	ClickUrl            string  `json:"click_url"`
	Commission          float64 `json:"commission"`
	CommissionNum       string  `json:"commission_num"`
	CommissionRate      string  `json:"commission_rate"`
	CommissionVolume    float64 `json:"commission_volume"`
	CouponEndTime       string  `json:"coupon_end_time"`
	CouponPrice         float64 `json:"coupon_price"`
	CouponRate          string  `json:"coupon_rate"`
	CouponStartTime     string  `json:"coupon_start_time"`
	ItemLocation        string  `json:"item_location"`
	KeywordClickUrl     string  `json:"keyword_click_url"`
	Nick                string  `json:"nick"`
	NumIid              int     `json:"num_iid"`
	PicUrl              string  `json:"pic_url"`
	Price               float64 `json:"price"`
	PromotionPrice      float64 `json:"promotion_price"`
	SellerCreditScore   int     `json:"seller_credit_score"`
	SellerId            int     `json:"seller_id"`
	ShopClickUrl        string  `json:"shop_click_url"`
	ShopType            string  `json:"shop_type"`
	TaobaokeCatClickUrl string  `json:"taobaoke_cat_click_url"`
	Title               string  `json:"title"`
	Volume              int     `json:"volume"`
}

/* 淘宝客商品详情 */
type TaobaokeItemDetail struct {
	ClickUrl          string `json:"click_url"`
	Item              *Item  `json:"item"`
	SellerCreditScore int    `json:"seller_credit_score"`
	ShopClickUrl      string `json:"shop_click_url"`
}

/* Item(商品)结构 */
type Item struct {
	AfterSaleId                int             `json:"after_sale_id"`
	ApproveStatus              string          `json:"approve_status"`
	AuctionPoint               int             `json:"auction_point"`
	AutoFill                   string          `json:"auto_fill"`
	ChangeProp                 string          `json:"change_prop"`
	Cid                        int             `json:"cid"`
	CodPostageId               int             `json:"cod_postage_id"`
	Created                    string          `json:"created"`
	CustomMadeTypeId           string          `json:"custom_made_type_id"`
	DelistTime                 string          `json:"delist_time"`
	DeliveryTime               *DeliveryTime   `json:"delivery_time"`
	Desc                       string          `json:"desc"`
	DescModuleInfo             *DescModuleInfo `json:"desc_module_info"`
	DescModules                string          `json:"desc_modules"`
	DetailUrl                  string          `json:"detail_url"`
	EmsFee                     float64         `json:"ems_fee"`
	ExpressFee                 float64         `json:"express_fee"`
	Features                   string          `json:"features"`
	FoodSecurity               *FoodSecurity   `json:"food_security"`
	FreightPayer               string          `json:"freight_payer"`
	GlobalStockType            string          `json:"global_stock_type"`
	HasDiscount                bool            `json:"has_discount"`
	HasInvoice                 bool            `json:"has_invoice"`
	HasShowcase                bool            `json:"has_showcase"`
	HasWarranty                bool            `json:"has_warranty"`
	Increment                  string          `json:"increment"`
	InnerShopAuctionTemplateId int             `json:"inner_shop_auction_template_id"`
	InputPids                  string          `json:"input_pids"`
	InputStr                   string          `json:"input_str"`
	Is3D                       bool            `json:"is_3D"`
	IsEx                       bool            `json:"is_ex"`
	IsFenxiao                  int             `json:"is_fenxiao"`
	IsLightningConsignment     bool            `json:"is_lightning_consignment"`
	IsPrepay                   bool            `json:"is_prepay"`
	IsTaobao                   bool            `json:"is_taobao"`
	IsTiming                   bool            `json:"is_timing"`
	IsVirtual                  bool            `json:"is_virtual"`
	IsXinpin                   bool            `json:"is_xinpin"`
	ItemImgs                   []*ItemImg      `json:"item_imgs"`
	ItemSize                   string          `json:"item_size"`
	ItemWeight                 string          `json:"item_weight"`
	ListTime                   string          `json:"list_time"`
	LocalityLife               *LocalityLife   `json:"locality_life"`
	Location                   *Location       `json:"location"`
	Modified                   string          `json:"modified"`
	Nick                       string          `json:"nick"`
	Num                        int             `json:"num"`
	NumIid                     int             `json:"num_iid"`
	OneStation                 bool            `json:"one_station"`
	OuterId                    string          `json:"outer_id"`
	OuterShopAuctionTemplateId int             `json:"outer_shop_auction_template_id"`
	PaimaiInfo                 *PaimaiInfo     `json:"paimai_info"`
	PicUrl                     string          `json:"pic_url"`
	PostFee                    float64         `json:"post_fee"`
	PostageId                  int             `json:"postage_id"`
	Price                      float64         `json:"price"`
	ProductId                  int             `json:"product_id"`
	PromotedService            string          `json:"promoted_service"`
	PropImgs                   []*PropImg      `json:"prop_imgs"`
	PropertyAlias              string          `json:"property_alias"`
	Props                      string          `json:"props"`
	PropsName                  string          `json:"props_name"`
	Score                      int             `json:"score"`
	SecondKill                 string          `json:"second_kill"`
	SellPoint                  string          `json:"sell_point"`
	SellPromise                bool            `json:"sell_promise"`
	SellerCids                 string          `json:"seller_cids"`
	Skus                       []*Sku          `json:"skus"`
	StuffStatus                string          `json:"stuff_status"`
	SubStock                   int             `json:"sub_stock"`
	TemplateId                 string          `json:"template_id"`
	Title                      string          `json:"title"`
	Type                       string          `json:"type"`
	ValidThru                  int             `json:"valid_thru"`
	Videos                     []*Video        `json:"videos"`
	Violation                  bool            `json:"violation"`
	WapDesc                    string          `json:"wap_desc"`
	WapDetailUrl               string          `json:"wap_detail_url"`
	WithHoldQuantity           int             `json:"with_hold_quantity"`
	WwStatus                   bool            `json:"ww_status"`
}

/* 发货时间数据结构 */
type DeliveryTime struct {
	DeliveryTime     string `json:"delivery_time"`
	DeliveryTimeType string `json:"delivery_time_type"`
	NeedDeliveryTime string `json:"need_delivery_time"`
}

/* 该数据结构保存宝贝描述对应的规范化信息 */
type DescModuleInfo struct {
	AnchorModuleIds string `json:"anchor_module_ids"`
	Type            int    `json:"type"`
}

/* 食品安全信息，包括：
生产许可证号、产品标准号、厂名、厂址等 */
type FoodSecurity struct {
	Contact          string `json:"contact"`
	DesignCode       string `json:"design_code"`
	Factory          string `json:"factory"`
	FactorySite      string `json:"factory_site"`
	FoodAdditive     string `json:"food_additive"`
	HealthProductNo  string `json:"health_product_no"`
	Mix              string `json:"mix"`
	Period           string `json:"period"`
	PlanStorage      string `json:"plan_storage"`
	PrdLicenseNo     string `json:"prd_license_no"`
	ProductDateEnd   string `json:"product_date_end"`
	ProductDateStart string `json:"product_date_start"`
	StockDateEnd     string `json:"stock_date_end"`
	StockDateStart   string `json:"stock_date_start"`
	Supplier         string `json:"supplier"`
}

/* ItemImg结构 */
type ItemImg struct {
	Created  string `json:"created"`
	Id       int    `json:"id"`
	Position int    `json:"position"`
	Url      string `json:"url"`
}

/* 本地生活垂直市场数据结构，修改宝贝时在参数empty_fields里设置locality_life可删除所有电子凭证信息 */
type LocalityLife struct {
	ChooseLogis           string `json:"choose_logis"`
	Expirydate            string `json:"expirydate"`
	Merchant              string `json:"merchant"`
	NetworkId             string `json:"network_id"`
	OnsaleAutoRefundRatio int    `json:"onsale_auto_refund_ratio"`
	RefundRatio           int    `json:"refund_ratio"`
	Verification          string `json:"verification"`
}

/* 用户地址 */
type Location struct {
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	District string `json:"district"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
}

/* 拍卖商品相关信息 */
type PaimaiInfo struct {
	Deposit     int     `json:"deposit"`
	Interval    int     `json:"interval"`
	Mode        int     `json:"mode"`
	Reserve     float64 `json:"reserve"`
	ValidHour   int     `json:"valid_hour"`
	ValidMinute int     `json:"valid_minute"`
}

/* 商品属性图片结构 */
type PropImg struct {
	Created    string `json:"created"`
	Id         int    `json:"id"`
	Position   int    `json:"position"`
	Properties string `json:"properties"`
	Url        string `json:"url"`
}

/* Sku结构 */
type Sku struct {
	ChangeProp       string `json:"change_prop"`
	Created          string `json:"created"`
	Iid              string `json:"iid"`
	Modified         string `json:"modified"`
	NumIid           int    `json:"num_iid"`
	OuterId          string `json:"outer_id"`
	Price            string `json:"price"`
	Properties       string `json:"properties"`
	PropertiesName   string `json:"properties_name"`
	Quantity         int    `json:"quantity"`
	SkuDeliveryTime  string `json:"sku_delivery_time"`
	SkuId            int    `json:"sku_id"`
	SkuSpecId        int    `json:"sku_spec_id"`
	Status           string `json:"status"`
	WithHoldQuantity int    `json:"with_hold_quantity"`
}

/* 商品视频关联记录 */
type Video struct {
	Created  string `json:"created"`
	Id       int    `json:"id"`
	Iid      string `json:"iid"`
	Modified string `json:"modified"`
	NumIid   int    `json:"num_iid"`
	Url      string `json:"url"`
	VideoId  int    `json:"video_id"`
}

/* 淘宝客订单 */
type TaobaokePayment struct {
	AppKey         string  `json:"app_key"`
	CategoryId     int     `json:"category_id"`
	CategoryName   string  `json:"category_name"`
	Commission     float64 `json:"commission"`
	CommissionRate string  `json:"commission_rate"`
	CreateTime     string  `json:"create_time"`
	ItemNum        int     `json:"item_num"`
	ItemTitle      string  `json:"item_title"`
	NumIid         int     `json:"num_iid"`
	OuterCode      string  `json:"outer_code"`
	PayPrice       float64 `json:"pay_price"`
	PayTime        string  `json:"pay_time"`
	RealPayFee     float64 `json:"real_pay_fee"`
	SellerNick     string  `json:"seller_nick"`
	ShopTitle      string  `json:"shop_title"`
	TradeId        int     `json:"trade_id"`
	TradeParentId  int     `json:"trade_parent_id"`
}

/* 淘宝客报表 */
type TaobaokeReport struct {
	TaobaokeReportMembers []*TaobaokeReportMember `json:"taobaoke_report_members"`
	TotalResults          int                     `json:"total_results"`
}

/* 淘宝客报表成员 */
type TaobaokeReportMember struct {
	AppKey         string  `json:"app_key"`
	CategoryId     int     `json:"category_id"`
	CategoryName   string  `json:"category_name"`
	Commission     float64 `json:"commission"`
	CommissionRate string  `json:"commission_rate"`
	CreateTime     string  `json:"create_time"`
	ItemNum        int     `json:"item_num"`
	ItemTitle      string  `json:"item_title"`
	NumIid         int     `json:"num_iid"`
	OuterCode      string  `json:"outer_code"`
	PayPrice       float64 `json:"pay_price"`
	PayTime        string  `json:"pay_time"`
	RealPayFee     float64 `json:"real_pay_fee"`
	SellerNick     string  `json:"seller_nick"`
	ShopTitle      string  `json:"shop_title"`
	TradeId        int     `json:"trade_id"`
	TradeParentId  int     `json:"trade_parent_id"`
}

/* 淘宝客店铺 */
type TaobaokeShop struct {
	AuctionCount   int    `json:"auction_count"`
	ClickUrl       string `json:"click_url"`
	CommissionRate string `json:"commission_rate"`
	SellerCredit   string `json:"seller_credit"`
	SellerNick     string `json:"seller_nick"`
	ShopTitle      string `json:"shop_title"`
	ShopType       string `json:"shop_type"`
	TotalAuction   string `json:"total_auction"`
	UserId         int    `json:"user_id"`
}
