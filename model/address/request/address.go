package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/address"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AddressSearch struct{
    address.Address
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
