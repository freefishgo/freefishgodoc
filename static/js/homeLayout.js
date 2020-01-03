$(function () {
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
        updateUrl($(this).innerText,url.split("?")[0])
    })
})
