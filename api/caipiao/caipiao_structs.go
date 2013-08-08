// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package caipiao

const VersionNo = "20130808"

/* 卖家商品信息 */
type LotteryWangcaiSellerGoodsInfo struct {
	ActBeginTime  string `json:"act_begin_time"`
	ActEndTime    string `json:"act_end_time"`
	GoodsId       int    `json:"goods_id"`
	LotteryTypeId int    `json:"lottery_type_id"`
	PresentType   int    `json:"present_type"`
}

/* 淘宝彩票彩种信息描述 */
type LotteryType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

/* 旺彩赠送DO */
type LotteryWangcaiPresent struct {
	AppName     string `json:"app_name"`
	LotteryName string `json:"lottery_name"`
	PresentDate string `json:"present_date"`
	PresentId   int    `json:"present_id"`
	StakeCount  int    `json:"stake_count"`
	UserNick    string `json:"user_nick"`
	WinFee      int    `json:"win_fee"`
}

/* 彩票赠送的统计数据DO */
type LotteryWangcaiPresentStat struct {
	DateId       int `json:"date_id"`
	PresentFee   int `json:"present_fee"`
	PresentStake int `json:"present_stake"`
	PresentUser  int `json:"present_user"`
	SellerId     int `json:"seller_id"`
}
