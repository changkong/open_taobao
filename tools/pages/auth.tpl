{{define "body"}}

<form action="/auth/save" method="post" class="form-horizontal">

<div class="row-fluid tool-bar">
	<input type="submit" class="btn btn-success" value="保存参数"/>
	<input type="button" class="btn btn-success"  onclick="window.location='/auth/do'" value="重新授权"/>&nbsp;&nbsp;&nbsp;
	<a href="http://open.taobao.com/doc/detail.htm?id=994">https免签调用方式详解</a>&nbsp;&nbsp;&nbsp;
	<a href="http://open.taobao.com/doc/detail.htm?id=118">OAuth2.0授权方式</a>
</div>
<div class="control-group">
	<label class="control-label" for="AppKey">App Key:</label>
    <div class="controls">
      <input type="text" id="AppKey" name="AppKey" class="span8" placeholder="淘宝授予, 对外唯一区分应用, 对外公开, 授权时必须" value="{{.GetConfMain.AppKey}}"/>
    </div>
</div>

<div class="control-group">
	<label class="control-label span2" for="AppSecret">App Secret:</label>
    <div class="controls">
      <input type="text" id="AppSecret" name="AppSecret" class="span8" placeholder="淘宝授予, 确保应用不被假冒, 对外不可公开, 授权时必须" value="{{.GetConfMain.AppSecret}}"/>
    </div>
</div>

<div class="control-group">
	<label class="control-label" for="RedirectUri">RedirectUri:</label>
    <div class="controls">
      	<input type="text" id="RedirectUri" name="RedirectUri" class="span8" placeholder="在淘宝中登记，通过该链接，淘宝将用户授权信息传递给应用，授权时必须" value="{{.GetConfMain.RedirectUri}}"/>
	  	<br><span class="text-info">当前测试程序要求 RedirectUri 须是 http://域名或IP[:端口]/auth/callback 的格式</span>
	  	<br><span class="text-warning">通过 RedirectUri 能访问到本服务, 用户授权才正常</span>
    </div>
</div>

<div class="control-group">
	<label class="control-label">AccessToken:</label>
    <div class="controls">
      	<span type="text" class="span8 uneditable-input">{{.GetConfMain.AccessToken}}
      		<span class="text-error">{{.GetConfMain.GetAuthMsg}}</span>
  		</span>
	  	<br><span class="text-info">用户授权后, 淘宝授予, 唯一确定用户对应用的信任关系</span>
	  	<br><span class="text-info">有时效性, 不可公开。后续API调用, 都需通过该参数</span>
    </div>
</div>

</form>

{{end}}