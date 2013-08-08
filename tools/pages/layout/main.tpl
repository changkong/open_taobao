{{define "main"}}
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link href="http://libs.baidu.com/bootstrap/2.3.2/css/bootstrap.min.css" rel="stylesheet">
	<link href="/assets/css/docs.css" rel="stylesheet">
	<script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
	<script src="http://libs.baidu.com/bootstrap/2.3.2/js/bootstrap.min.js"></script>
</head>
<body>
  	<header class="jumbotron subhead" >
  		<h2>淘宝开放平台SDK-Go语言版 测试编译工具</h2>
  	</header>
	<div class="container-fluid">
	  <div class="row-fluid">
	    <div class="span2 bs-docs-sidebar">
	     	<ul class="nav nav-list bs-docs-sidenav">
  				<li {{if .IsShowAuth}}class="active"{{end}}><a href="/auth/">
  					<i class="icon-chevron-right"></i>淘宝授权</a>
  				</li>
  				<li {{if .IsShowTest}}class="active"{{end}}><a href="/test/">
  					<i class="icon-chevron-right"></i>测试API</a>
  				</li>
  				<li {{if .IsShowMake}}class="active"{{end}}><a href="/make/">
  					<i class="icon-chevron-right"></i>编译SDK</a>
  				</li>
			</ul>
	  		<div class="row-fluid">
	  			<div class="ver-box">
	  				<dl>
	  					<dt>当前程序版本</dt>
	  					<dd>Ver {{.GetVersionNo}}</dd>
	  				</dl>	  				<dl>
	  					<dt>运行时SDK主库</dt>
	  					<dd>Ver {{.GetPkgTaobaoVersionNo}}</dd>
	  				</dl>
	  				<dl>
	  					<dt>运行时user子库</dt>
	  					<dd>Ver {{.GetPkgUserVersionNo}}</dd>
	  				</dl>
	  			</div>
	  		</div>			
	    </div>
	    <div class="span10">
		<div class="row-fluid info-box" >	
	    	<div class="alert {{if .IsMsgNo}}hide{{else}}{{if .IsMsgInfo}}alert-success{{else}}alert-error{{end}}{{end}}">
	    		<button type="button" class="close" data-dismiss="alert">&times;</button>
{{$page:=.}}
{{range $index, $msg := .GetResultMsg}}
{{$msg}}{{if $page.MsgNeedBr $index}}<br>{{end}}
{{end}}&nbsp;
			</div>
		</div>
		{{template "body" .}}

	    </div>
	  </div>
	</div>




	<br>

</body>
</html>
{{end}}