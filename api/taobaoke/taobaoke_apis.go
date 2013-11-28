// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

import (
	"github.com/yaofangou/open_taobao"
)





/* 批量查询卖家是否支持返利 */
type TaobaokeRebateAuthGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* nick或seller_id或num_iid，最大输入40个.格式如:"value1,value2,value3" 用","号分隔 */
func (r *TaobaokeRebateAuthGetRequest) SetParams(value string) {
	r.SetValue("params", value)
}

/* 类型：1-按nick查询，2-按seller_id查询，3-按num_iid查询 */
func (r *TaobaokeRebateAuthGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *TaobaokeRebateAuthGetRequest) GetResponse(accessToken string) (*TaobaokeRebateAuthGetResponse, []byte, error) {
	var resp TaobaokeRebateAuthGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.rebate.auth.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeRebateAuthGetResponse struct {
	Results []*TaobaokeAuthorize `json:"results"`
}

type TaobaokeRebateAuthGetResponseResult struct {
	Response *TaobaokeRebateAuthGetResponse `json:"taobaoke_rebate_auth_get_response"`
}





/* 查询卖家是否支持返利 */
type TaobaokeRebateAuthorizeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 卖家淘宝会员昵称.注：指的是淘宝的会员登录名 */
func (r *TaobaokeRebateAuthorizeGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 商品数字ID */
func (r *TaobaokeRebateAuthorizeGetRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 卖家淘宝会员数字ID. */
func (r *TaobaokeRebateAuthorizeGetRequest) SetSellerId(value string) {
	r.SetValue("seller_id", value)
}


func (r *TaobaokeRebateAuthorizeGetRequest) GetResponse(accessToken string) (*TaobaokeRebateAuthorizeGetResponse, []byte, error) {
	var resp TaobaokeRebateAuthorizeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.rebate.authorize.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeRebateAuthorizeGetResponse struct {
	Rebate bool `json:"rebate"`
}

type TaobaokeRebateAuthorizeGetResponseResult struct {
	Response *TaobaokeRebateAuthorizeGetResponse `json:"taobaoke_rebate_authorize_get_response"`
}





/* 根据输入开始时间，时间跨度，查询90天以内订单成功且支持返利淘宝客报表 */
type TaobaokeRebateReportGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表.可选值:TaobaokePayment淘宝客订单构体中的所有字段;字段之间用","分隔. */
func (r *TaobaokeRebateReportGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 当前页数 */
func (r *TaobaokeRebateReportGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页返回结果数，最小每页40条，默认每页40条，最大每页100条 */
func (r *TaobaokeRebateReportGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询报表的时间跨度，单位秒。 
以用户输入的start_time时间为起始时间，start_time+span为结束时间，查询该时间段内的订单。span最小值为60秒，最大值为600秒，默认值为60秒 */
func (r *TaobaokeRebateReportGetRequest) SetSpan(value string) {
	r.SetValue("span", value)
}

/* 需要查询报表的开始日期，有效的日期为从当前日期开始起90天以内的订单 */
func (r *TaobaokeRebateReportGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *TaobaokeRebateReportGetRequest) GetResponse(accessToken string) (*TaobaokeRebateReportGetResponse, []byte, error) {
	var resp TaobaokeRebateReportGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.rebate.report.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeRebateReportGetResponse struct {
	TaobaokePayments []*TaobaokePayment `json:"taobaoke_payments"`
}

type TaobaokeRebateReportGetResponseResult struct {
	Response *TaobaokeRebateReportGetResponse `json:"taobaoke_rebate_report_get_response"`
}



