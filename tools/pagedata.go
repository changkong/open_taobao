// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strings"
)

type PageDataT struct {
	IsAuth      bool
	AuthUrl     string
	ResultClass string
	ResultMsg   []string

	UserBuyerName  string
	UserBuyerJson  string
	UserSellerName string
	UserSellerJson string
}

func NewPageData() *PageDataT {
	var p PageDataT
	p.ResultClass = "none"
	return &p
}

func (p *PageDataT) MsgNeedBr(index int) bool {
	return len(p.ResultMsg) > 1 &&
		(index+1) < len(p.ResultMsg)
}

func (p *PageDataT) setErr(msg string) {
	p.ResultClass = "error"
	if msg == "" {
		p.ResultMsg = []string{}
	} else {
		p.ResultMsg = []string{msg}
	}
	fmt.Println("[error]", msg)
}

func (p *PageDataT) addErr(msg string) {
	p.ResultClass = "error"
	p.ResultMsg = append(p.ResultMsg, msg)
}

func (p *PageDataT) setInfo(msg string) {
	msg = strings.TrimSpace(msg)
	if msg == "" {
		p.ResultClass = "none"
		p.ResultMsg = []string{}
	} else {
		p.ResultClass = "info"
		p.ResultMsg = []string{msg}
	}
	fmt.Println("[info]", msg)
}
