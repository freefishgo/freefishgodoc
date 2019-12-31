function queryString(data){
    let params='';
    for(let index in data){
        params+=index+'='+data[index]+'&';
    }
    return params;
}