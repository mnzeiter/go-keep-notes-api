package main

import (
    "github.com/gin-gonic/gin"
    "go-keep-notes/internal/db"
    "go-keep-notes/internal/routes"
)

func main() {
    db.Connect()

    r := gin.Default()
    routes.Register(r)

    r.Run(":8080")
}