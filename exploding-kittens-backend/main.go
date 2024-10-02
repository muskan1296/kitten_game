package main

import (
    "exploding-kittens-backend/db"
    "exploding-kittens-backend/handlers"
    "github.com/gin-gonic/gin"
    "os"
)

func main() {
    // Connect to MongoDB
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017"
    }
    db.ConnectMongoDB(mongoURI)

    // Set up Gin router
    r := gin.Default()

    // Define routes
    r.GET("/draw", handlers.DrawCard)
    r.POST("/shuffle", handlers.ShuffleDeck)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
