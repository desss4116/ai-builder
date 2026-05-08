package memory

var ConversationMemory []string

func Save(item string) {

	ConversationMemory =
		append(
			ConversationMemory,
			item,
		)

	if len(ConversationMemory) > 20 {

		ConversationMemory =
			ConversationMemory[1:]
	}
}

func Get() []string {

	return ConversationMemory
}
