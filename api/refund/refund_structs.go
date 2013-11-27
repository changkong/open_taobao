// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package refund

const VersionNo = "20131127"


/* 退款结构 */
type Refund struct {
	Address string `json:"address"`
	AdvanceStatus int `json:"advance_status"`
	AlipayNo string `json:"alipay_no"`
	BuyerNick string `json:"buyer_nick"`
	CompanyName string `json:"company_name"`
	Created string `json:"created"`
	CsStatus int `json:"cs_status"`
	Desc string `json:"desc"`
	GoodReturnTime string `json:"good_return_time"`
	GoodStatus string `json:"good_status"`
	HasGoodReturn bool `json:"has_good_return"`
	Iid string `json:"iid"`
	Modified string `json:"modified"`
	Num int `json:"num"`
	NumIid int `json:"num_iid"`
	Oid int `json:"oid"`
	OrderStatus string `json:"order_status"`
	Payment string `json:"payment"`
	Price string `json:"price"`
	Reason string `json:"reason"`
	RefundFee string `json:"refund_fee"`
	RefundId int `json:"refund_id"`
	RefundRemindTimeout *RefundRemindTimeout `json:"refund_remind_timeout"`
	SellerNick string `json:"seller_nick"`
	ShippingType string `json:"shipping_type"`
	Sid string `json:"sid"`
	SplitSellerFee string `json:"split_seller_fee"`
	SplitTaobaoFee string `json:"split_taobao_fee"`
	Status string `json:"status"`
	Tid int `json:"tid"`
	Title string `json:"title"`
	TotalFee string `json:"total_fee"`

}

/* 退款超时结构 */
type RefundRemindTimeout struct {
	ExistTimeout bool `json:"exist_timeout"`
	RemindType int `json:"remind_type"`
	Timeout string `json:"timeout"`

}

/* 留言/凭证数据结构 */
type RefundMessage struct {
	Content string `json:"content"`
	Created string `json:"created"`
	Id int `json:"id"`
	MessageType string `json:"message_type"`
	OwnerId int `json:"owner_id"`
	OwnerNick string `json:"owner_nick"`
	OwnerRole string `json:"owner_role"`
	PicUrls []*PicUrl `json:"pic_urls"`
	RefundId int `json:"refund_id"`

}

/* 图片链接 */
type PicUrl struct {
	Url string `json:"url"`

}

/* 批量异步任务结果 */
type Task struct {
	CheckCode string `json:"check_code"`
	Created string `json:"created"`
	DownloadUrl string `json:"download_url"`
	Method string `json:"method"`
	Schedule string `json:"schedule"`
	Status string `json:"status"`
	Subtasks []*Subtask `json:"subtasks"`
	TaskId int `json:"task_id"`

}

/* 批量异步任务的子任务结果 */
type Subtask struct {
	IsSuccess bool `json:"is_success"`
	SubTaskRequest string `json:"sub_task_request"`
	SubTaskResult string `json:"sub_task_result"`

}

