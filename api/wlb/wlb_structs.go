// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wlb

const VersionNo = "20131207"


/*  */
type IpcInventoryDetailDoListObject struct {
	IpcInventoryDetailDo []*IpcInventoryDetailDo `json:"ipc_inventory_detail_do"`

}

/* 库存明细 */
type IpcInventoryDetailDo struct {
	BizOrderId int `json:"biz_order_id"`
	BizSubOrderId int `json:"biz_sub_order_id"`
	Flag int `json:"flag"`
	OccupyQuantity int `json:"occupy_quantity"`
	OwnerNick string `json:"owner_nick"`
	ReserveQuantity int `json:"reserve_quantity"`
	ScItemId int `json:"sc_item_id"`
	StoreCode string `json:"store_code"`

}

/*  */
type WlbInventoryListObject struct {
	WlbInventory []*WlbInventory `json:"wlb_inventory"`

}

/* 库存详情对象。其中包括货主ID，仓库编码，库存，库存类型等属性 */
type WlbInventory struct {
	ItemId int `json:"item_id"`
	LockQuantity int `json:"lock_quantity"`
	Quantity int `json:"quantity"`
	StoreCode string `json:"store_code"`
	Type string `json:"type"`
	UserId int `json:"user_id"`

}

/*  */
type WlbItemInventoryLogListObject struct {
	WlbItemInventoryLog []*WlbItemInventoryLog `json:"wlb_item_inventory_log"`

}

/* 库存变更记录 */
type WlbItemInventoryLog struct {
	BatchCode string `json:"batch_code"`
	GmtCreate string `json:"gmt_create"`
	Id int `json:"id"`
	InventType string `json:"invent_type"`
	ItemId int `json:"item_id"`
	OpType string `json:"op_type"`
	OpUserId int `json:"op_user_id"`
	OrderCode string `json:"order_code"`
	OrderItemId int `json:"order_item_id"`
	Quantity int `json:"quantity"`
	Remark string `json:"remark"`
	ResultQuantity int `json:"result_quantity"`
	StoreCode string `json:"store_code"`
	UserId int `json:"user_id"`

}

/*  */
type WlbAuthorizationListObject struct {
	WlbAuthorization []*WlbAuthorization `json:"wlb_authorization"`

}

/* 授权关系 */
type WlbAuthorization struct {
	AuthorizeEndTime string `json:"authorize_end_time"`
	AuthorizeId int `json:"authorize_id"`
	AuthorizeStartTime string `json:"authorize_start_time"`
	ConsignUserId int `json:"consign_user_id"`
	ItemId int `json:"item_id"`
	Name string `json:"name"`
	OwnerUserId int `json:"owner_user_id"`
	Quantity int `json:"quantity"`
	RuleCode string `json:"rule_code"`
	Status string `json:"status"`

}

/*  */
type WlbItemBatchInventoryListObject struct {
	WlbItemBatchInventory []*WlbItemBatchInventory `json:"wlb_item_batch_inventory"`

}

/* 商品的库存信息和批次信息 */
type WlbItemBatchInventory struct {
	ItemBatch *WlbItemBatchListObject `json:"item_batch"`
	ItemId int `json:"item_id"`
	ItemInventorys *WlbItemInventoryListObject `json:"item_inventorys"`
	TotalQuantity int `json:"total_quantity"`

}

/*  */
type WlbItemBatchListObject struct {
	WlbItemBatch []*WlbItemBatch `json:"wlb_item_batch"`

}

/* 批次库存查询结果记录 */
type WlbItemBatch struct {
	BatchCode string `json:"batch_code"`
	Creator string `json:"creator"`
	DefectQuantity int `json:"defect_quantity"`
	DueDate string `json:"due_date"`
	GmtCreate string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
	GuaranteePeriod string `json:"guarantee_period"`
	GuaranteeUnit int `json:"guarantee_unit"`
	Id int `json:"id"`
	IsDeleted bool `json:"is_deleted"`
	ItemId int `json:"item_id"`
	LastModifier string `json:"last_modifier"`
	ProduceArea string `json:"produce_area"`
	ProduceCode string `json:"produce_code"`
	ProduceDate string `json:"produce_date"`
	Quantity int `json:"quantity"`
	ReceiveDate string `json:"receive_date"`
	Remarks string `json:"remarks"`
	Status string `json:"status"`
	StoreCode string `json:"store_code"`
	UserId int `json:"user_id"`
	Version int `json:"version"`

}

/*  */
type WlbItemInventoryListObject struct {
	WlbItemInventory []*WlbItemInventory `json:"wlb_item_inventory"`

}

/* 物流宝商品库存 */
type WlbItemInventory struct {
	ItemId int `json:"item_id"`
	LockQuantity int `json:"lock_quantity"`
	Quantity int `json:"quantity"`
	StoreCode string `json:"store_code"`
	Type string `json:"type"`

}

/*  */
type WlbConsignMentListObject struct {
	WlbConsignMent []*WlbConsignMent `json:"wlb_consign_ment"`

}

/* 物流宝代销关系 */
type WlbConsignMent struct {
	Id int `json:"id"`
	ItemId int `json:"item_id"`
	Number int `json:"number"`
	TgtItemId int `json:"tgt_item_id"`
	TgtUserId int `json:"tgt_user_id"`
	UserId int `json:"user_id"`

}

/* 物流宝商品 */
type WlbItem struct {
	BrandId int `json:"brand_id"`
	Color string `json:"color"`
	Creator string `json:"creator"`
	Flag string `json:"flag"`
	GmtCreate string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
	GoodsCat string `json:"goods_cat"`
	Height int `json:"height"`
	Id int `json:"id"`
	IsDangerous bool `json:"is_dangerous"`
	IsFriable bool `json:"is_friable"`
	IsSku bool `json:"is_sku"`
	ItemCode string `json:"item_code"`
	LastModifier string `json:"last_modifier"`
	Length int `json:"length"`
	Name string `json:"name"`
	PackageMaterial string `json:"package_material"`
	ParentId int `json:"parent_id"`
	Price int `json:"price"`
	PricingCat string `json:"pricing_cat"`
	Properties string `json:"properties"`
	PublishVersion int `json:"publish_version"`
	Remark string `json:"remark"`
	Status string `json:"status"`
	Title string `json:"title"`
	Type string `json:"type"`
	UserId int `json:"user_id"`
	UserNick string `json:"user_nick"`
	Volume int `json:"volume"`
	Weight int `json:"weight"`
	Width int `json:"width"`

}

/*  */
type OutEntityItemListObject struct {
	OutEntityItem []*OutEntityItem `json:"out_entity_item"`

}

/* 外部商品实体 */
type OutEntityItem struct {
	EntityId string `json:"entity_id"`
	EntityType string `json:"entity_type"`

}

/*  */
type WlbItemListObject struct {
	WlbItem []*WlbItem `json:"wlb_item"`

}

/*  */
type WlbMessageListObject struct {
	WlbMessage []*WlbMessage `json:"wlb_message"`

}

/* 通道消息 */
type WlbMessage struct {
	GmtCreate string `json:"gmt_create"`
	Id int `json:"id"`
	MsgCode string `json:"msg_code"`
	MsgContent string `json:"msg_content"`
	MsgDescription string `json:"msg_description"`
	Status string `json:"status"`
	UserId int `json:"user_id"`

}

/*  */
type WlbOrderListObject struct {
	WlbOrder []*WlbOrder `json:"wlb_order"`

}

/* 物流宝订单对象 */
type WlbOrder struct {
	BuyerNick string `json:"buyer_nick"`
	CancelOrderStatus int `json:"cancel_order_status"`
	ConfirmStatus string `json:"confirm_status"`
	ExpectEndTime string `json:"expect_end_time"`
	ExpectStartTime string `json:"expect_start_time"`
	InvoiceInfo string `json:"invoice_info"`
	ItemKindsCount int `json:"item_kinds_count"`
	OperateType string `json:"operate_type"`
	OrderCode string `json:"order_code"`
	OrderFlag int `json:"order_flag"`
	OrderSource string `json:"order_source"`
	OrderSourceCode string `json:"order_source_code"`
	OrderStatus string `json:"order_status"`
	OrderStatusReason string `json:"order_status_reason"`
	OrderSubType string `json:"order_sub_type"`
	OrderType string `json:"order_type"`
	PrevOrderCode string `json:"prev_order_code"`
	RealKindsCount int `json:"real_kinds_count"`
	ReceivableAmount int `json:"receivable_amount"`
	ReceiverAddress string `json:"receiver_address"`
	ReceiverArea string `json:"receiver_area"`
	ReceiverCity string `json:"receiver_city"`
	ReceiverMail string `json:"receiver_mail"`
	ReceiverMobile string `json:"receiver_mobile"`
	ReceiverName string `json:"receiver_name"`
	ReceiverPhone string `json:"receiver_phone"`
	ReceiverProvince string `json:"receiver_province"`
	ReceiverWangwang string `json:"receiver_wangwang"`
	ReceiverZipCode string `json:"receiver_zip_code"`
	Remark string `json:"remark"`
	ScheduleDay string `json:"schedule_day"`
	ScheduleEnd string `json:"schedule_end"`
	ScheduleSpeed int `json:"schedule_speed"`
	ScheduleStart string `json:"schedule_start"`
	SenderAddress string `json:"sender_address"`
	SenderArea string `json:"sender_area"`
	SenderCity string `json:"sender_city"`
	SenderEmail string `json:"sender_email"`
	SenderMobile string `json:"sender_mobile"`
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
	SenderProvince string `json:"sender_province"`
	SenderZipCode string `json:"sender_zip_code"`
	ServiceFee int `json:"service_fee"`
	ShippingType string `json:"shipping_type"`
	StoreCode string `json:"store_code"`
	TmsTpCode string `json:"tms_tp_code"`
	TotalAmount int `json:"total_amount"`
	UserId int `json:"user_id"`
	UserNick string `json:"user_nick"`

}

/*  */
type WlbOrderDetailListObject struct {
	WlbOrderDetail []*WlbOrderDetail `json:"wlb_order_detail"`

}

/* 物流宝订单，并且包含订单详情 */
type WlbOrderDetail struct {
	BuyerNick string `json:"buyer_nick"`
	CreateTime string `json:"create_time"`
	IsCompleted bool `json:"is_completed"`
	ModifyTime string `json:"modify_time"`
	OperateType string `json:"operate_type"`
	OrderCode string `json:"order_code"`
	OrderItemList *WlbOrderItemListObject `json:"order_item_list"`
	OrderSource string `json:"order_source"`
	OrderSourceCode string `json:"order_source_code"`
	OrderStatus string `json:"order_status"`
	OrderSubType string `json:"order_sub_type"`
	OrderType string `json:"order_type"`
	Remark string `json:"remark"`
	StoreCode string `json:"store_code"`
	UserId int `json:"user_id"`
	UserNick string `json:"user_nick"`

}

/*  */
type WlbOrderItemListObject struct {
	WlbOrderItem []*WlbOrderItem `json:"wlb_order_item"`

}

/* 物流宝订单商品 */
type WlbOrderItem struct {
	BatchRemark string `json:"batch_remark"`
	ConfirmStatus string `json:"confirm_status"`
	ExtEntityId string `json:"ext_entity_id"`
	ExtEntityType string `json:"ext_entity_type"`
	Flag int `json:"flag"`
	Id int `json:"id"`
	InventoryType string `json:"inventory_type"`
	ItemCode string `json:"item_code"`
	ItemId int `json:"item_id"`
	ItemName string `json:"item_name"`
	ItemPrice int `json:"item_price"`
	OrderCode string `json:"order_code"`
	OrderId int `json:"order_id"`
	OrderSub2code string `json:"order_sub_2code"`
	OrderSubCode string `json:"order_sub_code"`
	OrderSubType string `json:"order_sub_type"`
	PictureUrl string `json:"picture_url"`
	PlanQuantity int `json:"plan_quantity"`
	ProviderTpId int `json:"provider_tp_id"`
	ProviderTpNick string `json:"provider_tp_nick"`
	PublishVersion int `json:"publish_version"`
	RealQuantity int `json:"real_quantity"`
	Remark string `json:"remark"`
	UserId int `json:"user_id"`
	UserNick string `json:"user_nick"`

}

/*  */
type WlbOrderScheduleRuleListObject struct {
	WlbOrderScheduleRule []*WlbOrderScheduleRule `json:"wlb_order_schedule_rule"`

}

/* 订单调度规则 */
type WlbOrderScheduleRule struct {
	AreaIds string `json:"area_ids"`
	BackupStoreId int `json:"backup_store_id"`
	DefaultStoreId int `json:"default_store_id"`
	Id int `json:"id"`
	Options string `json:"options"`
	PresellStoreId int `json:"presell_store_id"`
	RuleId int `json:"rule_id"`
	UserId int `json:"user_id"`
	UserNick string `json:"user_nick"`

}

/*  */
type WlbProcessStatusListObject struct {
	WlbProcessStatus []*WlbProcessStatus `json:"wlb_process_status"`

}

/* 物流宝订单流转信息对象 */
type WlbProcessStatus struct {
	Content string `json:"content"`
	OperateTime string `json:"operate_time"`
	Operater string `json:"operater"`
	OrderCode string `json:"order_code"`
	Remark string `json:"remark"`
	StatusCode string `json:"status_code"`
	StoreCode string `json:"store_code"`
	StoreTpCode string `json:"store_tp_code"`
	TmsOrderCode string `json:"tms_order_code"`
	TmsTpCode string `json:"tms_tp_code"`

}

/*  */
type WlbReplenishListObject struct {
	WlbReplenish []*WlbReplenish `json:"wlb_replenish"`

}

/* 物流宝补货统计对象 */
type WlbReplenish struct {
	EstimateValue string `json:"estimate_value"`
	HistoryValue string `json:"history_value"`
	ItemId int `json:"item_id"`
	RetrievalCount int `json:"retrieval_count"`
	SellCount int `json:"sell_count"`
	StoreCode string `json:"store_code"`
	TransportCount int `json:"transport_count"`
	UserId int `json:"user_id"`
	WarnCount int `json:"warn_count"`

}

/*  */
type WlbSellerSubscriptionListObject struct {
	WlbSellerSubscription []*WlbSellerSubscription `json:"wlb_seller_subscription"`

}

/* 卖家定购的服务 */
type WlbSellerSubscription struct {
	EndDate string `json:"end_date"`
	GmtCreate string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
	Id int `json:"id"`
	IsOwnService int `json:"is_own_service"`
	ParentId int `json:"parent_id"`
	ProviderUserId int `json:"provider_user_id"`
	Remark string `json:"remark"`
	ServiceAlias string `json:"service_alias"`
	ServiceCode string `json:"service_code"`
	ServiceId int `json:"service_id"`
	ServiceName string `json:"service_name"`
	ServiceType string `json:"service_type"`
	StartDate string `json:"start_date"`
	Status string `json:"status"`
	SubscriberUserId int `json:"subscriber_user_id"`
	SubscriberUserNick string `json:"subscriber_user_nick"`
	WlbPartnerAddress *WlbPartnerAddress `json:"wlb_partner_address"`
	WlbPartnerContact *WlbPartnerContact `json:"wlb_partner_contact"`

}

/* 联系人地址信息 */
type WlbPartnerAddress struct {
	Address string `json:"address"`
	Borough string `json:"borough"`
	City string `json:"city"`
	Province string `json:"province"`
	Zip string `json:"zip"`

}

/* 联系人联系详情 */
type WlbPartnerContact struct {
	Name string `json:"name"`
	Phone string `json:"phone"`

}

/*  */
type WlbTmsOrderListObject struct {
	WlbTmsOrder []*WlbTmsOrder `json:"wlb_tms_order"`

}

/* 物流订单运单信息 */
type WlbTmsOrder struct {
	Code string `json:"code"`
	CompanyCode string `json:"company_code"`
	OrderCode string `json:"order_code"`
	UserId int `json:"user_id"`

}

