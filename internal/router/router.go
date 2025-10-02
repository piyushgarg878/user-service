package router

import (
    "user-service/internal/handler"

    "github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
    r := gin.Default()

    r.POST("/signup", userHandler.Register)
    r.GET("/users/profile/:id", userHandler.GetProfile)
	r.GET("/users/user/:name",userHandler.GetUserByName)
	r.PUT("/users/updateprofile/:id",userHandler.UpdateProfile)

    return r
}