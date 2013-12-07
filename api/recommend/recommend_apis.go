// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package recommend

import (
	"github.com/yaofangou/open_taobao"
)





/* 根据类目信息推荐相关联的宝贝集 */
type CategoryrecommendItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 传入叶子类目ID */
func (r *CategoryrecommendItemsGetRequest) SetCategoryId(value string) {
	r.SetValue("category_id", value)
}

/* 请求个数，建议获取20个 */
func (r *CategoryrecommendItemsGetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 额外参数 */
func (r *CategoryrecommendItemsGetRequest) SetExt(value string) {
	r.SetValue("ext", value)
}

/* 请求类型，1：类目下热门商品推荐。其他值当非法值处理 */
func (r *CategoryrecommendItemsGetRequest) SetRecommendType(value string) {
	r.SetValue("recommend_type", value)
}


func (r *CategoryrecommendItemsGetRequest) GetResponse(accessToken string) (*CategoryrecommendItemsGetResponse, []byte, error) {
	var resp CategoryrecommendItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.categoryrecommend.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CategoryrecommendItemsGetResponse struct {
	FavoriteItems *FavoriteItemListObject `json:"favorite_items"`
}

type CategoryrecommendItemsGetResponseResult struct {
	Response *CategoryrecommendItemsGetResponse `json:"categoryrecommend_items_get_response"`
}





/* 根据推荐类型获取推荐的关联关系商品 */
type ItemrecommendItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 请求返回宝贝的个数，建议取20个 */
func (r *ItemrecommendItemsGetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 额外的参数信息 */
func (r *ItemrecommendItemsGetRequest) SetExt(value string) {
	r.SetValue("ext", value)
}

/* 商品ID */
func (r *ItemrecommendItemsGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 查询类型标识符，可传入1-3，1：同类商品推荐，2：异类商品推荐， 3：同店商品推荐。其他值当非法值处理 */
func (r *ItemrecommendItemsGetRequest) SetRecommendType(value string) {
	r.SetValue("recommend_type", value)
}


func (r *ItemrecommendItemsGetRequest) GetResponse(accessToken string) (*ItemrecommendItemsGetResponse, []byte, error) {
	var resp ItemrecommendItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.itemrecommend.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemrecommendItemsGetResponse struct {
	Values *FavoriteItemListObject `json:"values"`
}

type ItemrecommendItemsGetResponseResult struct {
	Response *ItemrecommendItemsGetResponse `json:"itemrecommend_items_get_response"`
}





/* 根据店铺信息推荐相关联的宝贝集 */
type ShoprecommendItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 请求个数，最大只能获取10个 */
func (r *ShoprecommendItemsGetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 额外参数 */
func (r *ShoprecommendItemsGetRequest) SetExt(value string) {
	r.SetValue("ext", value)
}

/* 请求类型，1：店内热门商品推荐。其他值当非法值处理 */
func (r *ShoprecommendItemsGetRequest) SetRecommendType(value string) {
	r.SetValue("recommend_type", value)
}

/* <p>传入卖家ID。这里的seller_id得通过<a href="http://api.taobao.com/apidoc/api.htm?path=cid:38-apiId:10449">taobao.taobaoke.shops.get</a>
跟<a href="http://api.taobao.com/apidoc/api.htm?path=cid:38-apiId:21419">taobao.taobaoke.widget.shops.convert</a>这两个接口去获取user_id字段。</p>
<p>如果是非淘客卖家，则无法获取，暂无替代方案。</p> */
func (r *ShoprecommendItemsGetRequest) SetSellerId(value string) {
	r.SetValue("seller_id", value)
}


func (r *ShoprecommendItemsGetRequest) GetResponse(accessToken string) (*ShoprecommendItemsGetResponse, []byte, error) {
	var resp ShoprecommendItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.shoprecommend.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ShoprecommendItemsGetResponse struct {
	FavoriteItems *FavoriteItemListObject `json:"favorite_items"`
}

type ShoprecommendItemsGetResponseResult struct {
	Response *ShoprecommendItemsGetResponse `json:"shoprecommend_items_get_response"`
}





/* 根据店铺信息推荐相关联的店铺集 */
type ShoprecommendShopsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 请求个数，建议获取16个 */
func (r *ShoprecommendShopsGetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 额外参数 */
func (r *ShoprecommendShopsGetRequest) SetExt(value string) {
	r.SetValue("ext", value)
}

/* 请求类型，1：关联店铺推荐。其他值当非法值处理 */
func (r *ShoprecommendShopsGetRequest) SetRecommendType(value string) {
	r.SetValue("recommend_type", value)
}

/* 传入卖家ID */
func (r *ShoprecommendShopsGetRequest) SetSellerId(value string) {
	r.SetValue("seller_id", value)
}


func (r *ShoprecommendShopsGetRequest) GetResponse(accessToken string) (*ShoprecommendShopsGetResponse, []byte, error) {
	var resp ShoprecommendShopsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.shoprecommend.shops.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ShoprecommendShopsGetResponse struct {
	FavoriteShops *FavoriteShopListObject `json:"favorite_shops"`
}

type ShoprecommendShopsGetResponseResult struct {
	Response *ShoprecommendShopsGetResponse `json:"shoprecommend_shops_get_response"`
}





/* 根据用户信息推荐相关联的宝贝集。仅支持widget入口调用，需要同时校验淘宝cookie登陆情况，以及cookie和session授权的一致性。调用入口为/widget/rest。签名方法简化为Hmac-md5,hmac(secret+‘app_key' ＋app_key +'timestamp' + timestamp+secret)。timestamp为60分钟内有效 
此API为组件API，调用方式需要参照：http://open.taobao.com/doc/detail.htm?id=988，以JS-SDK调用 */
type UserrecommendItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 请求个数，建议取20个 */
func (r *UserrecommendItemsGetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 额外参数 */
func (r *UserrecommendItemsGetRequest) SetExt(value string) {
	r.SetValue("ext", value)
}

/* 请求类型，1：用户购买意图。其他值当非法值处理 */
func (r *UserrecommendItemsGetRequest) SetRecommendType(value string) {
	r.SetValue("recommend_type", value)
}


func (r *UserrecommendItemsGetRequest) GetResponse(accessToken string) (*UserrecommendItemsGetResponse, []byte, error) {
	var resp UserrecommendItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.userrecommend.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UserrecommendItemsGetResponse struct {
	FavoriteItems *FavoriteItemListObject `json:"favorite_items"`
}

type UserrecommendItemsGetResponseResult struct {
	Response *UserrecommendItemsGetResponse `json:"userrecommend_items_get_response"`
}



