// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package notice

import (
	"github.com/changkong/open_taobao"
)





/* 获取一个appkey的哪些用户丢失了消息 */
type CometDiscardinfoGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 指定截止日志，如果不传则为服务端当前时间 */
func (r *CometDiscardinfoGetRequest) SetEnd(value string) {
	r.SetValue("end", value)
}

/* 用户nick */
func (r *CometDiscardinfoGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 指定从那个时间开始找丢弃的消息 */
func (r *CometDiscardinfoGetRequest) SetStart(value string) {
	r.SetValue("start", value)
}

/* 指定多个消息类型 */
func (r *CometDiscardinfoGetRequest) SetTypes(value string) {
	r.SetValue("types", value)
}

/* 指定查看那个用户的丢弃消息 */
func (r *CometDiscardinfoGetRequest) SetUserId(value string) {
	r.SetValue("user_id", value)
}


func (r *CometDiscardinfoGetRequest) GetResponse(accessToken string) (*CometDiscardinfoGetResponse, []byte, error) {
	var resp CometDiscardinfoGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.comet.discardinfo.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CometDiscardinfoGetResponse struct {
	DiscardInfoList []*DiscardInfo `json:"discard_info_list"`
}

type CometDiscardinfoGetResponseResult struct {
	Response *CometDiscardinfoGetResponse `json:"comet_discardinfo_get_response"`
}





/* 通用的用于获取用户授权的消息数据 */
type IncrementAuthorizeMessageGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 消息的结束时间，消息所对应的操作时间的最大值。和start_modified搭配使用能过滤消通知消息的时间段。不传时：如果设置了start_modified，默认为与start_modified同一天的23:59:59；否则默认为调用接口当天的23:59:59。（格式：yyyy-MM-dd HH:mm:ss） 
注意：start_modified和end_modified的日期必须在必须在同一天内，比如：start_modified设置2000-01-01 00:00:00，则end_modified必须设置为2000-01-01这个日期 */
func (r *IncrementAuthorizeMessageGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 消息所属的用户nick，他用型应用必传，自用型不传 */
func (r *IncrementAuthorizeMessageGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码，取值范围:大于零的整数; 默认值:1,即返回第一页数据。 */
func (r *IncrementAuthorizeMessageGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数，取值范围:大于零的整数;最大值:200;默认值:40 */
func (r *IncrementAuthorizeMessageGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 消息的开始时间，消息所对应的操作时间的最小值和end_modified搭配使用能过滤消通知消息的时间段。不传时：如果设置了end_modified，默认为与 end_modified同一天的00:00:00，否则默认为调用接口当天的00:00:00。（格式：yyyy-MM-dd HH:mm:ss）最早可取6天内的数据。
注意：start_modified和end_modified的日期必须在必须在同一天内，比如：start_modified设置2000-01-01 00:00:00，则end_modified必须设置为2000-01-01这个日期 */
func (r *IncrementAuthorizeMessageGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}

/* 指定消息的状态，比如：ItemAdd，如果不填次参数，默认查询所有状态的数据， */
func (r *IncrementAuthorizeMessageGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 消息类型，比如：item */
func (r *IncrementAuthorizeMessageGetRequest) SetTopic(value string) {
	r.SetValue("topic", value)
}


func (r *IncrementAuthorizeMessageGetRequest) GetResponse(accessToken string) (*IncrementAuthorizeMessageGetResponse, []byte, error) {
	var resp IncrementAuthorizeMessageGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.increment.authorize.message.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type IncrementAuthorizeMessageGetResponse struct {
	HasNext bool `json:"has_next"`
	Messages []string `json:"messages"`
}

type IncrementAuthorizeMessageGetResponseResult struct {
	Response *IncrementAuthorizeMessageGetResponse `json:"increment_authorize_message_get_response"`
}





/* 提供app为自己的用户开通增量消息服务功能 */
type IncrementCustomerPermitRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 表示与topics相对应的消息状态。各个消息主题间用";"分隔，各个状态间用","分隔，消息主题必须和topics一致。如果为all时，表示开通应用订阅的所有的消息。 
如果不填写,会默认开通应用所订阅的所有消息。 */
func (r *IncrementCustomerPermitRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 应用为用户开通的主动通知的消息主题（或消息类别），该主题必须是应用订阅的主题。如果应用未订阅该主题，则系统会自动过滤掉该主题。各个主题间用";"分隔。 如果不填写，会采用应用所订阅的消息，下次应用消息修改时此用户的消息会自动修改。 如果填写了消息类型，在应用修改增加消息类型时，用户的消息类型是不会随着增加相应类型，除非通过重新permit增加。 */
func (r *IncrementCustomerPermitRequest) SetTopics(value string) {
	r.SetValue("topics", value)
}

/* 用户需要开通的功能。值可为get,notify和syn分别表示增量api取消息，主动发送消息和同步数据功能。这三个值可任意无序组合。应用在为用户开通相应功能前需先订阅相应的功能；get和notify在主动通知界面中订阅，syn在mysql数据同步步界面中订阅。 
当重复开通时，只会在已经开通的功能上，开通新的功能，不会覆盖旧的开通。例如：在开通get功能后，再次开通notify功能的结果是开通了get和notify功能；而不会只开通notify功能。 
在开通时，type里面的参数会根据应用订阅的类型进行相应的过虑。如应用只订阅主动通知，则默认值过滤后为get,notify，如果应用只订阅数据同步服务，默认值过滤后为syn。 */
func (r *IncrementCustomerPermitRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *IncrementCustomerPermitRequest) GetResponse(accessToken string) (*IncrementCustomerPermitResponse, []byte, error) {
	var resp IncrementCustomerPermitResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.increment.customer.permit", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type IncrementCustomerPermitResponse struct {
	AppCustomer *AppCustomer `json:"app_customer"`
}

type IncrementCustomerPermitResponseResult struct {
	Response *IncrementCustomerPermitResponse `json:"increment_customer_permit_response"`
}





/* 供应用关闭其用户的增量消息服务功能，这样可以节省ISV的流量。 */
type IncrementCustomerStopRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 应用要关闭增量消息服务的用户昵称 */
func (r *IncrementCustomerStopRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 应用需要关闭用户的功能。取值可为get,notify和syn分别表示增量api取消息，主动发送消息和同步数据功能。用户关闭相应功能前,需应用已为用户经开通了相应的功能。这三个参数可无序任意组合。在关闭时，type里面的参数会根据应用订阅的类型进行相应的过虑。如应用只订阅主动通知，则默认值过滤后为get,notify，如果应用只订阅数据同步服务，默认值过滤后为syn。 */
func (r *IncrementCustomerStopRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *IncrementCustomerStopRequest) GetResponse(accessToken string) (*IncrementCustomerStopResponse, []byte, error) {
	var resp IncrementCustomerStopResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.increment.customer.stop", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type IncrementCustomerStopResponse struct {
	IsSuccess bool `json:"is_success"`
}

type IncrementCustomerStopResponseResult struct {
	Response *IncrementCustomerStopResponse `json:"increment_customer_stop_response"`
}





/* 提供查询应用为自身用户所开通的增量消息服务信息。这个接口有性能问题,服务端有流量限制。在丢消息后，不要每次都调用这个api获取用户列表，应用可自己保存一份主动通知用户名单。 */
type IncrementCustomersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需要返回的字段。可填写的字段参见AppCustomer中的返回字段。如：nick,created,status,type,subscriptions。 */
func (r *IncrementCustomersGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 查询用户的昵称。当为空时通过分页方式查询appkey开通的所有用户,最多填入20个昵称。 */
func (r *IncrementCustomersGetRequest) SetNicks(value string) {
	r.SetValue("nicks", value)
}

/* 分页查询时，查询的页码。此参数只有nicks为空时起作用。 */
func (r *IncrementCustomersGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分布查询时，页的大小。此参数只有当nicks为空时起作用。 */
func (r *IncrementCustomersGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询用户开通的功能。值可为get,notify和syn分别表示增量api取消息，主动发送消息和同步数据功能。这三个值不分次序。在查询时，type里面的参数会根据应用订阅的类型进行相应的过虑。如应用只订阅主动通知，则默认值过滤后为get,notify，如果应用只订阅数据同步服务，默认值过滤后为syn。 */
func (r *IncrementCustomersGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *IncrementCustomersGetRequest) GetResponse(accessToken string) (*IncrementCustomersGetResponse, []byte, error) {
	var resp IncrementCustomersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.increment.customers.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type IncrementCustomersGetResponse struct {
	AppCustomers []*AppCustomer `json:"app_customers"`
	TotalResults int `json:"total_results"`
}

type IncrementCustomersGetResponseResult struct {
	Response *IncrementCustomersGetResponse `json:"increment_customers_get_response"`
}





/* 开通主动通知业务的APP可以通过该接口获取商品变更通知信息 
<font color="red">建议获取增量消息的时间间隔是：半个小时</font> */
type IncrementItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 消息所对应的操作时间的最大值。和start_modified搭配使用能过滤消通知消息的时间段。不传时：如果设置了start_modified，默认为与start_modified同一天的23:59:59；否则默认为调用接口当天的23:59:59。（格式：yyyy-MM-dd HH:mm:ss）<br>
<font color="red">注意：start_modified和end_modified的日期必须在必须在同一天内，比如：start_modified设置2000-01-01 00:00:00，则end_modified必须设置为2000-01-01这个日期</font> */
func (r *IncrementItemsGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 消息所属于的用户的昵称。设置此参数，返回的消息会根据传入nick的进行过滤。自用型AppKey的昵称默认为自己的绑定昵称，此参数无效。 */
func (r *IncrementItemsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1,即返回第一页数据。 */
func (r *IncrementItemsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数;最大值:200;默认值:40。 */
func (r *IncrementItemsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 消息所对应的操作时间的最小值和end_modified搭配使用能过滤消通知消息的时间段。不传时：如果设置了end_modified，默认为与 end_modified同一天的00:00:00，否则默认为调用接口当天的00:00:00。（格式：yyyy-MM-dd HH:mm:ss）<br> 
最早可取6天内的数据。 
<font color="red">注意：start_modified和end_modified的日期必须在必须在同一天内，比如：start_modified设置2000-01-01 00:00:00，则end_modified必须设置为2000-01-01这个日期</font> */
func (r *IncrementItemsGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}

/* 商品操作状态，默认查询所有状态的数据，除了默认值外，每次可查询多种状态，每种状态间用英语逗号分隔。具体类型列表见： 
ItemAdd（新增商品）  
ItemUpshelf（上架商品，自动上架商品不能获取到增量信息）  
ItemDownshelf（下架商品）  
ItemDelete（删除商品）  
ItemUpdate（更新商品）  
ItemRecommendDelete（取消橱窗推荐商品）  
ItemRecommendAdd（橱窗推荐商品）  
ItemZeroStock（商品卖空） 
ItemPunishDelete（小二删除商品） 
ItemPunishDownshelf（小二下架商品） 
ItemPunishCc（小二CC商品） 
ItemSkuZeroStock（商品SKU卖空） 
ItemStockChanged（修改商品库存） */
func (r *IncrementItemsGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *IncrementItemsGetRequest) GetResponse(accessToken string) (*IncrementItemsGetResponse, []byte, error) {
	var resp IncrementItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.increment.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type IncrementItemsGetResponse struct {
	NotifyItems []*NotifyItem `json:"notify_items"`
	TotalResults int `json:"total_results"`
}

type IncrementItemsGetResponseResult struct {
	Response *IncrementItemsGetResponse `json:"increment_items_get_response"`
}





/* 开通主动通知业务的APP可以通过该接口获取用户的退款变更通知信息 
<font color="red">建议在获取增量消息的时间间隔是：半个小时</font> */
type IncrementRefundsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 消息所对应的操作时间的最大值。和start_modified搭配使用能过滤消通知消息的时间段。不传时：如果设置了start_modified，默认为与start_modified同一天的23:59:59；否则默认为调用接口当天的23:59:59。（格式：yyyy-MM-dd HH:mm:ss）<br>
<font color="red">注意：start_modified和end_modified的日期必须在必须在同一天内，比如：start_modified设置2000-01-01 00:00:00，则end_modified必须设置为2000-01-01这个日期</font> */
func (r *IncrementRefundsGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 消息所属于的用户的昵称。设置此参数，返回的消息会根据传入nick的进行过滤。自用型appKey的昵称默认为自己的绑定昵称，此参数无效。 */
func (r *IncrementRefundsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1,即返回第一页数据。 */
func (r *IncrementRefundsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数;最大值:200;默认值:40。 */
func (r *IncrementRefundsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 消息所对应的操作时间的最小值。和end_modified搭配使用能过滤消通知消息的时间段。不传时：如果设置了end_modified，默认为与 end_modified同一天的00:00:00，否则默认为调用接口当天的00:00:00。（格式：yyyy-MM-dd HH:mm:ss）<br>
<font color="red">注意：start_modified和end_modified的日期必须在必须在同一天内，比如：start_modified设置2000-01-01 00:00:00，则end_modified必须设置为2000-01-01这个日期</font> */
func (r *IncrementRefundsGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}

/* 退款操作状态，默认查询所有状态的数据，除了默认值外每次只能查询一种状态。具体字段说明请见 退款消息状态 */
func (r *IncrementRefundsGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *IncrementRefundsGetRequest) GetResponse(accessToken string) (*IncrementRefundsGetResponse, []byte, error) {
	var resp IncrementRefundsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.increment.refunds.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type IncrementRefundsGetResponse struct {
	NotifyRefunds []*NotifyRefund `json:"notify_refunds"`
	TotalResults int `json:"total_results"`
}

type IncrementRefundsGetResponseResult struct {
	Response *IncrementRefundsGetResponse `json:"increment_refunds_get_response"`
}





/* 开通主动通知业务的APP可以通过该接口获取用户的交易和评价变更通知信息 
<font color="red">建议在获取增量消息的时间间隔是：半个小时</font> */
type IncrementTradesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 消息所对应的操作时间的最大值。和start_modified搭配使用能过滤消通知消息的时间段。不传时：如果设置了start_modified，默认为与start_modified同一天的23:59:59；否则默认为调用接口当天的23:59:59。（格式：yyyy-MM-dd HH:mm:ss）<br>
<font color="red">注意：start_modified和end_modified的日期必须在必须在同一天内，比如：start_modified设置2000-01-01 00:00:00，则end_modified必须设置为2000-01-01这个日期</font> */
func (r *IncrementTradesGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 消息所属于的用户的昵称。设置此参数，返回的消息会根据传入nick的进行过滤。自用型appKey的昵称默认为自己的绑定昵称，此参数无效。 */
func (r *IncrementTradesGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1,即返回第一页数据。 */
func (r *IncrementTradesGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数;最大值:200;默认值:40。 */
func (r *IncrementTradesGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 消息所对应的操作时间的最小值。和end_modified搭配使用能过滤消通知消息的时间段。不传时：如果设置了end_modified，默认为与end_modified同一天的00:00:00，否则默认为调用接口当天的00:00:00。（格式：yyyy-MM-dd HH:mm:ss）<br>
<font color="red">注意：start_modified和end_modified的日期必须在必须在同一天内，比如：start_modified设置2000-01-01 00:00:00，则end_modified必须设置为2000-01-01这个日期</font> */
func (r *IncrementTradesGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}

/* 交易或评价的状态，默认查询所有状态的数据，除了默认值外每次只能查询一种状态。除了交易超时提醒消息没有type分类以外，交易的其他消息都有type分类。 
可选值  
TradeCreate（创建交易）  
TradeModifyFee（修改交易费用）  
TradeCloseAndModifyDetailOrder（关闭或修改子订单）  
TradeClose（关闭交易）  
TradeBuyerPay（买家付款）  
TradeSellerShip（卖家发货）  
TradeDelayConfirmPay（延长收货时间）  
TradePartlyRefund（子订单退款成功）  
TradePartlyConfirmPay（子订单打款成功）  
TradeSuccess（交易成功）  
TradeTimeoutRemind（交易超时提醒）  
TradeRated（交易评价变更）  
TradeMemoModified（交易备注修改）  
TradeLogisticsAddressChanged（修改交易收货地址）  
TradeChanged（修改订单信息（SKU等）） */
func (r *IncrementTradesGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 交易所属的类型，默认查询所有类型的数据，除了默认值外每次只能查询一种类型。交易超时提醒类型的消息没有消息类型，固定返回“timeout”。 
可选值： 
ec（直冲）  
guarantee_trade（一口价、拍卖）  
auto_delivery（自动发货）  
cod（货到付款）  
independent_shop_trade（旺店标准版交易）  
independent_simple_trade（旺店入门版交易）  
shopex_trade（ShopEX版)  
netcn_trade（淘宝与万网合作版网）  
travel（旅游产品交易）  
fenxiao（分销平台交易）  
game_equipment（网游虚拟交易） */
func (r *IncrementTradesGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *IncrementTradesGetRequest) GetResponse(accessToken string) (*IncrementTradesGetResponse, []byte, error) {
	var resp IncrementTradesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.increment.trades.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type IncrementTradesGetResponse struct {
	NotifyTrades []*NotifyTrade `json:"notify_trades"`
	TotalResults int `json:"total_results"`
}

type IncrementTradesGetResponseResult struct {
	Response *IncrementTradesGetResponse `json:"increment_trades_get_response"`
}





/* 异步获取增量消息数据 */
type TopatsIncrementMessagesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 消息结束时间，格式：yyyy-MM-dd HH:mm:ss，其中start < end < 现在，并且start和end在同一天。 */
func (r *TopatsIncrementMessagesGetRequest) SetEnd(value string) {
	r.SetValue("end", value)
}

/* 消息开始时间，格式：yyyy-MM-dd HH:mm:ss，其中start >= 前天零点 */
func (r *TopatsIncrementMessagesGetRequest) SetStart(value string) {
	r.SetValue("start", value)
}

/* 消息类型，多个类型之间用半角逗号分隔，可选值为：item,trade,refund。 */
func (r *TopatsIncrementMessagesGetRequest) SetTopics(value string) {
	r.SetValue("topics", value)
}


func (r *TopatsIncrementMessagesGetRequest) GetResponse(accessToken string) (*TopatsIncrementMessagesGetResponse, []byte, error) {
	var resp TopatsIncrementMessagesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.topats.increment.messages.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TopatsIncrementMessagesGetResponse struct {
	Task *Task `json:"task"`
}

type TopatsIncrementMessagesGetResponseResult struct {
	Response *TopatsIncrementMessagesGetResponse `json:"topats_increment_messages_get_response"`
}



