<!DOCTYPE HTML>
<html>
<head>
    <!-- Include Editor style. -->
    <title>创建文章</title>
</head>
<body>
<!-- Include Editor style. -->
<link href="/static/css/froala_editor.pkgd.min.css" rel="stylesheet" type="text/css" />
<!-- Create a tag that we will use as the editable area. -->
<!-- You can use a div tag as well. -->
<textarea id="edit"></textarea>
<div id="edit1"></div>
<button id="save">获取代码</button>
<!-- Include Editor JS files. -->
<script src="/static/js/jquery-3.3.1.min.js"></script>
<script type="text/javascript" src="/static/js/froala_editor.pkgd.min.js"></script>

<!-- Initialize the editor. -->
<script>
    $(function() {
        editor=new FroalaEditor('#edit');

        $("#save").click(function(){
            $("#edit1").html(editor.html.get())
            console.log(editor.html.get())
        })
    });

</script>
</body>