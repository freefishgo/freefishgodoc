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
<textarea></textarea>


<!-- Include Editor JS files. -->
<script src="/static/js/jquery-3.3.1.min.js"></script>
<script type="text/javascript" src="/static/js/froala_editor.pkgd.min.js"></script>

<!-- Initialize the editor. -->
<script>
    f=new FroalaEditor('textarea');

    $(function() {
        // var editCont = '<p>adddddp</p>';
        // $.FroalaEditor.DefineIcon('color', {SRC: 'txt_color@2x.png', ALT: 'Image button', template: 'image'}); //自定义图标
        // $('#edit').froalaEditor({
        //     autofocus: true,
        //     toolbarButtonsXS: ['undo', 'redo', '|', 'bold', 'italic', 'underline', '|', 'fontSize', 'align', 'color'],
        //     language: 'zh_cn',
        //     colorsHEXInput: false,
        //     colorsBackground: ['#2E2E2E', '#767676', '#DF281B', '#F4821C', '#46AC43', '#2E5BF7', '#A249B3', 'REMOVE'],
        //     colorsText: ['#2E2E2E', '#767676', '#DF281B', '#F4821C', '#46AC43', '#2E5BF7', '#A249B3', 'REMOVE'],
        //     fontSize: ['14', '16', '18', '20'],
        //     fontSizeDefaultSelection: '16',
        //     height: 220,
        //     htmlAllowComments: false,
        //     pasteAllowedStyleProps: ['font-size', 'color'],
        //     placeholderText: '请输入内容',
        //     charCounterMax: 500
        // }).froalaEditor('html.set', editCont);
        $("#save").click(function(){
            $('#edit').froalaEditor('html.get', true); //返回富文本编辑里面的html代码
        })
    });

</script>
</body>