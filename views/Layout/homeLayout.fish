<!DOCTYPE html>
<html>
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta charset="UTF-8">
    <title>FreeFishGo</title>
    <link rel="shortcut icon" href="/img/fish.ico" type="image/x-icon" />
    <link rel="stylesheet" href="/bootstrap/css/bootstrap.css" />
    <link rel="stylesheet" href="/tree/style.css" />
    <link rel="stylesheet" href="/css/homeLayout.css"/>
</head>
<script src="/js/jquery-3.3.1.min.js"></script>
<script src="/js/share.js"></script>
<script src="/js/homeLayout.js"></script>
<script src="/bootstrap/js/bootstrap.min.js"></script>
<body>
    <div class="row" id="headDiv">
        <div class="container  main-container">
            <nav>
                <div>
                    <h1 id="logoStyle">FreeFishGo</h1>
                    <ul id="head" style="display: inline">
                        {{{range $k, $v := .homeHeadLi}}} <li><h4 {{{if $v.Active}}} class="headActive" {{{end}}} href="{{{$v.Path}}}?type=xhr">{{{$v.Name}}}</h4></li> {{{end}}}
                    </ul>
                </div>
            </nav>
        </div>
    </div>
    <div class="row" style="margin-top: 80px"></div>
    <div id="homeContent" class="container main-container">
        ((.LayoutContent))
    </div>
</body>
</html>