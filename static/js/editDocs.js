/// <reference path="../js/jquery-3.3.1.min.js" />
/// <reference path="../js/share.js" />
/// <reference path="../froala_editor_3.1.0/js/froala_editor.pkgd.min.js"/>
var froalaOption ={
    language: 'zh_cn',
    height: 500,
    tableStyles: {
        'fr-dashed-borders':'dashed-borders',
        'fr-alternate-rows':'alternate-rows',       
      },
    tableCellStyles: {
        'fr-highlighted': 'Highlighted',
        'fr-thick': 'Thick',
        'noTableCellBorder':'noTableCellBorder',
        'noTableCellBorder-left':'noTableCellBorder-left',
        'noTableCellBorder-right':'noTableCellBorder-right',
        'noTableCellBorder-top':'noTableCellBorder-top',
        'noTableCellBorder-bottom':'noTableCellBorder-bottom',
    },
    linkStyles: {
        'fr-green': 'Green',
        'fr-strong': 'Thick',
        'treeContentClick':'treeContentClick'
      },
      linkAttributes: {
        title: 'Title',
        onclick:"return true",
      },
    // disable quick insert
    quickInsertTags: [],

    // toolbar buttons
    //toolbarButtons: ['fullscreen', 'bold', 'italic', 'underline', 'strikeThrough', '|', 'paragraphFormat', 'fontSize', 'color', '|', 'align', 'formatOL', 'formatUL', 'outdent', 'indent', 'quote', '-', 'insertLink', 'insertFile', 'insertImage', 'insertVideo', 'embedly', 'insertTable', '|', 'insertHR', 'selectAll', 'clearFormatting', '|', 'spellChecker', 'help', 'html', '|', 'undo', 'redo'],
    // toolbarButtons: {
    //     'moreText': {
    //       'buttons': ['bold', 'italic', 'underline', 'strikeThrough', 'subscript', 'superscript', 'fontFamily', 'fontSize', 'textColor', 'backgroundColor', 'inlineClass', 'inlineStyle', 'clearFormatting']
    //     },
    //     'moreParagraph': {
    //       'buttons': ['alignLeft', 'alignCenter', 'formatOLSimple', 'alignRight', 'alignJustify', 'formatOL', 'formatUL', 'paragraphFormat', 'paragraphStyle', 'lineHeight', 'outdent', 'indent', 'quote']
    //     },
    //     'moreRich': {
    //       'buttons': ['insertLink', 'insertImage', 'insertVideo', 'insertTable', 'emoticons', 'fontAwesome', 'specialCharacters', 'embedly', 'insertFile', 'insertHR']
    //     },
    //     'moreMisc': {
    //       'buttons': ['undo', 'redo', 'fullscreen', 'print', 'getPDF', 'spellChecker', 'selectAll', 'html', 'help'],
    //       'align': 'right',
    //       'buttonsVisible': 2
    //     }
    //   },
    // upload file
    fileUploadParam: 'file',
    fileUploadURL: '/media/UploadFile',
    fileUploadMethod: 'POST',
    fileMaxSize: 20 * 1024 * 1024,
    fileAllowedTypes: ['*'],

    // upload image
    // imageUploadParams:{type1:"img"},
    imageUploadParam: 'file',
    imageUploadURL: '/media/UploadFile?type=img',
    imageUploadMethod: 'POST',
    imageMaxSize: 20 * 1024 * 1024,
    imageAllowedTypes: ['jpeg', 'jpg', 'png', 'gif'],

    // upload video
    videoUploadParam: 'file',
    videoUploadURL: '/media/UploadFile',
    videoUploadMethod: 'POST',
    videoMaxSize: 50 * 1024 * 1024,
    videoAllowedTypes: ['avi', 'mov', 'mp4', 'm4v', 'mpeg', 'mpg', 'wmv', 'ogv'],
}
$(function(){
    var nowIndex=""
    $(".treeClick").dblclick(function () {
        nowIndex= $(this).attr("index");
        $("#docName").html($(this).text());
        $('#myModal').modal('show');
    });
    $("#dtldoc").click(function(){
        $.get("/docsOperation/DeleteDoc?index="+nowIndex+"&docsname="+$("#docName").text(),function(data,std){
            var list=nowIndex.split("/");
            nowIndex=""
            for (i=0;i<list.length-1;i++){
                nowIndex+=list[i]+"/";
            }
            location.href="/docs/"+nowIndex.trim("/");
            $('#myModal').modal('hide');
        });

    });
    $("#updateDocName").click(function(){
        $.get("/docsOperation/UpdateDocName?index="+nowIndex+"&docsname="+$("#updateDocNameVal").val(),function(data,std){
            location.href="/docs"
            $('#myModal').modal('hide');
        });
    });
    $("#updoc").click(function(){
        var obj={};
        obj.up=true;
        obj.index=nowIndex;
        obj.docsname=$("#docName").text();
        $.get("/docsOperation/UpOrDownDoc",obj,function(data,std){
            location.href="/docs/"+nowIndex
            $('#myModal').modal('hide');
        });
    });
    $("#downdoc").click(function(){
        var obj={};
        obj.up=false;
        obj.index=nowIndex;
        obj.docsname=$("#docName").text();
        $.get("/docsOperation/UpOrDownDoc",obj,function(data,std){
            location.href="/docs/"+nowIndex
            $('#myModal').modal('hide');
        });
    });
    $("#docsonnamebtn").click(function(){
        var obj={};
        obj.index=nowIndex;
        obj.docsname=$("#docsonname").val();
        $.get("/docsOperation/AddSonDoc",obj,function(data,std){
            location.href="/docs/"+nowIndex
            $('#myModal').modal('hide');
        });
    });

    $("#editAtc").click(function(){
        var html='<textarea id="ArticleCententArea"></textarea><button id="save" class="btn btn-default">保存</button>';
        $("#ArticleCentent").html(html);
        editor=new FroalaEditor('#ArticleCententArea',froalaOption);
        $('#myModal').modal('hide');
        $.get("/docs/"+nowIndex+"?type=xhrSon",function(data,std){
            editor.html.set(data)
        });
        $("#save").click(function(){
            var obj={};
            obj.index=nowIndex;
            obj.docsname=$("#docName").text();
            obj.content=editor.html.get().replace(/<a class="treeContentClick" href="/g,'<a class="treeContentClick" onclick="return false" href="');
            $.post("/docsOperation/UpdateDocContent",obj,function(data,std){
                editor.destroy();
                location.href="/docs/"+nowIndex
                $('#myModal').modal('hide');
            });
        });    
    });
    $('#myModal').on('show.bs.modal', function(){
        var $this = $(this);
        var $modal_dialog = $this.find('.modal-dialog');
        $this.css('display', 'block');
        $modal_dialog.css({'margin-top': Math.max(0, ($(window).height() - $modal_dialog.height()) / 2) });       
    });
    $("#bianj").click(function(){
        $("#bianj").hide();
        $(".contentMargin").removeClass("contentMargin");
        var html='<textarea id="contentArticle"></textarea><button id="saveContent" class="btn btn-default">保存</button>';
        $("#content").html(html);
        editor=new FroalaEditor('#contentArticle',froalaOption);
        var href =(document.location+"").split('?')[0];
        $.get(href+"/GetEditContent",function(data,std){
            editor.html.set(data)
        });
        $("#saveContent").click(function(){
            var obj={};
            obj.content=editor.html.get();
            $.post(href+"/save",obj,function(data,std){
                editor.destroy()
                location.href=href;
            });
        }); 
    });
});