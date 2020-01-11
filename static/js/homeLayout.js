$(function () {
    $("#logoStyle").dblclick(function(){
        $.get("/user/getedit",function(data,std){
            $("#editDocs").html(data);
            if(data!=""&&data!=undefined){
                $("#bianj").show()
            } else {
                $("#bianj").hide()
            }
        })
    });
    $("#head .heada").click(function () {
        var h=$("#head .heada")
        h.removeClass("headActive")
        $(this).addClass("headActive")
        url=$(this).attr("href")
        $.get(url+"?type=xhr",function (result,std) {
            $("#homeContent").html(result)
        })
        updateUrl($(this).text(),url.split("?")[0])
        document.title=$(this).text();
    });
    function changeHev(){
        $("#headHiv").css("margin-top",$("#headDivHei").outerHeight()+28+"px");
    }
    $("#headDivHei").click(function(){
        setTimeout(changeHev,370)
    });
    changeHev();
})
