package memory

type VectorMemory struct {
	Query   string
	Content string
}

var Memory []VectorMemory

func Store(query string, content string) {

	Memory = append(
		Memory,
		VectorMemory{
			Query:   query,
			Content: content,
		},
	)
}
