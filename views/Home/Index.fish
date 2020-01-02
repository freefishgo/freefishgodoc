<!DOCTYPE html>
<html>
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta charset="UTF-8">
    <title>FreeFishGo</title>
    <link rel="shortcut icon" href="/static/img/fish.ico" type="image/x-icon" />
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.css">
    <link rel="stylesheet" href="/static/tree/style.css" >
</head>
<style type="text/css">
    #headDiv{
        position: fixed;
        left: 0px;
        top:0px;
        width: 110%;
        border:lightgray 1px solid;
        z-index: 9999;
        background-color: whitesmoke;
    }
    #logoStyle{
        display: inline;
        margin-left: 25%;
    }
    #head li {
        display:inline;
        margin: 0;
        padding: 0;
        color: #428bca;
    }
    #head li h4{
        display:inline;
        margin-left: 20px;
        padding: 0;
        cursor:pointer ;
    }
    body{
        /*background-color: #f5f5f5;*/
    }
    .headActive{
        color: red;
    }
</style>
<body>
<div class="col-md-offset-1 col-md-10">
    <nav id="headDiv" class="row">
        <div>
            <h1 id="logoStyle">FreeFishGo</h1>
            <ul id="head" style="display: inline">
                <li><h4>首页</h4></li>
                <li><h4>快速入门</h4></li>
                <li><h4>开发者社区</h4></li>
                <li><h4  class="headActive" href="/docs">开发文档</h4></li>
                <li><h4>视频教程</h4></li>
                <li><h4>产品案例</h4></li>
                <li><h4>官方博客</h4></li>
            </ul>
        </div>
    </nav>
    <div class="row" style="margin-top: 80px"></div>
    <div class="row">
        <div class="col-md-3">
            <div class="tree">
                <ul>
                    <li>
                        <span>JavaScript<i class="fa fa-plus fa-fw"></i></span>
                        <ul>
                            <li>
                                <span>语言基础<i class="fa fa-plus fa-fw"></i></span>
                                <ul>
                                    <li>原型链</li>
                                    <li>作用域链和闭包</li>
                                    <li>this</li>
                                    <li>AJAX和Promise/Deferred</li>
                                    <li>数组操作</li>
                                </ul>
                            </li>
                            <li>
                                <span>类与继承<i class="fa fa-plus fa-fw"></i></span>
                                <ul>
                                    <li>Object.create</li>
                                    <li>klass等模拟类</li>
                                </ul>
                            </li>
                            <li>jQuery、Bootstrap等框架</li>
                            <li>
                                <span>MV*<i class="fa fa-plus fa-fw"></i></span>
                                <ul style="display: none;">
                                    <li>代码分层</li>
                                    <li>Angular.js</li>
                                    <li>Backbone.js</li>
                                </ul>
                            </li>
                            <li>
                                <span>模块化<i class="fa fa-plus fa-fw"></i></span>
                                <ul >
                                    <li>AMD和Require.js</li>
                                    <li>CMD和Sea.js</li>
                                    <li>ES6 Harmony的模块</li>
                                </ul>
                            </li>
                            <li>
                                <span>DOM和BOM操作<i class="fa fa-plus fa-fw"></i></span>
                                <ul>
                                    <li>节点的增删查改</li>
                                    <li>事件绑定与事件代理</li>
                                </ul>
                            </li>
                            <li>
                                <span>性能优化<i class="fa fa-plus fa-fw"></i></span>
                                <ul>
                                    <li>语法优化</li>
                                    <li>AJAX缓存和备忘录模式</li>
                                    <li>DOM操作优化</li>
                                    <li>作用域链和原型链上的优化</li>
                                </ul>
                            </li>
                            <li>浏览器兼容性</li>
                        </ul>
                    </li>
                </ul>
            </div>
        </div>
        <div class="col-md-9" style="border: lightgray 1px solid;background-color: #f5f5f5;border-radius: 3px;height: 1000px">
            fafsdfafdafafdsaf
        </div>
    </div>
</div>
<script src="/static/js/jquery-3.3.1.min.js"></script>
<script src="/static/tree/index.js"></script>
<script src="/static/bootstrap/js/bootstrap.min.js"></script>
<script>
</script>
</body>
</html>