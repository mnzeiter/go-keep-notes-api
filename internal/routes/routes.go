package routes

import (
    "github.com/gin-gonic/gin"
    "go-keep-notes/internal/handlers"
)

func Register(r *gin.Engine) {
    r.POST("/notes", handlers.CreateNote)
    r.GET("/notes", handlers.GetNotes)
    r.PUT("/notes/:id", handlers.UpdateNote)
    r.DELETE("/notes/:id", handlers.DeleteNote)
}