// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package alipy

const VersionNo = "20130729"

/* 小额支付单笔支付 */
type SinglePayDetail struct {
	AlipayOrderNo      string `json:"alipay_order_no"`
	Amount             string `json:"amount"`
	CreateTime         string `json:"create_time"`
	ModifiedTime       string `json:"modified_time"`
	PayUrl             string `json:"pay_url"`
	ReceiveUserId      string `json:"receive_user_id"`
	TransferOrderNo    string `json:"transfer_order_no"`
	TransferOutOrderNo string `json:"transfer_out_order_no"`
}

/* 小额支付冻结订单详情 */
type MicroPayOrderDetail struct {
	AlipayOrderNo   string `json:"alipay_order_no"`
	AlipayUserId    string `json:"alipay_user_id"`
	AvailableAmount string `json:"available_amount"`
	CreateTime      string `json:"create_time"`
	ExpireTime      string `json:"expire_time"`
	FreezeAmount    string `json:"freeze_amount"`
	Memo            string `json:"memo"`
	MerchantOrderNo string `json:"merchant_order_no"`
	ModifiedTime    string `json:"modified_time"`
	OrderStatus     string `json:"order_status"`
	PayAmount       string `json:"pay_amount"`
	PayConfirm      string `json:"pay_confirm"`
}

/* 冻结订单详情 */
type UnfreezeOrderDetail struct {
	AlipayOrderNo   string `json:"alipay_order_no"`
	CreateTime      string `json:"create_time"`
	Memo            string `json:"memo"`
	MerchantOrderNo string `json:"merchant_order_no"`
	ModifiedTime    string `json:"modified_time"`
	OrderAmount     string `json:"order_amount"`
	OrderStatus     string `json:"order_status"`
	UnfreezeAmount  string `json:"unfreeze_amount"`
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

/* 支付宝用户冻结明细信息 */
type AccountFreeze struct {
	FreezeAmount float64 `json:"freeze_amount"`
	FreezeName   string  `json:"freeze_name"`
	FreezeType   string  `json:"freeze_type"`
}

/* 支付宝用户账户信息 */
type AlipayAccount struct {
	AlipayUserId    string  `json:"alipay_user_id"`
	AvailableAmount float64 `json:"available_amount"`
	FreezeAmount    float64 `json:"freeze_amount"`
	TotalAmount     float64 `json:"total_amount"`
}

/* 用户订购信息 */
type AlipayContract struct {
	AlipayUserId    string `json:"alipay_user_id"`
	ContractContent string `json:"contract_content"`
	EndTime         string `json:"end_time"`
	PageUrl         string `json:"page_url"`
	StartTime       string `json:"start_time"`
	Subscribe       bool   `json:"subscribe"`
}

/* 支付宝会员信息详情 */
type AlipayUserDetail struct {
	AlipayUserId string `json:"alipay_user_id"`
	Certified    bool   `json:"certified"`
	LogonId      string `json:"logon_id"`
	RealName     string `json:"real_name"`
	Sex          string `json:"sex"`
	UserStatus   string `json:"user_status"`
	UserType     string `json:"user_type"`
}

/* 支付宝交易明细 */
type TradeRecord struct {
	AlipayOrderNo   string  `json:"alipay_order_no"`
	CreateTime      string  `json:"create_time"`
	InOutType       string  `json:"in_out_type"`
	MerchantOrderNo string  `json:"merchant_order_no"`
	ModifiedTime    string  `json:"modified_time"`
	OppositeLogonId string  `json:"opposite_logon_id"`
	OppositeName    string  `json:"opposite_name"`
	OppositeUserId  string  `json:"opposite_user_id"`
	OrderFrom       string  `json:"order_from"`
	OrderStatus     string  `json:"order_status"`
	OrderTitle      string  `json:"order_title"`
	OrderType       string  `json:"order_type"`
	OwnerLogonId    string  `json:"owner_logon_id"`
	OwnerName       string  `json:"owner_name"`
	OwnerUserId     string  `json:"owner_user_id"`
	ServiceCharge   float64 `json:"service_charge"`
	TotalAmount     float64 `json:"total_amount"`
}
