<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>login</title>
    <link rel="shortcut icon" href="/img/fish.ico" type="image/x-icon" />
    <style type="text/css">
        *{
            margin: 0;
            padding: 0;
        }
        #wrap {
            height: 719px;
            width: 100;
            /*background-image: url(4.jpg);*/
            background-repeat: no-repeat;
            background-position: center center;
            position: relative;
        }
        #head {
            height: 120px;
            width: 100;
            background-color: #66CCCC;
            text-align: center;
            position: relative;
        }
        #foot {
            width: 100;
            height: 126px;
            background-color: #CC9933;
            position: relative;
        }
        #wrap .logGet {
            height: 408px;
            width: 368px;
            position: absolute;
            background-color: #FFFFFF;
            top: 20%;
            right: 15%;
        }
        .logC a button {
            width: 100%;
            height: 45px;
            background-color: #ee7700;
            border: none;
            color: white;
            font-size: 18px;
        }
        .logGet .logD.logDtip .p1 {
            display: inline-block;
            font-size: 28px;
            margin-top: 30px;
            width: 86%;
        }
        #wrap .logGet .logD.logDtip {
            width: 86%;
            border-bottom: 1px solid #ee7700;
            margin-bottom: 60px;
            margin-top: 0px;
            margin-right: auto;
            margin-left: auto;
        }
        .logGet .lgD img {
            position: absolute;
            top: 12px;
            left: 8px;
        }
        .logGet .lgD input {
            width: 100%;
            height: 42px;
            text-indent: 2.5rem;
        }
        #wrap .logGet .lgD {
            width: 86%;
            position: relative;
            margin-bottom: 30px;
            margin-top: 30px;
            margin-right: auto;
            margin-left: auto;
        }
        #wrap .logGet .logC {
            width: 86%;
            margin-top: 0px;
            margin-right: auto;
            margin-bottom: 0px;
            margin-left: auto;
        }


        .title {
            font-family: "宋体";
            color: #FFFFFF;
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);  /* 使用css3的transform来实现 */
            font-size: 36px;
            height: 40px;
            width: 30%;
        }
    </style>
</head>

<body>
<div class="header" id="head">
    <div class="title">FreeFishGo</div>

</div>

<div class="wrap" id="wrap">
    <div id="app" class="logGet">
        <!-- 头部提示信息 -->
        <div class="logD logDtip" style="border-bottom: 0px">
            <p class="p1" >登录</p>
        </div>
        <!-- 输入框 -->
        <div class="lgD">
            <img src="/img/fish.jpg" width="20" height="20" alt=""/>
            <input type="text"
                   placeholder="输入用户名" v-model="username" />
        </div>
        <div class="lgD">
            <img src="/img/fish.jpg" width="20" height="20" alt=""/>
            <input type="password"
                   placeholder="输入用户密码" v-model="pwd" />
        </div>
        <div class="logC">
            <a><button  @click="doLogin()" >登 录</button></a>
        </div>
    </div>
</div>
</body>
<script src="/js/vue.min.js"></script>
<script src="/js/axios.min.js"></script>
<script src="/js/share.js"></script>
<script>
    new Vue({
        el: '#app',
        data: {
            username: "",
            pwd:"",
            code:""
        },
        methods: {
            reverseMessage: function () {
                this.message = this.message.split('').reverse().join('')
            },
            doLogin:function () {
                axios({
                    method: "post",
                    url: "/user/dologin",
                    // headers: {
                    // "content-type": "application/x-www-form-urlencoded"
                    // },
                    data:queryString(this.$data)
                })
                .then(response => {
                    var data=response.data
                    if(data.islogin){
                        window.location.href="/back/index"
                    } else {
                        alert("密码错误")
                    }
                })
                .catch(function (error) { // 请求失败处理
                    console.log(error);
                });
            }
        }
    })
</script>
</html>