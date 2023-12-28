package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/banner"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/member"
	"github.com/flipped-aurora/gin-vue-admin/server/router/order"
	"github.com/flipped-aurora/gin-vue-admin/server/router/pay"
	"github.com/flipped-aurora/gin-vue-admin/server/router/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/tag"
	"github.com/flipped-aurora/gin-vue-admin/server/router/task"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	PkgTest pkgTest.RouterGroup
	Banner  banner.RouterGroup
	Order   order.RouterGroup
	Tag     tag.RouterGroup
	Pay     pay.RouterGroup
	Member  member.RouterGroup
	Task    task.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
