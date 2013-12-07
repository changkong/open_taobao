// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package favorite

const VersionNo = "20131207"


/*  */
type CollectItemListObject struct {
	CollectItem []*CollectItem `json:"collect_item"`

}

/* 商品或店铺的信息 */
type CollectItem struct {
	ItemNumid int `json:"item_numid"`
	ItemOwnerNick string `json:"item_owner_nick"`
	Title string `json:"title"`

}

