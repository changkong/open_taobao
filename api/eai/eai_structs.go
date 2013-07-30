// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package eai

const VersionNo = "20130729"

/* 退款单 */
type RefundBill struct {
	ActualRefundFee     int           `json:"actual_refund_fee"`
	AlipayNo            string        `json:"alipay_no"`
	Attribute           string        `json:"attribute"`
	BillType            string        `json:"bill_type"`
	BuyerNick           string        `json:"buyer_nick"`
	Created             string        `json:"created"`
	CsStatus            string        `json:"cs_status"`
	CurrentPhaseTimeout string        `json:"current_phase_timeout"`
	ItemList            []*RefundItem `json:"item_list"`
	Modified            string        `json:"modified"`
	Oid                 int           `json:"oid"`
	OperationConstraint string        `json:"operation_constraint"`
	Reason              string        `json:"reason"`
	RefundFee           int           `json:"refund_fee"`
	RefundId            int           `json:"refund_id"`
	RefundPhase         string        `json:"refund_phase"`
	RefundType          string        `json:"refund_type"`
	RefundVersion       int           `json:"refund_version"`
	SellerNick          string        `json:"seller_nick"`
	Status              string        `json:"status"`
	TagList             []*Tag        `json:"tag_list"`
	Tid                 int           `json:"tid"`
	TradeStatus         string        `json:"trade_status"`
}

/* 退款商品信息 */
type RefundItem struct {
	Num     int    `json:"num"`
	NumIid  int    `json:"num_iid"`
	OuterId string `json:"outer_id"`
	Price   int    `json:"price"`
	Sku     string `json:"sku"`
}

/* 标签信息 */
type Tag struct {
	TagKey  string `json:"tag_key"`
	TagName string `json:"tag_name"`
	TagType string `json:"tag_type"`
}

/* 退货单 */
type ReturnBill struct {
	BillType      string        `json:"bill_type"`
	CompanyName   string        `json:"company_name"`
	Created       string        `json:"created"`
	ItemList      []*RefundItem `json:"item_list"`
	Modified      string        `json:"modified"`
	Oid           int           `json:"oid"`
	OperationLog  string        `json:"operation_log"`
	Reason        string        `json:"reason"`
	RefundId      int           `json:"refund_id"`
	RefundPhase   string        `json:"refund_phase"`
	RefundVersion int           `json:"refund_version"`
	Sid           string        `json:"sid"`
	Status        string        `json:"status"`
	TagList       []*Tag        `json:"tag_list"`
	Tid           int           `json:"tid"`
}

/* 退款留言 */
type TmallRefundMessage struct {
	Created        string `json:"created"`
	MessageContent string `json:"message_content"`
	RefundId       int    `json:"refund_id"`
	RefundPhase    string `json:"refund_phase"`
	UserNick       string `json:"user_nick"`
	UserRole       string `json:"user_role"`
	VoucherUrls    string `json:"voucher_urls"`
}
