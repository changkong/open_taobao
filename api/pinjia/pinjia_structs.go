// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pinjia

const VersionNo = "20130729"

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

/* 评价大印象返回的印象词接口 */
type ImprItemDO struct {
	AttributeId int    `json:"attribute_id"`
	Count       int    `json:"count"`
	Status      int    `json:"status"`
	Title       string `json:"title"`
}

/* 单条交易子订单语义标签对象 */
type ImprFeedIdDO struct {
	FeedInfoList []*ImprFeedInfoDO `json:"feed_info_list"`
	Nick         string            `json:"nick"`
	Rate         int               `json:"rate"`
}

/* 大家印象标签对应评价信息对象 */
type ImprFeedInfoDO struct {
	BizType     int      `json:"biz_type"`
	Feedback    string   `json:"feedback"`
	GmtCreate   string   `json:"gmt_create"`
	GmtModified string   `json:"gmt_modified"`
	ImprWords   []string `json:"impr_words"`
}
