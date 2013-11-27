// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tools

import (
	"github.com/changkong/open_taobao"
)





/* 获取ISV发起请求服务器IP */
type AppipGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *AppipGetRequest) GetResponse(accessToken string) (*AppipGetResponse, []byte, error) {
	var resp AppipGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.appip.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AppipGetResponse struct {
	Ip string `json:"ip"`
}

type AppipGetResponseResult struct {
	Response *AppipGetResponse `json:"appip_get_response"`
}





/* 添加用户反馈信息 */
type FeedbackAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 具体反馈的信息，一个json构成的字符串 */
func (r *FeedbackAddRequest) SetInfo(value string) {
	r.SetValue("info", value)
}

/* 反馈信息的类型，例如是同步服务的或者其他系统的 */
func (r *FeedbackAddRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *FeedbackAddRequest) GetResponse(accessToken string) (*FeedbackAddResponse, []byte, error) {
	var resp FeedbackAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.feedback.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FeedbackAddResponse struct {
	Result string `json:"result"`
}

type FeedbackAddResponseResult struct {
	Response *FeedbackAddResponse `json:"feedback_add_response"`
}





/* 对输入的文本信息进行禁忌关键词匹配，返回匹配的结果 */
type KfcKeywordSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 应用点，分为一级应用点、二级应用点。其中一级应用点通常是指某一个系统或产品，比如淘宝的商品应用（taobao_auction）；二级应用点，是指一级应用点下的具体的分类，比如商品标题(title)、商品描述(content)。不同的二级应用可以设置不同关键词。 
 
这里的apply参数是由一级应用点与二级应用点合起来的字符（一级应用点+"."+二级应用点），如taobao_auction.title。 
 
 
通常apply参数是不需要传递的。如有特殊需求（比如特殊的过滤需求，需要自己维护一套自己词库），需传递此参数。 */
func (r *KfcKeywordSearchRequest) SetApply(value string) {
	r.SetValue("apply", value)
}

/* 需要过滤的文本信息 */
func (r *KfcKeywordSearchRequest) SetContent(value string) {
	r.SetValue("content", value)
}

/* 发布信息的淘宝会员名，可以不传 */
func (r *KfcKeywordSearchRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *KfcKeywordSearchRequest) GetResponse(accessToken string) (*KfcKeywordSearchResponse, []byte, error) {
	var resp KfcKeywordSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.kfc.keyword.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type KfcKeywordSearchResponse struct {
	KfcSearchResult *KfcSearchResult `json:"kfc_search_result"`
}

type KfcKeywordSearchResponseResult struct {
	Response *KfcKeywordSearchResponse `json:"kfc_keyword_search_response"`
}





/* 获取淘宝系统当前时间 */
type TimeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *TimeGetRequest) GetResponse(accessToken string) (*TimeGetResponse, []byte, error) {
	var resp TimeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.time.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TimeGetResponse struct {
	Time string `json:"time"`
}

type TimeGetResponseResult struct {
	Response *TimeGetResponse `json:"time_get_response"`
}





/* 获取异步任务结果。 
<br/>异步API使用方法，请查看：<a href="http://open.taobao.com/doc/detail.htm?id=30">异步API使用说明</a> 
<br/>1. 只能获取AppKey自身创建的异步任务结果 
<br/>2. 如果任务状态为done，则返回任务结果（如taobao.topats.trades.fullinfo.get）或下载地址（如taobao.topats.trades.sold.get） 
<br/>3. 任务结果下载地址只能使用一次，需要重复下载可重新调用此接口获取下载地址 
<br/>4. 任务结果中的check_code字段为待下载文件的md5 sum，可通过此校验码验证文件下载的完整性 */
type TopatsResultGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 任务id号，创建任务时返回的task_id */
func (r *TopatsResultGetRequest) SetTaskId(value string) {
	r.SetValue("task_id", value)
}


func (r *TopatsResultGetRequest) GetResponse(accessToken string) (*TopatsResultGetResponse, []byte, error) {
	var resp TopatsResultGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.topats.result.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TopatsResultGetResponse struct {
	Task *Task `json:"task"`
}

type TopatsResultGetResponseResult struct {
	Response *TopatsResultGetResponse `json:"topats_result_get_response"`
}





/* 可用于取消已经创建的ATS任务。</br> 
条件限制：1)一次只可以取消一个任务</br> 
         2）只能取消自己创建的任务</br> */
type TopatsTaskDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需要取消的任务ID */
func (r *TopatsTaskDeleteRequest) SetTaskId(value string) {
	r.SetValue("task_id", value)
}


func (r *TopatsTaskDeleteRequest) GetResponse(accessToken string) (*TopatsTaskDeleteResponse, []byte, error) {
	var resp TopatsTaskDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.topats.task.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TopatsTaskDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TopatsTaskDeleteResponseResult struct {
	Response *TopatsTaskDeleteResponse `json:"topats_task_delete_response"`
}



