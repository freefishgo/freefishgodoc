/// <reference path="../js/jquery-3.3.1.min.js" />
/// <reference path="../js/share.js" />
$(function(){
    var nowIndex=""
    $(".treeClick").dblclick(function () {
        nowIndex= $(this).attr("index");
        $("#docName").html($(this).text());
        $("#dtldoc").click(function(){
            $.get("/docs/DeleteDoc?index="+nowIndex,function(data,std){
                if (data) {

                }
            });
        });
        $('#myModal').modal('show');
    })
    $('#myModal').on('show.bs.modal', function(){
        var $this = $(this);
        var $modal_dialog = $this.find('.modal-dialog');
        $this.css('display', 'block');
        $modal_dialog.css({'margin-top': Math.max(0, ($(window).height() - $modal_dialog.height()) / 2) });       
    });
});