package database

func AddKnowledge(
	category string,
	title string,
	content string,
) {

	DB.Exec(`
	INSERT INTO knowledge_base(
	category,
	title,
	content
	)
	VALUES(?,?,?)
	`,
		category,
		title,
		content,
	)
}
