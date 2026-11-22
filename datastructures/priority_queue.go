package datastructures

import "DSA_Playlist_BE/models"

type QueueNode struct {
	Song *models.Song
	Next *QueueNode
}

type PriorityQueue struct {
	Front *QueueNode
	Rear  *QueueNode
	Size  int
}

func (q *PriorityQueue) Enqueue(song *models.Song) {
	newNode := &QueueNode{
		Song: song,
	}

	if q.Front == nil {
		q.Front = newNode
		q.Rear = newNode
	} else {
		q.Rear.Next = newNode
		q.Rear = newNode
	}

	q.Size++
}

func (q *PriorityQueue) Dequeue() *models.Song {
	if q.Front == nil {
		return nil
	}

	song := q.Front.Song

	q.Front = q.Front.Next
	q.Size--

	if q.Front == nil {
		q.Rear = nil
	}

	return song
}
