package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/pay"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PaysSearch struct{
    pay.Pays
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
