// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package jipiao

import (
	"github.com/changkong/open_taobao"
)





/* 产品批量添加,传入文件大小限制在1M(一般1w条记录不会超过1M),每五分钟只允许调用一次 */
type JipiaoPoliciesAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* (ZIP压缩UTF-8编码JSON)，压缩前格式为:[{数据结构BatchPolicy的json格式},{数据结构BatchPolicy的json格式},...] 示例： [{ "attributes": "rfc=1;ipt=1;fdtod=0000;ldtod=2359", "source": null, "airline": "CZ", "arrAirports": "CSX,CTU,CAN", "autoHkFlag": true, "autoTicketFlag": true, "cabinRules": "[{\"matcher\":{\"mode\":\"ALL\",\"list\":[\"Y\"]},\"priceStrategy\":{\"mode\":\"DISCOUNT\",\"modeBaseValue\":null,\"retention\":5000,\"rebase\":-5},\"backMatcher\":null}]", "changeRule": null, "dayOfWeeks": "1234", "depAirports": "PEK", "ei": "ei", "excludeDate": null, "firstSaleAdvanceDay": null, "flags": null, "flightInfo": "+CA1234,CZ2345", "lastSaleAdvanceDay": null, "memo": "memoUpdate", "officeId": "RRR003", "outProductId": "46f9b762-ea50-4c71-877b-45fa2936277e", "policyType": 8, "refundRule": null, "reissueRule": null, "saleEndDate": "2013-06-19 00:00:00", "saleStartDate": "2013-06-09 00:00:00", "seatInfo": null, "shareSupport": false, "specialRule": null, "travelEndDate": "2013-06-19 00:00:00", "travelStartDate": "2013-06-09 00:00:00" } ] */
func (r *JipiaoPoliciesAddRequest) SetCompressedPolicies(value string) {
	r.SetValue("compressed_policies", value)
}


func (r *JipiaoPoliciesAddRequest) GetResponse(accessToken string) (*JipiaoPoliciesAddResponse, []byte, error) {
	var resp JipiaoPoliciesAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.policies.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoPoliciesAddResponse struct {
	InvokeId string `json:"invoke_id"`
	IsSuccess bool `json:"is_success"`
}

type JipiaoPoliciesAddResponseResult struct {
	Response *JipiaoPoliciesAddResponse `json:"jipiao_policies_add_response"`
}





/* 政策全量增加,会删除政策中字段source值为TOP的记录，并且把新增的记录插入，传入文件大小限制在5M(一般10w条记录不会超过5M),每半小时只允许调用一次 */
type JipiaoPoliciesFulladdRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* (ZIP压缩UTF-8编码JSON)，压缩前格式为:[{数据结构BatchPolicy的json格式},{数据结构BatchPolicy的json格式},...] 示例： 
[{ 
        "attributes": "rfc=1;ipt=1;fdtod=0000;ldtod=2359", 
        "source": null, 
        "airline": "CZ", 
        "arrAirports": "CSX,CTU,CAN", 
        "autoHkFlag": true, 
        "autoTicketFlag": true, 
        "cabinRules": "[{\"matcher\":{\"mode\":\"ALL\",\"list\":[\"Y\"]},\"priceStrategy\":{\"mode\":\"DISCOUNT\",\"modeBaseValue\":null,\"retention\":5000,\"rebase\":-5},\"backMatcher\":null}]", 
        "changeRule": null, 
        "dayOfWeeks": "1234", 
        "depAirports": "PEK", 
        "ei": "ei", 
        "excludeDate": null, 
        "firstSaleAdvanceDay": null, 
        "flags": null, 
        "flightInfo": "+CA1234,CZ2345", 
        "lastSaleAdvanceDay": null, 
        "memo": "memoUpdate", 
        "officeId": "RRR003", 
        "outProductId": "46f9b762-ea50-4c71-877b-45fa2936277e", 
        "policyType": 8, 
        "refundRule": null, 
        "reissueRule": null, 
        "saleEndDate": "2013-06-19 00:00:00", 
        "saleStartDate": "2013-06-09 00:00:00", 
        "seatInfo": null, 
        "shareSupport": false, 
        "specialRule": null, 
        "travelEndDate": "2013-06-19 00:00:00", 
        "travelStartDate": "2013-06-09 00:00:00" 
    } 
] */
func (r *JipiaoPoliciesFulladdRequest) SetCompressedPolicies(value string) {
	r.SetValue("compressed_policies", value)
}


func (r *JipiaoPoliciesFulladdRequest) GetResponse(accessToken string) (*JipiaoPoliciesFulladdResponse, []byte, error) {
	var resp JipiaoPoliciesFulladdResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.policies.fulladd", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoPoliciesFulladdResponse struct {
	Empty *BatchPolicy `json:"empty"`
	InvokeId string `json:"invoke_id"`
	IsSuccess bool `json:"is_success"`
}

type JipiaoPoliciesFulladdResponseResult struct {
	Response *JipiaoPoliciesFulladdResponse `json:"jipiao_policies_fulladd_response"`
}





/* 根据条件大批量更新产品，目前只支持删除,每十分钟只允许调用一次,更新记录数不能超过10W个 */
type JipiaoPoliciesstatusUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 航空公司二字码 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetAirline(value string) {
	r.SetValue("airline", value)
}

/* 到达机场三字码,此项必需与出发机场同时为空或不为空 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetArrAirport(value string) {
	r.SetValue("arr_airport", value)
}

/* 出发机场三字码,此项必需与到达机场同时为空或不为空 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetDepAirport(value string) {
	r.SetValue("dep_airport", value)
}

/* 外部产品id集,最多支持1000个,后续调大,其中的out_product_id含有空格将不会处理 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetOutProductIds(value string) {
	r.SetValue("out_product_ids", value)
}

/* 产品id集,最多支持1000个，后续调大，其中单个的policy_id含有留空格或不是数字将会忽略不处理 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetPolicyIds(value string) {
	r.SetValue("policy_ids", value)
}

/* 发布日期 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetPublishDate(value string) {
	r.SetValue("publish_date", value)
}

/* 发布来源, 通过接口taobao.jipiao.policy.process添加的政策会自动加上source为TOP,代理商后台页面录入的source为PC,excel上传的source为UPLOAD,通过接口taobao.jipiao.policies.fulladd,taobao.jipiao.policies.add的自定义source也可以 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 0：按policy_ids更新；1：按out_product_ids更新；2:按其它条件更新 */
func (r *JipiaoPoliciesstatusUpdateRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *JipiaoPoliciesstatusUpdateRequest) GetResponse(accessToken string) (*JipiaoPoliciesstatusUpdateResponse, []byte, error) {
	var resp JipiaoPoliciesstatusUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.policiesstatus.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoPoliciesstatusUpdateResponse struct {
	InvokeId string `json:"invoke_id"`
	IsSuccess bool `json:"is_success"`
}

type JipiaoPoliciesstatusUpdateResponseResult struct {
	Response *JipiaoPoliciesstatusUpdateResponse `json:"jipiao_policiesstatus_update_response"`
}





/* 根据政策id或者政策外部id查询政策信息 */
type JipiaoPolicyGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* type外0，表示机票政策id；type为1，表示机票政策out_product_id */
func (r *JipiaoPolicyGetRequest) SetPolicyId(value string) {
	r.SetValue("policy_id", value)
}

/* 0，表示按政策id进行查询；1，表示按政策外部id进行查询 */
func (r *JipiaoPolicyGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *JipiaoPolicyGetRequest) GetResponse(accessToken string) (*JipiaoPolicyGetResponse, []byte, error) {
	var resp JipiaoPolicyGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.policy.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoPolicyGetResponse struct {
	Policy *Policy `json:"policy"`
}

type JipiaoPolicyGetResponseResult struct {
	Response *JipiaoPolicyGetResponse `json:"jipiao_policy_get_response"`
}





/* 添加/修改一条机票政策 */
type JipiaoPolicyProcessRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 政策所支持的航空公司二字码 */
func (r *JipiaoPolicyProcessRequest) SetAirline(value string) {
	r.SetValue("airline", value)
}

/* 政策支持的到达机场列表，逗号分隔的机场三字码， 
注：  
1.让利(8)政策，当到达支持多个(最多25)时，出发只能支持一个；  
2.特价(6)/特殊(10)政策，出发/到达城市只支持一个 */
func (r *JipiaoPolicyProcessRequest) SetArrAirports(value string) {
	r.SetValue("arr_airports", value)
}

/* 扩展字段： 
    "rfc" 对应值 1:不退不改不签,2:收费退改签（退、改、签中任意一个收费即为收费退改签）3:免费退改签 
    "ipt" 对应值 1:等额行程单 2:不提供发票5:等额“行程单+发票”（对应旧的2） 6:等额发票（对应旧的1）,例如：rfc=1;ipt=1 
    "fdtod"对应值起飞时间“0000”为0时0分  
    "ldtod"对应值：到达时间“2359”为23时59分 */
func (r *JipiaoPolicyProcessRequest) SetAttributes(value string) {
	r.SetValue("attributes", value)
}

/* 政策是否支持自动HK，默认为不支持 */
func (r *JipiaoPolicyProcessRequest) SetAutoHkFlag(value string) {
	r.SetValue("auto_hk_flag", value)
}

/* 政策是否支持自动出票，默认不支持， 
注：自动出票时，必须同时自动HK */
func (r *JipiaoPolicyProcessRequest) SetAutoTicketFlag(value string) {
	r.SetValue("auto_ticket_flag", value)
}

/* @1: [{"matcher":{"mode":"ALL","list":["D"]},"priceStrategy":{"mode":
"TICKET_PRICE","modeBaseValue":500,"retention":500,"rebase":-5}}]

@2: [{"matcher"
:{"mode":"ALL","list":["D"]},"priceStrategy":{"mode":"DISCOUNT"
,"modeBaseValue"
:null,"retention":300,"rebase":-5}}]

@3: [{"matcher":{"mode":"ALL","list":["D"]},"priceStrategy":{"mode":"Y_DISCOUNT","modeBaseValue":75,"retention":500,"rebase":-5}}]

规则说明： cabin_rules分二个字段：matcher:适应舱位（又包含两个字段：mode：匹配模式枚举(填INCLUDE），list：适应舱位列表），priceStrategy:价格策略（mode:价格模式（让利产品：DISCOUNT，特价特殊产品有四种模式：票面价用TICKET_PRICE，Y舱折扣用Y_DISCOUNT，C舱折扣用C_DISCOUNT，F舱折扣用F_DISCOUNT）；modeBaseValue:当价格模式mode为DISCOUNT，modeBaseValue不填；当价格模式mode为TICKET_PRICE，modeBaseValue填票面价；当价格模式mode为Y_DISCOUNT/C_DISCOUNT/F_DISCOUNT，modeBaseValue填对应的折扣；retention为返点的百分比；rebase为留钱）
注意：特殊政策不需要舱位时需要传入22；retention、rebase传入小数时，会取整数部分 */
func (r *JipiaoPolicyProcessRequest) SetCabinRules(value string) {
	r.SetValue("cabin_rules", value)
}

/* 改签规则 */
func (r *JipiaoPolicyProcessRequest) SetChangeRule(value string) {
	r.SetValue("change_rule", value)
}

/* 政策适用的星期几，1-7分别表示周一到周日 
注：特殊政策不能填写该字段；其它政策填写时， 
包含全部时需要设置为1234567 */
func (r *JipiaoPolicyProcessRequest) SetDayOfWeeks(value string) {
	r.SetValue("day_of_weeks", value)
}

/* 政策支持的出发机场列表，逗号分隔的机场三字码， 
注： 
1.让利(8)政策，当出发支持多个(最多25)时，到达只能支持一个； 
2.特价(6)/特殊(10)政策，出发/到达城市只支持一个 */
func (r *JipiaoPolicyProcessRequest) SetDepAirports(value string) {
	r.SetValue("dep_airports", value)
}

/* ei项，自动HK，自动出票设定时必需 */
func (r *JipiaoPolicyProcessRequest) SetEi(value string) {
	r.SetValue("ei", value)
}

/* 政策旅行有效日期中排除指定日期，设定日期将不在搜索结果页面展现 
注：最多排除20天，特殊政策无此设置 */
func (r *JipiaoPolicyProcessRequest) SetExcludeDate(value string) {
	r.SetValue("exclude_date", value)
}

/* 政策销售日期最早提前天数，默认-1表示无效 */
func (r *JipiaoPolicyProcessRequest) SetFirstSaleAdvanceDay(value string) {
	r.SetValue("first_sale_advance_day", value)
}

/* flags是二进制标志 
从右到左数，第1个位表示：不PAT自动HK  
第2个位表示：儿童可与成人同价 
比如：“儿童可与成人同价”=true ,“不PAT自动HK”=false,则表示成二进制字符串"10",换算成十进制flags=2 */
func (r *JipiaoPolicyProcessRequest) SetFlags(value string) {
	r.SetValue("flags", value)
}

/* 包含/排除的航班号，注意格式： 
+CA1234,CZ3166，表示包含列表 
-CA1234,CZ3166，表示排除列表 
默认包含所有航班； 
不支持同时包含和排除 */
func (r *JipiaoPolicyProcessRequest) SetFlightInfo(value string) {
	r.SetValue("flight_info", value)
}

/* 政策销售日期最晚提前天数，默认-1表示无效 */
func (r *JipiaoPolicyProcessRequest) SetLastSaleAdvanceDay(value string) {
	r.SetValue("last_sale_advance_day", value)
}

/* 代理商对政策的备注信息 */
func (r *JipiaoPolicyProcessRequest) SetMemo(value string) {
	r.SetValue("memo", value)
}

/* 政策设置为支持自动HK，自动出票时必需 */
func (r *JipiaoPolicyProcessRequest) SetOfficeId(value string) {
	r.SetValue("office_id", value)
}

/* 政策的外部id，用于关联代理商自身维护的政策id，提供外部产品id时，可以在查询和修改时作为唯一条件使用 */
func (r *JipiaoPolicyProcessRequest) SetOutProductId(value string) {
	r.SetValue("out_product_id", value)
}

/* 0，type为0时，不需要设置； 
1，type为1时，表示policy_id为政策id 
2，type为2时，表示policy_id为政策out_product_id */
func (r *JipiaoPolicyProcessRequest) SetPolicyId(value string) {
	r.SetValue("policy_id", value)
}

/* 政策类型：6，特价政策；8，让利政策；10，特殊政策 */
func (r *JipiaoPolicyProcessRequest) SetPolicyType(value string) {
	r.SetValue("policy_type", value)
}

/* 退票规则 */
func (r *JipiaoPolicyProcessRequest) SetRefundRule(value string) {
	r.SetValue("refund_rule", value)
}

/* 签转规则 */
func (r *JipiaoPolicyProcessRequest) SetReissueRule(value string) {
	r.SetValue("reissue_rule", value)
}

/* 政策销售有效日期的截止日期(注意，格式如：20120-03-10 10:30:31 等价2012-03-11) */
func (r *JipiaoPolicyProcessRequest) SetSaleEndDate(value string) {
	r.SetValue("sale_end_date", value)
}

/* 政策销售有效日期的开始日期(注意，格式如：20120-03-10 10:30:31 等价2012-03-11) */
func (r *JipiaoPolicyProcessRequest) SetSaleStartDate(value string) {
	r.SetValue("sale_start_date", value)
}

/* 政策类型为10时，用于设置政策的每天的库存信息 */
func (r *JipiaoPolicyProcessRequest) SetSeatInfo(value string) {
	r.SetValue("seat_info", value)
}

/* 政策是否支持共享航班，默认为不支持；设置不支持情况下，当航班的实际承运方不是当前航空公司，则搜索结果页不能展现 */
func (r *JipiaoPolicyProcessRequest) SetShareSupport(value string) {
	r.SetValue("share_support", value)
}

/* 特殊说明,此字段不再使用,填写内容将会报相应的填写错误 */
func (r *JipiaoPolicyProcessRequest) SetSpecialRule(value string) {
	r.SetValue("special_rule", value)
}

/* 政策旅行有效日期的结束日期(注意，格式如：20120-03-10 10:30:31 等价2012-03-11) */
func (r *JipiaoPolicyProcessRequest) SetTravelEndDate(value string) {
	r.SetValue("travel_end_date", value)
}

/* 政策旅行有效日期的开始日期(注意，格式，如：2012-03-10 10:30:31 等价2012-03-10) */
func (r *JipiaoPolicyProcessRequest) SetTravelStartDate(value string) {
	r.SetValue("travel_start_date", value)
}

/* 0，表示添加政策； 
1，表示按id修改政策； 
2，表示按out_product_id修改政策 */
func (r *JipiaoPolicyProcessRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *JipiaoPolicyProcessRequest) GetResponse(accessToken string) (*JipiaoPolicyProcessResponse, []byte, error) {
	var resp JipiaoPolicyProcessResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.policy.process", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoPolicyProcessResponse struct {
	IsSuccess bool `json:"is_success"`
	PolicyId int `json:"policy_id"`
}

type JipiaoPolicyProcessResponseResult struct {
	Response *JipiaoPolicyProcessResponse `json:"jipiao_policy_process_response"`
}





/* 根据外部产品id，进行政策状态更新，挂起/解挂/删除 */
type JipiaoPolicystatusUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* type为0，表示机票政策id；type为1，表示机票政策out_product_id；最大支持政策数100，注意不要如果不要超出字符串的长度限制，超出的话，请调小批量的个数 */
func (r *JipiaoPolicystatusUpdateRequest) SetPolicyId(value string) {
	r.SetValue("policy_id", value)
}

/* 政策的状态: 0，挂起；1，解挂；2，删除 */
func (r *JipiaoPolicystatusUpdateRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 0，表示按政策id进行查询；1，表示按政策外部id进行查询 */
func (r *JipiaoPolicystatusUpdateRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *JipiaoPolicystatusUpdateRequest) GetResponse(accessToken string) (*JipiaoPolicystatusUpdateResponse, []byte, error) {
	var resp JipiaoPolicystatusUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.jipiao.policystatus.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type JipiaoPolicystatusUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type JipiaoPolicystatusUpdateResponseResult struct {
	Response *JipiaoPolicystatusUpdateResponse `json:"jipiao_policystatus_update_response"`
}





/* 国内机票代理商行程单信息回填 */
type TripJipiaoAgentItinerarySendRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流公司代码CODE，如不清楚，请找运营提供 */
func (r *TripJipiaoAgentItinerarySendRequest) SetCompanyCode(value string) {
	r.SetValue("company_code", value)
}

/* 邮寄单号，长度不能超过32字节 */
func (r *TripJipiaoAgentItinerarySendRequest) SetExpressCode(value string) {
	r.SetValue("express_code", value)
}

/* 淘宝系统行程单唯一键 */
func (r *TripJipiaoAgentItinerarySendRequest) SetItineraryId(value string) {
	r.SetValue("itinerary_id", value)
}

/* 行程单号 */
func (r *TripJipiaoAgentItinerarySendRequest) SetItineraryNo(value string) {
	r.SetValue("itinerary_no", value)
}

/* 邮寄日期，格式yyyy-mm-dd */
func (r *TripJipiaoAgentItinerarySendRequest) SetSendDate(value string) {
	r.SetValue("send_date", value)
}


func (r *TripJipiaoAgentItinerarySendRequest) GetResponse(accessToken string) (*TripJipiaoAgentItinerarySendResponse, []byte, error) {
	var resp TripJipiaoAgentItinerarySendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.itinerary.send", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentItinerarySendResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentItinerarySendResponseResult struct {
	Response *TripJipiaoAgentItinerarySendResponse `json:"trip_jipiao_agent_itinerary_send_response"`
}





/* 国内机票代理商确认订单接口 */
type TripJipiaoAgentOrderConfirmRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 国内机票订单id */
func (r *TripJipiaoAgentOrderConfirmRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* hk（占座）时需要的信息列表，元素结构：乘机人姓名;pnr (以分号进行分隔) */
func (r *TripJipiaoAgentOrderConfirmRequest) SetPnrInfo(value string) {
	r.SetValue("pnr_info", value)
}


func (r *TripJipiaoAgentOrderConfirmRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderConfirmResponse, []byte, error) {
	var resp TripJipiaoAgentOrderConfirmResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.confirm", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderConfirmResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentOrderConfirmResponseResult struct {
	Response *TripJipiaoAgentOrderConfirmResponse `json:"trip_jipiao_agent_order_confirm_response"`
}





/* 国内机票代理商失败订单接口 */
type TripJipiaoAgentOrderFailRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 失败类型为0，说明备注原因 */
func (r *TripJipiaoAgentOrderFailRequest) SetFailMemo(value string) {
	r.SetValue("fail_memo", value)
}

/* 失败原因：1．客户要求失败订单；2．此舱位已售完（经济舱或特舱）；3．剩余座位少于用户购买数量；4．特价管理里录入的特价票已经售完；5．假舱（请及时通过旺旺或者电话反馈给淘宝工作人员）；0．其它（请在备注中说明原因） */
func (r *TripJipiaoAgentOrderFailRequest) SetFailType(value string) {
	r.SetValue("fail_type", value)
}

/* 国内机票订单id */
func (r *TripJipiaoAgentOrderFailRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}


func (r *TripJipiaoAgentOrderFailRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderFailResponse, []byte, error) {
	var resp TripJipiaoAgentOrderFailResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.fail", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderFailResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentOrderFailResponseResult struct {
	Response *TripJipiaoAgentOrderFailResponse `json:"trip_jipiao_agent_order_fail_response"`
}





/* 根据淘宝机票政策id搜索订单 */
type TripJipiaoAgentOrderFindRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 创建订单时间范围的开始时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认开始时间为当前时间往前推三天 （具体天数可能调整） */
func (r *TripJipiaoAgentOrderFindRequest) SetBeginTime(value string) {
	r.SetValue("begin_time", value)
}

/* 创建订单时间范围的结束时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认为当前时间 （具体天数可能调整） */
func (r *TripJipiaoAgentOrderFindRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 页码，默认第一页；注意：页大小固定，搜索结果中返回页大小pageSize，和是否包含下一页hasNext */
func (r *TripJipiaoAgentOrderFindRequest) SetPage(value string) {
	r.SetValue("page", value)
}

/* 淘宝机票政策id */
func (r *TripJipiaoAgentOrderFindRequest) SetPolicyId(value string) {
	r.SetValue("policy_id", value)
}


func (r *TripJipiaoAgentOrderFindRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderFindResponse, []byte, error) {
	var resp TripJipiaoAgentOrderFindResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.find", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderFindResponse struct {
	SearchResult *SearchOrderResult `json:"search_result"`
}

type TripJipiaoAgentOrderFindResponseResult struct {
	Response *TripJipiaoAgentOrderFindResponse `json:"trip_jipiao_agent_order_find_response"`
}





/* 根据淘宝系统订单号获取订单详情信息 */
type TripJipiaoAgentOrderGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 淘宝政策id列表，当前支持列表长度为1，即当前只支持单个订单详情查询 */
func (r *TripJipiaoAgentOrderGetRequest) SetOrderIds(value string) {
	r.SetValue("order_ids", value)
}


func (r *TripJipiaoAgentOrderGetRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderGetResponse, []byte, error) {
	var resp TripJipiaoAgentOrderGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderGetResponse struct {
	Orders []*AtOrder `json:"orders"`
}

type TripJipiaoAgentOrderGetResponseResult struct {
	Response *TripJipiaoAgentOrderGetResponse `json:"trip_jipiao_agent_order_get_response"`
}





/* 国内机票代理商手工hk订单（未付款前，手工填写pnr） */
type TripJipiaoAgentOrderHkRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 国内机票订单id */
func (r *TripJipiaoAgentOrderHkRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* hk（占座）时需要的信息信息列表，元素结构：乘机人姓名;pnr (以分号进行分隔). */
func (r *TripJipiaoAgentOrderHkRequest) SetPnrInfo(value string) {
	r.SetValue("pnr_info", value)
}


func (r *TripJipiaoAgentOrderHkRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderHkResponse, []byte, error) {
	var resp TripJipiaoAgentOrderHkResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.hk", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderHkResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentOrderHkResponseResult struct {
	Response *TripJipiaoAgentOrderHkResponse `json:"trip_jipiao_agent_order_hk_response"`
}





/* 根据条件查询淘宝订单id列表 */
type TripJipiaoAgentOrderSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 创建订单时间范围的开始时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认开始时间为当前时间往前推三天 （具体天数可能调整） */
func (r *TripJipiaoAgentOrderSearchRequest) SetBeginTime(value string) {
	r.SetValue("begin_time", value)
}

/* 创建订单时间范围的结束时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认为当前时间 （具体天数可能调整） */
func (r *TripJipiaoAgentOrderSearchRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 是否需要行程单，true表示需要行程单；false表示不许要 （必须设置，且不能为null） */
func (r *TripJipiaoAgentOrderSearchRequest) SetHasItinerary(value string) {
	r.SetValue("has_itinerary", value)
}

/* 页码，默认第一页；注意：页大小固定，搜索结果中返回页大小pageSize，和是否包含下一页hasNext */
func (r *TripJipiaoAgentOrderSearchRequest) SetPage(value string) {
	r.SetValue("page", value)
}

/* 订单状态，默认为空，查询所有状态的订单 */
func (r *TripJipiaoAgentOrderSearchRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 航程类型： 0.单程和普通往返；2.多程（暂时没有使用）；3.特价往返 */
func (r *TripJipiaoAgentOrderSearchRequest) SetTripType(value string) {
	r.SetValue("trip_type", value)
}


func (r *TripJipiaoAgentOrderSearchRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderSearchResponse, []byte, error) {
	var resp TripJipiaoAgentOrderSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderSearchResponse struct {
	SearchResult *SearchOrderResult `json:"search_result"`
}

type TripJipiaoAgentOrderSearchResponseResult struct {
	Response *TripJipiaoAgentOrderSearchResponse `json:"trip_jipiao_agent_order_search_response"`
}





/* 国内机票订单接口，确认特殊产品能否支付 */
type TripJipiaoAgentOrderSpecialConfirmRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 能否支付 */
func (r *TripJipiaoAgentOrderSpecialConfirmRequest) SetCanPay(value string) {
	r.SetValue("can_pay", value)
}

/* can_pay=false,fail_type=0时，必需提供失败原因 */
func (r *TripJipiaoAgentOrderSpecialConfirmRequest) SetFailMemo(value string) {
	r.SetValue("fail_memo", value)
}

/* can_pay=false时，必需提供失败原因；失败原因：11,座位已售完;12,座位申请不成功;13,航班价格变动;14,买家要求失败订单;0,其它(必须在备注中说明原因) */
func (r *TripJipiaoAgentOrderSpecialConfirmRequest) SetFailType(value string) {
	r.SetValue("fail_type", value)
}

/* 国内机票订单id */
func (r *TripJipiaoAgentOrderSpecialConfirmRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* can_pay=true，必需提供最迟支付时间 */
func (r *TripJipiaoAgentOrderSpecialConfirmRequest) SetPayLatestTime(value string) {
	r.SetValue("pay_latest_time", value)
}


func (r *TripJipiaoAgentOrderSpecialConfirmRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderSpecialConfirmResponse, []byte, error) {
	var resp TripJipiaoAgentOrderSpecialConfirmResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.special.confirm", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderSpecialConfirmResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentOrderSpecialConfirmResponseResult struct {
	Response *TripJipiaoAgentOrderSpecialConfirmResponse `json:"trip_jipiao_agent_order_special_confirm_response"`
}





/* 淘宝机票代理商成功/解冻订单 */
type TripJipiaoAgentOrderSuccessRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 淘宝系统订单id */
func (r *TripJipiaoAgentOrderSuccessRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* 成功订单参数：列表元素结构为——旧乘机人姓名;新乘机人姓名;pnr;票号 (以分号进行分隔)。说明：有时用户输入的乘机人姓名输入错误或者有生僻字，代理商必须输入新的名字以保证验真通过；即旧乘机人姓名和新乘机人姓名不需要变化时可以相同 */
func (r *TripJipiaoAgentOrderSuccessRequest) SetSuccessInfo(value string) {
	r.SetValue("success_info", value)
}


func (r *TripJipiaoAgentOrderSuccessRequest) GetResponse(accessToken string) (*TripJipiaoAgentOrderSuccessResponse, []byte, error) {
	var resp TripJipiaoAgentOrderSuccessResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.trip.jipiao.agent.order.success", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TripJipiaoAgentOrderSuccessResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TripJipiaoAgentOrderSuccessResponseResult struct {
	Response *TripJipiaoAgentOrderSuccessResponse `json:"trip_jipiao_agent_order_success_response"`
}



