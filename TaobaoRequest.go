// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package open_taobao

import (
	"net/url"
)

type TaobaoRequestI interface {
	GetValues() url.Values
	GetReqUrl() string
}

type TaobaoRequest struct {
	reqUrl string
	values url.Values
}

func (t *TaobaoRequest) SetValue(key, value string) {
	if t.values == nil {
		t.values = url.Values{}
	}
	t.values.Set(key, value)
}

func (t *TaobaoRequest) GetValue(key string) string {
	return t.values.Get(key)
}

func (t *TaobaoRequest) GetValues() url.Values {
	return t.values
}

func (t *TaobaoRequest) SetReqUrl(resUrl string) {
	t.reqUrl = resUrl
}

func (t *TaobaoRequest) GetReqUrl() string {
	return t.reqUrl
}
