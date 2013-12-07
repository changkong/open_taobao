// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package notice

const VersionNo = "20131207"


/*  */
type DiscardInfoListObject struct {
	DiscardInfo []*DiscardInfo `json:"discard_info"`

}

/* 用户丢失消息的数据结构 */
type DiscardInfo struct {
	End int `json:"end"`
	Nick string `json:"nick"`
	Start int `json:"start"`
	SubscribeKey string `json:"subscribe_key"`
	SubscribeValue string `json:"subscribe_value"`
	Type string `json:"type"`
	UserId int `json:"user_id"`

}

/* 开通增量消息服务的应用用户 */
type AppCustomer struct {
	Created string `json:"created"`
	Modified string `json:"modified"`
	Nick string `json:"nick"`
	Status string `json:"status"`
	Subscriptions *SubscriptionListObject `json:"subscriptions"`
	Type []string `json:"type"`
	UserId int `json:"user_id"`

}

/*  */
type SubscriptionListObject struct {
	Subscription []*Subscription `json:"subscription"`

}

/* 订阅类型 */
type Subscription struct {
	Status string `json:"status"`
	Topic string `json:"topic"`

}

/*  */
type AppCustomerListObject struct {
	AppCustomer []*AppCustomer `json:"app_customer"`

}

/*  */
type NotifyItemListObject struct {
	NotifyItem []*NotifyItem `json:"notify_item"`

}

/* 商品通知消息 */
type NotifyItem struct {
	ChangedFields string `json:"changed_fields"`
	Iid string `json:"iid"`
	Increment int `json:"increment"`
	Modified string `json:"modified"`
	Nick string `json:"nick"`
	Num int `json:"num"`
	NumIid int `json:"num_iid"`
	Price float64 `json:"price"`
	SkuId int `json:"sku_id"`
	SkuNum int `json:"sku_num"`
	Status string `json:"status"`
	Title string `json:"title"`
	UserId int `json:"user_id"`

}

/*  */
type NotifyRefundListObject struct {
	NotifyRefund []*NotifyRefund `json:"notify_refund"`

}

/* 退款通知消息 */
type NotifyRefund struct {
	BuyerNick string `json:"buyer_nick"`
	Modified string `json:"modified"`
	Nick string `json:"nick"`
	Oid int `json:"oid"`
	RefundFee float64 `json:"refund_fee"`
	Rid int `json:"rid"`
	SellerNick string `json:"seller_nick"`
	Status string `json:"status"`
	Tid int `json:"tid"`
	UserId int `json:"user_id"`

}

/*  */
type NotifyTradeListObject struct {
	NotifyTrade []*NotifyTrade `json:"notify_trade"`

}

/* 交易通知消息 */
type NotifyTrade struct {
	BuyerNick string `json:"buyer_nick"`
	Modified string `json:"modified"`
	Nick string `json:"nick"`
	Oid int `json:"oid"`
	Payment float64 `json:"payment"`
	SellerNick string `json:"seller_nick"`
	Status string `json:"status"`
	Tid int `json:"tid"`
	TradeMark string `json:"trade_mark"`
	Type string `json:"type"`
	UserId int `json:"user_id"`

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

