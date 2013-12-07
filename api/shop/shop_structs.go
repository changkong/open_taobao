// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package shop

const VersionNo = "20131207"


/* 店铺内卖家自定义类目 */
type SellerCat struct {
	Cid int `json:"cid"`
	Created string `json:"created"`
	Modified string `json:"modified"`
	Name string `json:"name"`
	ParentCid int `json:"parent_cid"`
	PicUrl string `json:"pic_url"`
	SortOrder int `json:"sort_order"`
	Type string `json:"type"`

}

/*  */
type SellerCatListObject struct {
	SellerCat []*SellerCat `json:"seller_cat"`

}

/* 店铺信息 */
type Shop struct {
	AllCount int `json:"all_count"`
	Bulletin string `json:"bulletin"`
	Cid int `json:"cid"`
	Created string `json:"created"`
	Desc string `json:"desc"`
	Modified string `json:"modified"`
	Nick string `json:"nick"`
	PicPath string `json:"pic_path"`
	RemainCount int `json:"remain_count"`
	ShopScore *ShopScore `json:"shop_score"`
	Sid int `json:"sid"`
	Title string `json:"title"`
	UsedCount int `json:"used_count"`

}

/* 店铺动态评分信息 */
type ShopScore struct {
	DeliveryScore string `json:"delivery_score"`
	ItemScore string `json:"item_score"`
	ServiceScore string `json:"service_score"`

}

/*  */
type ShopCatListObject struct {
	ShopCat []*ShopCat `json:"shop_cat"`

}

/* 店铺类目 */
type ShopCat struct {
	Cid int `json:"cid"`
	IsParent bool `json:"is_parent"`
	Name string `json:"name"`
	ParentCid int `json:"parent_cid"`

}

