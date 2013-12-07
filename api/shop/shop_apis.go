// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package shop

import (
	"github.com/yaofangou/open_taobao"
)





/* 此API添加卖家店铺内自定义类目  
父类目parent_cid值等于0：表示此类目为店铺下的一级类目，值不等于0：表示此类目有父类目  
注：因为缓存的关系,添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布 */
type SellercatsListAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 卖家自定义类目名称。不超过20个字符 */
func (r *SellercatsListAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 父类目编号，如果类目为店铺下的一级类目：值等于0，如果类目为子类目，调用获取taobao.sellercats.list.get父类目编号 */
func (r *SellercatsListAddRequest) SetParentCid(value string) {
	r.SetValue("parent_cid", value)
}

/* 链接图片URL地址。(绝对地址，格式：http://host/image_path) */
func (r *SellercatsListAddRequest) SetPictUrl(value string) {
	r.SetValue("pict_url", value)
}

/* 该类目在页面上的排序位置,取值范围:大于零的整数 */
func (r *SellercatsListAddRequest) SetSortOrder(value string) {
	r.SetValue("sort_order", value)
}


func (r *SellercatsListAddRequest) GetResponse(accessToken string) (*SellercatsListAddResponse, []byte, error) {
	var resp SellercatsListAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.sellercats.list.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SellercatsListAddResponse struct {
	SellerCat *SellerCat `json:"seller_cat"`
}

type SellercatsListAddResponseResult struct {
	Response *SellercatsListAddResponse `json:"sellercats_list_add_response"`
}





/* 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家） */
type SellercatsListGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 卖家昵称 */
func (r *SellercatsListGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SellercatsListGetRequest) GetResponse(accessToken string) (*SellercatsListGetResponse, []byte, error) {
	var resp SellercatsListGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.sellercats.list.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SellercatsListGetResponse struct {
	SellerCats *SellerCatListObject `json:"seller_cats"`
}

type SellercatsListGetResponseResult struct {
	Response *SellercatsListGetResponse `json:"sellercats_list_get_response"`
}





/* 此API更新卖家店铺内自定义类目  
注：因为缓存的关系，添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布 */
type SellercatsListUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 卖家自定义类目编号 */
func (r *SellercatsListUpdateRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 卖家自定义类目名称。不超过20个字符 */
func (r *SellercatsListUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 链接图片URL地址 */
func (r *SellercatsListUpdateRequest) SetPictUrl(value string) {
	r.SetValue("pict_url", value)
}

/* 该类目在页面上的排序位置,取值范围:大于零的整数 */
func (r *SellercatsListUpdateRequest) SetSortOrder(value string) {
	r.SetValue("sort_order", value)
}


func (r *SellercatsListUpdateRequest) GetResponse(accessToken string) (*SellercatsListUpdateResponse, []byte, error) {
	var resp SellercatsListUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.sellercats.list.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SellercatsListUpdateResponse struct {
	SellerCat *SellerCat `json:"seller_cat"`
}

type SellercatsListUpdateResponseResult struct {
	Response *SellercatsListUpdateResponse `json:"sellercats_list_update_response"`
}





/* 获取卖家店铺的基本信息 */
type ShopGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔 */
func (r *ShopGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 卖家昵称 */
func (r *ShopGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *ShopGetRequest) GetResponse(accessToken string) (*ShopGetResponse, []byte, error) {
	var resp ShopGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.shop.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ShopGetResponse struct {
	Shop *Shop `json:"shop"`
}

type ShopGetResponseResult struct {
	Response *ShopGetResponse `json:"shop_get_response"`
}





/* 获取卖家店铺剩余橱窗数量，已用橱窗数量，总橱窗数量 */
type ShopRemainshowcaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *ShopRemainshowcaseGetRequest) GetResponse(accessToken string) (*ShopRemainshowcaseGetResponse, []byte, error) {
	var resp ShopRemainshowcaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.shop.remainshowcase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ShopRemainshowcaseGetResponse struct {
	Shop *Shop `json:"shop"`
}

type ShopRemainshowcaseGetResponseResult struct {
	Response *ShopRemainshowcaseGetResponse `json:"shop_remainshowcase_get_response"`
}





/* 目前只支持标题、公告和描述的更新 */
type ShopUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 店铺公告。不超过1024个字符 */
func (r *ShopUpdateRequest) SetBulletin(value string) {
	r.SetValue("bulletin", value)
}

/* 店铺描述。10～2000个字符以内 */
func (r *ShopUpdateRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 店铺标题。不超过30个字符；过滤敏感词，如淘咖啡、阿里巴巴等。title, bulletin和desc至少必须传一个 */
func (r *ShopUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *ShopUpdateRequest) GetResponse(accessToken string) (*ShopUpdateResponse, []byte, error) {
	var resp ShopUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.shop.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ShopUpdateResponse struct {
	Shop *Shop `json:"shop"`
}

type ShopUpdateResponseResult struct {
	Response *ShopUpdateResponse `json:"shop_update_response"`
}





/* 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异） */
type ShopcatsListGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent */
func (r *ShopcatsListGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}


func (r *ShopcatsListGetRequest) GetResponse(accessToken string) (*ShopcatsListGetResponse, []byte, error) {
	var resp ShopcatsListGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.shopcats.list.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ShopcatsListGetResponse struct {
	ShopCats *ShopCatListObject `json:"shop_cats"`
}

type ShopcatsListGetResponseResult struct {
	Response *ShopcatsListGetResponse `json:"shopcats_list_get_response"`
}



