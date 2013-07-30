// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConfMainT struct {
	AppKey      string
	AppSecret   string
	AccessToken string
	RedirectUri string

	ConfFile string `json:"-"` // 不需要输出
}

func (c *ConfMainT) IsNoAccessToken() bool {
	return c.AccessToken == ""
}

func NewConfMain(confFile string) (*ConfMainT, error) {
	fmt.Println("read", confFile)
	var confMain ConfMainT
	confMain.ConfFile = confFile

	body, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &confMain)
	if err != nil {
		return nil, err
	}

	return &confMain, nil
}

func (c *ConfMainT) Save() error {
	body, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.ConfFile, body, 0600)
}
