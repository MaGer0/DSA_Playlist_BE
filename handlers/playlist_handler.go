package handlers

import (
	"DSA_Playlist_BE/models"
	"DSA_Playlist_BE/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlaylistHandler struct {
	Service *services.PlaylistService
}

func NewPlaylistHandler(service *services.PlaylistService) *PlaylistHandler {
	return &PlaylistHandler{
		Service: service,
	}
}

func (h *PlaylistHandler) GetPlaylist(c *gin.Context) {
	var playlist []models.Song

	current := h.Service.Playlist.Head

	for current != nil {
		playlist = append(playlist, *current.Song)
		current = current.Next
	}

	c.JSON(http.StatusOK, gin.H{
		"playlist": playlist,
		"Songs":    h.Service.Songs,
	})
}

func (h *PlaylistHandler) AddToPlaylist(c *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := h.Service.AddToPlaylist(request.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Song added to playlist",
	})
}

func (h *PlaylistHandler) DeleteFromPlaylist(c *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := h.Service.DeleteFromPlaylist(request.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Song deleted from playlist",
	})
}

func (h *PlaylistHandler) MoveSong(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid song ID",
		})
		return
	}

	var request struct {
		ID          int `json:"id"`
		TargetIndex int `json:"target_index"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	if request.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Song ID does not match URL parameter",
		})
		return
	}

	err = h.Service.MoveSong(id, request.TargetIndex)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Song position updated",
	})
}

func (h *PlaylistHandler) AddToPriority(c *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := h.Service.AddToPriority(request.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Song added to priority queue",
	})
}

func (h *PlaylistHandler) Next(c *gin.Context) {
	song, err := h.Service.Next()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, song)
}

func (h *PlaylistHandler) Prev(c *gin.Context) {
	song, err := h.Service.Prev()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, song)
}

func (h *PlaylistHandler) PlaySong(c *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	song, err := h.Service.PlaySong(request.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, song)
}
