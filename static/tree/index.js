/// <reference path="../js/jquery-3.3.1.min.js" />
/// <reference path="../js/share.js" />
// $(function() {
// 	$(".tree span").append($("<i class='fa fa-plus fa-fw'></i>"));
// 	//$(".tree>ul ul").hide();
// 	$(".tree span").click(function(event) {
// 		var $this = $(this);
// 		var $ul = $this.parent().children("ul");
//
// 		if ($ul.length > 0) {
// 			if($ul.is(":visible")){
// 				$ul.slideUp();
// 				$this.children("i").removeClass("fa-minus").addClass("fa-plus");
// 			} else {
// 				$ul.slideDown();
// 				$this.children("i").removeClass("fa-plus").addClass("fa-minus");
// 			}
// 		}
// 	});
// });
var treeClick=function(){
    $(".treeContentClick").click(function(){
        var href= $(this).attr("href")
        $.get(href+"?type=xhrSon",function(data,std){
            $("#ArticleCentent").html(data);
            treeClick();
            var h=$(".treeClick")
            h.each(function () {
                if($(this).attr("href")==href){
                    $(this).addClass("treeActive")
                    var title=$(this).text();
                    updateUrl(title,href.split("?")[0])
                    document.title=title;
                } else{
                    $(this).removeClass("treeActive")
                }
            })
            return false;
        });
    });
}
$(function () {
    $(".treeClick").click(function () {
        var href= $(this).attr("href")
        $.get(href+"?type=xhrSon",function(data,std){
            $("#ArticleCentent").html(data);
            treeClick();
        });
        var h=$(".treeClick")
        h.each(function () {
            if($(this).attr("href")==href){
                $(this).addClass("treeActive")
                var title=$(this).text();
                updateUrl(title,href.split("?")[0])
                document.title=title;
            } else{
                $(this).removeClass("treeActive")
            }
        });
        return false;
    });
    treeClick();
})