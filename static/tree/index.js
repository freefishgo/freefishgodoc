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
$(function () {
    $(".treeClick").click(function () {
        var href= $(this).attr("href")
        $.get(href,function(data,std){
            $("#ArticleCentent").html(data)
        });
        var title=$(this).text();
        // if (title=="FreeFishGo"){
        //     title="开发文档";
        // }
        updateUrl(title,href.split("?")[0])
        document.title=title;
        var h=$(".treeClick")
        h.each(function () {
            $(this).removeClass("treeActive")
        })
        $(this).addClass("treeActive")
    })
})