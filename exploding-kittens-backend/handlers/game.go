package handlers

import (
	"context"
	"exploding-kittens-backend/db"
	"exploding-kittens-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DrawCard - Endpoint to draw a card from the deck
func DrawCard(c *gin.Context) {
    collection := db.GetCollection("exploding_kitten", "cards")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var card models.Card
    err := collection.FindOneAndDelete(ctx, bson.M{}).Decode(&card)
    if err == mongo.ErrNoDocuments {
        c.JSON(http.StatusNotFound, gin.H{"error": "No cards left in the deck"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, card)
}

// ShuffleDeck - Endpoint to shuffle the deck and refill it with cards
func ShuffleDeck(c *gin.Context) {
    collection := db.GetCollection("exploding_kittens", "deck")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Remove all existing cards
    _, err := collection.DeleteMany(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Insert new cards (example deck with different types)
    cards := []interface{}{
        models.Card{Type: "Cat"},
        models.Card{Type: "Exploding Kitten"},
        models.Card{Type: "Defusing"},
        models.Card{Type: "Shuffle"},
    }

    _, err = collection.InsertMany(ctx, cards)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Deck shuffled successfully!"})
}
