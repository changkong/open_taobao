// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package open_taobao

import (
	"errors"
)

type TaobaoMethodRequest struct {
	TaobaoRequest
}

func (t *TaobaoMethodRequest) GetResponse(accessToken, apiMethodName string, resp interface{}) ([]byte, error) {
	if accessToken == "" {
		return nil, errors.New("[" + apiMethodName + "] AccessToken is null")
	}
	t.SetReqUrl("https://eco.taobao.com/router/rest")

	t.SetValue("method", apiMethodName)
	t.SetValue("format", "json")
	t.SetValue("access_token", accessToken)
	t.SetValue("v", "2.0")

	return executeRequest(t, resp, InsecureSkipVerify, DisableCompression)
}
