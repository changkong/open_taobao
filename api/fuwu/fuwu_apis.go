// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fuwu

import (
	"github.com/yaofangou/open_taobao"
)





/* 服务商通过使用该接口来产生营销链接，通过把这种链接发送给商家来做自定义人群的服务营销<br> 
注：session是param_str这个参数串创建者生成的session，这个创建者与入参中的nick是不一致的。例如：A开发者创建了一个param_str的字符串，要为B商家生成一个营销链接，session必须是A开发者创建的session。 */
type FuwuSaleLinkGenRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 用户需要营销的目标人群中的用户nick */
func (r *FuwuSaleLinkGenRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 从服务商后台，营销链接功能中生成的参数串直接复制使用。不要修改，否则抛错。 */
func (r *FuwuSaleLinkGenRequest) SetParamStr(value string) {
	r.SetValue("param_str", value)
}


func (r *FuwuSaleLinkGenRequest) GetResponse(accessToken string) (*FuwuSaleLinkGenResponse, []byte, error) {
	var resp FuwuSaleLinkGenResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fuwu.sale.link.gen", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FuwuSaleLinkGenResponse struct {
	Url string `json:"url"`
}

type FuwuSaleLinkGenResponseResult struct {
	Response *FuwuSaleLinkGenResponse `json:"fuwu_sale_link_gen_response"`
}





/* 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。 
建议用于查询前一日的历史记录，不适合用作实时数据查询。 
现在只能查询90天以内的数据 
该接口限制每分钟所有appkey调用总和只能有800次。 */
type VasOrderSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
func (r *VasOrderSearchRequest) SetArticleCode(value string) {
	r.SetValue("article_code", value)
}

/* 订单号 */
func (r *VasOrderSearchRequest) SetBizOrderId(value string) {
	r.SetValue("biz_order_id", value)
}

/* 订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部 */
func (r *VasOrderSearchRequest) SetBizType(value string) {
	r.SetValue("biz_type", value)
}

/* 订单创建时间（订购时间）结束值 */
func (r *VasOrderSearchRequest) SetEndCreated(value string) {
	r.SetValue("end_created", value)
}

/* 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
func (r *VasOrderSearchRequest) SetItemCode(value string) {
	r.SetValue("item_code", value)
}

/* 淘宝会员名 */
func (r *VasOrderSearchRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 子订单号 */
func (r *VasOrderSearchRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* 页码 */
func (r *VasOrderSearchRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 一页包含的记录数 */
func (r *VasOrderSearchRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据） */
func (r *VasOrderSearchRequest) SetStartCreated(value string) {
	r.SetValue("start_created", value)
}


func (r *VasOrderSearchRequest) GetResponse(accessToken string) (*VasOrderSearchResponse, []byte, error) {
	var resp VasOrderSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.vas.order.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type VasOrderSearchResponse struct {
	ArticleBizOrders []*ArticleBizOrder `json:"article_biz_orders"`
	TotalItem int `json:"total_item"`
}

type VasOrderSearchResponseResult struct {
	Response *VasOrderSearchResponse `json:"vas_order_search_response"`
}





/* 用于ISV查询自己名下的应用及收费项目的订购记录 */
type VasSubscSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
func (r *VasSubscSearchRequest) SetArticleCode(value string) {
	r.SetValue("article_code", value)
}

/* 是否自动续费，true=自动续费 false=非自动续费 空=全部 */
func (r *VasSubscSearchRequest) SetAutosub(value string) {
	r.SetValue("autosub", value)
}

/* 到期时间结束值 */
func (r *VasSubscSearchRequest) SetEndDeadline(value string) {
	r.SetValue("end_deadline", value)
}

/* 是否到期提醒，true=到期提醒 false=非到期提醒 空=全部 */
func (r *VasSubscSearchRequest) SetExpireNotice(value string) {
	r.SetValue("expire_notice", value)
}

/* 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
func (r *VasSubscSearchRequest) SetItemCode(value string) {
	r.SetValue("item_code", value)
}

/* 淘宝会员名 */
func (r *VasSubscSearchRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *VasSubscSearchRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 一页包含的记录数 */
func (r *VasSubscSearchRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据） */
func (r *VasSubscSearchRequest) SetStartDeadline(value string) {
	r.SetValue("start_deadline", value)
}

/* 订购记录状态，1=有效 2=过期 空=全部 */
func (r *VasSubscSearchRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *VasSubscSearchRequest) GetResponse(accessToken string) (*VasSubscSearchResponse, []byte, error) {
	var resp VasSubscSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.vas.subsc.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type VasSubscSearchResponse struct {
	ArticleSubs []*ArticleSub `json:"article_subs"`
	TotalItem int `json:"total_item"`
}

type VasSubscSearchResponseResult struct {
	Response *VasSubscSearchResponse `json:"vas_subsc_search_response"`
}





/* 用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况 */
type VasSubscribeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
func (r *VasSubscribeGetRequest) SetArticleCode(value string) {
	r.SetValue("article_code", value)
}

/* 淘宝会员名 */
func (r *VasSubscribeGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *VasSubscribeGetRequest) GetResponse(accessToken string) (*VasSubscribeGetResponse, []byte, error) {
	var resp VasSubscribeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.vas.subscribe.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type VasSubscribeGetResponse struct {
	ArticleUserSubscribes []*ArticleUserSubscribe `json:"article_user_subscribes"`
}

type VasSubscribeGetResponseResult struct {
	Response *VasSubscribeGetResponse `json:"vas_subscribe_get_response"`
}



