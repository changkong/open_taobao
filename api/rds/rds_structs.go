// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rds

const VersionNo = "20130808"

/* rds创建、查询、删除返回结果数据结构 */
type RdsDbInfo struct {
	Charset      string `json:"charset"`
	Comment      string `json:"comment"`
	DbId         string `json:"db_id"`
	DbName       string `json:"db_name"`
	DbStatus     string `json:"db_status"`
	DbType       string `json:"db_type"`
	InstanceId   string `json:"instance_id"`
	InstanceName string `json:"instance_name"`
	InstanceType string `json:"instance_type"`
	MaxAccount   string `json:"max_account"`
	Password     string `json:"password"`
	Uid          string `json:"uid"`
	UserName     string `json:"user_name"`
}
