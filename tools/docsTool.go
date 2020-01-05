package tools

import (
	"freefishgodoc/models"
	"strconv"
)

// EachDocsTree 遍历文档树生成html字符串
func EachDocsTree(tree models.DocsTree, index string) string {
	treeStr := ""
	if len(tree.NextDocsTree) == 0 {
		treeStr = "<li index='" + index + "' class='treeClick' href='" + tree.Path + "?type=xhrSon' >" + tree.Name + "</li>"
	} else {
		treeStr = `<li><span index='` + index + `' class='treeClick'  path="` + tree.Path + `?type=xhrSon" >` + tree.Name + `<i class="fa fa-plus fa-fw"></i></span>
		<ul>`
		for k, v := range tree.NextDocsTree {
			treeStr += EachDocsTree(v, index+"_"+strconv.Itoa(k))
		}
		treeStr += `</ul></li>`
	}
	return treeStr
}
