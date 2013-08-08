// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pinjia

const VersionNo = "20130808"

/* 评价列表 */
type TradeRate struct {
	Content    string  `json:"content"`
	Created    string  `json:"created"`
	ItemPrice  float64 `json:"item_price"`
	ItemTitle  string  `json:"item_title"`
	Nick       string  `json:"nick"`
	NumIid     int     `json:"num_iid"`
	Oid        int     `json:"oid"`
	RatedNick  string  `json:"rated_nick"`
	Reply      string  `json:"reply"`
	Result     string  `json:"result"`
	Role       string  `json:"role"`
	Tid        int     `json:"tid"`
	ValidScore bool    `json:"valid_score"`
}
