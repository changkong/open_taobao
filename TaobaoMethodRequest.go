// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package open_taobao

import (
	"errors"
	"time"
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
	t.SetValue("v", "2.0")
	if taobaoConfig.isTop {
		t.SetValue("session", accessToken)
		t.SetValue("timestamp", time.Now().Format("2006-01-02 15:04:05"))
		t.SetValue("app_key", taobaoConfig.appKey)
		t.SetValue("sign_method", "md5")
		t.SetReqUrl("http://gw.api.taobao.com/router/rest")

		if taobaoConfig.requestUrl != "" {
			t.SetReqUrl(taobaoConfig.requestUrl)
		}
	} else {
		t.SetValue("access_token", accessToken)
	}

	return executeRequest(t, resp, InsecureSkipVerify, DisableCompression)
}
