{{define "body"}}

<form action="/make/do" method="post">

	<div class="row-fluid tool-bar">
		<input type="submit" class="btn btn-success"  value="根据配置重新生成 SDK 源码并编译"/>
		&nbsp;&nbsp;&nbsp;
		<a href="http://open.taobao.com/doc/category_list.htm?id=102">淘宝API详细列表</a>
	</div>

	<table>
		<tr>
			<td class="caption">配置文件</td>
			<td colspan=2>
				{{.GetData.ConfFile}}
				&nbsp;&nbsp;&nbsp;
				版本{{.GetData.MetaVersionNo}}
				&nbsp;&nbsp;
				<a href="http://api.taobao.com/myresources/standardSdk.htm?spm=0.0.0.0.7L33au">可从这里下载更新</a>
				&nbsp;&nbsp;
				注意更新后需重启服务
			</td>
		</tr>
		<tr>
			<td colspan=3>&nbsp;</td>
		</tr>
		<tr>
			<td class="caption">核心包</td>
			<td colspan=2>
				github.com/changkong/open_taobao
			</td>
		</tr>
		{{range .GetConfPackage.Mx}}
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

{{end}}