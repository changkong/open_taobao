// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package marketing

const VersionNo = "20130729"

/* 优惠信息对象 */
type PromotionDisplayTop struct {
	PromotionInItem []*PromotionInItem `json:"promotion_in_item"`
	PromotionInShop []*PromotionInShop `json:"promotion_in_shop"`
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

/* 商品优惠策略详情 */
type Promotion struct {
	DecreaseNum    int    `json:"decrease_num"`
	DiscountType   string `json:"discount_type"`
	DiscountValue  string `json:"discount_value"`
	EndDate        string `json:"end_date"`
	NumIid         int    `json:"num_iid"`
	PromotionDesc  string `json:"promotion_desc"`
	PromotionId    int    `json:"promotion_id"`
	PromotionTitle string `json:"promotion_title"`
	StartDate      string `json:"start_date"`
	Status         string `json:"status"`
	TagId          int    `json:"tag_id"`
}

/* 人群标签 */
type UserTag struct {
	CreateDate  string `json:"create_date"`
	Description string `json:"description"`
	TagId       int    `json:"tag_id"`
	TagName     string `json:"tag_name"`
}

/* 活动数据结构 */
type Activity struct {
	ActivityId       int    `json:"activity_id"`
	ActivityUrl      string `json:"activity_url"`
	AppliedCount     int    `json:"applied_count"`
	CouponId         int    `json:"coupon_id"`
	CreateUser       string `json:"create_user"`
	PersonLimitCount int    `json:"person_limit_count"`
	Status           string `json:"status"`
	TotalCount       int    `json:"total_count"`
}

/* 发放成功的优惠券的信息，包括couponNumber和buyerNick */
type CouponResult struct {
	BuyerNick    string `json:"buyer_nick"`
	CouponNumber int    `json:"coupon_number"`
}

/* 未发放成功买家昵称，发放失败原因 */
type ErrorMessage struct {
	BuyerNick string `json:"buyer_nick"`
	Reason    string `json:"reason"`
}

/* 优惠券详细信息 */
type CouponDetail struct {
	BuyerNick    string `json:"buyer_nick"`
	Channel      string `json:"channel"`
	CouponNumber int    `json:"coupon_number"`
	State        string `json:"state"`
}

/* 优惠券数据结构 */
type Coupon struct {
	Condition     int    `json:"condition"`
	CouponId      int    `json:"coupon_id"`
	CreatTime     string `json:"creat_time"`
	CreateChannel string `json:"create_channel"`
	Denominations int    `json:"denominations"`
	EndTime       string `json:"end_time"`
}

/* 限时打折详情 */
type LimitDiscountDetail struct {
	EndTime           string  `json:"end_time"`
	ItemDiscount      float64 `json:"item_discount"`
	ItemId            int     `json:"item_id"`
	LimitDiscountName string  `json:"limit_discount_name"`
	LimitNum          int     `json:"limit_num"`
	StartTime         string  `json:"start_time"`
}

/* 限时打折 */
type LimitDiscount struct {
	EndTime           string `json:"end_time"`
	ItemNum           int    `json:"item_num"`
	LimitDiscountId   int    `json:"limit_discount_id"`
	LimitDiscountName string `json:"limit_discount_name"`
	StartTime         string `json:"start_time"`
}

/* 搭配套餐类。 */
type Meal struct {
	ItemList  string  `json:"item_list"`
	MealId    int     `json:"meal_id"`
	MealMemo  string  `json:"meal_memo"`
	MealName  string  `json:"meal_name"`
	MealPrice float64 `json:"meal_price"`
	PostageId int     `json:"postage_id"`
	Status    string  `json:"status"`
}

/* 营销工具的范围类！ */
type Range struct {
	ParticipateId   int `json:"participate_id"`
	ParticipateType int `json:"participate_type"`
}
