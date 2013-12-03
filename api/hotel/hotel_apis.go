// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hotel

import (
	"github.com/yaofangou/open_taobao"
)





/* 此接口用于新增一个酒店，酒店的发布者是当前会话的用户。 
该接口发出的是一个酒店申请，需要淘宝小二审核。 */
type HotelAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 酒店地址。长度不能超过120 */
func (r *HotelAddRequest) SetAddress(value string) {
	r.SetValue("address", value)
}

/* 城市编码。参见：http://kezhan.trip.taobao.com/area.html，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取 */
func (r *HotelAddRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* domestic为true时，固定China； 
domestic为false时，必须传定义的海外国家编码值。参见：http://kezhan.trip.taobao.com/countrys.html */
func (r *HotelAddRequest) SetCountry(value string) {
	r.SetValue("country", value)
}

/* 装修年份。长度不能超过4。 */
func (r *HotelAddRequest) SetDecorateTime(value string) {
	r.SetValue("decorate_time", value)
}

/* 酒店介绍。不超过25000个汉字 */
func (r *HotelAddRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 区域（县级市）编码。参见：http://kezhan.trip.taobao.com/area.html */
func (r *HotelAddRequest) SetDistrict(value string) {
	r.SetValue("district", value)
}

/* 是否国内酒店。可选值：true，false */
func (r *HotelAddRequest) SetDomestic(value string) {
	r.SetValue("domestic", value)
}

/* 酒店级别。可选值：A,B,C,D,E,F。代表客栈公寓、经济连锁、二星级/以下、三星级/舒适、四星级/高档、五星级/豪华 */
func (r *HotelAddRequest) SetLevel(value string) {
	r.SetValue("level", value)
}

/* 酒店名称。不能超过60 */
func (r *HotelAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开业年份。长度不能超过4。 */
func (r *HotelAddRequest) SetOpeningTime(value string) {
	r.SetValue("opening_time", value)
}

/* 酒店定位。可选值：T,B。代表旅游度假、商务出行 */
func (r *HotelAddRequest) SetOrientation(value string) {
	r.SetValue("orientation", value)
}

/* 酒店图片。最大长度:500K。支持的文件类型：gif,jpg,png */
func (r *HotelAddRequest) SetPic(value string) {
	r.SetValue("pic", value)
}

/* 省份编码。参见：http://kezhan.trip.taobao.com/area.html，domestic为false时默认为0 */
func (r *HotelAddRequest) SetProvince(value string) {
	r.SetValue("province", value)
}

/* 房间数。长度不能超过4。 */
func (r *HotelAddRequest) SetRooms(value string) {
	r.SetValue("rooms", value)
}

/* 交通距离与设施服务。JSON格式。cityCenterDistance、railwayDistance、airportDistance分别代表距离市中心、距离火车站、距离机场公里数，为不超过3位正整数，默认-1代表无数据。 
其他key值true代表有此服务，false代表没有。 
parking：停车场，airportShuttle：机场接送，rentCar：租车，meetingRoom：会议室，businessCenter：商务中心，swimmingPool：游泳池，fitnessClub：健身中心，laundry：洗衣服务，morningCall：叫早服务，bankCard：接受银行卡，creditCard：接受信用卡，chineseRestaurant：中餐厅，westernRestaurant：西餐厅，cafe：咖啡厅，bar：酒吧，ktv：KTV。 */
func (r *HotelAddRequest) SetService(value string) {
	r.SetValue("service", value)
}

/* 接入卖家数据主键 */
func (r *HotelAddRequest) SetSiteParam(value string) {
	r.SetValue("site_param", value)
}

/* 楼层数。长度不能超过4。 */
func (r *HotelAddRequest) SetStoreys(value string) {
	r.SetValue("storeys", value)
}

/* 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886 */
func (r *HotelAddRequest) SetTel(value string) {
	r.SetValue("tel", value)
}


func (r *HotelAddRequest) GetResponse(accessToken string) (*HotelAddResponse, []byte, error) {
	var resp HotelAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelAddResponse struct {
	Hotel *Hotel `json:"hotel"`
}

type HotelAddResponseResult struct {
	Response *HotelAddResponse `json:"hotel_add_response"`
}





/* 根据国家编码查询该国家下的城市 */
type HotelCityGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 国家编码 */
func (r *HotelCityGetRequest) SetCountry(value string) {
	r.SetValue("country", value)
}


func (r *HotelCityGetRequest) GetResponse(accessToken string) (*HotelCityGetResponse, []byte, error) {
	var resp HotelCityGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.city.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelCityGetResponse struct {
	Result string `json:"result"`
}

type HotelCityGetResponseResult struct {
	Response *HotelCityGetResponse `json:"hotel_city_get_response"`
}





/* 此接口用于查询一个酒店，根据传入的酒店hid查询酒店信息。 */
type HotelGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* true表示查询酒店审核状态，false表示查询已审核通过酒店详细信息 */
func (r *HotelGetRequest) SetCheckAudit(value string) {
	r.SetValue("check_audit", value)
}

/* 要查询的酒店id。必须为数字 */
func (r *HotelGetRequest) SetHid(value string) {
	r.SetValue("hid", value)
}

/* 是否需要返回该酒店的房型列表。可选值：true，false。 */
func (r *HotelGetRequest) SetNeedRoomType(value string) {
	r.SetValue("need_room_type", value)
}


func (r *HotelGetRequest) GetResponse(accessToken string) (*HotelGetResponse, []byte, error) {
	var resp HotelGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelGetResponse struct {
	Hotel *Hotel `json:"hotel"`
}

type HotelGetResponseResult struct {
	Response *HotelGetResponse `json:"hotel_get_response"`
}





/* 酒店图片上传 */
type HotelImageUploadRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 酒店id */
func (r *HotelImageUploadRequest) SetHid(value string) {
	r.SetValue("hid", value)
}

/* 上传的图片 */
func (r *HotelImageUploadRequest) SetPic(value string) {
	r.SetValue("pic", value)
}


func (r *HotelImageUploadRequest) GetResponse(accessToken string) (*HotelImageUploadResponse, []byte, error) {
	var resp HotelImageUploadResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.image.upload", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelImageUploadResponse struct {
	HotelImage *HotelImage `json:"hotel_image"`
}

type HotelImageUploadResponseResult struct {
	Response *HotelImageUploadResponse `json:"hotel_image_upload_response"`
}





/* 用于回传hotel匹配结果(该接口即将废除请勿使用) */
type HotelMatchFeedbackRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需进行匹配的酒店id */
func (r *HotelMatchFeedbackRequest) SetHaid(value string) {
	r.SetValue("haid", value)
}

/* 匹配命中的酒店id */
func (r *HotelMatchFeedbackRequest) SetHid(value string) {
	r.SetValue("hid", value)
}

/* 匹配结果 0:未匹配，1:匹配 */
func (r *HotelMatchFeedbackRequest) SetMatchResult(value string) {
	r.SetValue("match_result", value)
}


func (r *HotelMatchFeedbackRequest) GetResponse(accessToken string) (*HotelMatchFeedbackResponse, []byte, error) {
	var resp HotelMatchFeedbackResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.match.feedback", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelMatchFeedbackResponse struct {
	Succ bool `json:"succ"`
}

type HotelMatchFeedbackResponseResult struct {
	Response *HotelMatchFeedbackResponse `json:"hotel_match_feedback_response"`
}





/* 此接口用于查询一个酒店，根据传入的酒店名称/别名查询酒店信息。 */
type HotelNameGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 城市编码。参见：http://kezhan.trip.taobao.com/area.html。 
domestic为true时，province,city,district不能同时为空或为0 */
func (r *HotelNameGetRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* domestic为true时，固定China； 
domestic为false时，必须传定义的海外国家编码值，是必填项。参见：http://kezhan.trip.taobao.com/countrys.html */
func (r *HotelNameGetRequest) SetCountry(value string) {
	r.SetValue("country", value)
}

/* 区域（县级市）编码。参见：http://kezhan.trip.taobao.com/area.html。 
domestic为true时，province,city,district不能同时为空或为0 */
func (r *HotelNameGetRequest) SetDistrict(value string) {
	r.SetValue("district", value)
}

/* 是否国内酒店。可选值：true，false */
func (r *HotelNameGetRequest) SetDomestic(value string) {
	r.SetValue("domestic", value)
}

/* 酒店全部名称/别名。不能超过60字节 */
func (r *HotelNameGetRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 省份编码。参见：http://kezhan.trip.taobao.com/area.html。 
domestic为true时，province,city,district不能同时为空或为0 */
func (r *HotelNameGetRequest) SetProvince(value string) {
	r.SetValue("province", value)
}


func (r *HotelNameGetRequest) GetResponse(accessToken string) (*HotelNameGetResponse, []byte, error) {
	var resp HotelNameGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.name.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelNameGetResponse struct {
	Hotel *Hotel `json:"hotel"`
}

type HotelNameGetResponseResult struct {
	Response *HotelNameGetResponse `json:"hotel_name_get_response"`
}





/* 下单结果回传(该接口即将废除请勿使用) */
type HotelOrderBookingFeedbackRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 失败原因,当result为failed时,此项为必填，最长200个字符 */
func (r *HotelOrderBookingFeedbackRequest) SetFailedReason(value string) {
	r.SetValue("failed_reason", value)
}

/* 指令消息中的messageid,最长128字符 */
func (r *HotelOrderBookingFeedbackRequest) SetMessageId(value string) {
	r.SetValue("message_id", value)
}

/* 酒店订单id */
func (r *HotelOrderBookingFeedbackRequest) SetOid(value string) {
	r.SetValue("oid", value)
}

/* 合作方订单号,最长250个字符 */
func (r *HotelOrderBookingFeedbackRequest) SetOutOid(value string) {
	r.SetValue("out_oid", value)
}

/* 在合作方退订时可能要用到的标识码，最长250个字符 */
func (r *HotelOrderBookingFeedbackRequest) SetRefundCode(value string) {
	r.SetValue("refund_code", value)
}

/* 预订结果  
S:成功 
F:失败 */
func (r *HotelOrderBookingFeedbackRequest) SetResult(value string) {
	r.SetValue("result", value)
}

/* 指令消息中的session_id */
func (r *HotelOrderBookingFeedbackRequest) SetSessionId(value string) {
	r.SetValue("session_id", value)
}


func (r *HotelOrderBookingFeedbackRequest) GetResponse(accessToken string) (*HotelOrderBookingFeedbackResponse, []byte, error) {
	var resp HotelOrderBookingFeedbackResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.order.booking.feedback", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelOrderBookingFeedbackResponse struct {
	IsSuccess bool `json:"is_success"`
}

type HotelOrderBookingFeedbackResponseResult struct {
	Response *HotelOrderBookingFeedbackResponse `json:"hotel_order_booking_feedback_response"`
}





/* 此接口用户到店支付（前台面付）卖家处理订单的入住情况，包括核实买家已入住和核实买家未入住 */
type HotelOrderFaceCheckRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 核实已入住或者未入住，true：已入住，false：未入住 */
func (r *HotelOrderFaceCheckRequest) SetChecked(value string) {
	r.SetValue("checked", value)
}

/* 实际入住时间 */
func (r *HotelOrderFaceCheckRequest) SetCheckinDate(value string) {
	r.SetValue("checkin_date", value)
}

/* 实际离店时间 */
func (r *HotelOrderFaceCheckRequest) SetCheckoutDate(value string) {
	r.SetValue("checkout_date", value)
}

/* 酒店订单id */
func (r *HotelOrderFaceCheckRequest) SetOid(value string) {
	r.SetValue("oid", value)
}


func (r *HotelOrderFaceCheckRequest) GetResponse(accessToken string) (*HotelOrderFaceCheckResponse, []byte, error) {
	var resp HotelOrderFaceCheckResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.order.face.check", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelOrderFaceCheckResponse struct {
	Result string `json:"result"`
}

type HotelOrderFaceCheckResponseResult struct {
	Response *HotelOrderFaceCheckResponse `json:"hotel_order_face_check_response"`
}





/* 该接口用于卖家确认到店支付订单或者取消订单。 */
type HotelOrderFaceDealRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 酒店订单oid */
func (r *HotelOrderFaceDealRequest) SetOid(value string) {
	r.SetValue("oid", value)
}

/* 操作类型，1：确认预订，2：取消订单 */
func (r *HotelOrderFaceDealRequest) SetOperType(value string) {
	r.SetValue("oper_type", value)
}

/* 取消订单时的取消原因备注信息 */
func (r *HotelOrderFaceDealRequest) SetReasonText(value string) {
	r.SetValue("reason_text", value)
}

/* 取消订单时的取消原因，可选值：1,2,3,4； 
1：无房，2：价格变动，3：买家原因，4：其它原因 */
func (r *HotelOrderFaceDealRequest) SetReasonType(value string) {
	r.SetValue("reason_type", value)
}


func (r *HotelOrderFaceDealRequest) GetResponse(accessToken string) (*HotelOrderFaceDealResponse, []byte, error) {
	var resp HotelOrderFaceDealResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.order.face.deal", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelOrderFaceDealResponse struct {
	Result string `json:"result"`
}

type HotelOrderFaceDealResponseResult struct {
	Response *HotelOrderFaceDealResponse `json:"hotel_order_face_deal_response"`
}





/* 此接口用于查询一个酒店订单，根据传入的订单号查询订单信息。 */
type HotelOrderGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 是否需要返回该订单的入住人列表。可选值：true，false。 */
func (r *HotelOrderGetRequest) SetNeedGuest(value string) {
	r.SetValue("need_guest", value)
}

/* 是否显示买家留言，可选值true、false */
func (r *HotelOrderGetRequest) SetNeedMessage(value string) {
	r.SetValue("need_message", value)
}

/* 酒店订单oid，必须为数字。oid，tid必须传一项，同时传递的情况下，作为查询条件的优先级由高到低依次为oid，tid。 */
func (r *HotelOrderGetRequest) SetOid(value string) {
	r.SetValue("oid", value)
}

/* 淘宝订单tid，必须为数字。oid，tid必须传一项，同时传递的情况下，作为查询条件的优先级由高到低依次为oid，tid。 */
func (r *HotelOrderGetRequest) SetTid(value string) {
	r.SetValue("tid", value)
}


func (r *HotelOrderGetRequest) GetResponse(accessToken string) (*HotelOrderGetResponse, []byte, error) {
	var resp HotelOrderGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.order.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelOrderGetResponse struct {
	HotelOrder *HotelOrder `json:"hotel_order"`
}

type HotelOrderGetResponseResult struct {
	Response *HotelOrderGetResponse `json:"hotel_order_get_response"`
}





/* 支付确认结果回传(该接口即将废除请勿使用) */
type HotelOrderPayFeedbackRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 失败原因,当result为failed时,此项为必填，最长200个字符 */
func (r *HotelOrderPayFeedbackRequest) SetFailedReason(value string) {
	r.SetValue("failed_reason", value)
}

/* 指令消息中的messageid,最长128字符 */
func (r *HotelOrderPayFeedbackRequest) SetMessageId(value string) {
	r.SetValue("message_id", value)
}

/* 酒店订单id */
func (r *HotelOrderPayFeedbackRequest) SetOid(value string) {
	r.SetValue("oid", value)
}

/* 合作方订单号,最长250个字符 */
func (r *HotelOrderPayFeedbackRequest) SetOutOid(value string) {
	r.SetValue("out_oid", value)
}

/* 预订结果  
S:成功 
F:失败 */
func (r *HotelOrderPayFeedbackRequest) SetResult(value string) {
	r.SetValue("result", value)
}

/* 指令消息中的session_id */
func (r *HotelOrderPayFeedbackRequest) SetSessionId(value string) {
	r.SetValue("session_id", value)
}


func (r *HotelOrderPayFeedbackRequest) GetResponse(accessToken string) (*HotelOrderPayFeedbackResponse, []byte, error) {
	var resp HotelOrderPayFeedbackResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.order.pay.feedback", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelOrderPayFeedbackResponse struct {
	IsSuccess bool `json:"is_success"`
}

type HotelOrderPayFeedbackResponseResult struct {
	Response *HotelOrderPayFeedbackResponse `json:"hotel_order_pay_feedback_response"`
}





/* 退订处理结果回传(该接口即将废除请勿使用) */
type HotelOrderRefundFeedbackRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 失败原因,当result为F时,此项为必填,最长200个字符 */
func (r *HotelOrderRefundFeedbackRequest) SetFailedReason(value string) {
	r.SetValue("failed_reason", value)
}

/* 指令消息中的messageid,最长128字符 */
func (r *HotelOrderRefundFeedbackRequest) SetMessageId(value string) {
	r.SetValue("message_id", value)
}

/* 合作方订单号,最长250个字符 */
func (r *HotelOrderRefundFeedbackRequest) SetOid(value string) {
	r.SetValue("oid", value)
}

/* 合作方订单号,最长250个字符 */
func (r *HotelOrderRefundFeedbackRequest) SetOutOid(value string) {
	r.SetValue("out_oid", value)
}

/* 预订结果  
S:成功 
F:失败 */
func (r *HotelOrderRefundFeedbackRequest) SetResult(value string) {
	r.SetValue("result", value)
}

/* 指令消息中的session_id */
func (r *HotelOrderRefundFeedbackRequest) SetSessionId(value string) {
	r.SetValue("session_id", value)
}


func (r *HotelOrderRefundFeedbackRequest) GetResponse(accessToken string) (*HotelOrderRefundFeedbackResponse, []byte, error) {
	var resp HotelOrderRefundFeedbackResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.order.refund.feedback", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelOrderRefundFeedbackResponse struct {
	IsSuccess bool `json:"is_success"`
}

type HotelOrderRefundFeedbackResponseResult struct {
	Response *HotelOrderRefundFeedbackResponse `json:"hotel_order_refund_feedback_response"`
}





/* 此接口用于查询多个酒店订单，根据传入的查询条件查询订单信息。 */
type HotelOrdersSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 入住时间查询结束时间，格式为：yyyy-MM-dd。不能早于checkin_date_start，间隔不能大于30 */
func (r *HotelOrdersSearchRequest) SetCheckinDateEnd(value string) {
	r.SetValue("checkin_date_end", value)
}

/* 入住时间查询起始时间，格式为：yyyy-MM-dd */
func (r *HotelOrdersSearchRequest) SetCheckinDateStart(value string) {
	r.SetValue("checkin_date_start", value)
}

/* 离店时间查询结束时间，格式为：yyyy-MM-dd。不能早于checkout_date_start，间隔不能大于30 */
func (r *HotelOrdersSearchRequest) SetCheckoutDateEnd(value string) {
	r.SetValue("checkout_date_end", value)
}

/* 离店时间查询起始时间，格式为：yyyy-MM-dd */
func (r *HotelOrdersSearchRequest) SetCheckoutDateStart(value string) {
	r.SetValue("checkout_date_start", value)
}

/* 订单创建时间查询结束时间，格式为：yyyy-MM-dd。不能早于created_start，间隔不能大于30 */
func (r *HotelOrdersSearchRequest) SetCreatedEnd(value string) {
	r.SetValue("created_end", value)
}

/* 订单创建时间查询起始时间，格式为：yyyy-MM-dd */
func (r *HotelOrdersSearchRequest) SetCreatedStart(value string) {
	r.SetValue("created_start", value)
}

/* 商品gid列表，多个gid用英文逗号隔开，一次不超过5个 */
func (r *HotelOrdersSearchRequest) SetGids(value string) {
	r.SetValue("gids", value)
}

/* 酒店hid列表，多个hid用英文逗号隔开，一次不超过5个 */
func (r *HotelOrdersSearchRequest) SetHids(value string) {
	r.SetValue("hids", value)
}

/* 是否需要返回该订单的入住人列表。可选值：true，false。 */
func (r *HotelOrdersSearchRequest) SetNeedGuest(value string) {
	r.SetValue("need_guest", value)
}

/* 是否显示买家留言，可选值true、false */
func (r *HotelOrdersSearchRequest) SetNeedMessage(value string) {
	r.SetValue("need_message", value)
}

/* 酒店订单oids列表，多个oid用英文逗号隔开，一次不超过20个。oids，tids，hids，rids，gids，（checkin_date_start，checkin_date_end），（checkout_date_start，checkout_date_end），（created_start，created_end）必须传入一项，括号表示需同时存在才做为查询条件。 
oids，tids，hids，rids，gids同时出现时，优先级按此顺序由高到低只取一项。其他同时出现时，并列使用。 */
func (r *HotelOrdersSearchRequest) SetOids(value string) {
	r.SetValue("oids", value)
}

/* 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20 */
func (r *HotelOrdersSearchRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 房型rid列表，多个rid用英文逗号隔开，一次不超过5个 */
func (r *HotelOrdersSearchRequest) SetRids(value string) {
	r.SetValue("rids", value)
}

/* 订单状态。A：等待买家付款。B：买家已付款待卖家发货。C：卖家已发货待买家确认。D：交易成功。E：交易关闭 */
func (r *HotelOrdersSearchRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 淘宝订单tid列表，多个tid用英文逗号隔开，一次不超过20个。 */
func (r *HotelOrdersSearchRequest) SetTids(value string) {
	r.SetValue("tids", value)
}


func (r *HotelOrdersSearchRequest) GetResponse(accessToken string) (*HotelOrdersSearchResponse, []byte, error) {
	var resp HotelOrdersSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.orders.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelOrdersSearchResponse struct {
	HotelOrders []*HotelOrder `json:"hotel_orders"`
	TotalResults int `json:"total_results"`
}

type HotelOrdersSearchResponseResult struct {
	Response *HotelOrdersSearchResponse `json:"hotel_orders_search_response"`
}





/* 此接口用于发布一个集市酒店商品，商品的发布者是当前会话的用户。如果该酒店、该房型、该用户所对应的商品在淘宝集市酒店系统中已经存在，则会返回错误提示。 */
type HotelRoomAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 面积。可选值：A,B,C,D。分别代表： 
A：15平米以下，B：16－30平米，C：31－50平米，D：50平米以上 */
func (r *HotelRoomAddRequest) SetArea(value string) {
	r.SetValue("area", value)
}

/* 宽带服务。A,B,C,D。分别代表： 
A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带 */
func (r *HotelRoomAddRequest) SetBbn(value string) {
	r.SetValue("bbn", value)
}

/* 床型。可选值：A,B,C,D,E,F,G,H,I。分别代表：A：单人床，B：大床，C：双床，D：双床/大床，E：子母床，F：上下床，G：圆形床，H：多床，I：其他床型 */
func (r *HotelRoomAddRequest) SetBedType(value string) {
	r.SetValue("bed_type", value)
}

/* 早餐。A,B,C,D,E。分别代表： 
A：无早，B：单早，C：双早，D：三早，E：多早 */
func (r *HotelRoomAddRequest) SetBreakfast(value string) {
	r.SetValue("breakfast", value)
}

/* 订金。0～99999900的正整数。在payment_type为订金时必须输入，存储的单位是分。不能带角分。 */
func (r *HotelRoomAddRequest) SetDeposit(value string) {
	r.SetValue("deposit", value)
}

/* 商品描述。不能超过25000个汉字（50000个字符）。 */
func (r *HotelRoomAddRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 手续费。0～99999900的正整数。在payment_type为手续费或手续费/间夜时必须输入，存储的单位是分。不能带角分。 */
func (r *HotelRoomAddRequest) SetFee(value string) {
	r.SetValue("fee", value)
}

/* 购买须知。不能超过300个字。 */
func (r *HotelRoomAddRequest) SetGuide(value string) {
	r.SetValue("guide", value)
}

/* 酒店商品是否提供发票 */
func (r *HotelRoomAddRequest) SetHasReceipt(value string) {
	r.SetValue("has_receipt", value)
}

/* 酒店id。必须为数字。 */
func (r *HotelRoomAddRequest) SetHid(value string) {
	r.SetValue("hid", value)
}

/* 为到店支付卖家特殊使用，代表多种支付类型的房态。room_quotas可选，如果有值，也会处理。 */
func (r *HotelRoomAddRequest) SetMultiRoomQuotas(value string) {
	r.SetValue("multi_room_quotas", value)
}

/* 支付类型。可选值：A,B,C,D,E。分别代表： 
A：全额支付，B：手续费，C：订金，D：手续费/间夜，E：前台面付 */
func (r *HotelRoomAddRequest) SetPaymentType(value string) {
	r.SetValue("payment_type", value)
}

/* 酒店商品图片。类型:JPG,GIF;最大长度:500K。支持的文件类型：gif,jpg,jpeg,png。发布的时候只能发布一张图片。如需再发图片，需要调用商品图片上传接口，1个商品最多5张图片。 */
func (r *HotelRoomAddRequest) SetPic(value string) {
	r.SetValue("pic", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
func (r *HotelRoomAddRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 价格类型。可选值：A,B。分别代表：A：参考预订价，B实时预订价 。未选该参数默认为参考预订价。选择实时预订价的情况下，支付类型必须选择为A(全额支付) */
func (r *HotelRoomAddRequest) SetPriceType(value string) {
	r.SetValue("price_type", value)
}

/* 发票说明，不能超过100个汉字,200个字符。 */
func (r *HotelRoomAddRequest) SetReceiptInfo(value string) {
	r.SetValue("receipt_info", value)
}

/* 发票类型为其他时的发票描述,不能超过30个汉字，60个字符。 */
func (r *HotelRoomAddRequest) SetReceiptOtherTypeDesc(value string) {
	r.SetValue("receipt_other_type_desc", value)
}

/* 发票类型。A,B。分别代表： A:酒店住宿发票,B:其他 */
func (r *HotelRoomAddRequest) SetReceiptType(value string) {
	r.SetValue("receipt_type", value)
}

/* 1. 全额支付类型必填 
2. t代表类别(1表示任意退;2表示不能退;3表示阶梯退)，p代表退款规则（数组）， d代表天数，r代表扣除手续费比率。示例代表的意思就是"阶梯退:提前3天内退订，收取订单总额10%的违约金;提前2天内退订，收取订单总额20%的违约金，提前1天内退订，收取订单总额30%的违约金"。 
3. 任意退、不能退不能指定退款规则;阶梯退不能没有退款规则;阶梯退规则最多10条,且每条规则天数、费率不能相同;阶梯退遵循天数越短,违约金越大的业务规则;天数需为整数,且90>天数>=0;费率需为整数且100<=费率<=0;阶梯退规则只有一条时,费率不能设置为100%;阶梯退规则只有一条时,不能设定0天收取0%; */
func (r *HotelRoomAddRequest) SetRefundPolicyInfo(value string) {
	r.SetValue("refund_policy_info", value)
}

/* 房型id。必须为数字。 */
func (r *HotelRoomAddRequest) SetRid(value string) {
	r.SetValue("rid", value)
}

/* 房态信息。可以传今天开始90天内的房态信息。日期必须连续。JSON格式传递。 
date：代表房态日期，格式为YYYY-MM-DD， 
price：代表当天房价，0～99999999，存储的单位是分， 
num：代表当天可售间数，0～999。 
如： 
[{"date":2011-01-28,"price":10000, "num":10},{"date":2011-01-29,"price":12000,"num":10}] */
func (r *HotelRoomAddRequest) SetRoomQuotas(value string) {
	r.SetValue("room_quotas", value)
}

/* 设施服务。JSON格式。 
value值true有此服务，false没有。 
bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 
如： 
{"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false} */
func (r *HotelRoomAddRequest) SetService(value string) {
	r.SetValue("service", value)
}

/* 接入卖家数据主键 */
func (r *HotelRoomAddRequest) SetSiteParam(value string) {
	r.SetValue("site_param", value)
}

/* 床宽。可选值：A,B,C,D,E,F,G,H。分别代表：A：1米及以下，B：1.1米，C：1.2米，D：1.35米，E：1.5米，F：1.8米，G：2米，H：2.2米及以上 */
func (r *HotelRoomAddRequest) SetSize(value string) {
	r.SetValue("size", value)
}

/* 楼层。长度不超过8 */
func (r *HotelRoomAddRequest) SetStorey(value string) {
	r.SetValue("storey", value)
}

/* 酒店商品名称。不能超过60字节 */
func (r *HotelRoomAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *HotelRoomAddRequest) GetResponse(accessToken string) (*HotelRoomAddResponse, []byte, error) {
	var resp HotelRoomAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.room.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelRoomAddResponse struct {
	Room *Room `json:"room"`
}

type HotelRoomAddResponseResult struct {
	Response *HotelRoomAddResponse `json:"hotel_room_add_response"`
}





/* 此接口用于查询一个商品，根据传入的gid查询商品信息。卖家只能查询自己的商品。 */
type HotelRoomGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 酒店房间商品gid。必须为数字。gid和item_id至少要传一个。 */
func (r *HotelRoomGetRequest) SetGid(value string) {
	r.SetValue("gid", value)
}

/* 酒店房间商品item_id。必须为数字。item_id和gid至少要传一个。 */
func (r *HotelRoomGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 是否需要返回该商品的酒店信息。可选值：true，false。 */
func (r *HotelRoomGetRequest) SetNeedHotel(value string) {
	r.SetValue("need_hotel", value)
}

/* 是否需要返回该商品的宝贝描述。可选值：true，false。 */
func (r *HotelRoomGetRequest) SetNeedRoomDesc(value string) {
	r.SetValue("need_room_desc", value)
}

/* 是否需要返回该商品的房态列表。可选值：true，false。 */
func (r *HotelRoomGetRequest) SetNeedRoomQuotas(value string) {
	r.SetValue("need_room_quotas", value)
}

/* 是否需要返回该商品的房型信息。可选值：true，false。 */
func (r *HotelRoomGetRequest) SetNeedRoomType(value string) {
	r.SetValue("need_room_type", value)
}


func (r *HotelRoomGetRequest) GetResponse(accessToken string) (*HotelRoomGetResponse, []byte, error) {
	var resp HotelRoomGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.room.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelRoomGetResponse struct {
	Room *Room `json:"room"`
}

type HotelRoomGetResponseResult struct {
	Response *HotelRoomGetResponse `json:"hotel_room_get_response"`
}





/* 此接口用于为商品删除商品图片。 */
type HotelRoomImgDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 酒店房间商品gid。必须为数字。 */
func (r *HotelRoomImgDeleteRequest) SetGid(value string) {
	r.SetValue("gid", value)
}

/* 图片序号，可选值：1，2，3，4，5。 
如果原图片个数小于等于1，则报错，不能删除图片。 
如果原图片个数小于待删除的图片序号，则报错，图片序号错误。 */
func (r *HotelRoomImgDeleteRequest) SetPosition(value string) {
	r.SetValue("position", value)
}


func (r *HotelRoomImgDeleteRequest) GetResponse(accessToken string) (*HotelRoomImgDeleteResponse, []byte, error) {
	var resp HotelRoomImgDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.room.img.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelRoomImgDeleteResponse struct {
	RoomImage *RoomImage `json:"room_image"`
}

type HotelRoomImgDeleteResponseResult struct {
	Response *HotelRoomImgDeleteResponse `json:"hotel_room_img_delete_response"`
}





/* 此接口用于为商品上传商品图片。 */
type HotelRoomImgUploadRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 酒店房间商品gid。必须为数字。 */
func (r *HotelRoomImgUploadRequest) SetGid(value string) {
	r.SetValue("gid", value)
}

/* 酒店商品图片。类型:JPG,GIF;最大长度:500K。支持的文件类型：gif,jpg,jpeg,png。 
如果原图片少于5张，若没传序号或序号大于原图片个数，则在原图片最后添加，否则按序号插入到原图片中去，自动后移。 
如果原图片大于5张，若没传序号，则替换最后一张图片，否则在序号位置插入，图片向后移，最后一张被删除。 */
func (r *HotelRoomImgUploadRequest) SetPic(value string) {
	r.SetValue("pic", value)
}

/* 图片序号，可选值：1，2，3，4，5 */
func (r *HotelRoomImgUploadRequest) SetPosition(value string) {
	r.SetValue("position", value)
}


func (r *HotelRoomImgUploadRequest) GetResponse(accessToken string) (*HotelRoomImgUploadResponse, []byte, error) {
	var resp HotelRoomImgUploadResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.room.img.upload", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelRoomImgUploadResponse struct {
	RoomImage *RoomImage `json:"room_image"`
}

type HotelRoomImgUploadResponseResult struct {
	Response *HotelRoomImgUploadResponse `json:"hotel_room_img_upload_response"`
}





/* 接入方房态查询结果返回(该接口即将废除请勿使用) */
type HotelRoomQuotasQueryFeedbackRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 选中日期范围内的最大可用房量 */
func (r *HotelRoomQuotasQueryFeedbackRequest) SetAvaliableRoomCount(value string) {
	r.SetValue("avaliable_room_count", value)
}

/* 入住酒店的日期 */
func (r *HotelRoomQuotasQueryFeedbackRequest) SetCheckinDate(value string) {
	r.SetValue("checkin_date", value)
}

/* 离开酒店的日期 */
func (r *HotelRoomQuotasQueryFeedbackRequest) SetCheckoutDate(value string) {
	r.SetValue("checkout_date", value)
}

/* 失败原因,当result为F时,此项为必填,最长200个字符 */
func (r *HotelRoomQuotasQueryFeedbackRequest) SetFailedReason(value string) {
	r.SetValue("failed_reason", value)
}

/* 指令消息中的messageid,最长128字符 */
func (r *HotelRoomQuotasQueryFeedbackRequest) SetMessageId(value string) {
	r.SetValue("message_id", value)
}

/* 预订结果  
S:成功 
F:失败 */
func (r *HotelRoomQuotasQueryFeedbackRequest) SetResult(value string) {
	r.SetValue("result", value)
}

/* 从入住时期到离店日期的每日一间房价与预定房量,JSON格式传递。 date：代表房态日期，格式为YYYY-MM-DD， price：代表当天房价，0～99999900，存储的单位是分，货币单位为人民币，num：代表当天剩余房量，0～999，所有日期的预订间数应该一致。 如： [{"date":2011-01-28,"price":10000, "num":10},{"date":2011-01-29,"price":12000,"num":10}],最长1500个字符 */
func (r *HotelRoomQuotasQueryFeedbackRequest) SetRoomQuotas(value string) {
	r.SetValue("room_quotas", value)
}

/* 订单总价。0～99999999的正整数。货币单位为人民币。 */
func (r *HotelRoomQuotasQueryFeedbackRequest) SetTotalRoomPrice(value string) {
	r.SetValue("total_room_price", value)
}


func (r *HotelRoomQuotasQueryFeedbackRequest) GetResponse(accessToken string) (*HotelRoomQuotasQueryFeedbackResponse, []byte, error) {
	var resp HotelRoomQuotasQueryFeedbackResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.room.quotas.query.feedback", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelRoomQuotasQueryFeedbackResponse struct {
	IsSuccess bool `json:"is_success"`
}

type HotelRoomQuotasQueryFeedbackResponseResult struct {
	Response *HotelRoomQuotasQueryFeedbackResponse `json:"hotel_room_quotas_query_feedback_response"`
}





/* 此接口用于更新一个集市酒店商品，根据传入的gid更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在淘宝集市酒店系统中不存在，则会返回错误提示。 */
type HotelRoomUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 面积。可选值：A,B,C,D。分别代表： 
A：15平米以下，B：16－30平米，C：31－50平米，D：50平米以上 */
func (r *HotelRoomUpdateRequest) SetArea(value string) {
	r.SetValue("area", value)
}

/* 宽带服务。A,B,C,D。分别代表： 
A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带 */
func (r *HotelRoomUpdateRequest) SetBbn(value string) {
	r.SetValue("bbn", value)
}

/* 床型。可选值：A,B,C,D,E,F,G,H,I。分别代表：A：单人床，B：大床，C：双床，D：双床/大床，E：子母床，F：上下床，G：圆形床，H：多床，I：其他床型 */
func (r *HotelRoomUpdateRequest) SetBedType(value string) {
	r.SetValue("bed_type", value)
}

/* 早餐。A,B,C,D,E。分别代表： 
A：无早，B：单早，C：双早，D：三早，E：多早 */
func (r *HotelRoomUpdateRequest) SetBreakfast(value string) {
	r.SetValue("breakfast", value)
}

/* 订金。0～99999900的正整数。在payment_type为订金时必须输入，存储的单位是分。不能带角分。 */
func (r *HotelRoomUpdateRequest) SetDeposit(value string) {
	r.SetValue("deposit", value)
}

/* 商品描述。不能超过25000个汉字（50000个字符）。 */
func (r *HotelRoomUpdateRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 手续费。0～99999900的正整数。在payment_type为手续费或手续费/间夜时必须输入，存储的单位是分。不能带角分。 */
func (r *HotelRoomUpdateRequest) SetFee(value string) {
	r.SetValue("fee", value)
}

/* 酒店房间商品gid。必须为数字。 */
func (r *HotelRoomUpdateRequest) SetGid(value string) {
	r.SetValue("gid", value)
}

/* 购买须知。不能超过300个字。 */
func (r *HotelRoomUpdateRequest) SetGuide(value string) {
	r.SetValue("guide", value)
}

/* 酒店商品是否提供发票 */
func (r *HotelRoomUpdateRequest) SetHasReceipt(value string) {
	r.SetValue("has_receipt", value)
}

/* 为到店支付卖家特殊使用，代表多种支付类型的房态。room_quotas可选，如果有值，也会处理。 */
func (r *HotelRoomUpdateRequest) SetMultiRoomQuotas(value string) {
	r.SetValue("multi_room_quotas", value)
}

/* 支付类型。可选值：A,B,C,D,E。分别代表： 
A：全额支付，B：手续费，C：订金，D：手续费/间夜，E：前台面付 */
func (r *HotelRoomUpdateRequest) SetPaymentType(value string) {
	r.SetValue("payment_type", value)
}

/* 酒店商品图片。类型:JPG,GIF;最大长度:500K。支持的文件类型：gif,jpg,jpeg,png。更新的时候只能更新一张图片，此图片覆盖原有所有图片。如果不传则使用原有所有图片。 
如需再发图片，需要调用商品图片上传接口，1个商品最多5张图片。 */
func (r *HotelRoomUpdateRequest) SetPic(value string) {
	r.SetValue("pic", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
func (r *HotelRoomUpdateRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 价格类型。可选值：A,B。分别代表：A：参考预订价，B实时预订价 。未选该参数默认为参考预订价。选择实时预订价的情况下，支付类型必须选择为A(全额支付) */
func (r *HotelRoomUpdateRequest) SetPriceType(value string) {
	r.SetValue("price_type", value)
}

/* 发票说明，不能超过100个汉字,200个字符。 */
func (r *HotelRoomUpdateRequest) SetReceiptInfo(value string) {
	r.SetValue("receipt_info", value)
}

/* 发票类型为其他时的发票描述,不能超过30个汉字，60个字符。 */
func (r *HotelRoomUpdateRequest) SetReceiptOtherTypeDesc(value string) {
	r.SetValue("receipt_other_type_desc", value)
}

/* 发票类型。A,B。分别代表： A:酒店住宿发票,B:其他 */
func (r *HotelRoomUpdateRequest) SetReceiptType(value string) {
	r.SetValue("receipt_type", value)
}

/* 1. 全额支付类型必填 2. t代表类别(1表示任意退;2表示不能退;3表示阶梯退)，p代表退款规则（数组）， d代表天数，r代表扣除手续费比率。示例代表的意思就是"阶梯退:提前3天内退订，收取订单总额10%的违约金;提前2天内退订，收取订单总额20%的违约金，提前1天内退订，收取订单总额30%的违约金"。 3. 任意退、不能退不能指定退款规则;阶梯退不能没有退款规则;阶梯退规则最多10条,且每条规则天数、费率不能相同;阶梯退遵循天数越短,违约金越大的业务规则;天数需为整数,且90>天数>=0;费率需为整数且100<=费率<=0;阶梯退规则只有一条时,费率不能设置为100%;阶梯退规则只有一条时,不能设定0天收取0%; */
func (r *HotelRoomUpdateRequest) SetRefundPolicyInfo(value string) {
	r.SetValue("refund_policy_info", value)
}

/* 房态信息。可以传今天开始90天内的房态信息。日期必须连续。JSON格式传递。 
date：代表房态日期，格式为YYYY-MM-DD， 
price：代表当天房价，0～99999999，存储的单位是分, 
num：代表当天可售间数，0～999。 
如： 
[{"date":2011-01-28,"price":10000, "num":10},{"date":2011-01-29,"price":12000,"num":10}] */
func (r *HotelRoomUpdateRequest) SetRoomQuotas(value string) {
	r.SetValue("room_quotas", value)
}

/* 设施服务。JSON格式。 
value值true有此服务，false没有。 
bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 
如： 
{"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false} */
func (r *HotelRoomUpdateRequest) SetService(value string) {
	r.SetValue("service", value)
}

/* 商品的site_param */
func (r *HotelRoomUpdateRequest) SetSiteParam(value string) {
	r.SetValue("site_param", value)
}

/* 床宽。可选值：A,B,C,D,E,F,G,H。分别代表：A：1米及以下，B：1.1米，C：1.2米，D：1.35米，E：1.5米，F：1.8米，G：2米，H：2.2米及以上 */
func (r *HotelRoomUpdateRequest) SetSize(value string) {
	r.SetValue("size", value)
}

/* 状态。可选值1，2，3。1：上架。2：下架。3：删除。传入相应状态代表去执行相应的操作。 */
func (r *HotelRoomUpdateRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 楼层。长度不超过8 */
func (r *HotelRoomUpdateRequest) SetStorey(value string) {
	r.SetValue("storey", value)
}

/* 酒店商品名称。不能超过60字节 */
func (r *HotelRoomUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *HotelRoomUpdateRequest) GetResponse(accessToken string) (*HotelRoomUpdateResponse, []byte, error) {
	var resp HotelRoomUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.room.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelRoomUpdateResponse struct {
	Room *Room `json:"room"`
}

type HotelRoomUpdateResponseResult struct {
	Response *HotelRoomUpdateResponse `json:"hotel_room_update_response"`
}





/* 此接口用于查询多个酒店商品，根据传入的参数查询商品信息。卖家只能查询自己的商品。 */
type HotelRoomsSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 酒店房间商品gid列表，多个gid用英文逗号隔开，一次不超过20个。gids，item_ids , hids，rids四项必须传一项，同时传递的情况下，作为查询条件的优先级由高到低依次为gids，item_ids , hids，rids。 */
func (r *HotelRoomsSearchRequest) SetGids(value string) {
	r.SetValue("gids", value)
}

/* 酒店hid列表，多个hid用英文逗号隔开，一次不超过5个。gids，item_ids , hids，rids四项必须传一项，同时传递的情况下，作为查询条件的优先级由高到低依次为gids，item_ids , hids，rids。 */
func (r *HotelRoomsSearchRequest) SetHids(value string) {
	r.SetValue("hids", value)
}

/* 酒店房间商品item_id列表，多个item_id用英文逗号隔开，一次不超过20个。gids，item_ids , hids，rids四项必须传一项，同时传递的情况下，作为查询条件的优先级由高到低依次为gids，item_ids , hids，rids。当item_ids参数值为-1，gids项不传值时，会返回卖家所有商品列表 */
func (r *HotelRoomsSearchRequest) SetItemIds(value string) {
	r.SetValue("item_ids", value)
}

/* 是否需要返回该商品的酒店信息。可选值：true，false。 */
func (r *HotelRoomsSearchRequest) SetNeedHotel(value string) {
	r.SetValue("need_hotel", value)
}

/* 是否需要返回该商品的宝贝描述。可选值：true，false。 */
func (r *HotelRoomsSearchRequest) SetNeedRoomDesc(value string) {
	r.SetValue("need_room_desc", value)
}

/* 是否需要返回该商品的房态列表。可选值：true，false。 */
func (r *HotelRoomsSearchRequest) SetNeedRoomQuotas(value string) {
	r.SetValue("need_room_quotas", value)
}

/* 是否需要返回该商品的房型信息。可选值：true，false。 */
func (r *HotelRoomsSearchRequest) SetNeedRoomType(value string) {
	r.SetValue("need_room_type", value)
}

/* 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20 */
func (r *HotelRoomsSearchRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 房型rid列表，多个rid用英文逗号隔开，一次不超过20个。gids，item_ids , hids，rids四项必须传一项，同时传递的情况下，作为查询条件的优先级由高到低依次为gids，item_ids , hids，rids。 */
func (r *HotelRoomsSearchRequest) SetRids(value string) {
	r.SetValue("rids", value)
}


func (r *HotelRoomsSearchRequest) GetResponse(accessToken string) (*HotelRoomsSearchResponse, []byte, error) {
	var resp HotelRoomsSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.rooms.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelRoomsSearchResponse struct {
	Rooms []*Room `json:"rooms"`
	TotalResults int `json:"total_results"`
}

type HotelRoomsSearchResponseResult struct {
	Response *HotelRoomsSearchResponse `json:"hotel_rooms_search_response"`
}





/* 此接口用于更新多个集市酒店商品房态信息，根据传入的gids更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在淘宝集市酒店系统中不存在，则会返回错误提示。是全量更新，非增量，会把之前的房态进行覆盖。 */
type HotelRoomsUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 多商品房态信息。json encode。每个商品房态参考单商品更新中的room_quota字段。反序列化后入：array(( 'gid'=>1, 'roomQuota'=>array(('date'=>'2011-01-29', 'price'=>100, 'num'=>1),('date'=>'2011-01-30', 'price'=>100, 'num'=>1)),( 'gid'=>2, 'roomQuota'=>array(('date'=>'2011-01-29', 'price'=>100, 'num'=>1),('date'=>'2011-01-30', 'price'=>100, 'num'=>1))) */
func (r *HotelRoomsUpdateRequest) SetGidRoomQuotaMap(value string) {
	r.SetValue("gid_room_quota_map", value)
}

/* 为到店支付卖家特殊使用，可传入多种支付类型的房态信息。 
该参数有值时，忽略gid_room_quota_map参数； 
该参数无值时，使用gid_room_quota_map参数 */
func (r *HotelRoomsUpdateRequest) SetMultiRoomQuotas(value string) {
	r.SetValue("multi_room_quotas", value)
}


func (r *HotelRoomsUpdateRequest) GetResponse(accessToken string) (*HotelRoomsUpdateResponse, []byte, error) {
	var resp HotelRoomsUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.rooms.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelRoomsUpdateResponse struct {
	Gids []string `json:"gids"`
}

type HotelRoomsUpdateResponseResult struct {
	Response *HotelRoomsUpdateResponse `json:"hotel_rooms_update_response"`
}





/* 1. 此接口用于查询该会话用户作为酒店发布者发布的酒店被审核通过的增量酒店信息。 
2. 查询提交的起始时间至今的增量酒店记录：start_modified：2011-7-1 16:00:00 
3. 返回数据结果为发布酒店时间的正序排列 */
type HotelSoldHotelsIncrementGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 【不推荐使用，现在总是返回从修改开始时间到目前为止的所有记录，与修改结束时间不再相关】查询修改结束时间，必须大于修改开始时间（修改时间跨度不能大于1天）。格式：yyyy-MM-dd HH:mm:ss。 */
func (r *HotelSoldHotelsIncrementGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据 */
func (r *HotelSoldHotelsIncrementGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页面大小，取值范围1-100，默认大小20 */
func (r *HotelSoldHotelsIncrementGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询修改开始时间（修改时间跨度不能大于1天）。格式：yyyy-MM-dd HH:mm:ss */
func (r *HotelSoldHotelsIncrementGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}

/* 【不推荐使用，现在返回结果总会包含总记录数和是否存在下一页】是否使用has_next的分页方式，如果指定true，则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的字段，效率比总记录数高 */
func (r *HotelSoldHotelsIncrementGetRequest) SetUseHasNext(value string) {
	r.SetValue("use_has_next", value)
}


func (r *HotelSoldHotelsIncrementGetRequest) GetResponse(accessToken string) (*HotelSoldHotelsIncrementGetResponse, []byte, error) {
	var resp HotelSoldHotelsIncrementGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.sold.hotels.increment.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelSoldHotelsIncrementGetResponse struct {
	HasNext bool `json:"has_next"`
	Hotels []*Hotel `json:"hotels"`
	TotalResults int `json:"total_results"`
}

type HotelSoldHotelsIncrementGetResponseResult struct {
	Response *HotelSoldHotelsIncrementGetResponse `json:"hotel_sold_hotels_increment_get_response"`
}





/* 1. 搜索当前会话用户作为卖家已卖出的增量交易数据  
2. 只能查询时间跨度为一天的增量交易记录：start_modified：2011-7-1 16:00:00 end_modified： 2011-7-2 15:59:59（注意不能写成16:00:00）  
3. 返回数据结果为创建订单时间的倒序 */
type HotelSoldOrdersIncrementGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询修改结束时间，必须大于修改开始时间（修改时间跨度不能大于1天）。格式：yyyy-MM-dd HH:mm:ss。 */
func (r *HotelSoldOrdersIncrementGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 是否需要返回该订单的入住人列表。可选值：true，false。 */
func (r *HotelSoldOrdersIncrementGetRequest) SetNeedGuest(value string) {
	r.SetValue("need_guest", value)
}

/* 是否返回买家留言，可选值true、false */
func (r *HotelSoldOrdersIncrementGetRequest) SetNeedMessage(value string) {
	r.SetValue("need_message", value)
}

/* 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。 */
func (r *HotelSoldOrdersIncrementGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页面大小，取值范围1-100，默认大小20。 */
func (r *HotelSoldOrdersIncrementGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询修改开始时间（修改时间跨度不能大于1天）。格式：yyyy-MM-dd HH:mm:ss */
func (r *HotelSoldOrdersIncrementGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}

/* 交易状态，默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。可选值：A：等待买家付款。B：买家已付款待卖家发货。C：卖家已发货待买家确认。D：交易成功。E：交易关闭 */
func (r *HotelSoldOrdersIncrementGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 是否使用has_next的分页方式，如果指定true，则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的字段，效率比总记录数高 */
func (r *HotelSoldOrdersIncrementGetRequest) SetUseHasNext(value string) {
	r.SetValue("use_has_next", value)
}


func (r *HotelSoldOrdersIncrementGetRequest) GetResponse(accessToken string) (*HotelSoldOrdersIncrementGetResponse, []byte, error) {
	var resp HotelSoldOrdersIncrementGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.sold.orders.increment.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelSoldOrdersIncrementGetResponse struct {
	HasNext bool `json:"has_next"`
	HotelOrders []*HotelOrder `json:"hotel_orders"`
	TotalResults int `json:"total_results"`
}

type HotelSoldOrdersIncrementGetResponseResult struct {
	Response *HotelSoldOrdersIncrementGetResponse `json:"hotel_sold_orders_increment_get_response"`
}





/* 1. 此接口用于查询该会话用户作为房型发布者发布的房型的审核情况。 
2. 查询提交的起始时间至今的增量房型记录：start_modified：2011-7-1 16:00:00 
3. 返回数据结果为发布房型时间正序排列 */
type HotelSoldTypesIncrementGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 【不推荐使用，现在总是返回从修改开始时间到目前为止的所有记录，与修改结束时间不再相关】查询修改结束时间，必须大于修改开始时间（修改时间跨度不能大于1天）。格式：yyyy-MM-dd HH:mm:ss。 */
func (r *HotelSoldTypesIncrementGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。 */
func (r *HotelSoldTypesIncrementGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页面大小，取值范围1-100，默认大小20。 */
func (r *HotelSoldTypesIncrementGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询修改开始时间（修改时间跨度不能大于1天）。格式：yyyy-MM-dd HH:mm:ss */
func (r *HotelSoldTypesIncrementGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}

/* 【不推荐使用，现在返回结果总会包含总记录数和是否存在下一页】是否使用has_next的分页方式，如果指定true，则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的字段，效率比总记录数高 */
func (r *HotelSoldTypesIncrementGetRequest) SetUseHasNext(value string) {
	r.SetValue("use_has_next", value)
}


func (r *HotelSoldTypesIncrementGetRequest) GetResponse(accessToken string) (*HotelSoldTypesIncrementGetResponse, []byte, error) {
	var resp HotelSoldTypesIncrementGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.sold.types.increment.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelSoldTypesIncrementGetResponse struct {
	HasNext bool `json:"has_next"`
	RoomTypes []*RoomType `json:"room_types"`
	TotalResults int `json:"total_results"`
}

type HotelSoldTypesIncrementGetResponseResult struct {
	Response *HotelSoldTypesIncrementGetResponse `json:"hotel_sold_types_increment_get_response"`
}





/* 此接口用于发布一个房型，房型的发布者是当前会话的用户。如果该房型在淘宝集市酒店下已经存在，则会返回错误提示。 
该接口发布的是一个新增房型申请，需要淘宝小二审核 */
type HotelTypeAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 酒店id。必须为数字 */
func (r *HotelTypeAddRequest) SetHid(value string) {
	r.SetValue("hid", value)
}

/* 房型名称。长度不能超过30 */
func (r *HotelTypeAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 接入卖家数据主键,格式为“接入方酒店id-接入方房型id” */
func (r *HotelTypeAddRequest) SetSiteParam(value string) {
	r.SetValue("site_param", value)
}


func (r *HotelTypeAddRequest) GetResponse(accessToken string) (*HotelTypeAddResponse, []byte, error) {
	var resp HotelTypeAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.type.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelTypeAddResponse struct {
	RoomType *RoomType `json:"room_type"`
}

type HotelTypeAddResponseResult struct {
	Response *HotelTypeAddResponse `json:"hotel_type_add_response"`
}





/* 此接口用于查询一个房型，根据传入的酒店hid，房型名称/别名查询房型信息。 */
type HotelTypeNameGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 要查询的酒店id。必须为数字 */
func (r *HotelTypeNameGetRequest) SetHid(value string) {
	r.SetValue("hid", value)
}

/* 房型全部名称/别名。不能超过60字节 */
func (r *HotelTypeNameGetRequest) SetName(value string) {
	r.SetValue("name", value)
}


func (r *HotelTypeNameGetRequest) GetResponse(accessToken string) (*HotelTypeNameGetResponse, []byte, error) {
	var resp HotelTypeNameGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.type.name.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelTypeNameGetResponse struct {
	RoomType *RoomType `json:"room_type"`
}

type HotelTypeNameGetResponseResult struct {
	Response *HotelTypeNameGetResponse `json:"hotel_type_name_get_response"`
}





/* 此接口用于更新一个酒店的信息，根据用户传入的hid更新酒店数据。如果该酒店在淘宝集市酒店不存在，则会返回错误提示。 
该接口发出的是一个更新酒店信息的申请，需要淘宝小二审核。 */
type HotelUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 酒店地址。长度不能超过120 */
func (r *HotelUpdateRequest) SetAddress(value string) {
	r.SetValue("address", value)
}

/* 城市编码。参见：http://kezhan.trip.taobao.com/area.html，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取 */
func (r *HotelUpdateRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* domestic为true时，固定China； 
domestic为false时，传海外国家编码。参见：http://kezhan.trip.taobao.com/countrys.html */
func (r *HotelUpdateRequest) SetCountry(value string) {
	r.SetValue("country", value)
}

/* 装修时间。长度不能超过4。 */
func (r *HotelUpdateRequest) SetDecorateTime(value string) {
	r.SetValue("decorate_time", value)
}

/* 酒店介绍。不超过25000个汉字 */
func (r *HotelUpdateRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 区域（县级市）编码。参见：http://kezhan.trip.taobao.com/area.html */
func (r *HotelUpdateRequest) SetDistrict(value string) {
	r.SetValue("district", value)
}

/* 是否国内酒店。可选值：true，false */
func (r *HotelUpdateRequest) SetDomestic(value string) {
	r.SetValue("domestic", value)
}

/* 酒店id。必须为数字。 */
func (r *HotelUpdateRequest) SetHid(value string) {
	r.SetValue("hid", value)
}

/* 酒店级别。可选值：A,B,C,D,E,F。代表客栈公寓、经济连锁、二星级/以下、三星级/舒适、四星级/高档、五星级/豪华 */
func (r *HotelUpdateRequest) SetLevel(value string) {
	r.SetValue("level", value)
}

/* 酒店名称。不能超过60字节 */
func (r *HotelUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开业时间。长度不能超过4。 */
func (r *HotelUpdateRequest) SetOpeningTime(value string) {
	r.SetValue("opening_time", value)
}

/* 酒店定位。可选值：T,B。代表旅游度假、商务出行 */
func (r *HotelUpdateRequest) SetOrientation(value string) {
	r.SetValue("orientation", value)
}

/* 酒店图片。类型:JPG,GIF;最大长度:500K。支持的文件类型：gif,jpg,jpeg,png。 
图片没有传，则代表不更新图片，使用原来的图片 */
func (r *HotelUpdateRequest) SetPic(value string) {
	r.SetValue("pic", value)
}

/* 省份编码。参见：http://kezhan.trip.taobao.com/area.html，domestic为false时默认为0 */
func (r *HotelUpdateRequest) SetProvince(value string) {
	r.SetValue("province", value)
}

/* 房间数。长度不能超过4。 */
func (r *HotelUpdateRequest) SetRooms(value string) {
	r.SetValue("rooms", value)
}

/* 交通距离与设施服务。JSON格式。cityCenterDistance、railwayDistance、airportDistance分别代表距离市中心、距离火车站、距离机场公里数，为不超过3位正整数，默认-1代表无数据。 
其他key值true代表有此服务，false代表没有。 
parking：停车场，airportShuttle：机场接送，rentCar：租车，meetingRoom：会议室，businessCenter：商务中心，swimmingPool：游泳池，fitnessClub：健身中心，laundry：洗衣服务，morningCall：叫早服务，bankCard：接受银行卡，creditCard：接受信用卡，chineseRestaurant：中餐厅，westernRestaurant：西餐厅，cafe：咖啡厅，bar：酒吧，ktv：KTV。 
如： 
{"airportShuttle":true,"parking":false,"fitnessClub":false,"chineseRestaurant":false,"rentCar":false,"laundry":false,"bankCard":false,"cityCenterDistance":-1,"creditCard":false,"westernRestaurant":false,"ktv":false,"railwayDistance":-1,"swimmingPool":false,"cafe":false,"businessCenter":false,"morningCall":false,"bar":false,"meetingRoom":false,"airportDistance":-1} */
func (r *HotelUpdateRequest) SetService(value string) {
	r.SetValue("service", value)
}

/* 楼层数。长度不能超过4。 */
func (r *HotelUpdateRequest) SetStoreys(value string) {
	r.SetValue("storeys", value)
}

/* 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886 */
func (r *HotelUpdateRequest) SetTel(value string) {
	r.SetValue("tel", value)
}


func (r *HotelUpdateRequest) GetResponse(accessToken string) (*HotelUpdateResponse, []byte, error) {
	var resp HotelUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotel.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelUpdateResponse struct {
	Hotel *Hotel `json:"hotel"`
}

type HotelUpdateResponseResult struct {
	Response *HotelUpdateResponse `json:"hotel_update_response"`
}





/* 此接口用于查询多个酒店，根据传入的参数查询酒店信息。 */
type HotelsSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 城市编码。参见：http://kezhan.trip.taobao.com/area.html。 
domestic为true时，province,city,district不能同时为空或为0 */
func (r *HotelsSearchRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* domestic为true时，固定China； 
domestic为false时，必须传定义的海外国家编码值，是必填项。参见：http://kezhan.trip.taobao.com/countrys.html */
func (r *HotelsSearchRequest) SetCountry(value string) {
	r.SetValue("country", value)
}

/* 区域（县级市）编码。参见：http://kezhan.trip.taobao.com/area.html。 
domestic为true时，province,city,district不能同时为空或为0 */
func (r *HotelsSearchRequest) SetDistrict(value string) {
	r.SetValue("district", value)
}

/* 是否国内酒店。可选值：true，false */
func (r *HotelsSearchRequest) SetDomestic(value string) {
	r.SetValue("domestic", value)
}

/* 酒店名称。不能超过60字节 */
func (r *HotelsSearchRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20 */
func (r *HotelsSearchRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 省份编码。参见：http://kezhan.trip.taobao.com/area.html。 
domestic为true时，province,city,district不能同时为空或为0 */
func (r *HotelsSearchRequest) SetProvince(value string) {
	r.SetValue("province", value)
}


func (r *HotelsSearchRequest) GetResponse(accessToken string) (*HotelsSearchResponse, []byte, error) {
	var resp HotelsSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hotels.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HotelsSearchResponse struct {
	Hotels []*Hotel `json:"hotels"`
	TotalResults int `json:"total_results"`
}

type HotelsSearchResponseResult struct {
	Response *HotelsSearchResponse `json:"hotels_search_response"`
}



