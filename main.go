package main

import (
	"DSA_Playlist_BE/handlers"
	"DSA_Playlist_BE/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize Service
	playlistService := services.NewPlaylistService()

	// Initialize Handler
	playlistHandler := handlers.NewPlaylistHandler(playlistService)

	// API Routes
	api := router.Group("/api")
	{
		// Playlist
		api.GET("/playlist", playlistHandler.GetPlaylist)
		api.POST("/playlist", playlistHandler.AddToPlaylist)
		api.PATCH("/playlist/:id", playlistHandler.MoveSong)
		api.DELETE("/playlist", playlistHandler.DeleteFromPlaylist)

		// Priority Queue
		api.POST("/priority", playlistHandler.AddToPriority)

		// Playback Action
		api.POST("/action/next", playlistHandler.Next)
		api.POST("/action/prev", playlistHandler.Prev)
		api.POST("/action", playlistHandler.PlaySong)
	}

	// Start Server
	log.Println("Server running on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
