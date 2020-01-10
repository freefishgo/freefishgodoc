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
    $("#head li h4").click(function () {
        var h=$("#head li h4")
        h.each(function () {
            $(this).removeClass("headActive")
        })
        $(this).addClass("headActive")
        url=$(this).attr("href")
        $.get(url,function (result,std) {
            $("#homeContent").html(result)
        })
        updateUrl($(this).text(),url.split("?")[0])
        document.title=$(this).text();
    })
})
