package controllers

import (
	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type userController struct {
	mvc.Controller
}

func init() {
	mvc.AddHandlers(&userController{})
}

func (user *userController) Login() {
	user.Response.RemoveSession()
	user.UseTplPath()
}

type userInfo struct {
	Code     string `json:"code"`
	Pwd      string `json:"pwd"`
	UserName string `json:"username"`
}

func (user *userController) DoLoginPost(u *userInfo) {
	tmp := struct {
		UserName string `json:"username"`
		IsLogin  bool   `json:"islogin"`
	}{IsLogin: false}
	if u.UserName == "huzhouyu" && u.Pwd == "123456" {
		user.Response.SetSession("userinfo", u)
		user.Response.SetSession("userip", user.Request.Host)
		tmp.UserName = u.UserName
		tmp.IsLogin = true
	}
	user.Response.WriteJson(tmp)
}

func (user *userController) GetEdit() {
	u := user.Response.GetSession("userinfo")
	if u != nil {
		user.Response.Write([]byte(`
		<!-- 模态框（Modal） -->
	<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<button type="button" class="close" data-dismiss="modal" aria-hidden="true">
						&times;
					</button>
					<h4 class="modal-title" id="myModalLabel">
						开发文档树操作
					</h4>
				</div>
				<div class="modal-body">
					<h2 id="docName"></h2>
					<label for="docname">修改当前文档树名字:</label><input id="updateDocNameVal" name="docname" type="text"/><button id="updateDocName">修改</button>
					<br/>
					<button id="dtldoc">删除当前文档树</button>
					<br/>
					<button id="updoc">上移当前文档树</button>
					<br/>
					<button id="downdoc">下移当前文档树</button>
					<br>
					<label for="docsonname">添加子文档树名字:</label><input id="docsonname" name="docsonname" type="text"/><button id="docsonnamebtn">添加</button>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
					<button type="button" id="editAtc" class="btn btn-default">编辑文章</button>
				</div>
			</div><!-- /.modal-content -->
		</div><!-- /.modal -->
	</div>
		<script src="/js/editDocs.js"></script>`))
	}
}
