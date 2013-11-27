// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rds

import (
	"github.com/changkong/open_taobao"
)





/* 在rds实例里创建数据库 */
type RdsDbCreateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 数据库名 */
func (r *RdsDbCreateRequest) SetDbName(value string) {
	r.SetValue("db_name", value)
}

/* rds的实例名 */
func (r *RdsDbCreateRequest) SetInstanceName(value string) {
	r.SetValue("instance_name", value)
}


func (r *RdsDbCreateRequest) GetResponse(accessToken string) (*RdsDbCreateResponse, []byte, error) {
	var resp RdsDbCreateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.rds.db.create", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type RdsDbCreateResponse struct {
	RdsDbInfo *RdsDbInfo `json:"rds_db_info"`
}

type RdsDbCreateResponseResult struct {
	Response *RdsDbCreateResponse `json:"rds_db_create_response"`
}





/* 通过api删除用户RDS的数据库 */
type RdsDbDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 数据库的name，可以通过 taobao.rds.db.get 获取 */
func (r *RdsDbDeleteRequest) SetDbName(value string) {
	r.SetValue("db_name", value)
}

/* rds的实例名 */
func (r *RdsDbDeleteRequest) SetInstanceName(value string) {
	r.SetValue("instance_name", value)
}


func (r *RdsDbDeleteRequest) GetResponse(accessToken string) (*RdsDbDeleteResponse, []byte, error) {
	var resp RdsDbDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.rds.db.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type RdsDbDeleteResponse struct {
	RdsDbInfo *RdsDbInfo `json:"rds_db_info"`
}

type RdsDbDeleteResponseResult struct {
	Response *RdsDbDeleteResponse `json:"rds_db_delete_response"`
}





/* 查询rds实例下的数据库 */
type RdsDbGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 数据库状态，默认值1 */
func (r *RdsDbGetRequest) SetDbStatus(value string) {
	r.SetValue("db_status", value)
}

/* rds的实例名 */
func (r *RdsDbGetRequest) SetInstanceName(value string) {
	r.SetValue("instance_name", value)
}


func (r *RdsDbGetRequest) GetResponse(accessToken string) (*RdsDbGetResponse, []byte, error) {
	var resp RdsDbGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.rds.db.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type RdsDbGetResponse struct {
	RdsDbInfos []*RdsDbInfo `json:"rds_db_infos"`
}

type RdsDbGetResponseResult struct {
	Response *RdsDbGetResponse `json:"rds_db_get_response"`
}



