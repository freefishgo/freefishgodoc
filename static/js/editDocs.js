/// <reference path="../js/jquery-3.3.1.min.js" />
/// <reference path="../js/share.js" />
/// <reference path="../js/froala_editor.pkgd.min.js"/>
$(function(){
    var nowIndex=""
    $(".treeClick").dblclick(function () {
        nowIndex= $(this).attr("index");
        $("#docName").html($(this).text());
        $('#myModal').modal('show');
    });
    $("#dtldoc").click(function(){
        $.get("/docsOperation/DeleteDoc?index="+nowIndex+"&docsname="+$("#docName").text(),function(data,std){
            location.href="/docs/"+nowIndex
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
        editor=new FroalaEditor('#ArticleCententArea');
        $.get("/docs/"+nowIndex+"?type=xhrSon",function(data,std){
            editor.html.set(data)
        });
        $('#myModal').modal('hide');
        $("#save").click(function(){
            var obj={};
            obj.index=nowIndex;
            obj.docsname=$("#docName").text();
            obj.content=editor.html.get();
            $.post("/docsOperation/UpdateDocContent",obj,function(data,std){
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
        var html='<textarea id="contentArticle"></textarea><button id="saveContent" class="btn btn-default">保存</button>';
        $("#content").html(html);
        editor=new FroalaEditor('#contentArticle');
        var href =(document.location+"").split('?')[0];
        $.get(href+"/GetEditContent",function(data,std){
            editor.html.set(data)
        });
        $("#saveContent").click(function(){
            var obj={};
            obj.content=editor.html.get();
            $.post(href+"/save",obj,function(data,std){
                location.href=href;
            });
        }); 
    });
});