// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

const VersionNo = "20130729"

/* 用户 */
type User struct {
	AlipayBind             string      `json:"alipay_bind"`
	AutoRepost             string      `json:"auto_repost"`
	Avatar                 string      `json:"avatar"`
	Birthday               string      `json:"birthday"`
	BuyerCredit            *UserCredit `json:"buyer_credit"`
	ConsumerProtection     bool        `json:"consumer_protection"`
	Created                string      `json:"created"`
	Email                  string      `json:"email"`
	HasMorePic             bool        `json:"has_more_pic"`
	HasShop                bool        `json:"has_shop"`
	HasSubStock            bool        `json:"has_sub_stock"`
	IsGoldenSeller         bool        `json:"is_golden_seller"`
	IsLightningConsignment bool        `json:"is_lightning_consignment"`
	ItemImgNum             int         `json:"item_img_num"`
	ItemImgSize            int         `json:"item_img_size"`
	LastVisit              string      `json:"last_visit"`
	Liangpin               bool        `json:"liangpin"`
	Location               *Location   `json:"location"`
	MagazineSubscribe      bool        `json:"magazine_subscribe"`
	Nick                   string      `json:"nick"`
	OnlineGaming           bool        `json:"online_gaming"`
	PromotedType           string      `json:"promoted_type"`
	PropImgNum             int         `json:"prop_img_num"`
	PropImgSize            int         `json:"prop_img_size"`
	SellerCredit           *UserCredit `json:"seller_credit"`
	Sex                    string      `json:"sex"`
	SignFoodSellerPromise  bool        `json:"sign_food_seller_promise"`
	Status                 string      `json:"status"`
	Type                   string      `json:"type"`
	Uid                    string      `json:"uid"`
	UserId                 int         `json:"user_id"`
	VerticalMarket         string      `json:"vertical_market"`
	VipInfo                string      `json:"vip_info"`
}

/* 用户信用 */
type UserCredit struct {
	GoodNum  int `json:"good_num"`
	Level    int `json:"level"`
	Score    int `json:"score"`
	TotalNum int `json:"total_num"`
}

/* 用户地址 */
type Location struct {
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	District string `json:"district"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
}
