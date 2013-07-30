// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wangwang

const VersionNo = "20130729"

/* 关键词列表 */
type WordList struct {
	Word string `json:"word"`
}

/* 聊天记录列表 */
type MsgList struct {
	Content   string           `json:"content"`
	Direction int              `json:"direction"`
	Length    int              `json:"length"`
	Time      string           `json:"time"`
	Type      int              `json:"type"`
	UrlLists  []*UrlList       `json:"url_lists"`
	WordLists []*WordCountList `json:"word_lists"`
}

/* url列表 */
type UrlList struct {
	Url string `json:"url"`
}

/* 关键词统计 */
type WordCountList struct {
	Count int    `json:"count"`
	Word  string `json:"word"`
}

/* 客户等待（客服）平均时长列表 */
type WaitingTimesOnDay struct {
	WaitingDate      string             `json:"waiting_date"`
	WaitingTimeByIds []*WaitingTimeById `json:"waiting_time_by_ids"`
}

/* 平均等待时长 */
type WaitingTimeById struct {
	AvgWaitingTimes int    `json:"avg_waiting_times"`
	ServiceStaffId  string `json:"service_staff_id"`
}

/* 聊天对象ID列表 */
type Chatpeer struct {
	Date string `json:"date"`
	Uid  string `json:"uid"`
}

/* 评价详细 */
type EvalDetail struct {
	EvalCode   int    `json:"eval_code"`
	EvalRecer  string `json:"eval_recer"`
	EvalSender string `json:"eval_sender"`
	EvalTime   string `json:"eval_time"`
	SendTime   string `json:"send_time"`
}

/* 客服评价统计列表(按天) */
type StaffEvalStatOnDay struct {
	EvalDate           string               `json:"eval_date"`
	StaffEvalStatByIds []*StaffEvalStatById `json:"staff_eval_stat_by_ids"`
}

/* 客服评价统计 */
type StaffEvalStatById struct {
	Evaluations    []*Evaluation `json:"evaluations"`
	ServiceStaffId string        `json:"service_staff_id"`
}

/* 客服评价 */
type Evaluation struct {
	EvaluationName string `json:"evaluation_name"`
	EvaluationNum  string `json:"evaluation_num"`
}

/* 组及其成员列表 */
type GroupMember struct {
	GroupId    int    `json:"group_id"`
	GroupName  string `json:"group_name"`
	MemberList string `json:"member_list"`
}

/* 登录日志 */
type LoginLog struct {
	Time string `json:"time"`
	Type string `json:"type"`
}

/* 未回复统计列表(按天) */
type NonReplyStatOnDay struct {
	NonreplyDate      string              `json:"nonreply_date"`
	NonreplyStatByIds []*NonreplyStatById `json:"nonreply_stat_by_ids"`
}

/* 客服未回复统计 */
type NonreplyStatById struct {
	NonReplyCustomId string `json:"non_reply_customId"`
	NonReplyNum      int    `json:"non_reply_num"`
	ServiceStaffId   string `json:"service_staff_id"`
}

/* 某天的客服在线时长列表 */
type OnlineTimesOnDay struct {
	OnlineDate      string            `json:"online_date"`
	OnlineTimeByIds []*OnlineTimeById `json:"online_time_by_ids"`
}

/* 在线时长 */
type OnlineTimeById struct {
	OnlineTimes    int    `json:"online_times"`
	ServiceStaffId string `json:"service_staff_id"`
}

/* (某天)回复统计列表 */
type ReplyStatOnDay struct {
	ReplyDate      string           `json:"reply_date"`
	ReplyStatByIds []*ReplyStatById `json:"reply_stat_by_ids"`
}

/* 客服回复统计 */
type ReplyStatById struct {
	ReplyNum int    `json:"reply_num"`
	UserId   string `json:"user_id"`
}

/* 分流权重 */
type StreamWeight struct {
	User   string `json:"user"`
	Weight int    `json:"weight"`
}
