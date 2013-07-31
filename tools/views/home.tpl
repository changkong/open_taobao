<html>
<head>
<style type="text/css">

body,td,input {
	font-size:12px;
}

.in {
	width:500px;
}
.title {
	text-align: center;
	padding:6px;
	color: white;
	background-color: #006000;
}
.caption {
	text-align: right;
	width:200px;
}
.explain {
	color:gray;
}
.error {
	padding:6px;
	color:red;
	background-color:#ffc0c0; 
	border:1px solid red;
}
.info {
	padding:6px;
	color:green;
	background-color:c0ffc0; 
	border:1px solid 40c000;
}
.none {
	padding:6px;
	color:black;
	background-color:white; 
	border:1px solid white;
}
.no_access_token {
	color:red;
	show:hidden;
}
</style>
</head>
<body>
	<form action="/save" method="post">
		<table>
			<tr>
				<td colspan=3 style="color:blue;" >淘宝go语言SDK测试
					<input type="button" value="刷新页面" onclick="window.location='/refresh'"/>
				</td>
			</tr>
			<tr>
				<td colspan=3 class="{{.PageData.ResultClass}}">
{{$pagedata := .PageData}}
{{range $index, $msg := .PageData.ResultMsg}}
{{$msg}}{{if $pagedata.MsgNeedBr $index}}<br>{{end}}
{{end}}
					&nbsp;
				</td>
			</tr>
			<tr>
				<td class="title">淘宝参数</td>
				<td colspan=2>
					<input type="submit" value="保存参数"/>
					<a href="http://open.taobao.com/doc/detail.htm?id=994">[https免签调用方式详解]</a>
					<a href="http://open.taobao.com/doc/detail.htm?id=118">[OAuth2.0授权方式]</a>
				</td>
			</tr>
			<tr>
				<td class="caption">App Key: </td>
				<td colspan=2><input type="text" name="AppKey" class="in" value="{{.ConfMain.AppKey}}"/>
				<span class="explain">淘宝授予, 对外唯一区分应用, 对外公开</span></td>
			</tr>
			<tr>
				<td class="caption">App Secret:</td>
				<td colspan=2><input type="text" name="AppSecret" class="in" value="{{.ConfMain.AppSecret}}"/>
				<span class="explain">淘宝授予, 确保应用不被假冒, 对外不可公开</span></td>
			</tr>
			<tr>
				<td class="caption">RedirectUri:</td>
				<td colspan=2><input type="text"  name="RedirectUri" class="in"  value="{{.ConfMain.RedirectUri}}"/>
				<span class="explain">在淘宝中登记，通过该链接，淘宝将用户授权信息传递给应用</span></td>
			</tr>
			<tr>
				<td class="caption">&nbsp;</td>
				<td colspan=2>
				<span style="color:blue;">注意, 当前测试程序, RedirectUri 须是 http://域名或IP[:端口]/callback 的格式, 且能访问到本服务, 用户授权才能正常</span></td>
			</tr>
			<tr>
				<td class="title">用户授权</td>
				<td colspan=2>
					<input type="button"  onclick="window.location='/auth'" value="重新授权"/>
				</td>
			</tr>
			<tr>
				<td class="caption">AccessToken: </td>
				<td colspan=2>
					{{.ConfMain.AccessToken}}
					{{if .ConfMain.IsNoAccessToken}}
					<span class="no_access_token{{.ConfMain.AccessToken}}">用户未授权</span>
					{{end}}
				</td>
			</tr>
			<tr>
				<td>&nbsp;</td>
				<td colspan=2>
					<span class="explain">用户授权后, 淘宝授予, 唯一确定用户对应用的信任关系, 有时效性, 不可公开, 后续API调用, 都需要通过 AccessToken</span>
				</td>
			</tr>
			<tr>
				<td class="title">测试API</td>
				<td colspan=2>
					运行时版本：open_taobao[ver {{.OpenTaobaoVersionNo}}] user[{{.PkgUserVersionNo}}]
				</td>
			</tr>
			<tr>
				<td class="caption">taobao.user.buyer.get:</td>
				<td colspan=2><input type="button" onclick="window.location='/user/buyer/get'" value="执行"/>&nbsp;查询买家信息</td>
			</tr>
			<tr>
				<td class="caption">返回昵称:</td>
				<td  colspan=2 style="color:green;">{{.PageData.UserBuyerName}}</td>
			</tr>
			<tr>
				<td class="caption">返回数据:</td>
				<td colspan=2>{{.PageData.UserBuyerJson}}</td>
			</tr>
			<tr>
				<td class="caption">taobao.user.seller.get:</td>
				<td colspan=2><input type="button" onclick="window.location='/user/seller/get'" value="执行"/>&nbsp;查询卖家信息</td>
			</tr>
			<tr>
				<td class="caption">返回昵称:</td>
				<td colspan=2 style="color:green;">{{.PageData.UserSellerName}}</td>
			</tr>
			<tr>
				<td class="caption">返回数据:</td>
				<td colspan=2>{{.PageData.UserSellerJson}}</td>
			</tr>
		</table>
	</form>
	<form action="/makeapi" method="post">
		<table>
			<tr>
				<td class="title">编译SDK</td>
				<td colspan=2>
					<input type="submit" value="根据 ApiMetadata.xml 重新生成 SDK 源代码并编译"/>
				</td>
			</tr>
			<tr>
				<td class="caption">配置文件</td>
				<td colspan=2>
					{{.Data.ConfFile}}&nbsp;版本{{.Data.MetaVersionNo}}&nbsp;
					<a href="http://api.taobao.com/myresources/standardSdk.htm?spm=0.0.0.0.7L33au">从淘宝这里下载最新配置</a>
					<span class="explain">文件替换后需重启服务</span>
				</td>
			</tr>
			<tr>
				<td class="caption">&nbsp;</td>
				<td colspan=2>
					<a href="http://open.taobao.com/doc/category_list.htm?id=102">在这里可以看到淘宝API详细列表</a>
				</td>
			</tr>
			<tr>
				<td class="caption">核心包</td>
				<td colspan=2>
					github.com/changkong/open_taobao
				</td>
			</tr>
			{{range .ConfPackage.Mx}}
			<tr>
				<td class="caption">{{.Caption}}包</td>
				<td colspan=2>
					<input name="PkgChoose" type=checkbox {{if .PkgChoose}}checked{{end}} value="{{.Name}}"/>
					{{.FullName}}
				</td>
			</tr>
			{{end}}
		</table>
	</form>
	<br>
</body>
</html>