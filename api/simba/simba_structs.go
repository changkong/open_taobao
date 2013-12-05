// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package simba

const VersionNo = "20131127"


/* 一页推荐词列表 */
type RecommendWordPage struct {
	PageNo int `json:"page_no"`
	PageSize int `json:"page_size"`
	RecommendWordList *RecommendWordListObject `json:"recommend_word_list"`
	TotalItem int `json:"total_item"`

}

/* 一页推荐词列表 */
type RecommendWordListObject struct {
	RecommendWord []*RecommendWord `json:"recommend_word"`

}

/* 推荐词 */
type RecommendWord struct {
	AveragePrice string `json:"average_price"`
	Pertinence string `json:"pertinence"`
	Pv string `json:"pv"`
	Word string `json:"word"`

}

/*  */
type INWordBaseListObject struct {
	INWordBase []*INWordBase `json:"i_n_word_base"`

}

/* 词基础数据对象 */
type INWordBase struct {
	InRecordBaseList *INRecordBaseListObject `json:"in_record_base_list"`
	Word string `json:"word"`

}

/* 词基础数据对象 */
type INRecordBaseListObject struct {
	INRecordBase []*INRecordBase `json:"i_n_record_base"`

}

/* 词基础数据对象 */
type INRecordBase struct {
	AvgPrice int `json:"avg_price"`
	Click int `json:"click"`
	Competition int `json:"competition"`
	Ctr float64 `json:"ctr"`
	Date string `json:"date"`
	Pv int `json:"pv"`

}

/*  */
type KeywordListObject struct {
	Keyword []*Keyword `json:"keyword"`

}

/* 关键词 */
type Keyword struct {
	AdgroupId int `json:"adgroup_id"`
	AuditDesc string `json:"audit_desc"`
	AuditStatus string `json:"audit_status"`
	CampaignId int `json:"campaign_id"`
	CreateTime string `json:"create_time"`
	IsDefaultPrice bool `json:"is_default_price"`
	IsGarbage bool `json:"is_garbage"`
	KeywordId int `json:"keyword_id"`
	MatchScope string `json:"match_scope"`
	MaxPrice int `json:"max_price"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`
	Qscore string `json:"qscore"`
	Word string `json:"word"`

}

