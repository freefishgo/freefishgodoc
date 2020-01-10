package tools

import (
	"freefishgodoc/models"
	"strconv"
	"strings"
)

// EachDocsTree 遍历文档树生成html字符串
func EachDocsTree(tree *models.DocsTree, index string, needActive string) string {
	treeStr := ""
	if len(tree.NextDocsTree) == 0 {
		if index == needActive {
			treeStr = "<li  class='treeClick treeActive' index='" + index + "' href='/docs/" + index + "?type=xhrSon' >" + tree.Name + "</li><script>document.title ='" + tree.Name + "';</script>"
		} else {
			treeStr = "<li  class='treeClick' index='" + index + "' href='/docs/" + index + "?type=xhrSon' >" + tree.Name + "</li>"
		}
	} else {
		if index == needActive {
			treeStr = `<li><span  class='treeClick treeActive' index='` + index + `' href="/docs/` + index + `?type=xhrSon" >` + tree.Name + `<i class="fa fa-plus fa-fw"></i></span>
			<ul><script>document.title ='` + tree.Name + `';</script>`
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

// 删掉指定索引上的节点
func DeleteTree(tree *models.DocsTree, pathIndex string, docsName string) bool {
	treeList := []*models.DocsTree{tree}
	pathList := strings.Split(pathIndex, "/")
	lens := len(pathList) - 1
	for k, v := range pathList {
		if v, err := strconv.Atoi(v); err != nil {
			return false
		} else {
			if lens == k {
				if len(treeList) > v {
					if treeList[v].Name == docsName {
						tree.NextDocsTree = append(treeList[0:v], treeList[v+1:]...)
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			}
			if len(treeList) > v {
				tree = treeList[v]
				treeList = treeList[v].NextDocsTree
			} else {
				return false
			}
		}
	}
	return false
}

// 删掉指定索引上的节点
func UpOrDownTree(tree *models.DocsTree, pathIndex string, docsName string, up bool) bool {
	treeList := []*models.DocsTree{tree}
	pathList := strings.Split(pathIndex, "/")
	lens := len(pathList) - 1
	for k, v := range pathList {
		if v, err := strconv.Atoi(v); err != nil {
			return false
		} else {
			if lens == k {
				if len(treeList) > v {
					if treeList[v].Name == docsName {
						tmpSlice := []*models.DocsTree{}
						if up {
							v1 := v - 1

							if v1 == -1 {
								return true
							} else {
								tmp := *treeList[v1]
								tmpSlice = append(treeList[0:v1], treeList[v])
								tmpSlice = append(tmpSlice, &tmp)
							}
							tmpSlice = append(tmpSlice, treeList[v+1:]...)
							tree.NextDocsTree = tmpSlice
						} else {
							tmp := *treeList[v]
							v1 := v - 1
							if v1 != -1 {
								tmpSlice = treeList[0:v1]
							}
							v2 := v + 1
							if v2 <= lens {
								tmpSlice = append(tmpSlice, treeList[v2])
							} else {
								return true
							}
							tmpSlice = append(tmpSlice, &tmp)
							tmpSlice = append(tmpSlice, treeList[v+2:]...)
							tree.NextDocsTree = tmpSlice
						}
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			}
			if len(treeList) > v {
				tree = treeList[v]
				treeList = treeList[v].NextDocsTree
			} else {
				return false
			}
		}
	}
	return false
}

// 获取指定节点的文章
func GetArticle(tree *models.DocsTree, pathIndex string) (string, bool) {
	treeList := []*models.DocsTree{tree}
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

// 设置指定节点的名字
func UpdateDocsName(tree *models.DocsTree, pathIndex string, docsName string) bool {
	treeList := []*models.DocsTree{tree}
	pathList := strings.Split(pathIndex, "/")
	for _, v := range pathList {
		if v, err := strconv.Atoi(v); err != nil {
			return false
		} else {
			if len(treeList) > v {
				tree = treeList[v]
				treeList = treeList[v].NextDocsTree
			} else {
				return false
			}
		}
	}
	tree.Name = docsName
	return true
}

// 添加子文档
func AddSonDoc(tree *models.DocsTree, pathIndex string, docsName string) bool {
	treeList := []*models.DocsTree{tree}
	pathList := strings.Split(pathIndex, "/")
	for _, v := range pathList {
		if v, err := strconv.Atoi(v); err != nil {
			return false
		} else {
			if len(treeList) > v {
				tree = treeList[v]
				treeList = treeList[v].NextDocsTree
			} else {
				return false
			}
		}
	}
	tree.NextDocsTree = append(tree.NextDocsTree, &models.DocsTree{Name: docsName, Content: docsName})
	return true
}

// 更新文章的内容
func UpdateDocContent(tree *models.DocsTree, pathIndex string, docsName string, content string) bool {
	treeList := []*models.DocsTree{tree}
	pathList := strings.Split(pathIndex, "/")
	for _, v := range pathList {
		if v, err := strconv.Atoi(v); err != nil {
			return false
		} else {
			if len(treeList) > v {
				tree = treeList[v]
				treeList = treeList[v].NextDocsTree
			} else {
				return false
			}
		}
	}
	if tree.Name == docsName {
		tree.Content = content
	} else {
		return false
	}
	return true
}
