{{define "body"}}


<table class="table table-striped">
	<tr>
		<td class="caption">taobao.user.buyer.get:</td>
		<td colspan=2><input type="button" class="btn btn-success" onclick="window.location='/test/user/buyer/get'" value="执行"/>&nbsp;查询买家信息</td>
	</tr>
	<tr>
		<td class="caption">返回昵称:</td>
		<td  colspan=2 style="color:green;">{{.UserBuyerName}}</td>
	</tr>
	<tr>
		<td class="caption">返回数据:</td>
		<td colspan=2>{{.UserBuyerJson}}</td>
	</tr>
	<tr>
		<td class="caption">taobao.user.seller.get:</td>
		<td colspan=2><input type="button" class="btn btn-primary" onclick="window.location='/test/user/seller/get'" value="执行"/>&nbsp;查询卖家信息</td>
	</tr>
	<tr>
		<td class="caption">返回昵称:</td>
		<td colspan=2 style="color:green;">{{.UserSellerName}}</td>
	</tr>
	<tr>
		<td class="caption">返回数据:</td>
		<td colspan=2>{{.UserSellerJson}}</td>
	</tr>
</table>

{{end}}