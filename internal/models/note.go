package models

type Note struct {
    ID      string `json:"id" bson:"_id,omitempty"`
    Title   string `json:"title" bson:"title"`
    Content string `json:"content" bson:"content"`
}