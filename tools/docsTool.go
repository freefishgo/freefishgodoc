package tools

import "freefishgodoc/models"

func EachDocsTree(tree models.DocsTree) string {
	treeStr := ""
	if len(tree.NextDocsTree) == 0 {
		treeStr = "<li><a href='" + tree.Path + "' >" + tree.Name + "</a></li>"
	} else {
		treeStr = `<li><span path="` + tree.Path + `" ><a href='` + tree.Path + `' >` + tree.Name + `</a><i class="fa fa-plus fa-fw"></i></span>
		<ul>`
		for _, v := range tree.NextDocsTree {
			treeStr += EachDocsTree(v)
		}
		treeStr += `</ul></li>`
	}
	return treeStr
}
