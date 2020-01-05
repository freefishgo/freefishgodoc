package models

// DocsTree 结构文档结构体
type DocsTree struct {
	Name         string
	Path         string
	NextDocsTree []DocsTree
}

// GetDocsTree 获取Freefish接口结构
func GetDocsTree() DocsTree {
	return DocsTree{Name: "FreeFishGo", Path: "/docs/index", NextDocsTree: []DocsTree{
		{Name: "简介", Path: "/docs/index", NextDocsTree: []DocsTree{
			{Name: "简介", Path: "/docs/index"},
			{Name: "安装升级", Path: "/docs/index"}}},
		{Name: "安装升级", Path: "/docs/index", NextDocsTree: []DocsTree{
			{Name: "简介", Path: "/docs/index"},
			{Name: "安装升级", Path: "/docs/index"}}}}}
}
