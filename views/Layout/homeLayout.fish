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
<div class="col-md-offset-1 col-md-10">
    <nav id="headDiv" class="row">
        <div>
            <h1 id="logoStyle">FreeFishGo</h1>
            <ul id="head" style="display: inline">
                {{{.homeHeadLi}}}
            </ul>
        </div>
    </nav>
    <div class="row" style="margin-top: 80px"></div>
    <div class="row">
        {{{.LayoutContent}}}
    </div>
</div>
<script src="/static/js/jquery-3.3.1.min.js"></script>
<script src="/static/tree/index.js"></script>
<script src="/static/bootstrap/js/bootstrap.min.js"></script>
<script>
</script>
</body>
</html>