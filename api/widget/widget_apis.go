// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package widget

import (
	"github.com/changkong/open_taobao"
)

/* app指定api名称获取此api的http调用方法、app是否有请求权限、是否需要授权等信息。仅支持widget入口调用。调用入口为/widget。签名方法简化为Hmac-md5,hmac(secret+‘app_key' ＋app_key +'timestamp' + timestamp+secret,secret)。timestamp为60分钟内有效 */
type WidgetAppapiruleCheckRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 要判断权限的api名称，如果指定的api不存在，报错invalid method */
func (r *WidgetAppapiruleCheckRequest) SetApiName(value string) {
	r.SetValue("api_name", value)
}

func (r *WidgetAppapiruleCheckRequest) GetResponse(accessToken string) (*WidgetAppapiruleCheckResponse, []byte, error) {
	var resp WidgetAppapiruleCheckResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.widget.appapirule.check", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WidgetAppapiruleCheckResponse struct {
	CallPermission string `json:"call_permission"`
	HttpMethod     string `json:"http_method"`
	NeedAuth       string `json:"need_auth"`
}

type WidgetAppapiruleCheckResponseResult struct {
	Response *WidgetAppapiruleCheckResponse `json:"widget_appapirule_check_response"`
}

/* 根据商品id获取sku选择面板需要的信息。session必传且用户当前浏览器必需已经在淘宝登陆，具体判断方法可以调用taobao.widget.loginstatus.get进行判断。会根据session生成购买链接。仅支持widget入口调用。调用入口为/widget。签名方法简化为Hmac-md5,hmac(secret+‘app_key' ＋app_key +'timestamp' + timestamp+secret, secret)。timestamp为60分钟内有效 */
type WidgetItempanelGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 参数fields为选填参数，表示需要返回哪些字段，默认为空：表示所有字段都返回。指定item_id返回item_id; 指定title返回title; 指定click_url返回click_url(如果此商品有淘宝客会默认返回转换过的淘宝客连接，绑定用户为appkey对应的用户); 指定price返回price(商品价格，如果有多个sku返回的是sku的价格区间); 指定quantify返回quantity(商品总数); 指定pic_url返回pic_url(商品主图地址); 指定item_pics返回item_pics(商品图片列表); 指定skus返回skus和sku_props组合; 指定shop_promotion_data返回shop_promotion_data(商品所属的店铺优惠信息); 指定item_promotion_data返回item_promotion_data(商品的优惠信息); 指定seller_nick返回seller_nick(卖家昵称); 指定is_mall返回is_mall(是否商城商品，true表示是商城商品);add_url不可选一定会返回 */
func (r *WidgetItempanelGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 要查询的商品的数字id，等同于Item的num_iid */
func (r *WidgetItempanelGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

func (r *WidgetItempanelGetRequest) GetResponse(accessToken string) (*WidgetItempanelGetResponse, []byte, error) {
	var resp WidgetItempanelGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.widget.itempanel.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WidgetItempanelGetResponse struct {
	AddUrl string      `json:"add_url"`
	Item   *WidgetItem `json:"item"`
}

type WidgetItempanelGetResponseResult struct {
	Response *WidgetItempanelGetResponse `json:"widget_itempanel_get_response"`
}

/* 获取当前浏览器用户在淘宝登陆状态。如果传了session并且此用户已在淘宝登陆，返回nick和userId。仅支持widget入口调用。调用入口为/widget。签名方法简化为Hmac-md5,hmac(secret+‘app_key' ＋app_key +'timestamp' + timestamp+secret, secret)。timestamp为60分钟内有效 */
type WidgetLoginstatusGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 指定判断当前浏览器登陆用户是否此昵称用户，并且返回是否登陆。如果用户不一致返回未登陆，如果用户一致且已登录返回已登陆 */
func (r *WidgetLoginstatusGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

func (r *WidgetLoginstatusGetRequest) GetResponse(accessToken string) (*WidgetLoginstatusGetResponse, []byte, error) {
	var resp WidgetLoginstatusGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.widget.loginstatus.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WidgetLoginstatusGetResponse struct {
	IsLogin bool   `json:"is_login"`
	Nick    string `json:"nick"`
	UserId  string `json:"user_id"`
}

type WidgetLoginstatusGetResponseResult struct {
	Response *WidgetLoginstatusGetResponse `json:"widget_loginstatus_get_response"`
}
