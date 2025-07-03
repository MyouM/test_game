package data

import (
	"test_game/internal/data/places"
	"test_game/internal/data/quests"
	"test_game/internal/data/user"
)

type Data struct {
	Places places.Places
	User   user.User
	Quest  quests.Quest
}

func InitData() *Data {
	return &Data{
		Places: places.InitPlaces(),
		User:   user.InitUser(),
		Quest:  quests.InitQuests(),
	}
}
