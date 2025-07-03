package items

var items = map[string]struct{}{
	"ключи":     struct{}{},
	"конспекты": struct{}{},
	"чай":       struct{}{},
}

var cloths = map[string]struct{}{
	"рюкзак": struct{}{},
}

func IsCloth(cloth string) bool {
	_, ok := cloths[cloth]
	return ok
}

func IsItem(item string) bool {
	_, ok := items[item]
	return ok
}
