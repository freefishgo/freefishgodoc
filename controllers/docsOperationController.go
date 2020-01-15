package controllers

import (
	"freefishgodoc/tools"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
)

type docsOperationController struct {
	mvc.Controller
}

func init() {
	mvc.AddHandlers(&docsOperationController{})
}

type docsTreeHelper struct {
	Index       string `json:"index"`
	Docsname    string `json:"docsname"`
	DocsContent string `json:"content"`
	Up          bool   `json:"up"`
}

// 删除开发文档的文档树
func (docs *docsOperationController) DeleteDoc(data *docsTreeHelper) {
	tree := tools.GetDocsTree()
	ok := tools.DeleteTree(tree, data.Index, data.Docsname)
	docs.Response.WriteJson(ok)
}

// 添加子文档
func (docs *docsOperationController) AddSonDoc(data *docsTreeHelper) {
	tree := tools.GetDocsTree()
	ok := tools.AddSonDoc(tree, data.Index, data.Docsname)
	docs.Response.WriteJson(ok)
}

// 更新开发文档的文档树文章
func (docs *docsOperationController) UpdateDocContentPost(data *docsTreeHelper) {
	tree := tools.GetDocsTree()
	ok := tools.UpdateDocContent(tree, data.Index, data.Docsname, data.DocsContent)
	docs.Response.WriteJson(ok)
}

// 更新开发文档的文档树名称
func (docs *docsOperationController) UpdateDocName(data *docsTreeHelper) {
	tree := tools.GetDocsTree()
	ok := tools.UpdateDocsName(tree, data.Index, data.Docsname)
	docs.Response.WriteJson(ok)
}

// 上移或者下移开发文档的文档树
func (docs *docsOperationController) UpOrDownDoc(data *docsTreeHelper) {
	var ok bool
	if data.Index != "" {
		tree := tools.GetDocsTree()
		ok = tools.UpOrDownTree(tree, data.Index, data.Docsname, data.Up)
	}
	docs.Response.WriteJson(ok)
}

// 控制器执行前调用
func (docsOperation *docsOperationController) Prepare() {
	userinfo, _ := docsOperation.Response.GetSession("userinfo")
	userip, _ := docsOperation.Response.GetSession("userip")
	if userinfo != nil && userip != nil {
		if userip != docsOperation.Request.Host {
			docsOperation.Response.Write([]byte("违规操作"))
			docsOperation.SkipController()
		}
	} else {
		docsOperation.Response.Write([]byte("登录过期了"))
		docsOperation.SkipController()
	}
}
