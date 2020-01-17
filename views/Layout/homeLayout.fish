<!DOCTYPE html>
<html>
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta charset="UTF-8">
    <title>freeFishGo</title>
    <link rel="shortcut icon" href="/img/fish.ico" type="image/x-icon" />
    <link rel="stylesheet" href="/bootstrap/css/bootstrap.min.css" />
    <link rel="stylesheet" href="/tree/style.css" />
    <link rel="stylesheet" href="/css/homeLayout.css"/>
    <link rel="stylesheet" href="/froala_editor_3.1.0/css/froala_editor.pkgd.css"/>

</head>
<script src="/js/jquery-3.3.1.min.js"></script>
<script src="/js/share.js"></script>
<script src="/js/homeLayout.js"></script>
<script src="/bootstrap/js/bootstrap.min.js"></script>
<body>
    <div class="row  navbar-default" id="headDiv">
        <div id="headDivHei" class="container  main-container">
            <nav class="navbar" style="margin-bottom: 0px;" role="navigation">
                <div class="container-fluid"> 
                <div class="navbar-header">
                    <button id="changeheadDivHei" type="button" class="navbar-toggle" data-toggle="collapse"
                            data-target="#example-navbar-collapse">
                        <span class="sr-only">切换导航</span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <a  id="logoStyle" style="padding-top: 0px;" class="navbar-brand" href="/" onclick="return false" >freeFishGo</a>
                </div>
                <div class="collapse navbar-collapse" id="example-navbar-collapse">
                    <ul  id="head"  class="nav navbar-nav">
                        {{{range $k, $v := .homeHeadLi}}}  {{{if $v.Active}}} <li><a class="headActive heada"  href="{{{$v.Path}}}"  onclick="return false">{{{$v.Name}}}</a> </li> <script>document.title ='{{{$v.Name}}}';</script>  {{{else}}} <li>  <a href="{{{$v.Path}}}"  class="heada" onclick="return false">{{{$v.Name}}}</a> {{{end}}}</li> {{{end}}}
                    </ul>
                </div>
                </div>
            </nav>
        </div>

    </div>
    <div class="row" id="headHiv" style="margin-top: 57px"></div>
    <div id="homeContent" class="container main-container fr-element fr-view">
        ((.LayoutContent))
    </div>
    <div id="editDocs" class="row"></div>
    <input id="bianj" style="display: none;" type="button" value="编辑" class="btn"/>
    <div id="footerTop"></div>
    <footer id="footer" class="container main-container">
        <div id="footerSon">
            <span>©2020&nbsp;huzhouyu&nbsp;</span>
            <a href="http://www.beian.miit.gov.cn" target="_blank">
                <span>渝ICP备17012865号</span>
            </a>
            <a href="https://github.com/freefishgo/freefishgodoc"  target="_blank">本网站采用freeFishGo</a>
        </div>
    </footer>
</body>
</html>