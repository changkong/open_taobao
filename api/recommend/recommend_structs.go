// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package recommend

const VersionNo = "20130729"

/* 推荐的关联商品 */
type FavoriteItem struct {
	ItemName       string  `json:"item_name"`
	ItemPictrue    string  `json:"item_pictrue"`
	ItemPrice      float64 `json:"item_price"`
	ItemUrl        string  `json:"item_url"`
	PromotionPrice float64 `json:"promotion_price"`
	SellCount      int     `json:"sell_count"`
	TrackIid       string  `json:"track_iid"`
}

/* 推荐关联店铺信息 */
type FavoriteShop struct {
	Rate       int    `json:"rate"`
	SellerId   int    `json:"seller_id"`
	SellerNick string `json:"seller_nick"`
	ShopId     int    `json:"shop_id"`
	ShopName   string `json:"shop_name"`
	ShopPic    string `json:"shop_pic"`
	ShopUrl    string `json:"shop_url"`
}
