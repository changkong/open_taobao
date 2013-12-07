// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package product

const VersionNo = "20131207"


/*  */
type AfterSaleListObject struct {
	AfterSale []*AfterSale `json:"after_sale"`

}

/* 卖家设置售后服务对象 */
type AfterSale struct {
	AfterSaleId int `json:"after_sale_id"`
	AfterSaleName string `json:"after_sale_name"`
	AfterSalePath string `json:"after_sale_path"`

}

/* Item(商品)结构 */
type Item struct {
	AfterSaleId int `json:"after_sale_id"`
	ApproveStatus string `json:"approve_status"`
	AuctionPoint int `json:"auction_point"`
	AutoFill string `json:"auto_fill"`
	ChangeProp string `json:"change_prop"`
	Cid int `json:"cid"`
	CodPostageId int `json:"cod_postage_id"`
	Created string `json:"created"`
	CustomMadeTypeId string `json:"custom_made_type_id"`
	DelistTime string `json:"delist_time"`
	DeliveryTime *DeliveryTime `json:"delivery_time"`
	Desc string `json:"desc"`
	DescModuleInfo *DescModuleInfo `json:"desc_module_info"`
	DescModules string `json:"desc_modules"`
	DetailUrl string `json:"detail_url"`
	EmsFee float64 `json:"ems_fee"`
	ExpressFee float64 `json:"express_fee"`
	Features string `json:"features"`
	FoodSecurity *FoodSecurity `json:"food_security"`
	FreightPayer string `json:"freight_payer"`
	GlobalStockCountry string `json:"global_stock_country"`
	GlobalStockType string `json:"global_stock_type"`
	HasDiscount bool `json:"has_discount"`
	HasInvoice bool `json:"has_invoice"`
	HasShowcase bool `json:"has_showcase"`
	HasWarranty bool `json:"has_warranty"`
	Increment string `json:"increment"`
	InnerShopAuctionTemplateId int `json:"inner_shop_auction_template_id"`
	InputPids string `json:"input_pids"`
	InputStr string `json:"input_str"`
	Is3D bool `json:"is_3D"`
	IsEx bool `json:"is_ex"`
	IsFenxiao int `json:"is_fenxiao"`
	IsLightningConsignment bool `json:"is_lightning_consignment"`
	IsPrepay bool `json:"is_prepay"`
	IsTaobao bool `json:"is_taobao"`
	IsTiming bool `json:"is_timing"`
	IsVirtual bool `json:"is_virtual"`
	IsXinpin bool `json:"is_xinpin"`
	ItemImgs *ItemImgListObject `json:"item_imgs"`
	ItemSize string `json:"item_size"`
	ItemWeight string `json:"item_weight"`
	ListTime string `json:"list_time"`
	LocalityLife *LocalityLife `json:"locality_life"`
	Location *Location `json:"location"`
	Modified string `json:"modified"`
	Nick string `json:"nick"`
	Num int `json:"num"`
	NumIid int `json:"num_iid"`
	OneStation bool `json:"one_station"`
	OuterId string `json:"outer_id"`
	OuterShopAuctionTemplateId int `json:"outer_shop_auction_template_id"`
	PaimaiInfo *PaimaiInfo `json:"paimai_info"`
	PicUrl string `json:"pic_url"`
	PostFee float64 `json:"post_fee"`
	PostageId int `json:"postage_id"`
	Price float64 `json:"price"`
	ProductId int `json:"product_id"`
	PromotedService string `json:"promoted_service"`
	PropImgs *PropImgListObject `json:"prop_imgs"`
	PropertyAlias string `json:"property_alias"`
	Props string `json:"props"`
	PropsName string `json:"props_name"`
	Score int `json:"score"`
	SecondKill string `json:"second_kill"`
	SellPoint string `json:"sell_point"`
	SellPromise bool `json:"sell_promise"`
	SellerCids string `json:"seller_cids"`
	Skus *SkuListObject `json:"skus"`
	StuffStatus string `json:"stuff_status"`
	SubStock int `json:"sub_stock"`
	TemplateId string `json:"template_id"`
	Title string `json:"title"`
	Type string `json:"type"`
	ValidThru int `json:"valid_thru"`
	Videos *VideoListObject `json:"videos"`
	Violation bool `json:"violation"`
	WapDesc string `json:"wap_desc"`
	WapDetailUrl string `json:"wap_detail_url"`
	WirelessDesc string `json:"wireless_desc"`
	WithHoldQuantity int `json:"with_hold_quantity"`
	WwStatus bool `json:"ww_status"`

}

/* 发货时间数据结构 */
type DeliveryTime struct {
	DeliveryTime string `json:"delivery_time"`
	DeliveryTimeType string `json:"delivery_time_type"`
	NeedDeliveryTime string `json:"need_delivery_time"`

}

/* 该数据结构保存宝贝描述对应的规范化信息 */
type DescModuleInfo struct {
	AnchorModuleIds string `json:"anchor_module_ids"`
	Type int `json:"type"`

}

/* 食品安全信息，包括： 
生产许可证号、产品标准号、厂名、厂址等 */
type FoodSecurity struct {
	Contact string `json:"contact"`
	DesignCode string `json:"design_code"`
	Factory string `json:"factory"`
	FactorySite string `json:"factory_site"`
	FoodAdditive string `json:"food_additive"`
	HealthProductNo string `json:"health_product_no"`
	Mix string `json:"mix"`
	Period string `json:"period"`
	PlanStorage string `json:"plan_storage"`
	PrdLicenseNo string `json:"prd_license_no"`
	ProductDateEnd string `json:"product_date_end"`
	ProductDateStart string `json:"product_date_start"`
	StockDateEnd string `json:"stock_date_end"`
	StockDateStart string `json:"stock_date_start"`
	Supplier string `json:"supplier"`

}

/*  */
type ItemImgListObject struct {
	ItemImg []*ItemImg `json:"item_img"`

}

/* ItemImg结构 */
type ItemImg struct {
	Created string `json:"created"`
	Id int `json:"id"`
	Position int `json:"position"`
	Url string `json:"url"`

}

/* 本地生活垂直市场数据结构，修改宝贝时在参数empty_fields里设置locality_life可删除所有电子凭证信息 */
type LocalityLife struct {
	ChooseLogis string `json:"choose_logis"`
	Eticket string `json:"eticket"`
	Expirydate string `json:"expirydate"`
	Merchant string `json:"merchant"`
	NetworkId string `json:"network_id"`
	OnsaleAutoRefundRatio int `json:"onsale_auto_refund_ratio"`
	RefundRatio int `json:"refund_ratio"`
	Refundmafee string `json:"refundmafee"`
	Verification string `json:"verification"`

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

/* 拍卖商品相关信息 */
type PaimaiInfo struct {
	Deposit int `json:"deposit"`
	Interval int `json:"interval"`
	Mode int `json:"mode"`
	Reserve float64 `json:"reserve"`
	ValidHour int `json:"valid_hour"`
	ValidMinute int `json:"valid_minute"`

}

/*  */
type PropImgListObject struct {
	PropImg []*PropImg `json:"prop_img"`

}

/* 商品属性图片结构 */
type PropImg struct {
	Created string `json:"created"`
	Id int `json:"id"`
	Position int `json:"position"`
	Properties string `json:"properties"`
	Url string `json:"url"`

}

/*  */
type SkuListObject struct {
	Sku []*Sku `json:"sku"`

}

/* Sku结构 */
type Sku struct {
	ChangeProp string `json:"change_prop"`
	Created string `json:"created"`
	Iid string `json:"iid"`
	Modified string `json:"modified"`
	NumIid int `json:"num_iid"`
	OuterId string `json:"outer_id"`
	Price string `json:"price"`
	Properties string `json:"properties"`
	PropertiesName string `json:"properties_name"`
	Quantity int `json:"quantity"`
	SkuDeliveryTime string `json:"sku_delivery_time"`
	SkuId int `json:"sku_id"`
	SkuSpecId int `json:"sku_spec_id"`
	Status string `json:"status"`
	WithHoldQuantity int `json:"with_hold_quantity"`

}

/*  */
type VideoListObject struct {
	Video []*Video `json:"video"`

}

/* 商品视频关联记录 */
type Video struct {
	Created string `json:"created"`
	Id int `json:"id"`
	Iid string `json:"iid"`
	Modified string `json:"modified"`
	NumIid int `json:"num_iid"`
	Url string `json:"url"`
	VideoId int `json:"video_id"`

}

/*  */
type IdsModuleListObject struct {
	IdsModule []*IdsModule `json:"ids_module"`

}

/* 用于保存宝贝描述规范化模块信息 */
type IdsModule struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type int `json:"type"`

}

/*  */
type ItemTemplateListObject struct {
	ItemTemplate []*ItemTemplate `json:"item_template"`

}

/* 宝贝详情页面信息 */
type ItemTemplate struct {
	ShopType int `json:"shop_type"`
	TemplateId int `json:"template_id"`
	TemplateName string `json:"template_name"`

}

/*  */
type ItemListObject struct {
	Item []*Item `json:"item"`

}

/* 产品结构 */
type Product struct {
	Binds string `json:"binds"`
	BindsStr string `json:"binds_str"`
	CatName string `json:"cat_name"`
	Cid int `json:"cid"`
	CollectNum int `json:"collect_num"`
	Created string `json:"created"`
	CspuFeature string `json:"cspu_feature"`
	CustomerProps string `json:"customer_props"`
	Desc string `json:"desc"`
	Level int `json:"level"`
	Modified string `json:"modified"`
	Name string `json:"name"`
	OuterId string `json:"outer_id"`
	PicPath string `json:"pic_path"`
	PicUrl string `json:"pic_url"`
	Price float64 `json:"price"`
	ProductExtraInfos *ProductExtraInfoListObject `json:"product_extra_infos"`
	ProductId int `json:"product_id"`
	ProductImgs *ProductImgListObject `json:"product_imgs"`
	ProductPropImgs *ProductPropImgListObject `json:"product_prop_imgs"`
	PropertyAlias string `json:"property_alias"`
	Props string `json:"props"`
	PropsStr string `json:"props_str"`
	RateNum int `json:"rate_num"`
	SaleNum int `json:"sale_num"`
	SaleProps string `json:"sale_props"`
	SalePropsStr string `json:"sale_props_str"`
	SellPt string `json:"sell_pt"`
	ShopPrice string `json:"shop_price"`
	StandardPrice string `json:"standard_price"`
	Status int `json:"status"`
	Tsc string `json:"tsc"`
	VerticalMarket int `json:"vertical_market"`

}

/*  */
type ProductExtraInfoListObject struct {
	ProductExtraInfo []*ProductExtraInfo `json:"product_extra_info"`

}

/* 产品扩展信息 */
type ProductExtraInfo struct {
	FieldKey string `json:"field_key"`
	FieldName string `json:"field_name"`
	FieldValue string `json:"field_value"`
	ProductId int `json:"product_id"`

}

/*  */
type ProductImgListObject struct {
	ProductImg []*ProductImg `json:"product_img"`

}

/* 产品图片 */
type ProductImg struct {
	Created string `json:"created"`
	Id int `json:"id"`
	Modified string `json:"modified"`
	Position int `json:"position"`
	ProductId int `json:"product_id"`
	Url string `json:"url"`

}

/*  */
type ProductPropImgListObject struct {
	ProductPropImg []*ProductPropImg `json:"product_prop_img"`

}

/* 产品属性图片 */
type ProductPropImg struct {
	Created string `json:"created"`
	Id int `json:"id"`
	Modified string `json:"modified"`
	Position int `json:"position"`
	ProductId int `json:"product_id"`
	Props string `json:"props"`
	Url string `json:"url"`

}

/*  */
type ProductListObject struct {
	Product []*Product `json:"product"`

}

/* 批量异步任务结果 */
type Task struct {
	CheckCode string `json:"check_code"`
	Created string `json:"created"`
	DownloadUrl string `json:"download_url"`
	Method string `json:"method"`
	Schedule string `json:"schedule"`
	Status string `json:"status"`
	Subtasks *SubtaskListObject `json:"subtasks"`
	TaskId int `json:"task_id"`

}

/*  */
type SubtaskListObject struct {
	Subtask []*Subtask `json:"subtask"`

}

/* 批量异步任务的子任务结果 */
type Subtask struct {
	IsSuccess bool `json:"is_success"`
	SubTaskRequest string `json:"sub_task_request"`
	SubTaskResult string `json:"sub_task_result"`

}

/* 优惠信息对象 */
type PromotionDisplayTop struct {
	PromotionInItem *PromotionInItemListObject `json:"promotion_in_item"`
	PromotionInShop *PromotionInShopListObject `json:"promotion_in_shop"`

}

/*  */
type PromotionInItemListObject struct {
	PromotionInItem []*PromotionInItem `json:"promotion_in_item"`

}

/* 单品级优惠信息 */
type PromotionInItem struct {
	Desc string `json:"desc"`
	EndTime string `json:"end_time"`
	ItemPromoPrice float64 `json:"item_promo_price"`
	Name string `json:"name"`
	OtherNeed string `json:"other_need"`
	OtherSend string `json:"other_send"`
	PromotionId string `json:"promotion_id"`
	SkuIdList []string `json:"sku_id_list"`
	SkuPriceList []float64 `json:"sku_price_list"`
	StartTime string `json:"start_time"`

}

/*  */
type PromotionInShopListObject struct {
	PromotionInShop []*PromotionInShop `json:"promotion_in_shop"`

}

/* 店铺级优惠信息 */
type PromotionInShop struct {
	Name string `json:"name"`
	PromotionDetailDesc string `json:"promotion_detail_desc"`
	PromotionId string `json:"promotion_id"`

}

/* 管控的类目以及品牌信息 */
type BrandCatControlInfo struct {
	BrandCatControls *BrandCatControlListObject `json:"brand_cat_controls"`

}

/*  */
type BrandCatControlListObject struct {
	BrandCatControl []*BrandCatControl `json:"brand_cat_control"`

}

/* 管控的品牌类目信息 */
type BrandCatControl struct {
	BrandId int `json:"brand_id"`
	BrandName string `json:"brand_name"`
	CatId int `json:"cat_id"`
	CatName string `json:"cat_name"`
	CertifiedData string `json:"certified_data"`

}

/* 属性输入特征DO */
type PropertyInputDO struct {
	IsAllowInput bool `json:"is_allow_input"`
	IsRootAllowInput bool `json:"is_root_allow_input"`
	IsSubProperty bool `json:"is_sub_property"`
	PropertyId int `json:"property_id"`

}

/*  */
type CatBrandSalePropListObject struct {
	CatBrandSaleProp []*CatBrandSaleProp `json:"cat_brand_sale_prop"`

}

/* 被管控的品牌和类目的所对应的销售属性 */
type CatBrandSaleProp struct {
	BrandId int `json:"brand_id"`
	CatId int `json:"cat_id"`
	DefMarketPropValue int `json:"def_market_prop_value"`
	IsNotSpec bool `json:"is_not_spec"`
	PropertyId int `json:"property_id"`

}

/* ProductSpec(产品规格)结构。 */
type ProductSpec struct {
	Barcode string `json:"barcode"`
	BrandId int `json:"brand_id"`
	CertifiedPics *CertPicInfoListObject `json:"certified_pics"`
	CertifiedTxts *CertTxtInfoListObject `json:"certified_txts"`
	ChangeProp string `json:"change_prop"`
	CustomePropsName string `json:"custome_props_name"`
	LabelPrice int `json:"label_price"`
	MarketTime string `json:"market_time"`
	PicUrl string `json:"pic_url"`
	ProductCode string `json:"product_code"`
	ProductId int `json:"product_id"`
	SpecId int `json:"spec_id"`
	SpecProps string `json:"spec_props"`
	SpecPropsAlias string `json:"spec_props_alias"`
	Status int `json:"status"`

}

/*  */
type CertPicInfoListObject struct {
	CertPicInfo []*CertPicInfo `json:"cert_pic_info"`

}

/* 产品资质认证图片信息，包括认证类型以及图片url */
type CertPicInfo struct {
	CertType int `json:"cert_type"`
	PicUrl string `json:"pic_url"`

}

/*  */
type CertTxtInfoListObject struct {
	CertTxtInfo []*CertTxtInfo `json:"cert_txt_info"`

}

/* 产品资质认证文本信息，包括认证类型以及文本信息 */
type CertTxtInfo struct {
	CertType int `json:"cert_type"`
	Text string `json:"text"`

}

/*  */
type ProductSpecListObject struct {
	ProductSpec []*ProductSpec `json:"product_spec"`

}

/*  */
type TicketListObject struct {
	Ticket []*Ticket `json:"ticket"`

}

/* 产品规格审核信息 */
type Ticket struct {
	AuditSellerId int `json:"audit_seller_id"`
	CreateUserId int `json:"create_user_id"`
	GmtCreate string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
	Memo string `json:"memo"`
	Reason string `json:"reason"`
	SpecId int `json:"spec_id"`
	Status int `json:"status"`

}

