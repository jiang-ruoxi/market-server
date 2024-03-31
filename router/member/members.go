package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MembersRouter struct {
}

// InitMembersRouter 初始化 zmUser表 路由信息
func (s *MembersRouter) InitMembersRouter(Router *gin.RouterGroup) {
	membersRouter := Router.Group("members").Use(middleware.OperationRecord())
	membersRouterWithoutRecord := Router.Group("members")
	var membersApi = v1.ApiGroupApp.MemberApiGroup.MembersApi
	{
		membersRouter.DELETE("deleteMembers", membersApi.DeleteMembers) // 删除zmUser表
		membersRouter.DELETE("deleteMembersByIds", membersApi.DeleteMembersByIds) // 批量删除zmUser表
	}
	{
		membersRouterWithoutRecord.GET("findMembers", membersApi.FindMembers)        // 根据ID获取zmUser表
		membersRouterWithoutRecord.GET("getMembersList", membersApi.GetMembersList)  // 获取zmUser表列表
	}
}
