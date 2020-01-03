<!DOCTYPE html>
<html>
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta charset="UTF-8">
    <title>FreeFishGo</title>
    <link rel="shortcut icon" href="/static/img/fish.ico" type="image/x-icon" />
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.css" />
    <link rel="stylesheet" href="/static/tree/style.css" />
    <link rel="stylesheet" href="/static/css/homeLayout.css"/>
</head>
<body>
    <div class="container" id="headDiv" style="border:lightgray 1px solid;height: 57px">
        <nav>
            <div>
                <h1 id="logoStyle">FreeFishGo</h1>
                <ul id="head" style="display: inline">
                    {{{range $k, $v := .homeHeadLi}}} <li><h4 {{{if $v.Active}}} class="headActive" {{{end}}} href="{{{$v.Path}}}?type=xhr">{{{$v.Name}}}</h4></li> {{{end}}}
                </ul>
            </div>
        </nav>
    </div>
    <div class="row" style="margin-top: 80px"></div>
    <div id="homeContent" class="container main-container">
        ((.LayoutContent))
    </div>
    <script src="/static/js/jquery-3.3.1.min.js"></script>
    <script src="/static/js/share.js"></script>
    <script src="/static/js/homeLayout.js"></script>
    <script src="/static/bootstrap/js/bootstrap.min.js"></script>
</body>
</html>