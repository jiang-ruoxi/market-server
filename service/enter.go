package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/address"
	"github.com/flipped-aurora/gin-vue-admin/server/service/bad"
	"github.com/flipped-aurora/gin-vue-admin/server/service/banner"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/member"
	"github.com/flipped-aurora/gin-vue-admin/server/service/order"
	"github.com/flipped-aurora/gin-vue-admin/server/service/pay"
	"github.com/flipped-aurora/gin-vue-admin/server/service/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/tag"
	"github.com/flipped-aurora/gin-vue-admin/server/service/task"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	PkgTestServiceGroup pkgTest.ServiceGroup
	BannerServiceGroup  banner.ServiceGroup
	OrderServiceGroup   order.ServiceGroup
	TagServiceGroup     tag.ServiceGroup
	PayServiceGroup     pay.ServiceGroup
	MemberServiceGroup  member.ServiceGroup
	TaskServiceGroup    task.ServiceGroup
	BadServiceGroup     bad.ServiceGroup
	AddressServiceGroup address.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
