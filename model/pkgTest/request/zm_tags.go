package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ZmTagsSearch struct{
    pkgTest.ZmTags
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
