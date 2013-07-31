// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
)

func main() {

	controller, err := NewController()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/refresh", controller.Refresh)
	http.HandleFunc("/save", controller.Save)
	http.HandleFunc("/auth", controller.Auth)
	http.HandleFunc("/makeapi", controller.MakeApi)
	http.HandleFunc("/user/buyer/get", controller.UserBuyerGet)
	http.HandleFunc("/user/seller/get", controller.UserSellerGet)
	http.HandleFunc("/callback", controller.Callback)
	http.HandleFunc("/", controller.Home)

	fmt.Println("start 80 ......")
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}
