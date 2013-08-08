// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package notice

const VersionNo = "20130808"

/* 用户丢失消息的数据结构 */
type DiscardInfo struct {
	End            int    `json:"end"`
	Nick           string `json:"nick"`
	Start          int    `json:"start"`
	SubscribeKey   string `json:"subscribe_key"`
	SubscribeValue string `json:"subscribe_value"`
	Type           string `json:"type"`
	UserId         int    `json:"user_id"`
}

/* 开通增量消息服务的应用用户 */
type AppCustomer struct {
	Created       string          `json:"created"`
	Modified      string          `json:"modified"`
	Nick          string          `json:"nick"`
	Status        string          `json:"status"`
	Subscriptions []*Subscription `json:"subscriptions"`
	Type          []string        `json:"type"`
	UserId        int             `json:"user_id"`
}

/* 订阅类型 */
type Subscription struct {
	Status string `json:"status"`
	Topic  string `json:"topic"`
}

/* 商品通知消息 */
type NotifyItem struct {
	ChangedFields string  `json:"changed_fields"`
	Iid           string  `json:"iid"`
	Increment     int     `json:"increment"`
	Modified      string  `json:"modified"`
	Nick          string  `json:"nick"`
	Num           int     `json:"num"`
	NumIid        int     `json:"num_iid"`
	Price         float64 `json:"price"`
	SkuId         int     `json:"sku_id"`
	SkuNum        int     `json:"sku_num"`
	Status        string  `json:"status"`
	Title         string  `json:"title"`
	UserId        int     `json:"user_id"`
}

/* 退款通知消息 */
type NotifyRefund struct {
	BuyerNick  string  `json:"buyer_nick"`
	Modified   string  `json:"modified"`
	Nick       string  `json:"nick"`
	Oid        int     `json:"oid"`
	RefundFee  float64 `json:"refund_fee"`
	Rid        int     `json:"rid"`
	SellerNick string  `json:"seller_nick"`
	Status     string  `json:"status"`
	Tid        int     `json:"tid"`
	UserId     int     `json:"user_id"`
}

/* 交易通知消息 */
type NotifyTrade struct {
	BuyerNick  string  `json:"buyer_nick"`
	Modified   string  `json:"modified"`
	Nick       string  `json:"nick"`
	Oid        int     `json:"oid"`
	Payment    float64 `json:"payment"`
	SellerNick string  `json:"seller_nick"`
	Status     string  `json:"status"`
	Tid        int     `json:"tid"`
	TradeMark  string  `json:"trade_mark"`
	Type       string  `json:"type"`
	UserId     int     `json:"user_id"`
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
