package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var users = make(map[string]User)

func CreateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    users[user.ID] = user
    c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
    id := c.Param("id")
    if user, exists := users[id]; exists {
        c.JSON(http.StatusOK, user)
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    }
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    users[id] = user
    c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    delete(users, id)
    c.JSON(http.StatusNoContent, gin.H{})
}
