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
	$("#head li h4").click(function () {
		var h=$("#head li h4")
		h.each(function () {
			$(this).removeClass("headActive")
		})
		$(this).addClass("headActive")
	})
})
