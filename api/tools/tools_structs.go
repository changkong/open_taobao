// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tools

const VersionNo = "20131127"


/* KFC 关键词过滤匹配结果 */
type KfcSearchResult struct {
	Content string `json:"content"`
	Level string `json:"level"`
	Matched bool `json:"matched"`

}

/* 批量异步任务结果 */
type Task struct {
	CheckCode string `json:"check_code"`
	Created string `json:"created"`
	DownloadUrl string `json:"download_url"`
	Method string `json:"method"`
	Schedule string `json:"schedule"`
	Status string `json:"status"`
	Subtasks []*Subtask `json:"subtasks"`
	TaskId int `json:"task_id"`

}

/* 批量异步任务的子任务结果 */
type Subtask struct {
	IsSuccess bool `json:"is_success"`
	SubTaskRequest string `json:"sub_task_request"`
	SubTaskResult string `json:"sub_task_result"`

}

