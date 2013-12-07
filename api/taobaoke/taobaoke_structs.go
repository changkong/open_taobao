// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

const VersionNo = "20131207"


/*  */
type TaobaokeAuthorizeListObject struct {
	TaobaokeAuthorize []*TaobaokeAuthorize `json:"taobaoke_authorize"`

}

/* 淘宝客返利授权 */
type TaobaokeAuthorize struct {
	Param string `json:"param"`
	Rebate bool `json:"rebate"`

}

/*  */
type TaobaokePaymentListObject struct {
	TaobaokePayment []*TaobaokePayment `json:"taobaoke_payment"`

}

/* 淘宝客订单 */
type TaobaokePayment struct {
	AppKey string `json:"app_key"`
	CategoryId int `json:"category_id"`
	CategoryName string `json:"category_name"`
	Commission float64 `json:"commission"`
	CommissionRate string `json:"commission_rate"`
	CreateTime string `json:"create_time"`
	ItemNum int `json:"item_num"`
	ItemTitle string `json:"item_title"`
	NumIid int `json:"num_iid"`
	OuterCode string `json:"outer_code"`
	PayPrice float64 `json:"pay_price"`
	PayTime string `json:"pay_time"`
	RealPayFee float64 `json:"real_pay_fee"`
	SellerNick string `json:"seller_nick"`
	ShopTitle string `json:"shop_title"`
	TradeId int `json:"trade_id"`
	TradeParentId int `json:"trade_parent_id"`

}

