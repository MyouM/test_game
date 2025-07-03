package actions

import (
	"strings"
	"test_game/internal/data"
	"test_game/internal/data/items"
	"test_game/internal/data/places"
)

type GameData struct {
	*data.Data
}

func InitGameData() *GameData {
	return &GameData{
		data.InitData(),
	}
}

// Обработчик пользовательских запросов, перенаправляющий к заданым функциям и
// следящий за продвижением квестов
func (d *GameData) Actions(cmnds []string) string {
	if len(cmnds) < 1 {
		return "введите команду"
	}
	if d.User.CheckQuest(d.Quest) {
		d.Quest.ChangeQuest()
	}
	switch cmnds[0] {
	case "осмотреться":
		return d.LookAround(cmnds)
	case "надеть":
		return d.Wear(cmnds)
	case "взять":
		return d.Take(cmnds)
	case "идти":
		return d.Move(cmnds)
	case "применить":
		return d.Use(cmnds)
	default:
		return "неизвестная команда"
	}
	return ""
}

// Функция для команды "применить"
func (d *GameData) Use(cmnds []string) string {
	if len(cmnds) < 2 {
		return "а что применить?"
	}
	if len(cmnds) < 3 {
		return "а к чему применить?"
	}
	if !d.User.FindItem(cmnds[1]) {
		return "нет предмета в инвентаре - " + cmnds[1]
	}

	for _, plc := range d.Places[d.User.Location].Enter {
		if d.Places[plc].Lock.Barrier == cmnds[2] &&
			d.Places[plc].Lock.UnlockItem == cmnds[1] {
			d.Places[plc].UnlockBarrier()
			return d.Places[plc].Lock.Barrier +
				d.Places[plc].Lock.UnlockStatus
		}
	}
	return "не к чему применить"
}

// Функция для команды "идти"
func (d *GameData) Move(cmnds []string) string {
	if len(cmnds) < 2 {
		return "а куда идти?"
	}
	if !places.IsRoom(cmnds[1]) && !places.IsGlobalPlace(cmnds[1]) {
		return "нет такого места"
	}
	if !d.Places[d.User.Location].FindEnter(cmnds[1]) {
		return "нет пути в " + cmnds[1]
	}
	if d.Places[cmnds[1]].Lock.Locked {
		return d.Places[cmnds[1]].Lock.Barrier +
			d.Places[cmnds[1]].Lock.LockStatus
	}
	d.User.Location = cmnds[1]
	var answer strings.Builder
	answer.WriteString(d.Places[cmnds[1]].WelcomeText + "можно пройти - ")
	for i, plc := range d.Places[cmnds[1]].Enter {
		answer.WriteString(plc)
		if i+1 < len(d.Places[cmnds[1]].Enter) {
			answer.WriteString(", ")
		}
	}
	return answer.String()
}

// Функция для команды "осмотреться"
func (d *GameData) LookAround(cmnds []string) string {
	var (
		answer   strings.Builder
		checkObj bool
	)
	place := d.Places[d.User.Location]
	answer.WriteString(place.LookText)
	for i, obj := range place.Objects {
		if obj.Items[0] == "" {
			continue
		}
		if !place.CheckItems(i) {
			continue
		}
		if i > 0 && place.CheckItems(i) {
			answer.WriteString(", ")
		}
		checkObj = true
		answer.WriteString(obj.Name)
		for j, item := range obj.Items {
			if item != "" {
				answer.WriteString(item)
				if j+1 < len(obj.Items) && obj.Items[j+1] != "" {
					answer.WriteString(", ")
				}
			}
		}

	}
	if !checkObj {
		answer.WriteString("пустая комната")
	}
	if d.User.Location == "кухня" {
		answer.WriteString(", " + d.Quest.Text)
	} else {
		answer.WriteString(". ")
	}
	answer.WriteString("можно пройти - ")
	for i, plc := range place.Enter {
		answer.WriteString(plc)
		if i+1 < len(place.Enter) {
			answer.WriteString(", ")
		}
	}
	return answer.String()
}

// Функция для команды "взять"
func (d *GameData) Take(cmnds []string) string {
	if len(cmnds) < 2 {
		return "а что взять?"
	}
	if !d.Places[d.User.Location].FindItem(cmnds[1]) {
		return "нет такого"
	}
	if !d.User.Bag {
		return "некуда класть"
	}
	d.User.Inventory = append(d.User.Inventory, cmnds[1])
	d.Places[d.User.Location].DeleteItem(cmnds[1])
	return "предмет добавлен в инвентарь: " + cmnds[1]
}

// Функция для команды "надеть"
func (d *GameData) Wear(cmnds []string) string {
	if len(cmnds) < 2 {
		return "а что надеть?"
	}
	if !d.Places[d.User.Location].FindItem(cmnds[1]) {
		return "нет такого"
	}
	if !items.IsCloth(cmnds[1]) {
		return "это не одежда"
	}
	if cmnds[1] == "рюкзак" {
		d.User.Bag = true
	}
	d.User.Inventory = append(d.User.Inventory, cmnds[1])
	d.Places[d.User.Location].DeleteItem(cmnds[1])
	return "вы надели: " + cmnds[1]
}
