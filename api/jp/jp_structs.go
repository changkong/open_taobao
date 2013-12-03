// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jp

const VersionNo = "20131202"


/* Appkey的站点导购数据 */
type SpmResult struct {
	AppKey string `json:"app_key"`
	Date string `json:"date"`
	SpmModules *TopSpmObject `json:"spm_modules"`
	SpmPages *TopSpmObject `json:"spm_pages"`
	SpmSite *TopSpm `json:"spm_site"`

}

/* Appkey的站点导购数据 */
type TopSpmObject struct {
	TopSpms []*TopSpm `json:"top_spm"`

}

/* 提供查询APPKey为淘宝带来成交以及流量 */
type TopSpm struct {
	AlipayDealAmount string `json:"alipay_deal_amount"`
	AlipayDealCount int `json:"alipay_deal_count"`
	AlipayUv int `json:"alipay_uv"`
	Module string `json:"module"`
	Page string `json:"page"`
	Pv int `json:"pv"`
	Uv int `json:"uv"`

}

/* 天猫搜索品牌信息 */
type TmallBrand struct {
	BrandId int `json:"brand_id"`
	BrandName string `json:"brand_name"`

}

/* 天猫搜索类目信息 */
type TmallCat struct {
	CatId int `json:"cat_id"`
	CatName string `json:"cat_name"`

}

/* 天猫搜索结果数据结构 */
type TmallSearchItem struct {
	FastPostFee float64 `json:"fast_post_fee"`
	IsCod int `json:"is_cod"`
	IsPromotion bool `json:"is_promotion"`
	ItemId int `json:"item_id"`
	Location string `json:"location"`
	Nick string `json:"nick"`
	PicPath string `json:"pic_path"`
	Price float64 `json:"price"`
	PriceWithRate float64 `json:"price_with_rate"`
	SellerLoc string `json:"seller_loc"`
	Shipping int `json:"shipping"`
	Sold string `json:"sold"`
	SpuId int `json:"spu_id"`
	Title string `json:"title"`
	Url string `json:"url"`
	UserId int `json:"user_id"`

}

/* 天猫搜索Minisite信息 */
type TmallMinisite struct {
	Id int `json:"id"`
	Promotions string `json:"promotions"`
	Title string `json:"title"`
	Type int `json:"type"`

}

/* 天猫精选商品列表 */
type SelectedItem struct {
	Cid int `json:"cid"`
	ItemScore string `json:"item_score"`
	ShopId int `json:"shop_id"`
	TrackIid string `json:"track_iid"`

}

/* 天猫品牌特卖搜索结果数据结构 */
type TmallSearchTmItem struct {
	BrandId int `json:"brand_id"`
	BrandName string `json:"brand_name"`
	CommentNum string `json:"comment_num"`
	CommissionRate string `json:"commission_rate"`
	DetailUrl string `json:"detail_url"`
	PicUrl string `json:"pic_url"`
	Price float64 `json:"price"`
	PromotionPrice float64 `json:"promotion_price"`
	TagHot string `json:"tag_hot"`
	TagLq string `json:"tag_lq"`
	TagNew string `json:"tag_new"`
	Title string `json:"title"`
	TrackIid string `json:"track_iid"`
	Volume int `json:"volume"`

}

/* 天猫搜索特卖类目信息（查询二级类目用） */
type TmallTmCat struct {
	SubCatId int `json:"sub_cat_id"`
	SubCatName string `json:"sub_cat_name"`

}

