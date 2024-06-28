package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Task struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Status string `json:"status"`
}

var tasks = make(map[string]Task)

func CreateTask(c *gin.Context) {
    var task Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    tasks[task.ID] = task
    c.JSON(http.StatusCreated, task)
}

func GetTask(c *gin.Context) {
    id := c.Param("id")
    if task, exists := tasks[id]; exists {
        c.JSON(http.StatusOK, task)
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
    }
}

func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var task Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    tasks[id] = task
    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    delete(tasks, id)
    c.JSON(http.StatusNoContent, gin.H{})
}
