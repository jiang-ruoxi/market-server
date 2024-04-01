package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/bad"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BadWordsSearch struct{
    bad.BadWords
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
