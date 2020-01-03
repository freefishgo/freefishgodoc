package models

type DocsTree struct {
	Name         string
	Path         string
	NextDocsTree []DocsTree
}

func GetDocsTree() DocsTree {
	return DocsTree{Name: "FreeFishGo", Path: "/docs/index", NextDocsTree: []DocsTree{
		{Name: "简介", Path: "/docs/index", NextDocsTree: []DocsTree{
			{Name: "简介", Path: "/docs/index"},
			{Name: "安装升级", Path: "/docs/index"}}},
		{Name: "安装升级", Path: "/docs/index", NextDocsTree: []DocsTree{
			{Name: "简介", Path: "/docs/index"},
			{Name: "安装升级", Path: "/docs/index"}}}}}
}
