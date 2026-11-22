package services

import (
	"DSA_Playlist_BE/datastructures"
	"DSA_Playlist_BE/models"
	"errors"
)

type PlaylistService struct {
	Songs         []*models.Song
	Playlist      datastructures.Playlist
	PriorityQueue datastructures.PriorityQueue
	History       datastructures.History

	CurrentSong       *models.Song
	PlaylistCurrentID int
}

func NewPlaylistService() *PlaylistService {
	songs := []*models.Song{
		{
			ID:     1,
			Judul:  "Lagu A",
			Artis:  "Artis A",
			Durasi: 200,
		},
		{
			ID:     2,
			Judul:  "Lagu B",
			Artis:  "Artis B",
			Durasi: 180,
		},
		{
			ID:     3,
			Judul:  "Lagu C",
			Artis:  "Artis C",
			Durasi: 220,
		},
		{
			ID:     4,
			Judul:  "Lagu D",
			Artis:  "Artis D",
			Durasi: 240,
		},
	}

	service := &PlaylistService{
		Songs: songs,
	}

	service.Playlist.Add(songs[0])
	service.Playlist.Add(songs[1])
	service.Playlist.Add(songs[2])

	return service
}

func (s *PlaylistService) FindSong(id int) (*models.Song, error) {
	for _, song := range s.Songs {
		if song.ID == id {
			return song, nil
		}
	}

	return nil, errors.New("song not found")
}

func (s *PlaylistService) AddToPlaylist(id int) error {
	song, err := s.FindSong(id)

	if err != nil {
		return err
	}

	if s.Playlist.Find(id) != nil {
		return errors.New("song already in playlist")
	}

	s.Playlist.Add(song)

	return nil
}

func (s *PlaylistService) DeleteFromPlaylist(id int) error {
	deleted := s.Playlist.Delete(id)

	if !deleted {
		return errors.New("song not found in playlist")
	}

	if s.PlaylistCurrentID == id {
		s.PlaylistCurrentID = 0
	}

	return nil
}

func (s *PlaylistService) MoveSong(id int, targetIndex int) error {
	if targetIndex < 0 || targetIndex >= s.Playlist.Size {
		return errors.New("invalid target index")
	}

	moved := s.Playlist.Move(id, targetIndex)

	if !moved {
		return errors.New("song not found in playlist")
	}

	return nil
}

func (s *PlaylistService) AddToPriority(id int) error {
	song, err := s.FindSong(id)

	if err != nil {
		return err
	}

	s.PriorityQueue.Enqueue(song)

	return nil
}

func (s *PlaylistService) PlaySong(id int) (*models.Song, error) {
	song, err := s.FindSong(id)

	if err != nil {
		return nil, err
	}

	s.CurrentSong = song
	s.History.Add(song)

	node := s.Playlist.Find(id)

	if node != nil {
		s.PlaylistCurrentID = id
	}

	return song, nil
}

func (s *PlaylistService) Next() (*models.Song, error) {

	// 1. Periksa apakah masih ada history berikutnya.
	if song := s.History.Next(); song != nil {
		s.CurrentSong = song

		return song, nil
	}

	// 2. Periksa Priority Queue.
	if s.PriorityQueue.Size > 0 {
		song := s.PriorityQueue.Dequeue()

		s.CurrentSong = song
		s.History.Add(song)

		return song, nil
	}

	// 3. Cari lagu berikutnya di Playlist.
	var nextNode *datastructures.PlaylistNode

	if s.PlaylistCurrentID == 0 {
		nextNode = s.Playlist.Head
	} else {
		currentNode := s.Playlist.Find(s.PlaylistCurrentID)

		if currentNode != nil {
			nextNode = currentNode.Next
		}
	}

	if nextNode == nil {
		return nil, errors.New("no next song available")
	}

	song := nextNode.Song

	s.PlaylistCurrentID = song.ID
	s.CurrentSong = song

	s.History.Add(song)

	return song, nil
}

func (s *PlaylistService) Prev() (*models.Song, error) {
	song := s.History.Previous()

	if song == nil {
		return nil, errors.New("no previous song available")
	}

	s.CurrentSong = song

	return song, nil
}
