// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/changkong/open_taobao"
)

/* 查询买家信息API，只能买家类应用调用。 */
type UserBuyerGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 只返回user_id,nick,sex,buyer_credit,avatar,has_shop,vip_info参数 */
func (r *UserBuyerGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

func (r *UserBuyerGetRequest) GetResponse(accessToken string) (*UserBuyerGetResponse, []byte, error) {
	var resp UserBuyerGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.user.buyer.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UserBuyerGetResponse struct {
	User *User `json:"user"`
}

type UserBuyerGetResponseResult struct {
	Response *UserBuyerGetResponse `json:"user_buyer_get_response"`
}

/* 得到单个用户 */
type UserGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表。可选值：User结构体中的所有字段；以半角逗号(,)分隔。需要用户授权才能获取用户对应的uid和user_id。 */
func (r *UserGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 用户昵称，如果昵称为中文，请使用UTF-8字符集对昵称进行URL编码。
<br><font color="red">注：在传入session的情况下,可以不传nick，表示取当前用户信息；否则nick必须传.<br>
自用型应用不需要传入nick
</font> */
func (r *UserGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

func (r *UserGetRequest) GetResponse(accessToken string) (*UserGetResponse, []byte, error) {
	var resp UserGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.user.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UserGetResponse struct {
	User *User `json:"user"`
}

type UserGetResponseResult struct {
	Response *UserGetResponse `json:"user_get_response"`
}

/* 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。 */
type UserSellerGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 只返回user_id,nick,sex,seller_credit,type,has_more_pic,item_img_num,item_img_size,prop_img_num,prop_img_size,auto_repost,promoted_type,status,alipay_bind,consumer_protection,avatar,liangpin,sign_food_seller_promise,has_shop,is_lightning_consignment,has_sub_stock,is_golden_seller,vip_info,magazine_subscribe,vertical_market,online_gaming参数 */
func (r *UserSellerGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

func (r *UserSellerGetRequest) GetResponse(accessToken string) (*UserSellerGetResponse, []byte, error) {
	var resp UserSellerGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.user.seller.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UserSellerGetResponse struct {
	User *User `json:"user"`
}

type UserSellerGetResponseResult struct {
	Response *UserSellerGetResponse `json:"user_seller_get_response"`
}

/* 传入多个淘宝会员帐号返回多个用户公开信息 */
type UsersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 查询字段：User数据结构的公开信息字段列表，以半角逗号(,)分隔 */
func (r *UsersGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 用户昵称，多个以半角逗号(,)分隔，最多40个 */
func (r *UsersGetRequest) SetNicks(value string) {
	r.SetValue("nicks", value)
}

func (r *UsersGetRequest) GetResponse(accessToken string) (*UsersGetResponse, []byte, error) {
	var resp UsersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.users.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UsersGetResponse struct {
	Users []*User `json:"users"`
}

type UsersGetResponseResult struct {
	Response *UsersGetResponse `json:"users_get_response"`
}
