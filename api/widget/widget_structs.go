// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package widget

const VersionNo = "20130808"

/* Widget获取到的商品信息 */
type WidgetItem struct {
	AddToCart         bool               `json:"add_to_cart"`
	ApproveStatus     string             `json:"approve_status"`
	ClickUrl          string             `json:"click_url"`
	IsMall            bool               `json:"is_mall"`
	ItemId            int                `json:"item_id"`
	ItemPics          []string           `json:"item_pics"`
	ItemPromotionData []*PromotionInItem `json:"item_promotion_data"`
	PicUrl            string             `json:"pic_url"`
	Price             string             `json:"price"`
	Quantity          int                `json:"quantity"`
	SellerNick        string             `json:"seller_nick"`
	ShopPromotionData []*PromotionInShop `json:"shop_promotion_data"`
	SkuProps          []*WidgetSkuProps  `json:"sku_props"`
	Skus              []*WidgetSku       `json:"skus"`
	Title             string             `json:"title"`
}

/* 单品级优惠信息 */
type PromotionInItem struct {
	Desc           string    `json:"desc"`
	EndTime        string    `json:"end_time"`
	ItemPromoPrice float64   `json:"item_promo_price"`
	Name           string    `json:"name"`
	OtherNeed      string    `json:"other_need"`
	OtherSend      string    `json:"other_send"`
	PromotionId    string    `json:"promotion_id"`
	SkuIdList      []string  `json:"sku_id_list"`
	SkuPriceList   []float64 `json:"sku_price_list"`
	StartTime      string    `json:"start_time"`
}

/* 店铺级优惠信息 */
type PromotionInShop struct {
	Name                string `json:"name"`
	PromotionDetailDesc string `json:"promotion_detail_desc"`
	PromotionId         string `json:"promotion_id"`
}

/* Widget使用的sku属性对应信息结构体 */
type WidgetSkuProps struct {
	Alias     string `json:"alias"`
	KeyName   string `json:"key_name"`
	PicUrl    string `json:"pic_url"`
	PropKey   int    `json:"prop_key"`
	PropValue int    `json:"prop_value"`
	ValueName string `json:"value_name"`
}

/* 组件sku信息 */
type WidgetSku struct {
	Price    float64 `json:"price"`
	Props    string  `json:"props"`
	Quantity int     `json:"quantity"`
	SkuId    int     `json:"sku_id"`
}
