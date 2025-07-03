package user

import "test_game/internal/data/quests"

type User struct {
	Inventory []string
	Location  string
	Bag       bool
}

func InitUser() User {
	return User{
		Location: "кухня",
		Bag:      false,
	}
}

func (u User) CheckQuest(q quests.Quest) bool {
	for _, qItem := range q.Need {
		check := false
		for _, item := range u.Inventory {
			if qItem == item {
				check = true
			}
		}
		if !check {
			return false
		}
	}
	return true
}

func (u User) FindItem(item string) bool {
	for _, thing := range u.Inventory {
		if thing == item {
			return true
		}
	}
	return false
}
