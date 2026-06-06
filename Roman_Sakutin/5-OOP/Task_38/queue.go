package main

type Queue struct {
	clients []*Client
	start   int
	end     int
}

func NewQueue(len int, start int, end int) *Queue {
	return &Queue{
		clients: make([]*Client, len),
		start:   start,
		end:     end,
	}
}

func (q *Queue) Enqueue(client *Client) {
	q.clients[q.end] = client
	q.end = (q.end + 1) % len(q.clients)
}

func (q *Queue) Dequeue() *Client {
	client := q.clients[q.start]
	q.start = (q.start + 1) % len(q.clients)
	return client
}
