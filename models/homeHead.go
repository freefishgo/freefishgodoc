package models

type HomeHead struct {
	Name   string
	Path   string
	Active bool
}

func GetHomeHeadList(needActive string) []HomeHead {
	homeList := []HomeHead{{Name: "首页", Path: "/home/index"}, {Name: "快速入门", Path: "/docs/index"},
		{Name: "开发者社区", Path: "/docs/index"}, {Name: "开发文档", Path: "/docs/index"}, {Name: "视频教程", Path: "/docs/index"},
		{Name: "产品案例", Path: "/docs/index"}, {Name: "官方博客", Path: "/docs/index"}}
	for k, v := range homeList {
		if v.Name == needActive {
			homeList[k].Active = true
			break
		}
	}
	return homeList
}
