package models

// HomeHead ...
type HomeHead struct {
	Name   string
	Path   string
	Active bool
}

// GetHomeHeadList ...
func GetHomeHeadList(needActive string) []HomeHead {
	homeList := []HomeHead{{Name: "首页", Path: "/home"}, {Name: "快速入门", Path: "/quickStart"},
		{Name: "开发者社区", Path: "/community"}, {Name: "开发文档", Path: "/docs"}, {Name: "视频教程", Path: "/video"},
		{Name: "产品案例", Path: "/products"}, {Name: "官方博客", Path: "blog"}}
	for k, v := range homeList {
		if v.Name == needActive {
			homeList[k].Active = true
			break
		}
	}
	return homeList
}
