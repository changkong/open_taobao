// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package simba

import (
	"github.com/yaofangou/open_taobao"
)





/* 取得一个推广组的推荐关键词列表 */
type SimbaKeywordsRecommendGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组ID */
func (r *SimbaKeywordsRecommendGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 主人昵称 */
func (r *SimbaKeywordsRecommendGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 排序方式:  
搜索量 search_volume  
市场平均价格 average_price  
相关度 relevance  
不排序 non  
默认为 non */
func (r *SimbaKeywordsRecommendGetRequest) SetOrderBy(value string) {
	r.SetValue("order_by", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaKeywordsRecommendGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,最大200 */
func (r *SimbaKeywordsRecommendGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 相关度 */
func (r *SimbaKeywordsRecommendGetRequest) SetPertinence(value string) {
	r.SetValue("pertinence", value)
}

/* 搜索量,设置此值后返回的就是大于此搜索量的词列表 */
func (r *SimbaKeywordsRecommendGetRequest) SetSearch(value string) {
	r.SetValue("search", value)
}


func (r *SimbaKeywordsRecommendGetRequest) GetResponse(accessToken string) (*SimbaKeywordsRecommendGetResponse, []byte, error) {
	var resp SimbaKeywordsRecommendGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywords.recommend.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordsRecommendGetResponse struct {
	RecommendWords *RecommendWordPage `json:"recommend_words"`
}

type SimbaKeywordsRecommendGetResponseResult struct {
	Response *SimbaKeywordsRecommendGetResponse `json:"simba_keywords_recommend_get_response"`
}





/* 词基础数据查询 */
type SimbaInsightWordsbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结果过滤。PV：返回展现量；CLICK：返回点击量；AVGCPC：返回平均出价；COMPETITION ：返回竞争宝贝数;CTR 点击率。filter可由,组合 */
func (r *SimbaInsightWordsbaseGetRequest) SetFilter(value string) {
	r.SetValue("filter", value)
}

/* 主人昵称 */
func (r *SimbaInsightWordsbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 时间格式(DAY: 最近一天； WEEK：最近一周。MONTH：最近一个月。3MONTH：最近三个月) */
func (r *SimbaInsightWordsbaseGetRequest) SetTime(value string) {
	r.SetValue("time", value)
}

/* 查询词组，最大长度170 */
func (r *SimbaInsightWordsbaseGetRequest) SetWords(value string) {
	r.SetValue("words", value)
}


func (r *SimbaInsightWordsbaseGetRequest) GetResponse(accessToken string) (*SimbaInsightWordsbaseGetResponse, []byte, error) {
	var resp SimbaInsightWordsbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.wordsbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightWordsbaseGetResponse struct {
	InWordBases *INWordBaseListObject `json:"in_word_bases"`
}

type SimbaInsightWordsbaseGetResponseResult struct {
	Response *SimbaInsightWordsbaseGetResponse `json:"simba_insight_wordsbase_get_response"`
}





/* 设置一批关键词的出价 */
type SimbaKeywordsPricevonSetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 关键词ID，出价和匹配方式json字符串，keywordId:词ID，整数。maxPrice：价格，是整数，以“分”为单位，不能小于5，不能大于日限额,当使用默认出价时必须将这个值设置为0。; isDefaultPrice：是否使用默认出价，只能是0，1(0代表不使用，1代表使用)。matchscope只能是1,2,4（1代表精确匹配，2代表子串匹配，4代表广泛匹配） */
func (r *SimbaKeywordsPricevonSetRequest) SetKeywordidPrices(value string) {
	r.SetValue("keywordid_prices", value)
}

/* 主人昵称 */
func (r *SimbaKeywordsPricevonSetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaKeywordsPricevonSetRequest) GetResponse(accessToken string) (*SimbaKeywordsPricevonSetResponse, []byte, error) {
	var resp SimbaKeywordsPricevonSetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywords.pricevon.set", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordsPricevonSetResponse struct {
	Keywords *KeywordListObject `json:"keywords"`
}

type SimbaKeywordsPricevonSetResponseResult struct {
	Response *SimbaKeywordsPricevonSetResponse `json:"simba_keywords_pricevon_set_response"`
}



