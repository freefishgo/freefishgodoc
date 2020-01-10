
// vue格式化请求参数
function queryString(data){
    let params='';
    for(let index in data){
        params+=index+'='+data[index]+'&';
    }
    return params;
}
// 更新浏览地址栏url
function updateUrl(title,newUrl) {
    var stateObject = {};
    history.pushState(stateObject,title,newUrl);
}
window.onpopstate = function (event) {
    location.href= this.document.location;
};