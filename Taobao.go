// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package open_taobao

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const VersionNo = "0.0.1"

var (
	InsecureSkipVerify = true  // 相信淘宝服务器，跳过对其检查，减少客户端配置问题
	DisableCompression = false // 默认压缩数据
	taobaoConfig       = &TaobaoConfig{}
)

type TaobaoConfig struct {
	appKey     string
	secret     string
	requestUrl string
	isTop      bool
}

func Init(key string, secret string, isTop bool) {
	taobaoConfig = &TaobaoConfig{appKey: key, secret: secret, isTop: isTop}
}

func SetRequestUrl(requestUrl string) {
	taobaoConfig.requestUrl = requestUrl
}

type keyValue struct{ key, value string }

type byKeyValue []keyValue

func (p byKeyValue) Len() int      { return len(p) }
func (p byKeyValue) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p byKeyValue) Less(i, j int) bool {
	return p[i].key < p[j].key
}

func (p byKeyValue) appendValues(values url.Values) byKeyValue {
	for k, vs := range values {
		for _, v := range vs {
			p = append(p, keyValue{k, v})
		}
	}
	return p
}

func executeRequest(req TaobaoRequestI, resp interface{}, insecureSkipVerify, disableCompression bool) ([]byte, error) {

	if taobaoConfig.isTop {
		params := make(byKeyValue, 0)
		params = params.appendValues(req.GetValues())
		sort.Sort(params)

		str := ""
		for _, keyValue := range params {
			str += keyValue.key + keyValue.value
		}

		str = taobaoConfig.secret + str + taobaoConfig.secret
		md5Util := md5.New()
		md5Util.Write([]byte(str))
		sign := hex.EncodeToString(md5Util.Sum(nil))

		req.GetValues().Add("sign", strings.ToUpper(sign))
	}

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
