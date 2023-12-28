package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/banner"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BannersSearch struct{
    banner.Banners
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
