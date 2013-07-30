// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fenxiao

const VersionNo = "20130729"

/* 合作分销关系 */
type Cooperation struct {
	AuthPayway      []string `json:"auth_payway"`
	CooperateId     int      `json:"cooperate_id"`
	DistributorId   int      `json:"distributor_id"`
	DistributorNick string   `json:"distributor_nick"`
	EndDate         string   `json:"end_date"`
	GradeId         int      `json:"grade_id"`
	ProductLine     string   `json:"product_line"`
	ProductLineName []string `json:"product_line_name"`
	StartDate       string   `json:"start_date"`
	Status          string   `json:"status"`
	SupplierId      int      `json:"supplier_id"`
	SupplierNick    string   `json:"supplier_nick"`
	TradeType       string   `json:"trade_type"`
}

/* 经销采购申请单 */
type DealerOrder struct {
	AppliedTime            string               `json:"applied_time"`
	ApplierNick            string               `json:"applier_nick"`
	AuditTimeApplier       string               `json:"audit_time_applier"`
	AuditTimeSupplier      string               `json:"audit_time_supplier"`
	CloseReason            string               `json:"close_reason"`
	DealerOrderDetails     []*DealerOrderDetail `json:"dealer_order_details"`
	DealerOrderId          int                  `json:"dealer_order_id"`
	DeliveredQuantityCount int                  `json:"delivered_quantity_count"`
	LogisticsFee           float64              `json:"logistics_fee"`
	LogisticsType          string               `json:"logistics_type"`
	ModifiedTime           string               `json:"modified_time"`
	OrderStatus            string               `json:"order_status"`
	PayType                string               `json:"pay_type"`
	QuantityCount          int                  `json:"quantity_count"`
	Receiver               *Receiver            `json:"receiver"`
	RefuseReasonApplier    string               `json:"refuse_reason_applier"`
	RefuseReasonSupplier   string               `json:"refuse_reason_supplier"`
	SupplierNick           string               `json:"supplier_nick"`
	TotalPrice             float64              `json:"total_price"`
}

/* 经销采购申请单商品明细 */
type DealerOrderDetail struct {
	DealerDetailId int     `json:"dealer_detail_id"`
	DealerOrderId  int     `json:"dealer_order_id"`
	FinalPrice     float64 `json:"final_price"`
	IsDeleted      bool    `json:"is_deleted"`
	OriginalPrice  float64 `json:"original_price"`
	PriceCount     float64 `json:"price_count"`
	ProductId      int     `json:"product_id"`
	ProductTitle   string  `json:"product_title"`
	Quantity       int     `json:"quantity"`
	SkuId          int     `json:"sku_id"`
	SkuNumber      string  `json:"sku_number"`
	SkuSpec        string  `json:"sku_spec"`
	SnapshotUrl    string  `json:"snapshot_url"`
}

/* 收货人详细信息 */
type Receiver struct {
	Address     string `json:"address"`
	CardId      string `json:"card_id"`
	City        string `json:"city"`
	District    string `json:"district"`
	MobilePhone string `json:"mobile_phone"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
}

/* 折扣信息 */
type Discount struct {
	Created    string            `json:"created"`
	Details    []*DiscountDetail `json:"details"`
	DiscountId int               `json:"discount_id"`
	Modified   string            `json:"modified"`
	Name       string            `json:"name"`
}

/* 折扣详情信息 */
type DiscountDetail struct {
	Created       string `json:"created"`
	DetailId      int    `json:"detail_id"`
	DiscountType  string `json:"discount_type"`
	DiscountValue int    `json:"discount_value"`
	Modified      string `json:"modified"`
	TargetId      int    `json:"target_id"`
	TargetName    string `json:"target_name"`
	TargetType    string `json:"target_type"`
}

/* 分销商档案信息 */
type DistributorArchive struct {
	ComplaintsCountPer    string `json:"complaints_count_per"`
	DownLoadRatio         string `json:"down_load_ratio"`
	GoodEvaluationRate    string `json:"good_evaluation_rate"`
	GoodEvaluationRateD   string `json:"good_evaluation_rate_d"`
	GoodsSum              string `json:"goods_sum"`
	IndusPerScole         string `json:"indus_per_scole"`
	Industry              string `json:"industry"`
	OpenashopTime         string `json:"openashop_time"`
	OrderMerchandiseScore string `json:"order_merchandise_score"`
	OrderShopRatio        string `json:"order_shop_ratio"`
	PunishedCount         string `json:"punished_count"`
	SaleConsignmentScore  string `json:"sale_consignment_score"`
	SellerRefundTime      string `json:"seller_refund_time"`
	SellerRefundTimePer   string `json:"seller_refund_time_per"`
	SellerStarName        string `json:"seller_star_name"`
	SellerType            string `json:"seller_type"`
	ServiceQualityScore   string `json:"service_quality_score"`
	ShopAddress           string `json:"shop_address"`
	ShopName              string `json:"shop_name"`
	UpSelfRatio           string `json:"up_self_ratio"`
	UpShopRatio           string `json:"up_shop_ratio"`
	UvShopRatio           string `json:"uv_shop_ratio"`
}

/* 分销商品下载记录 */
type FenxiaoItemRecord struct {
	Created       string `json:"created"`
	DistributorId int    `json:"distributor_id"`
	ItemId        int    `json:"item_id"`
	Modified      string `json:"modified"`
	ProductId     int    `json:"product_id"`
	TradeType     string `json:"trade_type"`
}

/* 分销商品流量 */
type DistributorItemFlow struct {
	ItemPv    string `json:"item_pv"`
	ItemUv    string `json:"item_uv"`
	ProductId int    `json:"product_id"`
}

/* 分销产品 */
type FenxiaoProduct struct {
	AlarmNumber         int             `json:"alarm_number"`
	CategoryId          string          `json:"category_id"`
	City                string          `json:"city"`
	CostPrice           float64         `json:"cost_price"`
	Created             string          `json:"created"`
	DealerCostPrice     float64         `json:"dealer_cost_price"`
	DescPath            string          `json:"desc_path"`
	Description         string          `json:"description"`
	DiscountId          int             `json:"discount_id"`
	HaveGuarantee       bool            `json:"have_guarantee"`
	HaveInvoice         bool            `json:"have_invoice"`
	Images              []*FenxiaoImage `json:"images"`
	InputProperties     string          `json:"input_properties"`
	IsAuthz             string          `json:"is_authz"`
	ItemId              int             `json:"item_id"`
	ItemsCount          int             `json:"items_count"`
	Modified            string          `json:"modified"`
	Name                string          `json:"name"`
	OrdersCount         int             `json:"orders_count"`
	OuterId             string          `json:"outer_id"`
	Pdus                []*FenxiaoPdu   `json:"pdus"`
	Pictures            string          `json:"pictures"`
	Pid                 int             `json:"pid"`
	PostageEms          float64         `json:"postage_ems"`
	PostageFast         float64         `json:"postage_fast"`
	PostageId           int             `json:"postage_id"`
	PostageOrdinary     float64         `json:"postage_ordinary"`
	PostageType         string          `json:"postage_type"`
	ProductcatId        int             `json:"productcat_id"`
	Properties          string          `json:"properties"`
	PropertyAlias       string          `json:"property_alias"`
	Prov                string          `json:"prov"`
	Quantity            int             `json:"quantity"`
	QueryItemId         int             `json:"query_item_id"`
	RetailPriceHigh     float64         `json:"retail_price_high"`
	RetailPriceLow      float64         `json:"retail_price_low"`
	ScitemId            int             `json:"scitem_id"`
	Skus                []*FenxiaoSku   `json:"skus"`
	StandardPrice       float64         `json:"standard_price"`
	StandardRetailPrice float64         `json:"standard_retail_price"`
	Status              string          `json:"status"`
	TradeType           string          `json:"trade_type"`
	UpshelfTime         string          `json:"upshelf_time"`
}

/* 产品的各种图片信息 */
type FenxiaoImage struct {
	ImageId       int    `json:"image_id"`
	ImagePosition int    `json:"image_position"`
	ImageUrl      string `json:"image_url"`
	Properties    string `json:"properties"`
	Type          string `json:"type"`
}

/* 产品分销商属性 */
type FenxiaoPdu struct {
	DistributorId   int    `json:"distributor_id"`
	DistributorName string `json:"distributor_name"`
	ProductId       int    `json:"product_id"`
	QuantityAgent   int    `json:"quantity_agent"`
	SkuProperties   string `json:"sku_properties"`
}

/* 分销产品SKU */
type FenxiaoSku struct {
	CostPrice       string `json:"cost_price"`
	DealerCostPrice string `json:"dealer_cost_price"`
	Id              int    `json:"id"`
	Name            string `json:"name"`
	OuterId         string `json:"outer_id"`
	Properties      string `json:"properties"`
	Quantity        int    `json:"quantity"`
	ScitemId        int    `json:"scitem_id"`
	StandardPrice   string `json:"standard_price"`
}

/* 分销API返回数据结构 */
type Distributor struct {
	AlipayAccount   string `json:"alipay_account"`
	Appraise        int    `json:"appraise"`
	CategoryId      int    `json:"category_id"`
	ContactPerson   string `json:"contact_person"`
	Created         string `json:"created"`
	DistributorId   int    `json:"distributor_id"`
	DistributorName string `json:"distributor_name"`
	Email           string `json:"email"`
	FullName        string `json:"full_name"`
	Level           int    `json:"level"`
	MobilePhone     string `json:"mobile_phone"`
	Phone           string `json:"phone"`
	ShopWebLink     string `json:"shop_web_link"`
	Starts          string `json:"starts"`
	UserId          int    `json:"user_id"`
}

/* 分销商等级 */
type FenxiaoGrade struct {
	Created  string `json:"created"`
	GradeId  int    `json:"grade_id"`
	Modified string `json:"modified"`
	Name     string `json:"name"`
}

/* 登录分销用户信息 */
type LoginUser struct {
	CreateTime string `json:"create_time"`
	Nick       string `json:"nick"`
	UserId     int    `json:"user_id"`
	UserType   string `json:"user_type"`
}

/* 采购单及子采购单信息 */
type PurchaseOrder struct {
	AlipayNo             string              `json:"alipay_no"`
	BuyerNick            string              `json:"buyer_nick"`
	BuyerPayment         float64             `json:"buyer_payment"`
	ConsignTime          string              `json:"consign_time"`
	Created              string              `json:"created"`
	DistributorFrom      string              `json:"distributor_from"`
	DistributorPayment   float64             `json:"distributor_payment"`
	DistributorUsername  string              `json:"distributor_username"`
	EndTime              string              `json:"end_time"`
	FenxiaoId            int                 `json:"fenxiao_id"`
	Id                   int                 `json:"id"`
	IsvCustomKey         []string            `json:"isv_custom_key"`
	IsvCustomValue       []string            `json:"isv_custom_value"`
	LogisticsCompanyName string              `json:"logistics_company_name"`
	LogisticsId          string              `json:"logistics_id"`
	Memo                 string              `json:"memo"`
	Modified             string              `json:"modified"`
	OrderMessages        []*OrderMessage     `json:"order_messages"`
	PayTime              string              `json:"pay_time"`
	PayType              string              `json:"pay_type"`
	PostFee              float64             `json:"post_fee"`
	Receiver             *Receiver           `json:"receiver"`
	Shipping             string              `json:"shipping"`
	SnapshotUrl          string              `json:"snapshot_url"`
	Status               string              `json:"status"`
	SubPurchaseOrders    []*SubPurchaseOrder `json:"sub_purchase_orders"`
	SupplierFlag         int                 `json:"supplier_flag"`
	SupplierFrom         string              `json:"supplier_from"`
	SupplierMemo         string              `json:"supplier_memo"`
	SupplierUsername     string              `json:"supplier_username"`
	TcOrderId            int                 `json:"tc_order_id"`
	TotalFee             float64             `json:"total_fee"`
	TradeType            string              `json:"trade_type"`
}

/* 采购单留言列表 */
type OrderMessage struct {
	MessageContent string `json:"message_content"`
	MessageTime    string `json:"message_time"`
	MessageTitle   string `json:"message_title"`
	PicUrl         string `json:"pic_url"`
}

/* 子采购单详细信息 */
type SubPurchaseOrder struct {
	AuctionPrice       float64 `json:"auction_price"`
	BillFee            float64 `json:"bill_fee"`
	BuyerPayment       float64 `json:"buyer_payment"`
	Created            string  `json:"created"`
	DistributorPayment float64 `json:"distributor_payment"`
	FenxiaoId          int     `json:"fenxiao_id"`
	Id                 int     `json:"id"`
	ItemId             int     `json:"item_id"`
	ItemOuterId        string  `json:"item_outer_id"`
	Num                int     `json:"num"`
	OldSkuProperties   string  `json:"old_sku_properties"`
	Order200Status     string  `json:"order_200_status"`
	Price              float64 `json:"price"`
	RefundFee          float64 `json:"refund_fee"`
	ScItemId           int     `json:"sc_item_id"`
	SkuId              int     `json:"sku_id"`
	SkuOuterId         string  `json:"sku_outer_id"`
	SkuProperties      string  `json:"sku_properties"`
	SnapshotUrl        string  `json:"snapshot_url"`
	Status             string  `json:"status"`
	TcAdjustFee        int     `json:"tc_adjust_fee"`
	TcDiscountFee      int     `json:"tc_discount_fee"`
	TcOrderId          int     `json:"tc_order_id"`
	TcPreferentialType string  `json:"tc_preferential_type"`
	Title              string  `json:"title"`
	TotalFee           float64 `json:"total_fee"`
}

/* 等级折扣数据结构 */
type GradeDiscount struct {
	DiscountId   int     `json:"discount_id"`
	DiscountType int     `json:"discount_type"`
	Price        float64 `json:"price"`
	SkuId        int     `json:"sku_id"`
	TradeType    int     `json:"trade_type"`
}

/* 产品线 */
type ProductCat struct {
	CostPercentAgent  string `json:"cost_percent_agent"`
	CostPercentDealer string `json:"cost_percent_dealer"`
	Id                int    `json:"id"`
	Name              string `json:"name"`
	ProductNum        int    `json:"product_num"`
	RetailHighPercent string `json:"retail_high_percent"`
	RetailLowPercent  string `json:"retail_low_percent"`
}

/* 采购单子单退款详情 */
type RefundDetail struct {
	BuyerRefund      *BuyerRefund `json:"buyer_refund"`
	DistributorNick  string       `json:"distributor_nick"`
	IsReturnGoods    bool         `json:"is_return_goods"`
	Modified         string       `json:"modified"`
	PaySupFee        float64      `json:"pay_sup_fee"`
	RefundCreateTime string       `json:"refund_create_time"`
	RefundDesc       string       `json:"refund_desc"`
	RefundFee        float64      `json:"refund_fee"`
	RefundReason     string       `json:"refund_reason"`
	RefundStatus     int          `json:"refund_status"`
	SubOrderId       int          `json:"sub_order_id"`
	SupplierNick     string       `json:"supplier_nick"`
}

/* 下游买家退款信息 */
type BuyerRefund struct {
	BizOrderId       int    `json:"biz_order_id"`
	BuyerNick        string `json:"buyer_nick"`
	GoodsStatusDesc  string `json:"goods_status_desc"`
	Modified         string `json:"modified"`
	NeedReturnGoods  bool   `json:"need_return_goods"`
	RefundCreateTime string `json:"refund_create_time"`
	RefundDesc       string `json:"refund_desc"`
	RefundId         int    `json:"refund_id"`
	RefundReason     string `json:"refund_reason"`
	RefundStatus     int    `json:"refund_status"`
	ReturnFee        int    `json:"return_fee"`
	SubOrderId       int    `json:"sub_order_id"`
	ToSellerFee      int    `json:"to_seller_fee"`
}

/* 合作申请 */
type Requisition struct {
	DistAppraise     int    `json:"dist_appraise"`
	DistCategory     int    `json:"dist_category"`
	DistCategoryName string `json:"dist_category_name"`
	DistIsXiaobao    int    `json:"dist_is_xiaobao"`
	DistLevel        int    `json:"dist_level"`
	DistMessage      string `json:"dist_message"`
	DistOpenDate     string `json:"dist_open_date"`
	DistShopAddress  string `json:"dist_shop_address"`
	DistributorId    int    `json:"distributor_id"`
	DistributorNick  string `json:"distributor_nick"`
	GmtCreate        string `json:"gmt_create"`
	RequisitionId    int    `json:"requisition_id"`
	Status           int    `json:"status"`
}

/* 经销订单监控记录信息 */
type TradeMonitor struct {
	Area               string `json:"area"`
	BuyAmount          int    `json:"buy_amount"`
	BuyerFullName      string `json:"buyer_full_name"`
	BuyerNick          string `json:"buyer_nick"`
	City               string `json:"city"`
	DistributorNick    string `json:"distributor_nick"`
	ItemId             int    `json:"item_id"`
	ItemNumber         string `json:"item_number"`
	ItemPrice          int    `json:"item_price"`
	ItemSkuName        string `json:"item_sku_name"`
	ItemSkuNumber      string `json:"item_sku_number"`
	ItemTitle          string `json:"item_title"`
	ItemTotalPrice     int    `json:"item_total_price"`
	PayTime            string `json:"pay_time"`
	ProductId          int    `json:"product_id"`
	ProductNumber      string `json:"product_number"`
	ProductSkuNumber   string `json:"product_sku_number"`
	ProductTitle       string `json:"product_title"`
	Provence           string `json:"provence"`
	RetailPriceHigh    int    `json:"retail_price_high"`
	RetailPriceLow     int    `json:"retail_price_low"`
	ShippingAddress    string `json:"shipping_address"`
	Status             string `json:"status"`
	SubTcOrderId       int    `json:"sub_tc_order_id"`
	SupplierNick       string `json:"supplier_nick"`
	TcAdjustFee        int    `json:"tc_adjust_fee"`
	TcDiscountFee      int    `json:"tc_discount_fee"`
	TcOrderId          int    `json:"tc_order_id"`
	TcPreferentialType string `json:"tc_preferential_type"`
	TradeMonitorId     int    `json:"trade_monitor_id"`
}

/* 提示信息对象 */
type TipInfo struct {
	Info     string `json:"info"`
	ScItemId string `json:"sc_item_id"`
}

/* 授权信息数据结构 */
type InventoryAuthorizeInfo struct {
	AuthorizeCode   string `json:"authorize_code"`
	ChannelFlag     int    `json:"channel_flag"`
	Index           int    `json:"index"`
	InventoryType   int    `json:"inventory_type"`
	NickName        string `json:"nick_name"`
	NickNameList    string `json:"nick_name_list"`
	OccupyQuantity  int    `json:"occupy_quantity"`
	Quantity        int    `json:"quantity"`
	QuotaQuantity   int    `json:"quota_quantity"`
	ReserveQuantity int    `json:"reserve_quantity"`
	ScItemCode      string `json:"sc_item_code"`
	ScItemId        int    `json:"sc_item_id"`
	StoreCode       string `json:"store_code"`
}

/* 商品库存对象 */
type InventorySum struct {
	InventoryType     int    `json:"inventory_type"`
	InventoryTypeName string `json:"inventory_type_name"`
	OccupyQuantity    int    `json:"occupy_quantity"`
	Quantity          int    `json:"quantity"`
	ReserveQuantity   int    `json:"reserve_quantity"`
	ScItemCode        string `json:"sc_item_code"`
	ScItemId          int    `json:"sc_item_id"`
	StoreCode         string `json:"store_code"`
}

/* 仓库对象 */
type Store struct {
	Address         string `json:"address"`
	AddressAreaName string `json:"address_area_name"`
	AliasName       string `json:"alias_name"`
	Contact         string `json:"contact"`
	Phone           string `json:"phone"`
	PostCode        string `json:"post_code"`
	StoreCode       string `json:"store_code"`
	StoreName       string `json:"store_name"`
	StoreType       string `json:"store_type"`
}

/* 后端商品 */
type ScItem struct {
	BarCode      string `json:"bar_code"`
	BrandId      int    `json:"brand_id"`
	BrandName    string `json:"brand_name"`
	Height       int    `json:"height"`
	IsAreaSale   int    `json:"is_area_sale"`
	IsCostly     bool   `json:"is_costly"`
	IsDangerous  bool   `json:"is_dangerous"`
	IsFriable    bool   `json:"is_friable"`
	IsWarranty   bool   `json:"is_warranty"`
	ItemId       int    `json:"item_id"`
	ItemName     string `json:"item_name"`
	ItemType     int    `json:"item_type"`
	Length       int    `json:"length"`
	MatterStatus string `json:"matter_status"`
	OuterCode    string `json:"outer_code"`
	Price        int    `json:"price"`
	Properties   string `json:"properties"`
	Remark       string `json:"remark"`
	Volume       int    `json:"volume"`
	Weight       int    `json:"weight"`
	Width        int    `json:"width"`
	WmsCode      string `json:"wms_code"`
}

/* 后端商品映射关系对象 */
type ScItemMap struct {
	ItemId       int    `json:"item_id"`
	MapType      int    `json:"map_type"`
	RelItemId    int    `json:"rel_item_id"`
	RelOuterCode string `json:"rel_outer_code"`
	RelUserId    int    `json:"rel_user_id"`
	RelUserNick  string `json:"rel_user_nick"`
	SkuId        int    `json:"sku_id"`
	UserId       int    `json:"user_id"`
	UserNick     string `json:"user_nick"`
}

/* 查询分页结构 */
type QueryPagination struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

/* 批量异步任务结果 */
type Task struct {
	CheckCode   string     `json:"check_code"`
	Created     string     `json:"created"`
	DownloadUrl string     `json:"download_url"`
	Method      string     `json:"method"`
	Schedule    string     `json:"schedule"`
	Status      string     `json:"status"`
	Subtasks    []*Subtask `json:"subtasks"`
	TaskId      int        `json:"task_id"`
}

/* 批量异步任务的子任务结果 */
type Subtask struct {
	IsSuccess      bool   `json:"is_success"`
	SubTaskRequest string `json:"sub_task_request"`
	SubTaskResult  string `json:"sub_task_result"`
}

/* 库存明细 */
type IpcInventoryDetailDo struct {
	BizOrderId      int    `json:"biz_order_id"`
	BizSubOrderId   int    `json:"biz_sub_order_id"`
	Flag            int    `json:"flag"`
	OccupyQuantity  int    `json:"occupy_quantity"`
	OwnerNick       string `json:"owner_nick"`
	ReserveQuantity int    `json:"reserve_quantity"`
	ScItemId        int    `json:"sc_item_id"`
	StoreCode       string `json:"store_code"`
}
