package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/immanuel-supanova/go-auth/controllers"
	"github.com/immanuel-supanova/go-auth/middleware"
)

func UserRoutes(r *gin.Engine) {
	r.Use(middleware.LoggerMiddleware(gin.LoggerConfig{}, "go-auth"))

	r.POST("/create", controllers.UserCreate)
	r.POST("/create-admin", controllers.UserCreateAdmin)
	r.POST("/create-staff", controllers.UserCreateStaff)
	r.GET("/read", controllers.UserRead)
	r.GET("/list", controllers.UserList)
	r.PUT("/isactive", controllers.UserIsActiveChange)
	r.PUT("/isstaff", controllers.UserIsStaffChange)
	r.PUT("/isadmin", controllers.UserIsAdminChange)

	r.PUT("/reset-password", controllers.UserResetPassword)
	r.DELETE("/delete", controllers.UserDelete)
}
