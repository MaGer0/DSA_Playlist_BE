package datastructures

import "DSA_Playlist_BE/models"

type PlaylistNode struct {
	Song *models.Song
	Next *PlaylistNode
}

type Playlist struct {
	Head *PlaylistNode
	Tail *PlaylistNode
	Size int
}

func (p *Playlist) Add(song *models.Song) {
	newNode := &PlaylistNode{
		Song: song,
	}

	if p.Head == nil {
		p.Head = newNode
		p.Tail = newNode
	} else {
		p.Tail.Next = newNode
		p.Tail = newNode
	}

	p.Size++
}

func (p *Playlist) Find(id int) *PlaylistNode {
	current := p.Head

	for current != nil {
		if current.Song.ID == id {
			return current
		}

		current = current.Next
	}

	return nil
}

func (p *Playlist) Delete(id int) bool {
	current := p.Head
	var previous *PlaylistNode

	for current != nil {
		if current.Song.ID == id {
			if previous == nil {
				p.Head = current.Next
			} else {
				previous.Next = current.Next
			}

			if current == p.Tail {
				p.Tail = previous
			}

			p.Size--

			return true
		}

		previous = current
		current = current.Next
	}

	return false
}

func (p *Playlist) Move(id int, targetIndex int) bool {
	if targetIndex < 0 || targetIndex >= p.Size {
		return false
	}

	if p.Head == nil {
		return false
	}

	node := p.Find(id)

	if node == nil {
		return false
	}

	song := node.Song

	p.Delete(id)

	newNode := &PlaylistNode{
		Song: song,
	}

	if targetIndex == 0 {
		newNode.Next = p.Head
		p.Head = newNode

		if p.Tail == nil {
			p.Tail = newNode
		}
	} else {
		current := p.Head

		for i := 0; i < targetIndex-1; i++ {
			current = current.Next
		}

		newNode.Next = current.Next
		current.Next = newNode

		if newNode.Next == nil {
			p.Tail = newNode
		}
	}

	p.Size++

	return true
}
