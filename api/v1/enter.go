package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/banner"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/member"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/order"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/pay"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/tag"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/task"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	PkgTestApiGroup pkgTest.ApiGroup
	BannerApiGroup  banner.ApiGroup
	OrderApiGroup   order.ApiGroup
	TagApiGroup     tag.ApiGroup
	PayApiGroup     pay.ApiGroup
	MemberApiGroup  member.ApiGroup
	TaskApiGroup    task.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
