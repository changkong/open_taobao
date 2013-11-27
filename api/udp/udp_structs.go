// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package udp

const VersionNo = "20131127"


/* TargetSearchTopResult */
type TargetSearchTopResult struct {
	Field string `json:"field"`
	PageSize int `json:"page_size"`
	ResultData string `json:"result_data"`
	TotalCount int `json:"total_count"`

}

