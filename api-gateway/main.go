package main

import (
    "io"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.Any("/user/*path", func(c *gin.Context) {
        proxy(c, "http://localhost:8081")
    })

    router.Any("/task/*path", func(c *gin.Context) {
        proxy(c, "http://localhost:8082")
    })

    router.Run(":8080")
}

func proxy(c *gin.Context, target string) {
    client := &http.Client{}
    req, err := http.NewRequest(c.Request.Method, target+c.Param("path"), c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    req.Header = c.Request.Header
    resp, err := client.Do(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()

    c.Status(resp.StatusCode)
    c.Writer.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

    if _, err := io.Copy(c.Writer, resp.Body); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    }
}
