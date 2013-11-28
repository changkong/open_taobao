// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wangwang

import (
	"github.com/yaofangou/open_taobao"
)





/* 增加关键词，只支持json返回 */
type WangwangAbstractAddwordRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 传入参数的字符集 */
func (r *WangwangAbstractAddwordRequest) SetCharset(value string) {
	r.SetValue("charset", value)
}

/* 关键词，长度大于0 */
func (r *WangwangAbstractAddwordRequest) SetWord(value string) {
	r.SetValue("word", value)
}


func (r *WangwangAbstractAddwordRequest) GetResponse(accessToken string) (*WangwangAbstractAddwordResponse, []byte, error) {
	var resp WangwangAbstractAddwordResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.abstract.addword", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangAbstractAddwordResponse struct {
	ErrorMsg string `json:"error_msg"`
	RetCode int `json:"ret_code"`
}

type WangwangAbstractAddwordResponseResult struct {
	Response *WangwangAbstractAddwordResponse `json:"wangwang_abstract_addword_response"`
}





/* 删除关键词，只支持json返回 */
type WangwangAbstractDeletewordRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 传入参数的字符集 */
func (r *WangwangAbstractDeletewordRequest) SetCharset(value string) {
	r.SetValue("charset", value)
}

/* 关键词，长度大于0 */
func (r *WangwangAbstractDeletewordRequest) SetWord(value string) {
	r.SetValue("word", value)
}


func (r *WangwangAbstractDeletewordRequest) GetResponse(accessToken string) (*WangwangAbstractDeletewordResponse, []byte, error) {
	var resp WangwangAbstractDeletewordResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.abstract.deleteword", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangAbstractDeletewordResponse struct {
	ErrorMsg string `json:"error_msg"`
	RetCode int `json:"ret_code"`
}

type WangwangAbstractDeletewordResponseResult struct {
	Response *WangwangAbstractDeletewordResponse `json:"wangwang_abstract_deleteword_response"`
}





/* 获取关键词列表，只支持json返回 */
type WangwangAbstractGetwordlistRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 传入参数的字符集 */
func (r *WangwangAbstractGetwordlistRequest) SetCharset(value string) {
	r.SetValue("charset", value)
}


func (r *WangwangAbstractGetwordlistRequest) GetResponse(accessToken string) (*WangwangAbstractGetwordlistResponse, []byte, error) {
	var resp WangwangAbstractGetwordlistResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.abstract.getwordlist", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangAbstractGetwordlistResponse struct {
	ErrorMsg string `json:"error_msg"`
	RetCode int `json:"ret_code"`
	WordLists []*WordList `json:"word_lists"`
}

type WangwangAbstractGetwordlistResponseResult struct {
	Response *WangwangAbstractGetwordlistResponse `json:"wangwang_abstract_getwordlist_response"`
}





/* 模糊查询服务初始化，只支持json返回 */
type WangwangAbstractInitializeRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 传入参数的字符集 */
func (r *WangwangAbstractInitializeRequest) SetCharset(value string) {
	r.SetValue("charset", value)
}


func (r *WangwangAbstractInitializeRequest) GetResponse(accessToken string) (*WangwangAbstractInitializeResponse, []byte, error) {
	var resp WangwangAbstractInitializeResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.abstract.initialize", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangAbstractInitializeResponse struct {
	ErrorMsg string `json:"error_msg"`
	RetCode int `json:"ret_code"`
}

type WangwangAbstractInitializeResponseResult struct {
	Response *WangwangAbstractInitializeResponse `json:"wangwang_abstract_initialize_response"`
}





/* 模糊聊天记录查询,仅支持json返回 */
type WangwangAbstractLogqueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 传入参数的字符集 */
func (r *WangwangAbstractLogqueryRequest) SetCharset(value string) {
	r.SetValue("charset", value)
}

/* 获取记录条数，默认值是1000 */
func (r *WangwangAbstractLogqueryRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 东八区时间 */
func (r *WangwangAbstractLogqueryRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 卖家id，有cntaobao前缀 */
func (r *WangwangAbstractLogqueryRequest) SetFromId(value string) {
	r.SetValue("from_id", value)
}

/* 设置了这个值，那么聊天记录就从这个点开始查询 */
func (r *WangwangAbstractLogqueryRequest) SetNextKey(value string) {
	r.SetValue("next_key", value)
}

/* 东八区时间 */
func (r *WangwangAbstractLogqueryRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}

/* 买家id，有cntaobao前缀 */
func (r *WangwangAbstractLogqueryRequest) SetToId(value string) {
	r.SetValue("to_id", value)
}


func (r *WangwangAbstractLogqueryRequest) GetResponse(accessToken string) (*WangwangAbstractLogqueryResponse, []byte, error) {
	var resp WangwangAbstractLogqueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.abstract.logquery", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangAbstractLogqueryResponse struct {
	ErrorMsg string `json:"error_msg"`
	FromId string `json:"from_id"`
	IsSliced int `json:"is_sliced"`
	MsgLists []*MsgList `json:"msg_lists"`
	NextKey string `json:"next_key"`
	RetCode int `json:"ret_code"`
	ToId string `json:"to_id"`
}

type WangwangAbstractLogqueryResponseResult struct {
	Response *WangwangAbstractLogqueryResponse `json:"wangwang_abstract_logquery_response"`
}





/* 根据客服ID和日期，获取该客服"当日接待的所有客户的平均等待时长"。  <br/> 
备注：  <br/> 
1、如果是操作者ID=被查者ID，返回被查者ID的"当日接待的所有客户的平均等待时长"。  <br/> 
2、如果操作者是组管理员，他可以查询他的组中的所有子帐号的"当日接待的所有客户的平均等待时长"。service_staff_id为所有子帐号，用 "," 隔开 <br/> 
3、如果操作者是主账户，他可以查询所有子帐号的"当日接待的所有客户的平均等待时长"。  service_staff_id为所有子帐号，用 "," 隔开<br/> 
 4、被查者ID可以是多个，用 "," 隔开，id数不能超过30。  <br/> 
 5、开始时间与结束时间之间的间隔不能超过7天  <br/> 
 6、不能查询90天以前的数据  <br/> 
 7、不能查询当天的记录 */
type WangwangEserviceAvgwaittimeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结束时间 */
func (r *WangwangEserviceAvgwaittimeGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 客服人员id：cntaobao+淘宝nick，例如cntaobaotest */
func (r *WangwangEserviceAvgwaittimeGetRequest) SetServiceStaffId(value string) {
	r.SetValue("service_staff_id", value)
}

/* 开始时间 */
func (r *WangwangEserviceAvgwaittimeGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}


func (r *WangwangEserviceAvgwaittimeGetRequest) GetResponse(accessToken string) (*WangwangEserviceAvgwaittimeGetResponse, []byte, error) {
	var resp WangwangEserviceAvgwaittimeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.eservice.avgwaittime.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangEserviceAvgwaittimeGetResponse struct {
	WaitingTimeListOnDays []*WaitingTimesOnDay `json:"waiting_time_list_on_days"`
}

type WangwangEserviceAvgwaittimeGetResponseResult struct {
	Response *WangwangEserviceAvgwaittimeGetResponse `json:"wangwang_eservice_avgwaittime_get_response"`
}





/* 获取聊天对象列表，只能获取最近1个月内的数据且查询时间段<=7天,只支持xml返回 。 */
type WangwangEserviceChatpeersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 字符集 */
func (r *WangwangEserviceChatpeersGetRequest) SetCharset(value string) {
	r.SetValue("charset", value)
}

/* 聊天用户ID：cntaobao+淘宝nick，例如cntaobaotest */
func (r *WangwangEserviceChatpeersGetRequest) SetChatId(value string) {
	r.SetValue("chat_id", value)
}

/* 查询结束日期。如2010-03-24，与起始日期跨度不能超过7天 */
func (r *WangwangEserviceChatpeersGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 查询起始日期。如2010-02-01，与当前日期间隔小于1个月。 */
func (r *WangwangEserviceChatpeersGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}


func (r *WangwangEserviceChatpeersGetRequest) GetResponse(accessToken string) (*WangwangEserviceChatpeersGetResponse, []byte, error) {
	var resp WangwangEserviceChatpeersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.eservice.chatpeers.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangEserviceChatpeersGetResponse struct {
	Chatpeers []*Chatpeer `json:"chatpeers"`
	Count int `json:"count"`
	Ret int `json:"ret"`
}

type WangwangEserviceChatpeersGetResponseResult struct {
	Response *WangwangEserviceChatpeersGetResponse `json:"wangwang_eservice_chatpeers_get_response"`
}





/* 根据用户id查询用户对应的评价详细情况， 
主账号id可以查询店铺内子账号的评价 
组管理员可以查询组内账号的评价 
非管理员的子账号可以查自己的评价 */
type WangwangEserviceEvalsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结束时间 */
func (r *WangwangEserviceEvalsGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 想要查询的账号列表 */
func (r *WangwangEserviceEvalsGetRequest) SetServiceStaffId(value string) {
	r.SetValue("service_staff_id", value)
}

/* 开始时间 */
func (r *WangwangEserviceEvalsGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}


func (r *WangwangEserviceEvalsGetRequest) GetResponse(accessToken string) (*WangwangEserviceEvalsGetResponse, []byte, error) {
	var resp WangwangEserviceEvalsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.eservice.evals.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangEserviceEvalsGetResponse struct {
	ResultCode int `json:"result_code"`
	ResultCount int `json:"result_count"`
	StaffEvalDetails []*EvalDetail `json:"staff_eval_details"`
}

type WangwangEserviceEvalsGetResponseResult struct {
	Response *WangwangEserviceEvalsGetResponse `json:"wangwang_eservice_evals_get_response"`
}





/* 根据操作者ID，返回被查者ID指定日期内每个帐号每日的"客服评价统计" 
 
备注：1、如果是操作者ID=被查者ID，返回被查者ID的"客服评价统计"。 
 
    2、如果操作者是组管理员，他可以查询他的组中的所有子帐号的"客服评价统计"。 
 
    3、如果操作者是主账户，他可以查询所有子帐号的"客服评价统计"。 
    4、被查者ID可以是多个，用 "," 隔开，id数不能超过30。 
    5、开始时间与结束时间之间的间隔不能超过7天 
    6、不能查询90天以前的数据 
    7、不能查询当天的记录 */
type WangwangEserviceEvaluationGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询结束日期 */
func (r *WangwangEserviceEvaluationGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 客服人员id：cntaobao+淘宝nick，例如cntaobaotest */
func (r *WangwangEserviceEvaluationGetRequest) SetServiceStaffId(value string) {
	r.SetValue("service_staff_id", value)
}

/* 查询开始日期 */
func (r *WangwangEserviceEvaluationGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}


func (r *WangwangEserviceEvaluationGetRequest) GetResponse(accessToken string) (*WangwangEserviceEvaluationGetResponse, []byte, error) {
	var resp WangwangEserviceEvaluationGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.eservice.evaluation.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangEserviceEvaluationGetResponse struct {
	StaffEvalStatOnDays []*StaffEvalStatOnDay `json:"staff_eval_stat_on_days"`
}

type WangwangEserviceEvaluationGetResponseResult struct {
	Response *WangwangEserviceEvaluationGetResponse `json:"wangwang_eservice_evaluation_get_response"`
}





/* 用某个组管理员账号查询，返回该组组名、和该组所有组成员ID（E客服的分流设置）。 
 
用旺旺主帐号查询，返回所有组的组名和该组所有组成员ID。 
 
返回的组成员ID可以是多个，用 "," 隔开。 
 
被查者ID只能传入一个。 
 
组成员中排名最靠前的ID是组管理员ID */
type WangwangEserviceGroupmemberGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主帐号ID：cntaobao+淘宝nick，例如cntaobaotest */
func (r *WangwangEserviceGroupmemberGetRequest) SetManagerId(value string) {
	r.SetValue("manager_id", value)
}


func (r *WangwangEserviceGroupmemberGetRequest) GetResponse(accessToken string) (*WangwangEserviceGroupmemberGetResponse, []byte, error) {
	var resp WangwangEserviceGroupmemberGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.eservice.groupmember.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangEserviceGroupmemberGetResponse struct {
	GroupMemberList []*GroupMember `json:"group_member_list"`
}

type WangwangEserviceGroupmemberGetResponseResult struct {
	Response *WangwangEserviceGroupmemberGetResponse `json:"wangwang_eservice_groupmember_get_response"`
}





/* 通过用户id查询用户自己或者子账户的登录日志： 
主账号可以查询自己和店铺子账户的登录日志（查询时需要输入子账号，多个用，隔开） 
组管理员可以查询自己和组内子账号的登录日志（查询时需要输入子账号，多个用，隔开） 
非组管理员的子账户只能查询自己的登录日志 
注：目前sdk只能支持到body内容的返回，body内数据结构的解析需要开发者自己写代码 */
type WangwangEserviceLoginlogsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询登录日志的结束时间，必须按示例的格式，否则会返回错误 */
func (r *WangwangEserviceLoginlogsGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 需要查询登录日志的账号列表 */
func (r *WangwangEserviceLoginlogsGetRequest) SetServiceStaffId(value string) {
	r.SetValue("service_staff_id", value)
}

/* 查询登录日志的开始日期，必须按示例的格式，否则会返回错误 */
func (r *WangwangEserviceLoginlogsGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}


func (r *WangwangEserviceLoginlogsGetRequest) GetResponse(accessToken string) (*WangwangEserviceLoginlogsGetResponse, []byte, error) {
	var resp WangwangEserviceLoginlogsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.eservice.loginlogs.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangEserviceLoginlogsGetResponse struct {
	Count int `json:"count"`
	Loginlogs []*LoginLog `json:"loginlogs"`
	UserId string `json:"user_id"`
}

type WangwangEserviceLoginlogsGetResponseResult struct {
	Response *WangwangEserviceLoginlogsGetResponse `json:"wangwang_eservice_loginlogs_get_response"`
}





/* 根据操作者ID，返回被查者ID指定日期内每个帐号每日的"未回复情况" 
 
备注：1、如果是操作者ID=被查者ID，返回被查者ID的"未回复情况"（未回复人数、未回复的ID）。 
 
    2、如果操作者是组管理员，他可以查询他的组中的所有子帐号的"未回复情况"。 
 
    3、如果操作者是主账户，他可以查询所有子帐号的"未回复情况"。 
    4、被查者ID可以是多个，用 "," 隔开，id数不能超过30。 
    5、开始时间与结束时间之间的间隔不能超过7天 
    6、不能查询90天以前的数据 
    7、不能查询当天的记录 */
type WangwangEserviceNoreplynumGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结束日期 */
func (r *WangwangEserviceNoreplynumGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 客服人员id：cntaobao+淘宝nick，例如cntaobaotest */
func (r *WangwangEserviceNoreplynumGetRequest) SetServiceStaffId(value string) {
	r.SetValue("service_staff_id", value)
}

/* 开始日期 */
func (r *WangwangEserviceNoreplynumGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}


func (r *WangwangEserviceNoreplynumGetRequest) GetResponse(accessToken string) (*WangwangEserviceNoreplynumGetResponse, []byte, error) {
	var resp WangwangEserviceNoreplynumGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.eservice.noreplynum.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangEserviceNoreplynumGetResponse struct {
	NonReplyStatOnDays []*NonReplyStatOnDay `json:"non_reply_stat_on_days"`
}

type WangwangEserviceNoreplynumGetResponseResult struct {
	Response *WangwangEserviceNoreplynumGetResponse `json:"wangwang_eservice_noreplynum_get_response"`
}





/* 描述：根据客服ID和日期，获取该客服"当日在线时长"。 
 
备注：1、如果是操作者ID=被查者ID，返回被查者ID的"当日在线时长"。 
 
    2、如果操作者是组管理员，他可以查询他的组中的所有子帐号的"当日在线时长"。 
 
    3、如果操作者是主账户，他可以查询所有子帐号的"当日在线时长"。 
    4、被查者ID可以是多个，用 "," 隔开，id数不能超过30。 
    5、日累计在线时长定义：当日该用户累计的旺旺在线时长 
 
    6、开始时间与结束时间之间的间隔不能超过7天 
    7、不能查询90天以前的数据 
    8、不能查询当天的记录 */
type WangwangEserviceOnlinetimeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结束日期 */
func (r *WangwangEserviceOnlinetimeGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 客服人员id：cntaobao+淘宝nick，例如cntaobaotest */
func (r *WangwangEserviceOnlinetimeGetRequest) SetServiceStaffId(value string) {
	r.SetValue("service_staff_id", value)
}

/* 开始日期 */
func (r *WangwangEserviceOnlinetimeGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}


func (r *WangwangEserviceOnlinetimeGetRequest) GetResponse(accessToken string) (*WangwangEserviceOnlinetimeGetResponse, []byte, error) {
	var resp WangwangEserviceOnlinetimeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.eservice.onlinetime.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangEserviceOnlinetimeGetResponse struct {
	OnlineTimesListOnDays []*OnlineTimesOnDay `json:"online_times_list_on_days"`
}

type WangwangEserviceOnlinetimeGetResponseResult struct {
	Response *WangwangEserviceOnlinetimeGetResponse `json:"wangwang_eservice_onlinetime_get_response"`
}





/* 根据操作者ID，返回被查者ID指定时间段内每个帐号的"已接待人数"<br/> 
备注：<br/> 
1、如果是操作者ID=被查者ID，返回被查者ID的"已接待人数"。<br/> 
2、如果操作者是组管理员，他可以查询他的组中的所有子帐号的"已接待人数"。<br/> 
3、如果操作者是主账户，他可以查询所有子帐号的"已接待人数"。<br/> 
（注意：这里说的是授权是主帐号，可以查询所有子帐号的数据，具体要查询哪些子账号的数据，需要在service_staff_id具体指定，而不是service_staff_id直接输入主帐号）<br/> 
 4、被查者ID可以是多个，用 "," 隔开，id数不能超过30。<br/> 
 5、规则：某客服在1天内和同一个客户交流了多次，已回复人数算1。<br/> 
6、"已接待人数"定义：买家、卖家彼此发过至少1条消息 ，不论谁先发都可以。<br/> 
  7、被查者ID可以是多个，用 "," 隔开，id数不能超过30。<br/> 
  8、开始时间与结束时间之间的间隔不能超过7天<br/> 
  9、不能查询90天以前的数据<br/> 
  10、不能查询当天的记录<br/> 
   11、查询日期精确到日<br/> */
type WangwangEserviceReceivenumGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询接待人数的结束日期 时间精确到日 时分秒可以直接传00:00:00 */
func (r *WangwangEserviceReceivenumGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 客服人员id：cntaobao+淘宝nick，例如cntaobaotest */
func (r *WangwangEserviceReceivenumGetRequest) SetServiceStaffId(value string) {
	r.SetValue("service_staff_id", value)
}

/* 查询接待人数的开始日期 时间精确到日 时分秒可直接传 00:00:00 */
func (r *WangwangEserviceReceivenumGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}


func (r *WangwangEserviceReceivenumGetRequest) GetResponse(accessToken string) (*WangwangEserviceReceivenumGetResponse, []byte, error) {
	var resp WangwangEserviceReceivenumGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.eservice.receivenum.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangEserviceReceivenumGetResponse struct {
	ReplyStatListOnDays []*ReplyStatOnDay `json:"reply_stat_list_on_days"`
}

type WangwangEserviceReceivenumGetResponseResult struct {
	Response *WangwangEserviceReceivenumGetResponse `json:"wangwang_eservice_receivenum_get_response"`
}





/* 获取当前登录用户自己的店铺内的分流权重设置，只支持xml 返回。 */
type WangwangEserviceStreamweigthsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *WangwangEserviceStreamweigthsGetRequest) GetResponse(accessToken string) (*WangwangEserviceStreamweigthsGetResponse, []byte, error) {
	var resp WangwangEserviceStreamweigthsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wangwang.eservice.streamweigths.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WangwangEserviceStreamweigthsGetResponse struct {
	ResultCode int `json:"result_code"`
	ResultCount int `json:"result_count"`
	StaffStreamWeights []*StreamWeight `json:"staff_stream_weights"`
	TotalWeight int `json:"total_weight"`
}

type WangwangEserviceStreamweigthsGetResponseResult struct {
	Response *WangwangEserviceStreamweigthsGetResponse `json:"wangwang_eservice_streamweigths_get_response"`
}



