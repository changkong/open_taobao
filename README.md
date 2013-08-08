open_taobao
===========

淘宝开放平台的 go 语言 SDK

自带Web小工具
=============
* 可测试基本流程
* 能根据淘宝 ApiMetadata.xml 重新生成SDK源码并编译, 保持与淘宝同步

版本说明
========

* **主包:** 0.0.1
* **子包:** 根据淘宝 20130808 的 ApiMetadata.xml 生成
* **工具程序:** 0.0.1

安装说明
========

**安装主包**

* go get github.com/changkong/open_taobao

**安装子包**

* go get github.com/changkong/open_taobao/api/user

**更多子包**

* **用户API:**  github.com/changkong/open_taobao/api/user
* **类目API:**  github.com/changkong/open_taobao/api/item
* **商品API:**  github.com/changkong/open_taobao/api/product
* **交易API:**  github.com/changkong/open_taobao/api/trade
* **评价API:**  github.com/changkong/open_taobao/api/pinjia
* **物流API:**  github.com/changkong/open_taobao/api/delivery
* **店铺API:**  github.com/changkong/open_taobao/api/shop
* **分销API:**  github.com/changkong/open_taobao/api/fenxiao
* **旺旺API:**  github.com/changkong/open_taobao/api/wangwang
* **淘客API:**  github.com/changkong/open_taobao/api/taobaoke
* **主动通知业务API:**  github.com/changkong/open_taobao/api/notice
* **工具类API:**  github.com/changkong/open_taobao/api/tools
* **物流宝API:**  github.com/changkong/open_taobao/api/wlb
* **直通车API:**  github.com/changkong/open_taobao/api/simba
* **收藏夹API:**  github.com/changkong/open_taobao/api/favorite
* **机票API:**  github.com/changkong/open_taobao/api/jipiao
* **营销API:**  github.com/changkong/open_taobao/api/marketing
* **数据API:**  github.com/changkong/open_taobao/api/udp
* **酒店API:**  github.com/changkong/open_taobao/api/hotel
* **店铺会员管理API:**  github.com/changkong/open_taobao/api/crm
* **多媒体平台API:**  github.com/changkong/open_taobao/api/dmt
* **子账号管理API:**  github.com/changkong/open_taobao/api/subuser
* **服务平台API:**  github.com/changkong/open_taobao/api/fuwu
* **退款API:**  github.com/changkong/open_taobao/api/refund
* **关联推荐API:**  github.com/changkong/open_taobao/api/recommend
* **组件API:**  github.com/changkong/open_taobao/api/widget
* **支付宝API:**  github.com/changkong/open_taobao/api/alipy
* **天猫精品库API:**  github.com/changkong/open_taobao/api/jp
* **聚石塔API:**  github.com/changkong/open_taobao/api/rds
* **旅行度假API:**  github.com/changkong/open_taobao/api/trvael
* **彩票API:**  github.com/changkong/open_taobao/api/caipiao
* **账务API:**  github.com/changkong/open_taobao/api/bill
* **天猫退款退货API:**  github.com/changkong/open_taobao/api/eai

测试说明
========

运行如下命令, 执行web小工具

* cd tools
* go build .
* tools

启动后，默认通过80端口访问, 提供如下功能：

* 参数修改保存
* 授权登陆测试
* 用户信息获取测试
* 根据淘宝最新 ApiMetadata.xml 重新生成SDK源码并编译, 保证代码最新
* ApiMetadata.xml 在 tools/conf 目录下，可从 [淘宝这里](http://api.taobao.com/myresources/standardSdk.htm) 下载最新配置

代码样例
========

* **获取授权url**(供用户访问、授权)

  * url, err := open_taobao.GetUrlForAuth(AppKey, RedirectUri, "")
  * 可参考 tools/controller.go 的 Auth 方法
  
* **获取用户授权后的处理**(接收淘宝回调信息)

  * var req open_taobao.TokenGetRequest
  * req.SetAppKey(AppKey)
  * req.SetAppSecret(AppSecret)
  * req.SetRedirectUri(RedirectUri)
  * req.SetCode(code)
  * req.SetState("")
  * resp, _, err := req.GetResponse()
  * 成功后 resp.AccessToken 就是访问令牌，可以用来调用功能API了
  * 可参考 tools/controller.go 的 Callback 方法

* **调用功能API**

  * var req user.UserBuyerGetRequest
  * req.SetFields("user_id, nick")
  * resp, data, err := req.GetResponse(AccessToken)
  * 成功后 resp 是结果对象, data 是返回的原始json串
  * 可参考 tools/controller.go 的 UserBuyerGet 或 UserSellerGet 方法

* **其他资源**

  * [淘宝API详细列表](http://open.taobao.com/doc/category_list.htm?id=102) 可以API更详细的解释
  * 对登陆授权的更多介绍，参见  [https免签调用方式详解](http://open.taobao.com/doc/detail.htm?id=994) [OAuth2.0授权方式](http://open.taobao.com/doc/detail.htm?id=118)

技术说明
========

* **调用方式:** 只支持 https 免签调用方式
* **业务字段:** 为了调用方便, 都生成相关字段或方法, 而不是 map 或 json 对象, 缺点是代码多, 包大
* **子包分拆:** 由于生成完整的Go包, 有20多M, 太大了，因此按淘宝官方的分类分拆成子包

  * API分类淘宝没有给出配置文件, 是根据[淘宝API列表](http://open.taobao.com/doc/category_list.htm?id=102), 手工维护的配置表, 原理根据 api 前缀分类
  * 配置文件是 tools/conf/package.json, 若分类有问题，可修改该配置文件, 然后重新生成SDK

* **调用参数:** 所有调用参数，只支持字符串类型
* **返回结果:** 做了 string bool int float struct 的判断转型, 调用方便
* **淘宝错误:** 对淘宝返回错误做了包装, 都会通过 error 对象返回
* **对象命名:** 尽量参照淘宝官方其他SDK的命名方式, 例如:

  * 获取卖家信息的淘宝API是 taobao.user.seller.get
  * open_taobao 中对应 user 子包的 UserSellerGetRequest 对象的 GetResponse 方法

* **前端框架:** bootstrap jquery
* **后端框架:**直接使用 go 自带功能

Good Enjoy
==========

(^_^)

LICENSE
=======
open_taobao is licensed under the MIT
