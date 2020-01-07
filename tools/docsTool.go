package tools

import (
	"freefishgodoc/models"
	"strconv"
	"strings"
)

// EachDocsTree 遍历文档树生成html字符串
func EachDocsTree(tree models.DocsTree, index string, needActive string) string {
	treeStr := ""
	if len(tree.NextDocsTree) == 0 {
		if index == needActive {
			treeStr = "<li  class='treeClick treeActive' index='" + index + "' href='/docs/" + index + "?type=xhrSon' >" + tree.Name + "</li>"
		} else {
			treeStr = "<li  class='treeClick' index='" + index + "' href='/docs/" + index + "?type=xhrSon' >" + tree.Name + "</li>"
		}
	} else {

		if index == needActive {
			treeStr = `<li><span  class='treeClick treeActive' index='` + index + `' href="/docs/` + index + `?type=xhrSon" >` + tree.Name + `<i class="fa fa-plus fa-fw"></i></span>
			<ul>`
		} else {
			treeStr = `<li><span  class='treeClick' index='` + index + `' href="/docs/` + index + `?type=xhrSon" >` + tree.Name + `<i class="fa fa-plus fa-fw"></i></span>
			<ul>`
		}
		for k, v := range tree.NextDocsTree {
			treeStr += EachDocsTree(v, index+"/"+strconv.Itoa(k), needActive)
		}
		treeStr += `</ul></li>`
	}
	return treeStr
}

func GetArticle(tree models.DocsTree, pathIndex string) (string, bool) {
	treeList := []models.DocsTree{tree}
	pathList := strings.Split(pathIndex, "/")
	for _, v := range pathList {
		if v, err := strconv.Atoi(v); err != nil {
			return "", false
		} else {
			if len(treeList) > v {
				tree = treeList[v]
				treeList = treeList[v].NextDocsTree
			} else {
				return "", false
			}
		}
	}
	return tree.Content, true
}
