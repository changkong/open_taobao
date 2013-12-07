// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fuwu

const VersionNo = "20131207"


/*  */
type ArticleBizOrderListObject struct {
	ArticleBizOrder []*ArticleBizOrder `json:"article_biz_order"`

}

/* 应用订单信息 */
type ArticleBizOrder struct {
	ArticleCode string `json:"article_code"`
	ArticleItemName string `json:"article_item_name"`
	ArticleName string `json:"article_name"`
	BizOrderId int `json:"biz_order_id"`
	BizType int `json:"biz_type"`
	Create string `json:"create"`
	Fee string `json:"fee"`
	ItemCode string `json:"item_code"`
	ItemName string `json:"item_name"`
	Nick string `json:"nick"`
	OrderCycle string `json:"order_cycle"`
	OrderCycleEnd string `json:"order_cycle_end"`
	OrderCycleStart string `json:"order_cycle_start"`
	OrderId int `json:"order_id"`
	PromFee string `json:"prom_fee"`
	RefundFee string `json:"refund_fee"`
	TotalPayFee string `json:"total_pay_fee"`

}

/*  */
type ArticleSubListObject struct {
	ArticleSub []*ArticleSub `json:"article_sub"`

}

/* 应用订购信息 */
type ArticleSub struct {
	ArticleCode string `json:"article_code"`
	ArticleName string `json:"article_name"`
	Autosub bool `json:"autosub"`
	Deadline string `json:"deadline"`
	ExpireNotice bool `json:"expire_notice"`
	ItemCode string `json:"item_code"`
	ItemName string `json:"item_name"`
	Nick string `json:"nick"`
	Status int `json:"status"`

}

/*  */
type ArticleUserSubscribeListObject struct {
	ArticleUserSubscribe []*ArticleUserSubscribe `json:"article_user_subscribe"`

}

/* 用户订购信息 */
type ArticleUserSubscribe struct {
	Deadline string `json:"deadline"`
	ItemCode string `json:"item_code"`

}

