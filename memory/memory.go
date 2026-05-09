package memory

var ConversationMemory []string

func Save(text string) {

	ConversationMemory =
		append(
			ConversationMemory,
			text,
		)

	if len(ConversationMemory) > 25 {

		ConversationMemory =
			ConversationMemory[1:]
	}
}

func Get() []string {
	return ConversationMemory
}
