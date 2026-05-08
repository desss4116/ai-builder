package memory

var ConversationMemory []string

func Save(data string) {

	ConversationMemory = append(
		ConversationMemory,
		data,
	)
}
