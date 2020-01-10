package models

// DocsTree 结构文档结构体
type DocsTree struct {
	Name         string
	Path         string
	Content      string `json:"-"`
	NextDocsTree []*DocsTree
}
