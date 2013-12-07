// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bill

const VersionNo = "20131207"


/*  */
type AccountListObject struct {
	Account []*Account `json:"account"`

}

/* 费用科目 */
type Account struct {
	AccountCode string `json:"account_code"`
	AccountId int `json:"account_id"`
	AccountName string `json:"account_name"`
	AccountType int `json:"account_type"`

}

/* 账单结构 */
type Bill struct {
	AccountId int `json:"account_id"`
	AlipayId string `json:"alipay_id"`
	AlipayMail string `json:"alipay_mail"`
	AlipayNo string `json:"alipay_no"`
	AlipayNotice string `json:"alipay_notice"`
	AlipayOutno string `json:"alipay_outno"`
	Amount int `json:"amount"`
	Bid int `json:"bid"`
	BizTime string `json:"biz_time"`
	BookTime string `json:"book_time"`
	GmtCreate string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
	NumIid string `json:"num_iid"`
	ObjAlipayId string `json:"obj_alipay_id"`
	ObjAlipayMail string `json:"obj_alipay_mail"`
	OrderId string `json:"order_id"`
	PayTime string `json:"pay_time"`
	Status int `json:"status"`
	TotalAmount int `json:"total_amount"`
	TradeId string `json:"trade_id"`

}

/*  */
type BillListObject struct {
	Bill []*Bill `json:"bill"`

}

/* 虚拟账户账单结构 */
type BookBill struct {
	AccountId int `json:"account_id"`
	Amount int `json:"amount"`
	Bid int `json:"bid"`
	BookTime string `json:"book_time"`
	Description string `json:"description"`
	GmtCreate string `json:"gmt_create"`
	JournalType int `json:"journal_type"`
	OtherAlipayId string `json:"other_alipay_id"`
	TaobaoAlipayId string `json:"taobao_alipay_id"`

}

/*  */
type BookBillListObject struct {
	BookBill []*BookBill `json:"book_bill"`

}

