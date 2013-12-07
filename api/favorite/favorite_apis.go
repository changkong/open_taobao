// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package favorite

import (
	"github.com/yaofangou/open_taobao"
)





/* 添加商品或店铺到收藏夹 */
type FavoriteAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* ITEM表示宝贝，SHOP表示店铺，只能传入这两者之一 */
func (r *FavoriteAddRequest) SetCollectType(value string) {
	r.SetValue("collect_type", value)
}

/* 如果收藏的是商品，就传num_iid，如果是店铺，就传入sid */
func (r *FavoriteAddRequest) SetItemNumid(value string) {
	r.SetValue("item_numid", value)
}

/* 该收藏是否公开 */
func (r *FavoriteAddRequest) SetShared(value string) {
	r.SetValue("shared", value)
}


func (r *FavoriteAddRequest) GetResponse(accessToken string) (*FavoriteAddResponse, []byte, error) {
	var resp FavoriteAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.favorite.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FavoriteAddResponse struct {
	Success bool `json:"success"`
}

type FavoriteAddResponseResult struct {
	Response *FavoriteAddResponse `json:"favorite_add_response"`
}





/* 查询淘宝用户收藏的商品或店铺信息 */
type FavoriteSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* ITEM表示宝贝，SHOP表示商铺，collect_type只能为这两者之一 */
func (r *FavoriteSearchRequest) SetCollectType(value string) {
	r.SetValue("collect_type", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1。一页20条记录。 */
func (r *FavoriteSearchRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}


func (r *FavoriteSearchRequest) GetResponse(accessToken string) (*FavoriteSearchResponse, []byte, error) {
	var resp FavoriteSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.favorite.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FavoriteSearchResponse struct {
	CollectItems *CollectItemListObject `json:"collect_items"`
	TotalResults int `json:"total_results"`
}

type FavoriteSearchResponseResult struct {
	Response *FavoriteSearchResponse `json:"favorite_search_response"`
}



