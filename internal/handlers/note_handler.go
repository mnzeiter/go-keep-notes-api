package handlers

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "go-keep-notes/internal/db"
    "go-keep-notes/internal/models"
)

func CreateNote(c *gin.Context) {
    var note models.Note
    if err := c.BindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    result, err := db.NotesCollection().InsertOne(ctx, note)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    note.ID = result.InsertedID.(primitive.ObjectID).Hex()
    c.JSON(http.StatusCreated, note)
}

func GetNotes(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := db.NotesCollection().Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var notes []models.Note
    for cursor.Next(ctx) {
        var note models.Note
        cursor.Decode(&note)
        notes = append(notes, note)
    }

    c.JSON(http.StatusOK, notes)
}

func UpdateNote(c *gin.Context) {
    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)

    var note models.Note
    if err := c.BindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    update := bson.M{"$set": bson.M{
        "title":   note.Title,
        "content": note.Content,
    }}

    _, err := db.NotesCollection().UpdateOne(ctx, bson.M{"_id": objID}, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    note.ID = id
    c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {
    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := db.NotesCollection().DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"deleted": id})
}
