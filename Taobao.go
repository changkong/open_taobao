// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// 淘宝开放平台的 go 语言 SDK
package open_taobao

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const VersionNo = "0.0.1"

var (
	InsecureSkipVerify = true  // 相信淘宝服务器，跳过对其检查，减少客户端配置问题
	DisableCompression = false // 默认压缩数据
)

func executeRequest(req TaobaoRequestI, resp interface{}, insecureSkipVerify, disableCompression bool) ([]byte, error) {
	body := strings.NewReader(req.GetValues().Encode())
	req1, err := http.NewRequest("POST", req.GetReqUrl(), body)
	if err != nil {
		return nil, err
	}

	// 下面这句必须要，否则Post参数不对
	req1.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: insecureSkipVerify},
		DisableCompression: disableCompression,
	}
	client := &http.Client{Transport: tr}

	resp1, err := client.Do(req1)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		return nil, err
	}

	var errResp TaobaoTopErrResponse
	err = json.Unmarshal(data, &errResp)
	if err != nil {
		return data, err
	}
	errstr := errResp.GetErr()
	if errstr != "" {
		return data, errors.New(errstr)
	}

	return data, json.Unmarshal(data, resp)
}
