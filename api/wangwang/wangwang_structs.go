// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wangwang

const VersionNo = "20131207"


/*  */
type WordListListObject struct {
	WordList []*WordList `json:"word_list"`

}

/* 关键词列表 */
type WordList struct {
	Word string `json:"word"`

}

/*  */
type MsgListListObject struct {
	MsgList []*MsgList `json:"msg_list"`

}

/* 聊天记录列表 */
type MsgList struct {
	Content string `json:"content"`
	Direction int `json:"direction"`
	Length int `json:"length"`
	Time string `json:"time"`
	Type int `json:"type"`
	UrlLists *UrlListListObject `json:"url_lists"`
	WordLists *WordCountListListObject `json:"word_lists"`

}

/*  */
type UrlListListObject struct {
	UrlList []*UrlList `json:"url_list"`

}

/* url列表 */
type UrlList struct {
	Url string `json:"url"`

}

/*  */
type WordCountListListObject struct {
	WordCountList []*WordCountList `json:"word_count_list"`

}

/* 关键词统计 */
type WordCountList struct {
	Count int `json:"count"`
	Word string `json:"word"`

}

/*  */
type WaitingTimesOnDayListObject struct {
	WaitingTimesOnDay []*WaitingTimesOnDay `json:"waiting_times_on_day"`

}

/* 客户等待（客服）平均时长列表 */
type WaitingTimesOnDay struct {
	WaitingDate string `json:"waiting_date"`
	WaitingTimeByIds *WaitingTimeByIdListObject `json:"waiting_time_by_ids"`

}

/*  */
type WaitingTimeByIdListObject struct {
	WaitingTimeById []*WaitingTimeById `json:"waiting_time_by_id"`

}

/* 平均等待时长 */
type WaitingTimeById struct {
	AvgWaitingTimes int `json:"avg_waiting_times"`
	ServiceStaffId string `json:"service_staff_id"`

}

/*  */
type ChatpeerListObject struct {
	Chatpeer []*Chatpeer `json:"chatpeer"`

}

/* 聊天对象ID列表 */
type Chatpeer struct {
	Date string `json:"date"`
	Uid string `json:"uid"`

}

/*  */
type EvalDetailListObject struct {
	EvalDetail []*EvalDetail `json:"eval_detail"`

}

/* 评价详细 */
type EvalDetail struct {
	EvalCode int `json:"eval_code"`
	EvalRecer string `json:"eval_recer"`
	EvalSender string `json:"eval_sender"`
	EvalTime string `json:"eval_time"`
	SendTime string `json:"send_time"`

}

/*  */
type StaffEvalStatOnDayListObject struct {
	StaffEvalStatOnDay []*StaffEvalStatOnDay `json:"staff_eval_stat_on_day"`

}

/* 客服评价统计列表(按天) */
type StaffEvalStatOnDay struct {
	EvalDate string `json:"eval_date"`
	StaffEvalStatByIds *StaffEvalStatByIdListObject `json:"staff_eval_stat_by_ids"`

}

/*  */
type StaffEvalStatByIdListObject struct {
	StaffEvalStatById []*StaffEvalStatById `json:"staff_eval_stat_by_id"`

}

/* 客服评价统计 */
type StaffEvalStatById struct {
	Evaluations *EvaluationListObject `json:"evaluations"`
	ServiceStaffId string `json:"service_staff_id"`

}

/*  */
type EvaluationListObject struct {
	Evaluation []*Evaluation `json:"evaluation"`

}

/* 客服评价 */
type Evaluation struct {
	EvaluationName string `json:"evaluation_name"`
	EvaluationNum string `json:"evaluation_num"`

}

/*  */
type GroupMemberListObject struct {
	GroupMember []*GroupMember `json:"group_member"`

}

/* 组及其成员列表 */
type GroupMember struct {
	GroupId int `json:"group_id"`
	GroupName string `json:"group_name"`
	MemberList string `json:"member_list"`

}

/*  */
type LoginLogListObject struct {
	LoginLog []*LoginLog `json:"login_log"`

}

/* 登录日志 */
type LoginLog struct {
	Time string `json:"time"`
	Type string `json:"type"`

}

/*  */
type NonReplyStatOnDayListObject struct {
	NonReplyStatOnDay []*NonReplyStatOnDay `json:"non_reply_stat_on_day"`

}

/* 未回复统计列表(按天) */
type NonReplyStatOnDay struct {
	NonreplyDate string `json:"nonreply_date"`
	NonreplyStatByIds *NonreplyStatByIdListObject `json:"nonreply_stat_by_ids"`

}

/*  */
type NonreplyStatByIdListObject struct {
	NonreplyStatById []*NonreplyStatById `json:"nonreply_stat_by_id"`

}

/* 客服未回复统计 */
type NonreplyStatById struct {
	NonReplyCustomId string `json:"non_reply_customId"`
	NonReplyNum int `json:"non_reply_num"`
	ServiceStaffId string `json:"service_staff_id"`

}

/*  */
type OnlineTimesOnDayListObject struct {
	OnlineTimesOnDay []*OnlineTimesOnDay `json:"online_times_on_day"`

}

/* 某天的客服在线时长列表 */
type OnlineTimesOnDay struct {
	OnlineDate string `json:"online_date"`
	OnlineTimeByIds *OnlineTimeByIdListObject `json:"online_time_by_ids"`

}

/*  */
type OnlineTimeByIdListObject struct {
	OnlineTimeById []*OnlineTimeById `json:"online_time_by_id"`

}

/* 在线时长 */
type OnlineTimeById struct {
	OnlineTimes int `json:"online_times"`
	ServiceStaffId string `json:"service_staff_id"`

}

/*  */
type ReplyStatOnDayListObject struct {
	ReplyStatOnDay []*ReplyStatOnDay `json:"reply_stat_on_day"`

}

/* (某天)回复统计列表 */
type ReplyStatOnDay struct {
	ReplyDate string `json:"reply_date"`
	ReplyStatByIds *ReplyStatByIdListObject `json:"reply_stat_by_ids"`

}

/*  */
type ReplyStatByIdListObject struct {
	ReplyStatById []*ReplyStatById `json:"reply_stat_by_id"`

}

/* 客服回复统计 */
type ReplyStatById struct {
	ReplyNum int `json:"reply_num"`
	UserId string `json:"user_id"`

}

/*  */
type StreamWeightListObject struct {
	StreamWeight []*StreamWeight `json:"stream_weight"`

}

/* 分流权重 */
type StreamWeight struct {
	User string `json:"user"`
	Weight int `json:"weight"`

}

