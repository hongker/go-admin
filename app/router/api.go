package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go.uber.org/dig"

	"go-admin/app/api"
)

func RegisterAPI(r *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(
		user *api.User,
		role *api.Role,
	) error {
		{
			r.POST("/register", user.Register)
			r.POST("/login", user.Login)
		}

		admin := r.Group("/admin")
		{
			admin.POST("/roles", role.CreateRole)
		}

		//--------------------------------API-----------------------------------
		apiV1 := r.Group("/api/v1")
		{
			apiV1.GET("/users/:id", user.GetUserByID)
		}
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return err
}
