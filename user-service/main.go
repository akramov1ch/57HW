package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    v1 := router.Group("/user")
    {
        v1.POST("/", CreateUser)
        v1.GET("/:id", GetUser)
        v1.PUT("/:id", UpdateUser)
        v1.DELETE("/:id", DeleteUser)
    }
    router.Run(":8081")
}
