// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/yaofangou/open_taobao"
)





/* 查询买家信息API，只能买家类应用调用。 */
type UserBuyerGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 只返回nick, avatar参数 */
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



