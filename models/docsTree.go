package models

// DocsTree 结构文档结构体
type DocsTree struct {
	Name         string
	Path         string
	Content      string
	NextDocsTree []DocsTree
}

// GetDocsTree 获取Freefish接口结构
func GetDocsTree() DocsTree {
	return DocsTree{Name: "FreeFishGo", Path: "/docs/index", Content: "FreeFishGo", NextDocsTree: []DocsTree{
		{Name: "简介", Path: "/docs/index", Content: "简介", NextDocsTree: []DocsTree{
			{Name: "简介", Path: "/docs/index", Content: "简介"},
			{Name: "安装升级", Path: "/docs/index", Content: "安装升级"}}},
		{Name: "安装升级", Path: "/docs/index", Content: "安装升级", NextDocsTree: []DocsTree{
			{Name: "简介", Path: "/docs/index", Content: "简介"},
			{Name: "安装升级", Path: "/docs/index", Content: "安装升级"}}}}}
}
