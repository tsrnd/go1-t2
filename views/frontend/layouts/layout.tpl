<!DOCTYPE html>
<html lang="">
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="shortcut icon" href="/static/img/favicon.png">
		<title>{{.Title }}</title>
		<link href="/static/css/bootstrap.css" rel="stylesheet">
		<link href='http://fonts.googleapis.com/css?family=Roboto:400,300,300italic,400italic,500,700,500italic,100italic,100' rel='stylesheet' type='text/css'>
		<link href="/static/css/font-awesome.min.css" rel="stylesheet">
		<link rel="stylesheet" href="/static/css/flexslider.css" type="text/css" media="screen"/>
		<link href="/static/css/sequence-looptheme.css" rel="stylesheet" media="all"/>
		<link href="/static/css/style.css" rel="stylesheet">
	</head>
	<body id="home">
		<div class="wrapper">
			<div class="header">
				{{.Header}}
			</div>
			<div class="clearfix"></div>
			<div class="hom-slider">
				{{.Slider}}
			</div>
			<div class="clearfix"></div>
			<div class="container_fullwidth">
				<div class="container">
					{{.LayoutContent}}
				</div>
			</div>
			<div class="clearfix"></div>
			<div class="footer">
				{{.Footer}}
			</div>
		</div>
		<script type="text/javascript" src="/static/js/jquery-1.10.2.min.js"></script>
		<script type="text/javascript" src="/static/js/jquery.easing.1.3.js"></script>
		<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
		<script type="text/javascript" src="/static/js/jquery.sequence-min.js"></script>
		<script type="text/javascript" src="/static/js/jquery.carouFredSel-6.2.1-packed.js"></script>
		<script defer src="/static/js/jquery.flexslider.js"></script>
		<script type="text/javascript" src="/static/js/script.min.js" ></script>
		<script type="text/javascript" src="/static/js/bootbox.min.js"></script>
		<script type="text/javascript" src="/static/js/popup-confirm.js"></script>
	</body>
</html>