package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    v1 := router.Group("/task")
    {
        v1.POST("/", CreateTask)
        v1.GET("/:id", GetTask)
        v1.PUT("/:id", UpdateTask)
        v1.DELETE("/:id", DeleteTask)
    }
    router.Run(":8082")
}
