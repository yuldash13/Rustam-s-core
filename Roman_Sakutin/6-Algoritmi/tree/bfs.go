package main

const dist = 1e9

type Queue struct {
	q []int
}

func bfs(graph map[int][]int) []int {
	distance := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		distance[i] = dist
	}

	q := NewQueue()
	distance[0] = 0
	q.q = append(q.q, 0)

	for len(q.q) > 0 {
		cur := q.q[0]
		q.q = q.Pop()
		for _, r := range graph[cur] {
			if distance[r] == dist {
				distance[r] = distance[cur] + 1
				q.q = append(q.q, r)
			}
		}
	}
	return distance
}

func NewQueue() *Queue {
	return &Queue{q: make([]int, 0)}
}

func (q *Queue) Pop() []int {
	q.q = q.q[1:]
	return q.q
}
