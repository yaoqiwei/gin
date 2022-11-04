package lib

import "github.com/gin-gonic/gin"

type RouterInfo struct {
	Uid         int64
	Path        string // 路径
	RouterName  string // 路由名称
	OperateType uint8  // 操作类型：1添加；2修改；3删除
}

var RouterInfoMap map[string]*RouterInfo

type RegisterRouterGroup struct {
	*gin.RouterGroup
}

func InitRouterInfoMap() {
	RouterInfoMap = map[string]*RouterInfo{}
}

func (r RegisterRouterGroup) Post(routerName string, operateType uint8, relativePath string, handlers ...gin.HandlerFunc) {
	r.RouterGroup.POST(relativePath, handlers...)
	if routerName != "" || operateType != 0 {
		RouterInfoMap[relativePath] = getRouterInfo(relativePath, routerName, operateType)
	}
}

func getRouterInfo(path, routerName string, operateType uint8) *RouterInfo {
	routerInfo := RouterInfo{}
	routerInfo.Path = path
	routerInfo.RouterName = routerName
	routerInfo.OperateType = operateType
	return &routerInfo
}
