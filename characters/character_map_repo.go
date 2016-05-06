package characters

type CharacterMapRepo interface {
	GetCharacterMap() map[string]int
}
