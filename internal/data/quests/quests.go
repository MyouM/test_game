package quests

type Quest struct {
	Text     string
	Need     []string
	QuestNum int
}

// Описание всех квестов в игре
var allQuests = []Quest{
	{
		Text:     "надо собрать рюкзак и идти в универ. ",
		Need:     []string{"рюкзак", "конспекты",},
		QuestNum: 0,
	},
	{
		Text:     "надо идти в универ. ",
		Need:     []string{"гранит",},
		QuestNum: 1,
	},
}

// Продвижение по квестам
func (q *Quest) ChangeQuest() {
	*q = allQuests[q.QuestNum+1]
}

func InitQuests() Quest {
	return allQuests[0]
}
