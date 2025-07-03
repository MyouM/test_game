package places

import (
	"slices"
)

type Places map[string]Env

type Lock struct {
	Locked       bool
	UnlockItem   string
	UnlockStatus string
	LockStatus   string
	Barrier      string
}

type Env struct {
	Objects     []Object
	Enter       []string
	LookText    string
	WelcomeText string
	Lock        *Lock
}

type Object struct {
	Name  string
	Items []string
}

var globalPlaces = map[string]string{
	"домой": "коридор",
	"улица": "улица",
}

var rooms = map[string]struct{}{
	"кухня":   struct{}{},
	"коридор": struct{}{},
	"комната": struct{}{},
	"улица":   struct{}{},
}

func IsRoom(plc string) bool {
	_, ok := rooms[plc]
	return ok
}

func IsGlobalPlace(plc string) bool {
	_, ok := globalPlaces[plc]
	return ok
}

func FromGlobToRoom(plc string) string {
	return globalPlaces[plc]
}

func InitPlaces() Places {
	return Places{
		"кухня": Env{
			Objects: []Object{
				{
					Name:  "на столе: ",
					Items: []string{"чай",},
				},
			},
			Enter:       []string{"коридор",},
			LookText:    "ты находишься на кухне, ",
			WelcomeText: "кухня, ничего интересного. ",
			Lock:        &Lock{Locked: false},
		},
		"коридор": Env{
			Objects:     []Object{},
			Enter:       []string{"кухня", "комната", "улица"},
			LookText:    "ты находишься в коридоре, ",
			WelcomeText: "ничего интересного. ",
			Lock:        &Lock{Locked: false},
		},
		"комната": Env{
			Objects: []Object{
				{
					Name:  "на столе: ",
					Items: []string{"ключи", "конспекты"},
				},
				{
					Name:  "на стуле: ",
					Items: []string{"рюкзак",},
				},
			},
			Enter:       []string{"коридор",},
			LookText:    "",
			WelcomeText: "ты в своей комнате. ",
			Lock:        &Lock{Locked: false},
		},
		"улица": Env{
			Objects:     []Object{},
			Enter:       []string{"домой",},
			LookText:    "ты на улице, ",
			WelcomeText: "на улице весна. ",
			Lock: &Lock{
				Locked:       true,
				UnlockItem:   "ключи",
				UnlockStatus: " открыта",
				LockStatus:   " закрыта",
				Barrier:      "дверь",
			},
		},
	}
}

func (e Env) CheckItems(i int) bool {
	check := false
	if i >= len(e.Objects) {
		return check
	}
	for _, item := range e.Objects[i].Items {
		if item != "" {
			check = true
		}
	}
	return check
}

func (e Env) UnlockBarrier() {
	e.Lock.Locked = false
}

func (e Env) IsBarrier(brr string) bool {
	if brr == e.Lock.Barrier {
		return true
	}
	return false
}

func (e Env) FindEnter(plc string) bool {
	for _, etr := range e.Enter {
		if plc == etr {
			return true
		}
	}
	return false
}

func (e Env) FindItem(item string) bool {
	for _, obj := range e.Objects {
		for _, thing := range obj.Items {
			if thing == item {
				return true
			}
		}
	}
	return false
}

func (e Env) DeleteItem(item string) {
	for i, obj := range e.Objects {
		for j, thing := range obj.Items {
			if thing == item {
				slices.Delete(e.Objects[i].Items, j, j+1)
				return
			}
		}
	}
	return
}
