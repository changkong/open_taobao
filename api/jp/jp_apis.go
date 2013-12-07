// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jp

import (
	"github.com/yaofangou/open_taobao"
)





/* 导购效果SPM报表查询，可查询某天某站点某频道某页面为淘宝带来的流量及成交情况（对于page和module最多返回10000条数据，请正确使用spm数据） 
只能查询距离今天3个月内的数据 */
type SpmeffectGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 日期 */
func (r *SpmeffectGetRequest) SetDate(value string) {
	r.SetValue("date", value)
}

/* 查询指定的SPM第四位的效果报表。默认值为false，不传视为不需要 */
func (r *SpmeffectGetRequest) SetModuleDetail(value string) {
	r.SetValue("module_detail", value)
}

/* 查询指定的SPM第三位的效果报表。默认值为false，不传视为不需要 */
func (r *SpmeffectGetRequest) SetPageDetail(value string) {
	r.SetValue("page_detail", value)
}


func (r *SpmeffectGetRequest) GetResponse(accessToken string) (*SpmeffectGetResponse, []byte, error) {
	var resp SpmeffectGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.spmeffect.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SpmeffectGetResponse struct {
	SpmResult *SpmResult `json:"spm_result"`
}

type SpmeffectGetResponseResult struct {
	Response *SpmeffectGetResponse `json:"spmeffect_get_response"`
}





/* 获取天猫折扣商品的精选结果。在得到商品ID列表后，再调用 taobao.taobaoke.widget.items.convert 获取有佣金的淘客推广链接。auction_tag不再支持天猫精品库。如有需要调用精选商品，请改为调用：tmall.selected.items.search */
type TmallItemsDiscountSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品标签。品牌特卖商品库：3458；天猫原创商品库：4801,不再支持 天猫精品库：8578； */
func (r *TmallItemsDiscountSearchRequest) SetAuctionTag(value string) {
	r.SetValue("auction_tag", value)
}

/* 品牌的id。点击某品牌类目的时候会出现。 */
func (r *TmallItemsDiscountSearchRequest) SetBrand(value string) {
	r.SetValue("brand", value)
}

/* 前台类目id，目前其他接口无法获取，只能自己去寻找。建议使用关键字获取数据。支持多选过滤，cat=catid1,catid2。 */
func (r *TmallItemsDiscountSearchRequest) SetCat(value string) {
	r.SetValue("cat", value)
}

/* 商品最高价格 */
func (r *TmallItemsDiscountSearchRequest) SetEndPrice(value string) {
	r.SetValue("end_price", value)
}

/* 是否包邮，-1为包邮 */
func (r *TmallItemsDiscountSearchRequest) SetPostFee(value string) {
	r.SetValue("post_fee", value)
}

/* 表示搜索的关键字，例如搜索query=nike。当输入关键字为中文时，将对他进行URLEncode的UTF-8格式编码，如 耐克，那么q=%E8%80%90%E5%85%8B。 */
func (r *TmallItemsDiscountSearchRequest) SetQ(value string) {
	r.SetValue("q", value)
}

/* 排序类型。类型包括： 
s: 人气排序 
p: 价格从低到高; 
pd: 价格从高到低; 
d: 月销量从高到低; 
td: 总销量从高到低; 
pt: 按发布时间排序. */
func (r *TmallItemsDiscountSearchRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 可以用该字段来实现分页功能。表示查询起始位置，默认从第0条开始，start=10,表示从第10条数据开始查询，start不得大于1000。 */
func (r *TmallItemsDiscountSearchRequest) SetStart(value string) {
	r.SetValue("start", value)
}

/* 商品最低价格 */
func (r *TmallItemsDiscountSearchRequest) SetStartPrice(value string) {
	r.SetValue("start_price", value)
}


func (r *TmallItemsDiscountSearchRequest) GetResponse(accessToken string) (*TmallItemsDiscountSearchResponse, []byte, error) {
	var resp TmallItemsDiscountSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.items.discount.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallItemsDiscountSearchResponse struct {
	BrandList *TmallBrandListObject `json:"brand_list"`
	CatList *TmallCatListObject `json:"cat_list"`
	ItemList *TmallSearchItemListObject `json:"item_list"`
	MinisiteList *TmallMinisiteListObject `json:"minisite_list"`
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	ParamValue string `json:"param_value"`
	SearchUrl string `json:"search_url"`
	TotalPage int `json:"total_page"`
	TotalResults string `json:"total_results"`
}

type TmallItemsDiscountSearchResponseResult struct {
	Response *TmallItemsDiscountSearchResponse `json:"tmall_items_discount_search_response"`
}





/* 获取天猫各类目下精选商品列表。在得到商品ID列表后，再调用taobao.item.get获取商品详情，再调用taobao.taobaoke.items.convert 获取有佣金的淘客推广链接 */
type TmallSelectedItemsSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 后台类目ID，支持父类目或叶子类目，可以通过taobao.itemcats.get 获取到后台类目ID列表 */
func (r *TmallSelectedItemsSearchRequest) SetCid(value string) {
	r.SetValue("cid", value)
}


func (r *TmallSelectedItemsSearchRequest) GetResponse(accessToken string) (*TmallSelectedItemsSearchResponse, []byte, error) {
	var resp TmallSelectedItemsSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.selected.items.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallSelectedItemsSearchResponse struct {
	ItemList *SelectedItemListObject `json:"item_list"`
}

type TmallSelectedItemsSearchResponseResult struct {
	Response *TmallSelectedItemsSearchResponse `json:"tmall_selected_items_search_response"`
}





/* 根据品牌特卖（temai.tmall.com）类目返回品牌特卖商品列表。在得到商品ID列表后，再调用taobao.item.get获取商品详情，再调用taobao.taobaoke.items.convert 获取有佣金的淘客推广链接 */
type TmallTemaiItemsSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 特卖前台类目id，传入的必须是特卖类目50100982或其下的二级类目。 */
func (r *TmallTemaiItemsSearchRequest) SetCat(value string) {
	r.SetValue("cat", value)
}

/* 排序类型。类型包括： s: 人气排序 p: 价格从低到高; pd: 价格从高到低; d: 月销量从高到低; pt: 按发布时间排序. */
func (r *TmallTemaiItemsSearchRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 表示查询起始位置: 
start=0:返回第1条记录到第48条记录（即第一页）； 
start=48:返回第48条记录到第96条记录（即第二页）； 
start=96，start=144，start=192...... 
依此类推，每次加start值加48(每页返回记录数固定为48条) */
func (r *TmallTemaiItemsSearchRequest) SetStart(value string) {
	r.SetValue("start", value)
}


func (r *TmallTemaiItemsSearchRequest) GetResponse(accessToken string) (*TmallTemaiItemsSearchResponse, []byte, error) {
	var resp TmallTemaiItemsSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.temai.items.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallTemaiItemsSearchResponse struct {
	ItemList *TmallSearchTmItemListObject `json:"item_list"`
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalPage int `json:"total_page"`
	TotalResults int `json:"total_results"`
}

type TmallTemaiItemsSearchResponseResult struct {
	Response *TmallTemaiItemsSearchResponse `json:"tmall_temai_items_search_response"`
}





/* 特卖子类目搜索 */
type TmallTemaiSubcatsSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 父类目ID，固定是特卖前台一级类目id：50100982 */
func (r *TmallTemaiSubcatsSearchRequest) SetCat(value string) {
	r.SetValue("cat", value)
}


func (r *TmallTemaiSubcatsSearchRequest) GetResponse(accessToken string) (*TmallTemaiSubcatsSearchResponse, []byte, error) {
	var resp TmallTemaiSubcatsSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.temai.subcats.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallTemaiSubcatsSearchResponse struct {
	CatList *TmallTmCatListObject `json:"cat_list"`
}

type TmallTemaiSubcatsSearchResponseResult struct {
	Response *TmallTemaiSubcatsSearchResponse `json:"tmall_temai_subcats_search_response"`
}



