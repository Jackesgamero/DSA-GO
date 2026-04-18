package graph

type (
	Vertex struct {
		Val   int
		Edges []*Vertex
	}
	queue     struct{ collection []*Vertex }
	stack     struct{ collection []*Vertex }
	container interface {
		Push(n *Vertex)
		Pop() *Vertex
		Len() int
	}
)

func IterativeTraversal(graph []*Vertex) ([]int, []int) {
	dfs := traverseAllNodes(new(stack), graph)
	bfs := traverseAllNodes(new(queue), graph)
	return bfs, dfs
}

func traverseAllNodes(c container, graph []*Vertex) []int {
	output := []int{}
	visited := make(map[*Vertex]struct{})
	c.Push(graph[0])

	for c.Len() != 0 {
		tmp := c.Pop()
		for _, neighbor := range tmp.Edges {
			c.Push(neighbor)
		}
		if _, ok := visited[tmp]; !ok {
			output = append(output, tmp.Val)
			visited[tmp] = struct{}{}
		}
	}
	return output
}

func newVertex(v int) *Vertex { return &Vertex{Val: v, Edges: []*Vertex{}} }

func (q *queue) Len() int       { return len(q.collection) }
func (q *queue) Push(n *Vertex) { q.collection = append(q.collection, n) }
func (q *queue) Pop() *Vertex {
	tmp := q.collection[0]
	q.collection = q.collection[1:len(q.collection)]
	return tmp
}

func (s *stack) Len() int       { return len(s.collection) }
func (s *stack) Push(n *Vertex) { s.collection = append(s.collection, n) }
func (s *stack) Pop() *Vertex {
	tmp := s.collection[len(s.collection)-1]
	s.collection = s.collection[:len(s.collection)-1]
	return tmp
}
