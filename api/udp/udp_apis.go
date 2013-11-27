// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package udp

import (
	"github.com/changkong/open_taobao"
)





/* 商品指标查询 */
type UdpItemGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 地区ID */
func (r *UdpItemGetRequest) SetArea(value string) {
	r.SetValue("area", value)
}

/* 开始时间 */
func (r *UdpItemGetRequest) SetBeginTime(value string) {
	r.SetValue("begin_time", value)
}

/* 结束时间 */
func (r *UdpItemGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 指标ID(参阅指标编号) */
func (r *UdpItemGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 商品ID */
func (r *UdpItemGetRequest) SetItemid(value string) {
	r.SetValue("itemid", value)
}

/* 多个宝贝列表 */
func (r *UdpItemGetRequest) SetItems(value string) {
	r.SetValue("items", value)
}

/* 排序指标 */
func (r *UdpItemGetRequest) SetOrderBy(value string) {
	r.SetValue("order_by", value)
}

/* 排序规则 */
func (r *UdpItemGetRequest) SetOrderRule(value string) {
	r.SetValue("order_rule", value)
}

/* 查询页码，0为第一页 */
func (r *UdpItemGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页记录数 */
func (r *UdpItemGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 备用 */
func (r *UdpItemGetRequest) SetParameters(value string) {
	r.SetValue("parameters", value)
}

/* 来源ID */
func (r *UdpItemGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}


func (r *UdpItemGetRequest) GetResponse(accessToken string) (*UdpItemGetResponse, []byte, error) {
	var resp UdpItemGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.udp.item.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UdpItemGetResponse struct {
	Content *TargetSearchTopResult `json:"content"`
}

type UdpItemGetResponseResult struct {
	Response *UdpItemGetResponse `json:"udp_item_get_response"`
}





/* 聚划算指标查询 */
type UdpJuhuasuanGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 开始时间 */
func (r *UdpJuhuasuanGetRequest) SetBeginTime(value string) {
	r.SetValue("begin_time", value)
}

/* 类目编号 */
func (r *UdpJuhuasuanGetRequest) SetCatid(value string) {
	r.SetValue("catid", value)
}

/* 结束时间 */
func (r *UdpJuhuasuanGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 指标ID(参阅指标编号) */
func (r *UdpJuhuasuanGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 2407364 */
func (r *UdpJuhuasuanGetRequest) SetItemid(value string) {
	r.SetValue("itemid", value)
}

/* 备用 */
func (r *UdpJuhuasuanGetRequest) SetParameters(value string) {
	r.SetValue("parameters", value)
}


func (r *UdpJuhuasuanGetRequest) GetResponse(accessToken string) (*UdpJuhuasuanGetResponse, []byte, error) {
	var resp UdpJuhuasuanGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.udp.juhuasuan.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UdpJuhuasuanGetResponse struct {
	Content *TargetSearchTopResult `json:"content"`
}

type UdpJuhuasuanGetResponseResult struct {
	Response *UdpJuhuasuanGetResponse `json:"udp_juhuasuan_get_response"`
}





/* 店铺指标查询 */
type UdpShopGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 地区ID(参阅地区编号) */
func (r *UdpShopGetRequest) SetArea(value string) {
	r.SetValue("area", value)
}

/* 开始时间 */
func (r *UdpShopGetRequest) SetBeginTime(value string) {
	r.SetValue("begin_time", value)
}

/* 结束时间 */
func (r *UdpShopGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 指标ID(参阅指标编号) */
func (r *UdpShopGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 备用 */
func (r *UdpShopGetRequest) SetParameters(value string) {
	r.SetValue("parameters", value)
}


func (r *UdpShopGetRequest) GetResponse(accessToken string) (*UdpShopGetResponse, []byte, error) {
	var resp UdpShopGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.udp.shop.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UdpShopGetResponse struct {
	Content *TargetSearchTopResult `json:"content"`
}

type UdpShopGetResponseResult struct {
	Response *UdpShopGetResponse `json:"udp_shop_get_response"`
}



