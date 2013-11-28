// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pinjia

import (
	"github.com/yaofangou/open_taobao"
)





/* 新增单个评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不能再通过该接口进行评价</font>) */
type TraderateAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 是否匿名,卖家评不能匿名。可选值:true(匿名),false(非匿名)。注意：如果交易订单匿名，则评价也匿名 */
func (r *TraderateAddRequest) SetAnony(value string) {
	r.SetValue("anony", value)
}

/* 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容 */
func (r *TraderateAddRequest) SetContent(value string) {
	r.SetValue("content", value)
}

/* 子订单ID */
func (r *TraderateAddRequest) SetOid(value string) {
	r.SetValue("oid", value)
}

/* 评价结果,可选值:good(好评),neutral(中评),bad(差评) */
func (r *TraderateAddRequest) SetResult(value string) {
	r.SetValue("result", value)
}

/* 评价者角色,可选值:seller(卖家),buyer(买家) */
func (r *TraderateAddRequest) SetRole(value string) {
	r.SetValue("role", value)
}

/* 交易ID */
func (r *TraderateAddRequest) SetTid(value string) {
	r.SetValue("tid", value)
}


func (r *TraderateAddRequest) GetResponse(accessToken string) (*TraderateAddResponse, []byte, error) {
	var resp TraderateAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.traderate.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TraderateAddResponse struct {
	TradeRate *TradeRate `json:"trade_rate"`
}

type TraderateAddResponseResult struct {
	Response *TraderateAddResponse `json:"traderate_add_response"`
}





/* 商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释） */
type TraderateExplainAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 子订单ID */
func (r *TraderateExplainAddRequest) SetOid(value string) {
	r.SetValue("oid", value)
}

/* 评价解释内容,最大长度: 500个汉字 */
func (r *TraderateExplainAddRequest) SetReply(value string) {
	r.SetValue("reply", value)
}


func (r *TraderateExplainAddRequest) GetResponse(accessToken string) (*TraderateExplainAddResponse, []byte, error) {
	var resp TraderateExplainAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.traderate.explain.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TraderateExplainAddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TraderateExplainAddResponseResult struct {
	Response *TraderateExplainAddResponse `json:"traderate_explain_add_response"`
}





/* 根据商品id获取这个商品id对应的大家印象 */
type TraderateImprImprwordByaucidGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 淘宝的商品id */
func (r *TraderateImprImprwordByaucidGetRequest) SetAuctionId(value string) {
	r.SetValue("auction_id", value)
}


func (r *TraderateImprImprwordByaucidGetRequest) GetResponse(accessToken string) (*TraderateImprImprwordByaucidGetResponse, []byte, error) {
	var resp TraderateImprImprwordByaucidGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.traderate.impr.imprword.byaucid.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TraderateImprImprwordByaucidGetResponse struct {
	ImprWordList []*ImprItemDO `json:"impr_word_list"`
}

type TraderateImprImprwordByaucidGetResponseResult struct {
	Response *TraderateImprImprwordByaucidGetResponse `json:"traderate_impr_imprword_byaucid_get_response"`
}





/* 根据卖家nick和交易id（如果是子订单，输入子订单id），查询到该条评价的大家印象标签 */
type TraderateImprImprwordByfeedidGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 交易订单id号（如果包含子订单，请使用子订单id号） */
func (r *TraderateImprImprwordByfeedidGetRequest) SetChildTradeId(value string) {
	r.SetValue("child_trade_id", value)
}


func (r *TraderateImprImprwordByfeedidGetRequest) GetResponse(accessToken string) (*TraderateImprImprwordByfeedidGetResponse, []byte, error) {
	var resp TraderateImprImprwordByfeedidGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.traderate.impr.imprword.byfeedid.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TraderateImprImprwordByfeedidGetResponse struct {
	ImprFeed *ImprFeedIdDO `json:"impr_feed"`
}

type TraderateImprImprwordByfeedidGetResponseResult struct {
	Response *TraderateImprImprwordByfeedidGetResponse `json:"traderate_impr_imprword_byfeedid_get_response"`
}





/* 根据淘宝后台类目的一级类目和叶子类目 */
type TraderateImprImprwordsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 淘宝叶子类目id */
func (r *TraderateImprImprwordsGetRequest) SetCatLeafId(value string) {
	r.SetValue("cat_leaf_id", value)
}

/* 淘宝一级类目id */
func (r *TraderateImprImprwordsGetRequest) SetCatRootId(value string) {
	r.SetValue("cat_root_id", value)
}


func (r *TraderateImprImprwordsGetRequest) GetResponse(accessToken string) (*TraderateImprImprwordsGetResponse, []byte, error) {
	var resp TraderateImprImprwordsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.traderate.impr.imprwords.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TraderateImprImprwordsGetResponse struct {
	ImprWords []string `json:"impr_words"`
}

type TraderateImprImprwordsGetResponseResult struct {
	Response *TraderateImprImprwordsGetResponse `json:"traderate_impr_imprwords_get_response"`
}





/* 针对父子订单新增批量评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不用再通过该接口进行评价</font>) */
type TraderateListAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 是否匿名，卖家评不能匿名。可选值:true(匿名),false(非匿名)。 注意：如果买家匿名购买，那么买家的评价默认匿名 */
func (r *TraderateListAddRequest) SetAnony(value string) {
	r.SetValue("anony", value)
}

/* 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容 */
func (r *TraderateListAddRequest) SetContent(value string) {
	r.SetValue("content", value)
}

/* 评价结果。可选值:good(好评),neutral(中评),bad(差评) */
func (r *TraderateListAddRequest) SetResult(value string) {
	r.SetValue("result", value)
}

/* 评价者角色。可选值:seller(卖家),buyer(买家) */
func (r *TraderateListAddRequest) SetRole(value string) {
	r.SetValue("role", value)
}

/* 交易ID */
func (r *TraderateListAddRequest) SetTid(value string) {
	r.SetValue("tid", value)
}


func (r *TraderateListAddRequest) GetResponse(accessToken string) (*TraderateListAddResponse, []byte, error) {
	var resp TraderateListAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.traderate.list.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TraderateListAddResponse struct {
	TradeRate *TradeRate `json:"trade_rate"`
}

type TraderateListAddResponseResult struct {
	Response *TraderateListAddResponse `json:"traderate_list_add_response"`
}





/* 搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。 */
type TraderatesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。 */
func (r *TraderatesGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔 */
func (r *TraderatesGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 商品的数字ID */
func (r *TraderatesGetRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 页码。取值范围:大于零的整数最大限制为200; 默认值:1 */
func (r *TraderatesGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页获取条数。默认值40，最小值1，最大值150。 */
func (r *TraderatesGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 评价类型。可选值:get(得到),give(给出) */
func (r *TraderatesGetRequest) SetRateType(value string) {
	r.SetValue("rate_type", value)
}

/* 评价结果。可选值:good(好评),neutral(中评),bad(差评) */
func (r *TraderatesGetRequest) SetResult(value string) {
	r.SetValue("result", value)
}

/* 评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。 */
func (r *TraderatesGetRequest) SetRole(value string) {
	r.SetValue("role", value)
}

/* 评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。 */
func (r *TraderatesGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}

/* 交易订单id，可以是父订单id号，也可以是子订单id号 */
func (r *TraderatesGetRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。 */
func (r *TraderatesGetRequest) SetUseHasNext(value string) {
	r.SetValue("use_has_next", value)
}


func (r *TraderatesGetRequest) GetResponse(accessToken string) (*TraderatesGetResponse, []byte, error) {
	var resp TraderatesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.traderates.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TraderatesGetResponse struct {
	HasNext bool `json:"has_next"`
	TotalResults int `json:"total_results"`
	TradeRates []*TradeRate `json:"trade_rates"`
}

type TraderatesGetResponseResult struct {
	Response *TraderatesGetResponse `json:"traderates_get_response"`
}



