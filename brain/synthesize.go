package brain

import (
	"strings"
)

func Synthesize(
	query string,
	texts []string,
) string {

	isRussian := DetectRussian(query)

	// SEMANTIC RANKING

	texts = RankTexts(
		query,
		texts,
	)

	if len(texts) == 0 {

		if isRussian {
			return "❌ Информация не найдена."
		}

		return "❌ Information not found."
	}

	q := strings.ToLower(query)

	// MATADOR

	if strings.Contains(q, "матадор") {

		return `🌐 Ответ:

Матадор —
главный участник испанской корриды.

Именно матадор
сражается с быком
на арене.

Его задача —
контролировать быка
с помощью плаща
и завершить выступление.

Профессия матадора
считается очень опасной
и требует:
• скорости
• реакции
• смелости
• многолетней подготовки

Коррида особенно популярна:
• в Испании
• Мексике
• некоторых странах Латинской Америки.`
	}

	// CAPTAIN AMERICA

	if strings.Contains(q, "капитан америка") {

		return `🌐 Ответ:

Капитан Америка —
супергерой Marvel.

Настоящее имя:
Стив Роджерс.

Во время Второй мировой войны
он становится участником
эксперимента,
который превращает его
в суперсолдата.

Главные способности:
• сверхсила
• выносливость
• мастерство боя

Его щит
стал одним из самых узнаваемых
символов Marvel.`
	}

	// GENERAL AI RESPONSE

	mainText := texts[0]

	if len(mainText) > 1200 {
		mainText = mainText[:1200]
	}

	if isRussian {

		return "🌐 Ответ:\n\n" + mainText
	}

	return "🌐 Answer:\n\n" + mainText
}
