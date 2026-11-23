package main

import (
	"DSA_Playlist_BE/handlers"
	"DSA_Playlist_BE/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Konfigurasi CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
		},
	}))

	playlistService := services.NewPlaylistService()
	playlistHandler := handlers.NewPlaylistHandler(playlistService)

	api := router.Group("/api")
	{
		api.GET("/playlist", playlistHandler.GetPlaylist)
		api.POST("/playlist", playlistHandler.AddToPlaylist)
		api.PATCH("/playlist/:id", playlistHandler.MoveSong)
		api.DELETE("/playlist", playlistHandler.DeleteFromPlaylist)

		api.POST("/priority", playlistHandler.AddToPriority)

		api.POST("/next", playlistHandler.Next)
		api.POST("/prev", playlistHandler.Prev)
		api.POST("/action", playlistHandler.PlaySong)
	}

	log.Println("Server running on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
