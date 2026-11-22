package datastructures

import "DSA_Playlist_BE/models"

type HistoryNode struct {
	Song *models.Song
	Prev *HistoryNode
	Next *HistoryNode
}

type History struct {
	Current *HistoryNode
}

func (h *History) Add(song *models.Song) {
	newNode := &HistoryNode{
		Song: song,
	}

	if h.Current == nil {
		h.Current = newNode
		return
	}

	// Hapus history setelah posisi current.
	h.Current.Next = nil

	newNode.Prev = h.Current
	h.Current.Next = newNode

	h.Current = newNode
}

func (h *History) Previous() *models.Song {
	if h.Current == nil || h.Current.Prev == nil {
		return nil
	}

	h.Current = h.Current.Prev

	return h.Current.Song
}

func (h *History) Next() *models.Song {
	if h.Current == nil || h.Current.Next == nil {
		return nil
	}

	h.Current = h.Current.Next

	return h.Current.Song
}
